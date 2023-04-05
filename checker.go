/*
 * Copyright (c) 2023 Brandon Jordan
 */

package main

func check() {
	for _, set := range tokenSets {
		var lastToken token
		for i, tok := range set.tokens {
			if tok != lastToken {
				if lastToken.typeof == INTEGER && tok.typeof == INTEGER {
					parsingError("Two integers with no separating operator", set.line, tok.col)
				}
				if lastToken.typeof == INTEGER && tok.typeof == STARTCLOSURE {
					parsingError("Integer and starting closure with no separating operator", set.line, tok.col)
				}
				if lastToken.typeof == ENDCLOSURE && tok.typeof == INTEGER {
					parsingError("Integer and ending closure with no separating operator", set.line, tok.col)
				}
				if lastToken.typeof == ENDCLOSURE && tok.typeof == STARTCLOSURE {
					parsingError("No separating operator for closure", set.line, tok.col)
				}
				if lastToken.typeof == STARTCLOSURE && tok.typeof == ENDCLOSURE {
					parsingError("Empty closure, must contain expression", set.line, tok.col)
				}
				if tok.typeof == STARTCLOSURE && nextToken(set, i).typeof != INTEGER {
					parsingError("Invalid closure", set.line, tok.col)
				}
				if isOperator(tok) && (lastToken.typeof != INTEGER && lastToken.typeof != ENDCLOSURE) {
					if nextToken(set, i).typeof != INTEGER && tok.typeof != MINUS {
						parsingError("Operator does not follow an integer or ending closure", set.line, tok.col)
					}
				}
				if isOperator(lastToken) && isOperator(tok) {
					parsingError("Two operators with no separating integer or closure", set.line, tok.col)
				}
				if isOperator(tok) && (nextToken(set, i).typeof != INTEGER && nextToken(set, i).typeof != STARTCLOSURE) {
					parsingError("Expecting integer or starting closure after operator", set.line, tok.col)
				}
			}
			lastToken = tok
		}
	}
}

func nextToken(set tokenSet, idx int) token {
	idx++
	if len(set.tokens) > idx {
		return set.tokens[idx]
	} else {
		return token{}
	}
}

func isOperator(token token) (isOperator bool) {
	isOperator = false
	switch token.typeof {
	case PLUS:
		isOperator = true
	case MINUS:
		isOperator = true
	case MULTIPLY:
		isOperator = true
	case DIVIDE:
		isOperator = true
	case MODULUS:
		isOperator = true
	}
	return
}
