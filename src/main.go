/*
 * Copyright (c) 2022 Brandon Jordan
 */

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"
)

var EOL = "\n"

func main() {
	if runtime.GOOS == "windows" {
		EOL = "\r\n"
	}
	if len(os.Args) > 1 {
		// file
		if _, err := os.Stat(os.Args[1]); errors.Is(err, os.ErrNotExist) {
			errorMsg(fmt.Sprintf("File %s does not exist!", os.Args[1]))
		}
		var fileContents, err = os.ReadFile(os.Args[1])
		if strings.Split(os.Args[1], ".")[1] != "aml" {
			errorMsg("Not an AML file!")
		}
		handle(err)
		contents = string(fileContents)
		parse()
		interpret()
	} else {
		// interactive shell
		fmt.Println(style("Enter \"q\" to quit.", CYAN) + "\n")
		for {
			tokenSets = []tokenSet{}
			fmt.Printf(style("> ", GREEN))
			var reader = bufio.NewReader(os.Stdin)
			var input, _, err = reader.ReadLine()
			handle(err)
			contents = string(input)
			if contents == "q" {
				break
			}
			contents += "\n"
			parse()
			interpret()
		}
	}
}

func errorMsg(message string) {
	fmt.Println(style("Error: "+message, RED))
	os.Exit(1)
}

func interpreterError(message string, line int, col int) {
	var dashes = "-------"
	fmt.Println("\n" + style("Error: "+message, RED) + "\n")
	if len(os.Args) > 1 {
		fmt.Printf(style(dashes, DIM)+" %s:%d:%d\n", os.Args[1], line+1, col+1)
	} else {
		fmt.Println(style(dashes, DIM))
	}
	// line before
	if line-1 > 0 {
		fmt.Printf(style(fmt.Sprintf("%d |%s\n", line, lines[line-1]), DIM))
	}
	// error line
	fmt.Printf("%d ", line+1)
	fmt.Printf(style("|", DIM))
	for i, char := range strings.Split(lines[line], "") {
		if i == col {
			fmt.Printf("%s", style(char, RED, BOLD, UNDERLINE))
		} else {
			fmt.Printf(char)
		}
	}
	fmt.Printf("\n")
	// highlight error
	fmt.Printf("  ")
	for i := 0; i <= col; i++ {
		fmt.Printf(" ")
	}
	fmt.Println(style("^", RED, BOLD))
	// line after
	if len(lines) > (line + 1) {
		fmt.Printf(style(fmt.Sprintf("%d |%s\n", line+2, lines[line+1]), DIM))
	}
	fmt.Println(style(dashes, DIM))
	os.Exit(1)
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}
