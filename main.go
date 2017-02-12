package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	arguments := os.Args[1:]

	if len(arguments) == 2 {
		start, err := strconv.Atoi(arguments[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(arguments[1])
		if err != nil {
			panic(err)
		}
		fmt.Printf("Displaying %v to %v \n", start, end)
		printUnicodeForRange(start, end)
	} else {
		fmt.Println("No arguments or incorrect arguments provided, displaying default 23 127 arguments")
		printUnicodeForRange(23, 127)
	}
}

func printUnicodeForRange(start int, end int) {
	for i := start; i < end; i++ {
		fmt.Printf("Binary: %b, \t Decimal: %d, \t Hexidecimal: %X, \t Unicode(UTF-8): %#U \n", i, i, i, i)
	}
}