/*
 * Copyright (c) 2023 Brandon Jordan
 */

package main

type tokenType string

const (
	INTEGER      tokenType = "0123456789.-^"
	PLUS         tokenType = "+"
	MINUS        tokenType = "-"
	MULTIPLY     tokenType = "*×"
	DIVIDE       tokenType = "/"
	MODULUS      tokenType = "%"
	STARTCLOSURE tokenType = "("
	ENDCLOSURE   tokenType = ")"
)

type token struct {
	col    int
	typeof tokenType
	value  float64
}

type tokenSet struct {
	line   int
	tokens []token
}

var tokenSets []tokenSet
