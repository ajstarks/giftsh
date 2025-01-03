// giftsh -- a little language for image editing using the Go image filtering toolkit (gift)
package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/disintegration/gift"
)

// commandTable documents commands and their arguments
var commandTable = map[string]string{
	"blur":         "value > 0",
	"brightness":   "value (-100, 100)",
	"colorbalance": "red green blue (-100, 500)",
	"colorize":     "hue (0-360) saturation (0-100) percentage (0-100)",
	"colorspace":   "l for linear->sRGB or s for sRGB->linear",
	"contrast":     "value (-100, 100)",
	"crop":         "x1 y1 x2 y2 (rectangle at (x1,y1) and (x2,y2)",
	"cropsize":     "width height",
	"edge":         "edge filter",
	"emboss":       "emboss filter",
	"fliph":        "flip horizontal",
	"flipv":        "flip vertical",
	"gamma":        "value (< 1 darken, > 1 lighten)",
	"gray":         "grayscale image",
	"help":         "show command set",
	"hue":          "value (-180, 180)",
	"invert":       "invert image",
	"max":          "local maximum size (odd positive integer)",
	"mean":         "local mean size (odd positive integer)",
	"median":       "local median size (odd positive integer)",
	"min":          "local minimum size (odd positive integer)",
	"opacity":      "value (0-100)",
	"pixelate":     "pixels",
	"read":         "imagefile (open source file)",
	"reset":        "discard image edits (watch mode only)",
	"resize":       "width height",
	"resizefill":   "width height",
	"resizefit":    "width height",
	"rotate":       "degrees counter-clockwise",
	"saturation":   "value (-100, 500)",
	"sepia":        "value (0-100)",
	"sigmoid":      "midpoint (0,1) factor (-10,10)",
	"sobel":        "sobel filter",
	"threshold":    "color threshold percentage (0-100)",
	"transpose":    "flip horizontally and rotate 90° counter-clockwise",
	"transverse":   "flip vertically and rotate 90° counter-clockwise",
	"unsharp":      "sigma (> 0) amount (0.5, 1.5) threshold (0, 0.05)",
}

// atof converts a string to a float32 value
func atof(s string) float32 {
	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0
	}
	return float32(v)
}

// perror prints an error message
func perror(s []string, linenumber int) {
	if len(s) < 1 {
		fmt.Fprintf(os.Stderr, "invalid command\n")
	}
	cmd := s[0]
	es, ok := commandTable[cmd]
	if !ok {
		es = "invalid command"
	}
	fmt.Fprintf(os.Stderr, "line %d: %s %s\n", linenumber, cmd, es)
}

