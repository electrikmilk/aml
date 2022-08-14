/*
 * Copyright (c) 2022 Brandon Jordan
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

var contents string
var lines []string

var c int
var chars []string
var currentChar string

func parse() {
	lines = strings.Split(contents, EOL)
	for l, line := range lines {
		if line != "" {
			var lineTokens []token
			c = -1
			chars = strings.Split(line, "")
			advance()
			for currentChar != "" {
				if currentChar == " " || currentChar == "\t" || currentChar == EOL {
					advance()
				} else if currentChar == "#" {
					waitForComment()
				} else if currentChar == PLUS {
					lineTokens = append(lineTokens, token{typeof: PLUS, col: c})
					advance()
				} else if currentChar == MINUS && next(1) == " " {
					lineTokens = append(lineTokens, token{typeof: MINUS, col: c})
					advance()
				} else if strings.ContainsAny(currentChar, MULTIPLY) {
					lineTokens = append(lineTokens, token{typeof: MULTIPLY, col: c})
					advance()
				} else if currentChar == DIVIDE {
					lineTokens = append(lineTokens, token{typeof: DIVIDE, col: c})
					advance()
				} else if currentChar == MODULUS {
					lineTokens = append(lineTokens, token{typeof: MODULUS, col: c})
					advance()
				} else if currentChar == STARTCLOSURE {
					lineTokens = append(lineTokens, token{typeof: STARTCLOSURE, col: c})
					advance()
				} else if currentChar == ENDCLOSURE {
					lineTokens = append(lineTokens, token{typeof: ENDCLOSURE, col: c})
					advance()
				} else if strings.Contains(string(INTEGER), currentChar) {
					lineTokens = append(lineTokens, tokenizeInteger())
				} else if currentChar == "v" && next(1) == "a" && next(2) == "r" {
					c += 3
					advance()
					collectVariable("var", l, c)
				} else if currentChar == "c" && next(1) == "o" {
					c += 5
					advance()
					collectVariable("const", l, c)
				} else {
					if constTok, found := constants[currentChar]; found {
						lineTokens = append(lineTokens, constTok)
						advance()
					} else if varTok, found := variables[currentChar]; found {
						lineTokens = append(lineTokens, varTok)
						advance()
					} else {
						interpreterError(fmt.Sprintf("Invalid character: %s", currentChar), l, c)
					}
				}
			}
			if len(lineTokens) > 0 {
				tokenSets = append(tokenSets, tokenSet{line: l, tokens: lineTokens})
			}
		}
	}
}

func waitForComment() {
	var comment string
	for currentChar != "" {
		if currentChar == EOL {
			break
		}
		comment += currentChar
		advance()
	}
	return
}

func collectVariable(kind string, line int, col int) {
	var identifier string
	for currentChar != " " {
		identifier += currentChar
		advance()
	}
	if kind == "const" {
		if _, found := variables[identifier]; found {
			interpreterError(fmt.Sprintf("Variable \"%s\" already exists!", identifier), line, col)
		}
	}
	if _, found := constants[identifier]; found {
		interpreterError(fmt.Sprintf("Constant \"%s\" already exists!", identifier), line, col)
	}
	advance()
	if currentChar != "=" {
		interpreterError("Missing equality operator", line, col)
	}
	c++
	advance()
	if kind == "const" {
		constants[identifier] = tokenizeInteger()
	} else {
		variables[identifier] = tokenizeInteger()
	}
}

func tokenizeInteger() token {
	var value string
	var saveValue string
	for currentChar != "" && strings.Contains(string(INTEGER), currentChar) {
		if currentChar == "^" {
			saveValue = value
			value = ""
		} else {
			value += currentChar
		}
		advance()
	}
	float, floatErr := strconv.ParseFloat(value, 64)
	handle(floatErr)
	if len(saveValue) > 0 {
		saveFloat, saveFloatErr := strconv.ParseFloat(saveValue, 64)
		handle(saveFloatErr)
		float = saveFloat
		var exponent, exponentErr = strconv.ParseFloat(value, 64)
		handle(exponentErr)
		var initValue float64 = float
		exponent--
		var i float64
		for i = 0; i < exponent; i++ {
			float *= initValue
		}
	}
	return token{typeof: INTEGER, value: float, col: c}
}

func advance() {
	c++
	if c < len(chars) {
		currentChar = chars[c]
	} else {
		currentChar = ""
	}
}

func next(mov int) (nextChar string) {
	if len(chars) > (c + mov) {
		nextChar = chars[c+mov]
	} else {
		nextChar = ""
	}
	return
}

func prev(mov int) (prevChar string) {
	if len(chars) < (c - mov) {
		prevChar = chars[c-mov]
	} else {
		prevChar = ""
	}
	return
}

func nextSkip(mov int) string {
	return seek(&mov, false)
}

func prevSkip(mov int) string {
	return seek(&mov, true)
}

func seek(mov *int, reverse bool) (seekedChar string) {
	var charPos int
	seekedChar = " "
	for seekedChar != " " {
		if reverse == true {
			charPos = c - *mov
		} else {
			charPos = c + *mov
		}
		seekedChar = chars[charPos]
	}
	return
}

func printCurrentChar() {
	var char string
	switch currentChar {
	case "\t":
		char = "TAB"
		break
	case " ":
		char = "SPACE"
		break
	case EOL:
		char = "EOL"
		break
	default:
		char = currentChar
	}
	fmt.Println(char)
}
