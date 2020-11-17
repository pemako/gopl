package main

import (
	"fmt"
	"github.com/pemako/gopl/codes/ch1"
	"os"
	"strconv"

	"github.com/pemako/gopl/codes/ch2/tempconv"
)

func testCh2Tempconv() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}

func testFetchAll(){
	ch1.Fetchall()
}