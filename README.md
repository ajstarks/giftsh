# giftsh -- a DSL and shell for editing images

![image-edit](dsh-sublime.png)

gift is a little language (aka domain-specific language) for editing images.

commands are read one line at time (optional comments are skipped)
and image opertations are performed in order on the image.

For example, running giftsh:

```giftsh < test.gsh > f.jpg```

with this script in test.gsh

```
read ajs.jpg
gray
pixelate 20
emboss
sobel
```

reads an image, converts it to grayscale, pixelates it, and applies emboss and sobel filters. At the end of the script the result is written to ```f.jpg```.

By default giftsh reads commands from standard in, and writes results to standard out.

## options

```
Usage of giftsh:
  -c filename (script file)
  -h show command set
  -o filename (result file)
  -w file     (file to watch changes interactively)
```

## command set

```
Command         Parameters
blur            value
brightness      value (-100, 100)
colorbalance    red green blue (percentages)
colorize        hue (0-360) saturation (0-100) percentage (0-100)
contrast        value (-100, 100)
crop            x1 y1 x2 y2 (rectangle at (x1,y1) and (x2,y2)
cropsize        width height
edge            edge filter
emboss          emboss image
fliph           flip horizontal
flipv           flip vertical
gamma           value
gray            grayscale image
help            show command set
hue             value (-180, 180)
invert          invert image
max             local maximum (kernel size)
mean            local mean filter (kernel size)
median          local median filter (kernel size)
min             local minimum (kernel size)
opacity         percentage (0-100)
pixelate        pixels
read            imagefile (open source file)
resize          width height
resizefill      width height
resizefit       width height
rotate          degrees counter-clockwise
saturation      value (-100, 500)
sepia           sepia percentage (0-100)
sigmoid         sigmoid contrast (midpoint factor)
sobel           sobel filter
threshold       color threshold percentage (0-100)
transpose       flip horizontally and rotate 90° counter-clockwise
transverse      flips vertically and rotate 90° counter-clockwise
unsharp         unsharp mask (sigma amount threshold)
```