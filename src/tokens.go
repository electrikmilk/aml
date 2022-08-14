/*
 * Copyright (c) 2022 Brandon Jordan
 */

package main

type tokenType string

const (
	INTEGER      tokenType = "0123456789.-^"
	PLUS                   = "+"
	MINUS                  = "-"
	MULTIPLY               = "*xX×"
	DIVIDE                 = "/"
	MODULUS                = "%"
	STARTCLOSURE           = "("
	ENDCLOSURE             = ")"
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
