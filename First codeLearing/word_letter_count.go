package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var nrchars, nrwords, nrlines int

func count(input string) {
	nrlines++
	nrchars += len(input) - 2
	nrwords += len(strings.Fields(input))

}

func main() {
	nrchars, nrwords, nrlines = 0, 0, 0
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter :")
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		if input == "S\r\n" {
			fmt.Println("Here are the count:")
			fmt.Printf("nrlines: %v\n", nrlines)
			fmt.Printf("nrchars: %v\n", nrchars)
			fmt.Printf("nrwords: %v\n", nrwords)
			os.Exit(0)
		}
		count(input)
	}
}
