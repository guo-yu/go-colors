package colors

const VERSION = "0.1.0"

var ColorsMap = map[string][]string{
	// Styles
	"bold":          []string{"\x1B[1m", "\x1B[22m"},
	"italic":        []string{"\x1B[3m", "\x1B[23m"},
	"underline":     []string{"\x1B[4m", "\x1B[24m"},
	"inverse":       []string{"\x1B[7m", "\x1B[27m"},
	"strikethrough": []string{"\x1B[9m", "\x1B[29m"},
	// Grayscale
	"white": []string{"\x1B[37m", "\x1B[39m"},
	"grey":  []string{"\x1B[90m", "\x1B[39m"},
	"black": []string{"\x1B[30m", "\x1B[39m"},
	// Colors
	"blue":    []string{"\x1B[34m", "\x1B[39m"},
	"cyan":    []string{"\x1B[36m", "\x1B[39m"},
	"green":   []string{"\x1B[32m", "\x1B[39m"},
	"magenta": []string{"\x1B[35m", "\x1B[39m"},
	"red":     []string{"\x1B[31m", "\x1B[39m"},
	"yellow":  []string{"\x1B[33m", "\x1B[39m"},
	// Background colors
	"whiteBG":   []string{"\x1B[47m", "\x1B[49m"},
	"greyBG":    []string{"\x1B[49;5;8m", "\x1B[49;25m"},
	"blackBG":   []string{"\x1B[40m", "\x1B[49m"},
	"blueBG":    []string{"\x1B[44m", "\x1B[49m"},
	"cyanBG":    []string{"\x1B[46m", "\x1B[49m"},
	"greenBG":   []string{"\x1B[42m", "\x1B[49m"},
	"magentaBG": []string{"\x1B[45m", "\x1B[49m"},
	"redBG":     []string{"\x1B[41m", "\x1B[49m"},
	"yellowBG":  []string{"\x1B[43m", "\x1B[49m"},
}

func Color(color, text string) (str string) {
	colorString, ok := ColorsMap[color]
	if !ok {
		str = text
	} else {
		str = colorString[0] + text + colorString[1]
	}
	return
}
