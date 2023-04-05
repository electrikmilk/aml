/*
 * Copyright (c) 2023 Brandon Jordan
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
var currentChar rune

func parse() {
	lines = strings.Split(contents, eol)
	for l, line := range lines {
		if line != "" {
			var lineTokens []token
			c = -1
			chars = strings.Split(line, "")
			advance()
			for currentChar != -1 {
				if currentChar == ' ' || currentChar == '\t' || currentChar == eolRune {
					advance()
				} else if currentChar == '#' {
					waitForComment()
				} else if isToken(PLUS) {
					lineTokens = append(lineTokens, token{typeof: PLUS, col: c})
					advance()
				} else if isToken(MINUS) && next(1) == ' ' {
					lineTokens = append(lineTokens, token{typeof: MINUS, col: c})
					advance()
				} else if strings.ContainsAny(string(currentChar), string(MULTIPLY)) {
					lineTokens = append(lineTokens, token{typeof: MULTIPLY, col: c})
					advance()
				} else if isToken(DIVIDE) {
					lineTokens = append(lineTokens, token{typeof: DIVIDE, col: c})
					advance()
				} else if isToken(MODULUS) {
					lineTokens = append(lineTokens, token{typeof: MODULUS, col: c})
					advance()
				} else if isToken(STARTCLOSURE) {
					lineTokens = append(lineTokens, token{typeof: STARTCLOSURE, col: c})
					advance()
				} else if isToken(ENDCLOSURE) {
					lineTokens = append(lineTokens, token{typeof: ENDCLOSURE, col: c})
					advance()
				} else if strings.Contains(string(INTEGER), string(currentChar)) {
					lineTokens = append(lineTokens, tokenizeInteger())
				} else if currentChar == 'v' && next(1) == 'a' && next(2) == 'r' {
					c += 3
					advance()
					collectVariable("var", l, c)
				} else if currentChar == 'c' && next(1) == 'o' {
					c += 5
					advance()
					collectVariable("const", l, c)
				} else {
					if constTok, found := constants[string(currentChar)]; found {
						lineTokens = append(lineTokens, constTok)
						advance()
					} else if varTok, found := variables[string(currentChar)]; found {
						lineTokens = append(lineTokens, varTok)
						advance()
					} else {
						parsingError(fmt.Sprintf("Invalid character: %s", string(currentChar)), l, c)
					}
				}
			}
			if len(lineTokens) > 0 {
				tokenSets = append(tokenSets, tokenSet{line: l, tokens: lineTokens})
			}
		}
	}
}

func isToken(token tokenType) bool {
	return string(currentChar) == string(token)
}

func waitForComment() {
	var comment string
	for currentChar != -1 {
		if currentChar == eolRune {
			break
		}
		comment += string(currentChar)
		advance()
	}
	return
}

func collectVariable(kind string, line int, col int) {
	var identifier string
	for currentChar != ' ' {
		identifier += string(currentChar)
		advance()
	}
	if kind == "const" {
		if _, found := variables[identifier]; found {
			parsingError(fmt.Sprintf("Variable \"%s\" already exists!", identifier), line, col)
		}
	}
	if _, found := constants[identifier]; found {
		parsingError(fmt.Sprintf("Constant \"%s\" already exists!", identifier), line, col)
	}
	advance()
	if currentChar != '=' {
		parsingError("Missing equality operator", line, col)
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
	for currentChar != -1 && strings.Contains(string(INTEGER), string(currentChar)) {
		if currentChar == '^' {
			saveValue = value
			value = ""
		} else {
			value += string(currentChar)
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
		currentChar = []rune(chars[c])[0]
	} else {
		currentChar = -1
	}
}

func next(mov int) (nextChar rune) {
	if len(chars) > (c + mov) {
		nextChar = []rune(chars[c+mov])[0]
	} else {
		nextChar = -1
	}
	return
}

func prev(mov int) (prevChar rune) {
	if len(chars) < (c - mov) {
		prevChar = []rune(chars[c-mov])[0]
	} else {
		prevChar = -1
	}
	return
}

func nextSkip(mov int) rune {
	return seek(&mov, false)
}

func prevSkip(mov int) rune {
	return seek(&mov, true)
}

func seek(mov *int, reverse bool) (seekedChar rune) {
	var charPos int
	seekedChar = ' '
	for seekedChar != ' ' {
		if reverse == true {
			charPos = c - *mov
		} else {
			charPos = c + *mov
		}
		seekedChar = []rune(chars[charPos])[0]
	}
	return
}

func printCurrentChar() {
	var char string
	switch currentChar {
	case '\t':
		char = "TAB"
		break
	case ' ':
		char = "SPACE"
		break
	case eolRune:
		char = "eol"
		break
	default:
		char = string(currentChar)
	}
	fmt.Println(char)
}
