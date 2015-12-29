package ipcalc

import (
	"strconv"
)

func Long2Ip(i int) string {
	first := strconv.Itoa(i >> 24)
	second := strconv.Itoa(i >> 16 & 255)
	third := strconv.Itoa(i >> 8 & 255)
	fourth := strconv.Itoa(i & 255)

	return first + "." + second + "." + third + "." + fourth
}
