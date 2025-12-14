// Package main provides utility functions for the production planner.
// This file contains input/output helpers.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(n string) (string, error) {
	fmt.Printf("Please enter %v: \n", n)

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(line), nil
}
