/*
 * Copyright (c) 2023 Brandon Jordan
 */

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/electrikmilk/ttuy"
)

var variables map[string]token
var constants map[string]token

func main() {
	variables = make(map[string]token)
	constants = make(map[string]token)
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
		fmt.Println(ttuy.Style("Enter \"q\" to quit.", ttuy.CyanText) + "\n")
		for {
			tokenSets = []tokenSet{}
			fmt.Print(ttuy.Style("> ", ttuy.GreenText))
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
	fmt.Println(ttuy.Style("Error: "+message, ttuy.RedText))
	os.Exit(1)
}

func parsingError(message string, line int, col int) {
	var dashes = "-------"
	fmt.Println("\n" + ttuy.Style("Error: "+message, ttuy.RedText) + "\n")
	if len(os.Args) > 1 {
		fmt.Printf(ttuy.Style(dashes, ttuy.Dim)+" %s:%d:%d\n", os.Args[1], line+1, col+1)
	} else {
		fmt.Println(ttuy.Style(dashes, ttuy.Dim))
	}
	// line before
	if line-1 > 0 {
		fmt.Print(ttuy.Style(fmt.Sprintf("%d |%s\n", line, lines[line-1]), ttuy.Dim))
	}
	// error line
	fmt.Printf("%d ", line+1)
	fmt.Print(ttuy.Style("|", ttuy.Dim))
	for i, char := range strings.Split(lines[line], "") {
		if i == col {
			fmt.Printf("%s", ttuy.Style(char, ttuy.RedText, ttuy.Bold, ttuy.Underlined))
		} else {
			fmt.Print(char)
		}
	}
	fmt.Print("\n")
	// highlight error
	fmt.Print("  ")
	for i := 0; i <= col; i++ {
		fmt.Print(" ")
	}
	fmt.Println(ttuy.Style("^", ttuy.RedText, ttuy.Bold))
	// line after
	if len(lines) > (line + 1) {
		fmt.Print(ttuy.Style(fmt.Sprintf("%d |%s\n", line+2, lines[line+1]), ttuy.Dim))
	}
	fmt.Println(ttuy.Style(dashes, ttuy.Dim))
	os.Exit(1)
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}
