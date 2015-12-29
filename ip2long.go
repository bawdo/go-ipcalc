package ipcalc

import (
	"math"
	"strconv"
	"strings"
)

func Ip2Long(s string) int {
	soctets := strings.Split(s, ".")

	octets := make([]float64, 4)
	for i, v := range soctets {
		octets[i], _ = strconv.ParseFloat(v, 64)
	}

	first := int(octets[0] * math.Pow(2, 24))
	second := int(octets[1] * math.Pow(2, 16))
	third := int(octets[2] * math.Pow(2, 8))
	fourth := int(octets[3])

	return first + second + third + fourth
}
