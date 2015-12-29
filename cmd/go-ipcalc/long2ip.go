package main

import (
	"fmt"
	"github.com/bawdo/go-ipcalc"
	"os"
	"strconv"
)

func main() {
	long, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("%s\n", ipcalc.Long2Ip(long))
}
