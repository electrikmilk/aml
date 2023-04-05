/*
 * Copyright (c) 2023 Brandon Jordan
 */

package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var currentTest string

func TestAML(t *testing.T) {
	var files, err = os.ReadDir("examples")
	if err != nil {
		t.Fail()
	}
	for _, file := range files {
		if strings.Contains(file.Name(), ".aml") {
			currentTest = "examples/" + file.Name()
			fmt.Printf("Testing %s...\n", currentTest)
			os.Args[1] = currentTest
			reset()
			main()
			fmt.Print("\033[32mPASSED\033[0m\n\n")
		}
	}
}

func reset() {
	contents = ""
	lines = []string{}
	chars = []string{}
	c = 0
	currentChar = 0
	tokenSets = []tokenSet{}
}
