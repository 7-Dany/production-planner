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
