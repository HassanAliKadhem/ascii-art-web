package ascii

import (
	"strconv"
	"strings"
)

func GetColor(color string) string {
	if strings.HasPrefix(color, "--color=#") {
		return hexColor(color)
	} else if strings.HasPrefix(color, "--color=rgb") {
		return rgbColor(color)
	} else if strings.HasPrefix(color, "--color=hsl") {
		return hslColor(color)
	} else {
		return ansiiColor(color)
	}
}

var ansiiColors = map[string]string{
	"black":   "\x1b[30m",
	"red":     "\x1b[31m",
	"green":   "\x1b[32m",
	"yellow":  "\x1b[33m",
	"blue":    "\x1b[34m",
	"magenta": "\x1b[35m",
	"cyan":    "\x1b[36m",
	"white":   "\x1b[37m",
	"gold":    "\x1b[38;5;3m",
	"orange":  "\033[38;2;255;165;0m",
}

func ansiiColor(ansiiCode string) string {
	color, ok := ansiiColors[strings.TrimPrefix(ansiiCode, "--color=")]
	if ok {
		return color
	} else {
		return ansiiColors["white"]
	}
}

func hexColor(hexCode string) string {
	red, err := strconv.ParseInt(hexCode[9:11], 16, 64)
	if err != nil {
		return ansiiColors["white"]
	}
	green, err := strconv.ParseInt(hexCode[11:13], 16, 64)
	if err != nil {
		return ansiiColors["white"]
	}
	blue, err := strconv.ParseInt(hexCode[13:15], 16, 64)
	if err != nil {
		return ansiiColors["white"]
	}
	return "\033[38;2;" + strconv.FormatInt(red, 10) + ";" + strconv.FormatInt(green, 10) + ";" + strconv.FormatInt(blue, 10) + "m"
}

func rgbColor(rgbCode string) string {
	justValues := strings.TrimSuffix(strings.TrimPrefix(rgbCode, "--color=rgb("), ")")
	rgbColor := "\033[38;2"
	for _, value := range strings.Split(justValues, ",") {
		rgbColor += ";" + strings.TrimSpace(value)
	}
	rgbColor += "m"
	return rgbColor
}

func hslColor(hslCode string) string {
	justValues := strings.TrimSuffix(strings.TrimPrefix(hslCode, "--color=hsl("), ")")
	rgbColor := "\033[38;2"
	for _, value := range strings.Split(justValues, ",") {
		i, err := strconv.Atoi(strings.TrimSpace(strings.TrimSuffix(value, "%")))
		if err != nil {
			return ansiiColors["white"]
		}
		code := i / 100 * 255
		rgbColor += ";" + strings.Split(strconv.Itoa(code), ".")[0]
	}
	rgbColor += "m"
	return rgbColor
}
