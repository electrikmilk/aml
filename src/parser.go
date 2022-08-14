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
				} else {
					interpreterError(fmt.Sprintf("Invalid character: %s", currentChar), l, c)
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
