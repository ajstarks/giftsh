#!/bin/sh
go build || exit 1
if test $# -lt 1
then
  echo "specify a file" 1>&2
  exit 2
fi
input="$1"
if test ! -r "$input"
then
	echo "cannot read $input" 1>&2
	exit 3
fi
enc=`echo "$input" | awk -F\. '{print $NF}'`

(echo r "$input"; echo blur 10) 							| ./giftsh > blur.$enc
(echo r "$input"; echo brightness 20) 				| ./giftsh > bright20.$enc
(echo r "$input"; echo brightness -20) 				| ./giftsh > bright-20.$enc
(echo r "$input"; echo contrast 30) 					| ./giftsh > contrast30.$enc
(echo r "$input"; echo contrast -30) 					| ./giftsh > contrast-30.$enc
(echo r "$input"; echo colorize 240 50 100) 	| ./giftsh > colorize.$enc
(echo r "$input"; echo colorspace l) 					| ./giftsh > colorspace-l.$enc
(echo r "$input"; echo colorspace s) 					| ./giftsh > colorspace-s.$enc
(echo r "$input"; echo colorbalance 20 -20 0) | ./giftsh > colorbalance.$enc
(echo r "$input"; echo crop 90 90 250 250) 		| ./giftsh > crop.$enc
(echo r "$input"; echo cropsize 100 100) 			| ./giftsh > cropsize.$enc
(echo r "$input"; echo edge) 									| ./giftsh > edge.$enc
(echo r "$input"; echo emboss) 								| ./giftsh > emboss.$enc
(echo r "$input"; echo fliph) 								| ./giftsh > fliph.$enc
(echo r "$input"; echo flipv) 								| ./giftsh > flipv.$enc
(echo r "$input"; echo gamma 1.5) 						| ./giftsh > gamma.$enc
(echo r "$input"; echo gray) 									| ./giftsh > gray.$enc
(echo r "$input"; echo hue 45) 								| ./giftsh > hue45.$enc
(echo r "$input"; echo hue -45) 							| ./giftsh > hue-45.$enc
(echo r "$input"; echo invert) 								| ./giftsh > invert.$enc
(echo r "$input"; echo max 9) 								| ./giftsh > max.$enc
(echo r "$input"; echo mean 9) 								| ./giftsh > mean.$enc
(echo r "$input"; echo median 9) 							| ./giftsh > median.$enc
(echo r "$input"; echo min 9) 								| ./giftsh > min.$enc
(echo r "$input"; echo opacity 50) 						| ./giftsh > opacity.$enc
(echo r "$input"; echo pixelate 50) 					| ./giftsh > pixelate.$enc
(echo r "$input"; echo resize 200 0) 					| ./giftsh > resize.$enc
(echo r "$input"; echo resizefit 100 100) 		| ./giftsh > resizefit.$enc
(echo r "$input"; echo resizefill 100 100) 		| ./giftsh > resizefill.$enc
(echo r "$input"; echo rotate 60) 						| ./giftsh > rotate60.$enc
(echo r "$input"; echo rotate 90) 						| ./giftsh > rotate90.$enc
(echo r "$input"; echo rotate 180) 						| ./giftsh > rotate180.$enc
(echo r "$input"; echo rotate 270) 						| ./giftsh > rotate270.$enc
(echo r "$input"; echo saturation 50) 				| ./giftsh > sat50.$enc
(echo r "$input"; echo saturation -50) 				| ./giftsh > sat-50.$enc
(echo r "$input"; echo sepia 100) 						| ./giftsh > sepia.$enc
(echo r "$input"; echo sigmoid 0.5 5.0) 			| ./giftsh > sigmoid.$enc
(echo r "$input"; echo sobel) 								| ./giftsh > sobel.$enc
(echo r "$input"; echo transpose) 						| ./giftsh > transpose.$enc
(echo r "$input"; echo transverse) 						| ./giftsh > transverse.$enc
(echo r "$input"; echo unsharp 1.0 1.5 0.0) 	| ./giftsh > unsharp.$enc
(echo r "$input"; echo threshold 50) 					| ./giftsh > threshold.$enc
