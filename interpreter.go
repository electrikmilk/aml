/*
 * Copyright (c) 2023 Brandon Jordan
 */

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/electrikmilk/ttuy"
)

var result float64
var operation tokenType

var saveOperation tokenType
var saveResult float64

var currentSet tokenSet
var currentToken token

func interpret() {
	check()
	operation = ""
	saveOperation = ""
	result = 0
	saveResult = 0
	for _, set := range tokenSets {
		currentSet = set
		for _, tok := range set.tokens {
			currentToken = tok
			switch tok.typeof {
			case INTEGER:
				if operation != "" {
					doOperation(operation, tok.value)
					operation = ""
				} else {
					result = tok.value
				}
			case PLUS:
				operation = PLUS
			case MINUS:
				operation = MINUS
			case MULTIPLY:
				operation = MULTIPLY
			case DIVIDE:
				operation = DIVIDE
			case MODULUS:
				operation = MODULUS
			case STARTCLOSURE:
				saveResult = result
				saveOperation = operation
				result = 0
			case ENDCLOSURE:
				doOperation(saveOperation, saveResult)
			default:
				operation = ""
				result = 0
			}
		}
		stripComments()
		replaceVariables()
		var equals string = ttuy.Style("= ", ttuy.Dim)
		if len(os.Args) > 1 {
			fmt.Printf(ttuy.Style(printModulus(lines[set.line]), ttuy.YellowText) + " ")
		}
		fmt.Print(equals)
		var answer string = strconv.FormatFloat(result, 'f', -1, 64)
		fmt.Print(ttuy.Style(answer, ttuy.GreenText) + "\n")
	}
}

// Percent replace for modulus operator to make printf happy
func printModulus(str string) string {
	if strings.Contains(str, "%") {
		return strings.ReplaceAll(str, "%", "%%")
	}
	return str
}

func stripComments() {
	for i, line := range lines {
		if len(line) > 0 {
			if strings.Contains(line, "#") {
				lines[i] = strings.TrimSpace(strings.Split(lines[i], "#")[0])
			}
		}
	}
}

func replaceVariables() {
	for i, line := range lines {
		if len(line) > 0 {
			for _, char := range strings.Split(line, "") {
				if tok, found := variables[char]; found {
					lines[i] = strings.ReplaceAll(lines[i], char, strconv.FormatFloat(tok.value, 'f', -1, 64))
				} else if tok, found := constants[char]; found {
					lines[i] = strings.ReplaceAll(lines[i], char, strconv.FormatFloat(tok.value, 'f', -1, 64))
				}
			}
		}
	}
}

func doOperation(operation tokenType, value float64) {
	switch operation {
	case PLUS:
		result += value
	case MINUS:
		result -= value
	case MULTIPLY:
		result *= value
	case DIVIDE:
		if result == 0 && value == 0 {
			parsingError("Divide by zero error", currentSet.line, currentToken.col)
		}
		result /= value
	case MODULUS:
		result = math.Mod(result, value)
	}
}
