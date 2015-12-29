package main

import (
	"fmt"
	"github.com/bawdo/go-ipcalc"
	"os"
)

func main() {
	fmt.Printf("%d\n", ipcalc.Ip2Long(os.Args[1]))
}
