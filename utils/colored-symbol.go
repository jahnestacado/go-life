package utils

import "strconv"

const symbol = "*"

var (
	Symbol coloredSymbol
	codes  map[string]int
)

type coloredSymbol struct {
	Red   string
	Green string
	Cyan  string
}

func init() {
	codes = map[string]int{
		"red":   31,
		"green": 32,
		"cyan":  36,
	}

	Symbol = coloredSymbol{
		Red:   getColoredSymbol("red"),
		Green: getColoredSymbol("green"),
		Cyan:  getColoredSymbol("cyan"),
	}
}

func getColoredSymbol(color string) string {
	return "\033[" + strconv.Itoa(codes[color]) + "m" + symbol + "\033[0m"
}
