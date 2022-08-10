/*
 * Copyright (c) 2022 Brandon Jordan
 */

package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if _, err := os.Stat("src/aml"); errors.Is(err, os.ErrNotExist) {
		panic("AML binary has not been built in src!\n\nRun:\ncd src\ngo build")
	}
	var tests, dirErr = os.ReadDir("tests")
	if dirErr != nil {
		panic(dirErr)
	}
	for _, test := range tests {
		var testCmd = exec.Command("./src/aml", "tests/"+test.Name())
		var output, testErr = testCmd.Output()
		fmt.Println(testCmd.Args[1] + ":")
		if testErr != nil {
			var stderr bytes.Buffer
			testCmd.Stderr = &stderr
			if testErr != nil {
				fmt.Println(stderr.String())
			}
		}
		fmt.Println(string(output))
	}
}
