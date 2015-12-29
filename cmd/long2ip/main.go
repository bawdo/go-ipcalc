package main

import (
	"fmt"
	"github.com/bawdo/go-ipcalc"
	"log"
	"os"
	"strconv"
)

func main() {
	long_input, _ := strconv.Atoi(os.Args[1])
	long, err := ipcalc.Long2Ip(long_input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", long)
}
