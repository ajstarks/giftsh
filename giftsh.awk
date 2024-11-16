# convert a giftsh script to a gift command line
BEGIN {
	cmdline=""
}

$1 == "r" || $1 == "read" {input=$2}

$1 == "blur"		{cmdline=sprintf("%s -%s=%s",			cmdline, $1, $2)} 
$1 == "brightness"	{cmdline=sprintf("%s -%s=%s",			cmdline, $1, $2)} 
$1 == "colorbalance"{cmdline=sprintf("%s -%s=%s,%s,%s",		cmdline, $1, $2, $3, $4)} 
$1 == "colorize"	{cmdline=sprintf("%s -%s=%s,%s,%s",		cmdline, $1, $2, $3, $4)}
$1 == "colorspace"	{cmdline=sprintf("%s -%s=%s",			cmdline, $1, $2)} 
$1 == "contrast"	{cmdline=sprintf("%s -%s=%s",			cmdline, $1, $2)} 
$1 == "crop"		{cmdline=sprintf("%s -%s=%s,%s,%s,%s",	cmdline, $1, $2, $3, $4, $5)} 
$1 == "cropsize"	{cmdline=sprintf("%s -%s=%s,%s",		cmdline, $1, $2, $3)} 
$1 == "edge"		{cmdline=sprintf("%s -%s",				cmdline, $1)} 
$1 == "emboss"		{cmdline=sprintf("%s -%s",				cmdline, $1)} 
$1 == "fliph"		{cmdline=sprintf("%s -%s",				cmdline, $1)} 
$1 == "flipv"		{cmdline=sprintf("%s -%s",				cmdline, $1)} 
$1 == "gamma"		{cmdline=sprintf("%s -%s=%s",			cmdline, $1, $2)} 
$1 == "gray"		{cmdline=sprintf("%s -%s",				cmdline, $1)} 
$1 == "hue"			{cmdline=sprintf("%s -%s=%s",			cmdline, $1, $2)} 
$1 == "invert"		{cmdline=sprintf("%s -%s",				cmdline, $1)} 
$1 == "max"			{cmdline=sprintf("%s -%s=%s",			cmdline, $1, $2)} 
$1 == "mean"		{cmdline=sprintf("%s -%s=%s",			cmdline, $1, $2)} 
$1 == "median"		{cmdline=sprintf("%s -%s=%s",			cmdline, $1, $2)} 
$1 == "min"			{cmdline=sprintf("%s -%s=%s",			cmdline, $1, $2)} 
$1 == "opacity"		{cmdline=sprintf("%s -%s=%s",			cmdline, $1, $2)} 
$1 == "pixelate"	{cmdline=sprintf("%s -%s=%s",			cmdline, $1, $2)}
$1 == "resize"		{cmdline=sprintf("%s -%s=%s,%s",		cmdline, $1, $2, $3)} 
$1 == "resizefill"	{cmdline=sprintf("%s -%s=%s,%s",		cmdline, $1, $2, $3)} 
$1 == "resizefit"	{cmdline=sprintf("%s -%s=%s,%s",		cmdline, $1, $2, $3)} 
$1 == "rotate"		{cmdline=sprintf("%s -%s=%s",			cmdline, $1, $2)} 
$1 == "saturation"	{cmdline=sprintf("%s -%s=%s",			cmdline, $1, $2)} 
$1 == "sepia"		{cmdline=sprintf("%s -%s=%s",			cmdline, $1, $2)} 
$1 == "sigmoid"		{cmdline=sprintf("%s -%s=%s",			cmdline, $1, $2)} 
$1 == "sobel"		{cmdline=sprintf("%s -%s",				cmdline, $1)} 
$1 == "threshold"	{cmdline=sprintf("%s -%s=%s",			cmdline, $1, $2)} 
$1 == "transpose"	{cmdline=sprintf("%s -%s",				cmdline, $1)} 
$1 == "transverse"	{cmdline=sprintf("%s -%s",				cmdline, $1)} 
$1 == "unsharp"		{cmdline=sprintf("%s -%s=%s,%s,%s",		cmdline, $1, $2, $3, $4)} 

END {
	printf "gift%s %s\n", cmdline, input
}