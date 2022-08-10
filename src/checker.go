/*
 * Copyright (c) 2022 Brandon Jordan
 */

package main

func check() {
	for _, set := range tokenSets {
		var lastToken token
		for i, tok := range set.tokens {
			if tok != lastToken {
				if lastToken.typeof == INTEGER && tok.typeof == INTEGER {
					interpreterError("Two integers with no separating operator", set.line, tok.col)
				}
				if lastToken.typeof == INTEGER && tok.typeof == STARTCLOSURE {
					interpreterError("Integer and starting closure with no separating operator", set.line, tok.col)
				}
				if lastToken.typeof == ENDCLOSURE && tok.typeof == INTEGER {
					interpreterError("Integer and ending closure with no separating operator", set.line, tok.col)
				}
				if lastToken.typeof == ENDCLOSURE && tok.typeof == STARTCLOSURE {
					interpreterError("No separating operator for closure", set.line, tok.col)
				}
				if lastToken.typeof == STARTCLOSURE && tok.typeof == ENDCLOSURE {
					interpreterError("Empty closure, must contain expression", set.line, tok.col)
				}
				if isOperator(tok) && (lastToken.typeof != INTEGER && lastToken.typeof != ENDCLOSURE) {
					if nextToken(set, i).typeof != INTEGER && tok.typeof != MINUS {
						interpreterError("Operator does not follow an integer or ending closure", set.line, tok.col)
					}
				}
				if isOperator(lastToken) && isOperator(tok) {
					interpreterError("Two operators with no separating integer or closure", set.line, tok.col)
				}
				if isOperator(tok) && (nextToken(set, i).typeof != INTEGER && nextToken(set, i).typeof != STARTCLOSURE) {
					interpreterError("Expecting integer or starting closure after operator", set.line, tok.col)
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
