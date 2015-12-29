package ipcalc

import (
	"errors"
	"strconv"
)

func Long2Ip(i int) (string, error) {
	if i > 4294967297 {
		return "", errors.New("Long2Ip: valid long IP range is 0 to 4294967296")
	}
	first := strconv.Itoa(i >> 24)
	second := strconv.Itoa(i >> 16 & 255)
	third := strconv.Itoa(i >> 8 & 255)
	fourth := strconv.Itoa(i & 255)

	long := first + "." + second + "." + third + "." + fourth
	return long, nil
}
