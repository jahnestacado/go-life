package utils

import (
	"strconv"
)

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
		Red:   GetColoredSymbol("red"),
		Green: GetColoredSymbol("green"),
		Cyan:  GetColoredSymbol("cyan"),
	}
}

func GetColoredSymbol(color string) string {
	return "\033[" + strconv.Itoa(codes[color]) + "m" + symbol + "\033[0m"
}
