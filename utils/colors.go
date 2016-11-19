package utils

import "strconv"

func ColorString(str, color string) string {
	codes := map[string]int{
		"red":   31,
		"green": 32,
		"cyan":  36,
	}
	return "\033[" + strconv.Itoa(codes[color]) + "m" + str + "\033[0m"
}