// help prints usage and command summaries
func help() {
	const helpfmt = "%-15s %s\n"
	keys := make([]string, 0, len(commandTable))
	for k := range commandTable {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Fprintf(os.Stderr, helpfmt, "Command", "Parameters")
	for _, k := range keys {
		fmt.Fprintf(os.Stderr, helpfmt, k, commandTable[k])
	}
}

// readimage opens a file, and returns the image and format
func readimage(s []string, linenumber int) (image.Image, string, error) {
	if len(s) < 1 {
		perror(s, linenumber)
	}
	fname := s[1]
	r, err := os.Open(fname)
	if err != nil {
		return nil, "", err
	}
	img, format, err := image.Decode(r)
	if err != nil {
		return nil, "", err
	}
	return img, format, nil
}

// writeimage writes the image data
func writeimage(w io.Writer, src image.Image, format string, g *gift.GIFT) {
	if src == nil {
		return
	}
	dst := image.NewRGBA(g.Bounds(src.Bounds()))
	g.Draw(dst, src)
	switch format {
	case "png":
		png.Encode(w, dst)
	case "jpeg":
		jpeg.Encode(w, dst, nil)
	case "gif":
		gif.Encode(w, dst, nil)
	}
}

// Image transformation functions

// blur blurs the image
func blur(s []string, g *gift.GIFT, linenumber int) {
	if len(s) < 2 {
		perror(s, linenumber)
		return
	}
	g.Add(gift.GaussianBlur(atof(s[1])))
}

// brightness adjusts the image brightness
func brightness(s []string, g *gift.GIFT, linenumber int) {
	if len(s) < 2 {
		perror(s, linenumber)
		return
	}
	brvalue := atof(s[1])
	if !(brvalue >= -100 && brvalue <= 100) {
		perror(s, linenumber)
		return
	}
	g.Add(gift.Brightness(brvalue))
}

// colorbalance updates the color balance
func colorbalance(s []string, g *gift.GIFT, linenumber int) {
	if len(s) < 4 {
		perror(s, linenumber)
		return
	}
	pctred, pctgreen, pctblue := atof(s[1]), atof(s[2]), atof(s[3])
	g.Add(gift.ColorBalance(pctred, pctgreen, pctblue))
}

// colorize colorizes the image
func colorize(s []string, g *gift.GIFT, linenumber int) {
	if len(s) < 4 {
		perror(s, linenumber)
		return
	}
	chue, csaturation, cpercent := atof(s[1]), atof(s[2]), atof(s[3])
	g.Add(gift.Colorize(chue, csaturation, cpercent))
}

// colorspace converts to and from linear and sRGB colorspaces
func colorspace(s []string, g *gift.GIFT, linenumber int) {
	if len(s) < 2 {
		perror(s, linenumber)
		return
	}
	switch s[1] {
	case "linear", "l":
		g.Add(gift.ColorspaceSRGBToLinear())
	case "sRGB", "s":
		g.Add(gift.ColorspaceLinearToSRGB())
	default:
		perror(s, linenumber)
	}
}

// opacity sets the image opacity (0-100%)
func opacity(s []string, g *gift.GIFT, linenumber int) {
	if len(s) < 2 {
		perror(s, linenumber)
		return
	}
	opacityvalue := atof(s[1])
	if !(opacityvalue >= 0 && opacityvalue <= 100) {
		perror(s, linenumber)
		return
	}
	g.Add(gift.ColorFunc(func(r0, g0, b0, a0 float32) (r, g, b, a float32) {
		return r0, g0, b0, opacityvalue / 100
	}))
}

// gamma set images's Gamma value
func gamma(s []string, g *gift.GIFT, linenumber int) {
	if len(s) < 2 {
		perror(s, linenumber)
		return
	}
	gammavalue := atof(s[1])
	g.Add(gift.Gamma(gammavalue))
}

// hue sets the image hue
func hue(s []string, g *gift.GIFT, linenumber int) {
	if len(s) < 2 {
		perror(s, linenumber)
		return
	}
	hvalue := atof(s[1])
	if !(hvalue >= -180 && hvalue <= 180) {
		perror(s, linenumber)
		return
	}
	g.Add(gift.Hue(hvalue))
}

// contrast sets the image contrast
func contrast(s []string, g *gift.GIFT, linenumber int) {
	if len(s) < 2 {
		perror(s, linenumber)
		return
	}
	value := atof(s[1])
	if !(value >= -100 && value <= 100) {
		perror(s, linenumber)
		return
	}
	g.Add(gift.Contrast(value))
}

// saturation sets the image saturation
func staturation(s []string, g *gift.GIFT, linenumber int) {
	if len(s) < 2 {
		perror(s, linenumber)
		return
	}
	value := atof(s[1])
	if !(value >= -100 && value <= 500) {
		perror(s, linenumber)
		return
	}
	g.Add(gift.Saturation(value))
}

// sepia colorizes with a sepia value (0-100)
func sepia(s []string, g *gift.GIFT, linenumber int) {
	if len(s) < 2 {
		perror(s, linenumber)
		return
	}
	value := atof(s[1])
	if !(value >= 0 && value <= 100) {
		perror(s, linenumber)
		return
	}
	g.Add(gift.Sepia(value))
}

// rotate rotates the image counter-clockwise degrees
func rotate(s []string, g *gift.GIFT, linenumber int) {
	if len(s) < 2 {
		perror(s, linenumber)
		return
	}
	rotvalue := atof(s[1])
	if rotvalue > 0 && rotvalue <= 360 {
		switch rotvalue {
		case 90:
			g.Add(gift.Rotate90())
		case 180:
			g.Add(gift.Rotate180())
		case 270:
			g.Add(gift.Rotate270())
		default:
			g.Add(gift.Rotate(rotvalue, color.White, gift.LinearInterpolation))
		}
	}
}

// localk sets min, max, median, and mean image local kernel values
func localk(s []string, g *gift.GIFT, linenumber int) {
	if len(s) < 2 {
		perror(s, linenumber)
		return
	}
	value, err := strconv.Atoi(s[1])
	if !(value > 0 && value%1 == 0) || err != nil {
		perror(s, linenumber)
		return
	}
	switch s[0] {
	case "minimum":
		g.Add(gift.Minimum(value, true))
	case "maximum":
		g.Add(gift.Maximum(value, true))
	case "median":
		g.Add(gift.Median(value, true))
	case "mean":
		g.Add(gift.Mean(value, true))
	}
}

// pixelate pixelates the image by the specified number of pixels
func pixelate(s []string, g *gift.GIFT, linenumber int) {
	if len(s) < 2 {
		perror(s, linenumber)
		return
	}
	pixelatevalue, err := strconv.Atoi(s[1])
	if err != nil {
		perror(s, linenumber)
		return
	}
	g.Add(gift.Pixelate(pixelatevalue))
}

// resize resizes/crops by width, height
func resize(s []string, g *gift.GIFT, linenumber int) {
	if len(s) < 3 {
		perror(s, linenumber)
		return
	}
	width, err := strconv.Atoi(s[1])
	if err != nil {
		perror(s, linenumber)
		return
	}
	height, err := strconv.Atoi(s[2])
	if err != nil {
		perror(s, linenumber)
		return
	}
	switch s[0] {
	case "cropsize":
		g.Add(gift.CropToSize(width, height, gift.CenterAnchor))
	case "resize":
		g.Add(gift.Resize(width, height, gift.LanczosResampling))
	case "resizefill":
		g.Add(gift.ResizeToFill(width, height, gift.LanczosResampling, gift.CenterAnchor))
	case "resizefit":
		g.Add(gift.ResizeToFit(width, height, gift.LanczosResampling))
	}
}

// crop crops the image by the rectangle defined by (x1,y1) and (x2,y2)
func crop(s []string, g *gift.GIFT, linenumber int) {
	if len(s) < 5 {
		perror(s, linenumber)
		return
	}
	x1, err := strconv.Atoi(s[1])
	if err != nil {
		perror(s, linenumber)
	}
	y1, err := strconv.Atoi(s[2])
	if err != nil {
		perror(s, linenumber)
	}
	x2, err := strconv.Atoi(s[3])
	if err != nil {
		perror(s, linenumber)
	}
	y2, err := strconv.Atoi(s[4])
	if err != nil {
		perror(s, linenumber)
	}
	g.Add(gift.Crop(image.Rect(x1, y1, x2, y2)))
}

// sigmoid applies a sigmod filter
func sigmoid(s []string, g *gift.GIFT, linenumber int) {
	if len(s) < 3 {
		perror(s, linenumber)
		return
	}
	midpoint, factor := atof(s[1]), atof(s[2])
	g.Add(gift.Sigmoid(midpoint, factor))
}

// threshold applies a threshold filter
func threshold(s []string, g *gift.GIFT, linenumber int) {
	if len(s) < 2 {
		perror(s, linenumber)
		return
	}
	threshvalue := atof(s[1])
	if !(threshvalue >= 0 && threshvalue <= 100) {
		perror(s, linenumber)
		return
	}
	g.Add(gift.Threshold(threshvalue))
}

// unsharp applies the unsharining image filter
func unsharp(s []string, g *gift.GIFT, linenumber int) {
	if len(s) < 4 {
		perror(s, linenumber)
		return
	}
	sigma, amount, threshold := atof(s[0]), atof(s[1]), atof(s[2])
	g.Add(gift.UnsharpMask(sigma, amount, threshold))
}

// parse reads line of data, performing image operations
func parse(s []string, g *gift.GIFT, linenumber int) {
	if len(s) < 1 {
		return
	}
	switch s[0] {
	case "help", "?", "h":
		help()
	case "blur":
		blur(s, g, linenumber)
	case "brightness":
		brightness(s, g, linenumber)
	case "colorbalance":
		colorbalance(s, g, linenumber)
	case "colorize":
		colorize(s, g, linenumber)
	case "colorspace":
		colorspace(s, g, linenumber)
	case "contrast":
		contrast(s, g, linenumber)
	case "crop":
		crop(s, g, linenumber)
	case "gray":
		g.Add(gift.Grayscale())
	case "fliph":
		g.Add(gift.FlipHorizontal())
	case "flipv":
		g.Add(gift.FlipVertical())
	case "invert", "neg":
		g.Add(gift.Invert())
	case "transpose":
		g.Add(gift.Transpose())
	case "transverse":
		g.Add(gift.Transverse())
	case "edge":
		g.Add(gift.Convolution([]float32{-1, -1, -1, -1, 8, -1, -1, -1, -1}, false, false, false, 0.0))
	case "emboss":
		g.Add(gift.Convolution([]float32{-1, -1, 0, -1, 1, 1, 0, 1, 1}, false, false, false, 0.0))
	case "gamma":
		gamma(s, g, linenumber)
	case "hue":
		hue(s, g, linenumber)
	case "min", "max", "mean", "median":
		localk(s, g, linenumber)
	case "opacity":
		opacity(s, g, linenumber)
	case "pixelate", "pix":
		pixelate(s, g, linenumber)
	case "cropsize", "resize", "resizefill", "resizefit":
		resize(s, g, linenumber)
	case "rotate", "rot":
		rotate(s, g, linenumber)
	case "saturation", "sat":
		staturation(s, g, linenumber)
	case "sepia":
		sepia(s, g, linenumber)
	case "sigmoid":
		sigmoid(s, g, linenumber)
	case "sobel":
		g.Add(gift.Sobel())
	case "threshold":
		threshold(s, g, linenumber)
	case "unsharp":
		unsharp(s, g, linenumber)
	}
}

// process reads lines, parsing giftsh commands, performs image operations
func process(w io.Writer, r io.Reader, srcfile, watchfile string, g *gift.GIFT) {
	scanner := bufio.NewScanner(r)
	lw := len(watchfile)
	var src image.Image
	var format string
	var err error

	// if a sourefile is specified, open it
	if len(srcfile) > 0 {
		src, format, err = readimage([]string{"read", srcfile}, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			return
		}
	}

	// loop over lines of input, processing commands and counting lines.
	for n := 1; scanner.Scan(); n++ {
		t := scanner.Text() // line of text

		// skip comments
		if strings.HasPrefix(t, "#") || strings.HasPrefix(t, "//") {
			continue
		}

		// break lines into commands and arguments
		line := strings.Fields(t)

		// open the source image, continue processing
		if len(line) > 1 && (line[0] == "read" || line[0] == "r") {
			src, format, err = readimage(line, n)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				return
			}
			continue
		}
		// parse commands, processing the image
		parse(line, g, n)

		// if watching, write the result, process resets
		if lw > 0 {
			wf, err := os.Create(watchfile)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				return
			}
			if len(line) > 0 && line[0] == "reset" {
				g = gift.New()
			}
			writeimage(wf, src, format, g)
		}

	}
	// write the final result (if not watching)
	if lw == 0 {
		writeimage(w, src, format, g)
	}
}

func main() {
	// command flags
	var scriptfile, outputfile, sourcefile, watchfile string
	var showhelp bool
	flag.StringVar(&scriptfile, "c", "", "script filename")
	flag.StringVar(&outputfile, "o", "", "output filename")
	flag.StringVar(&sourcefile, "f", "", "source image filename")
	flag.StringVar(&watchfile, "w", "", "file to watch changes interactively")
	flag.BoolVar(&showhelp, "h", false, "show command set")
	flag.Parse()

	var cmdr io.Reader = os.Stdin
	var outw io.Writer = os.Stdout
	var err error

	// if specifies show help and exit
	if showhelp {
		flag.Usage()
		fmt.Fprintln(os.Stderr)
		help()
		os.Exit(1)
	}

	// process input for the script
	if len(scriptfile) > 0 {
		cmdr, err = os.Open(scriptfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(2)
		}
	}

	// process output file destination
	if len(outputfile) > 0 {
		outw, err = os.Create(outputfile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(3)
		}
	}

	// process
	process(outw, cmdr, sourcefile, watchfile, gift.New())
}
