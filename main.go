package main

import (
	"fmt"
)

func main() {
	compReg := NewComponentRegistry()
	bomReg := NewBOMRegistry()
run:
	for {
		menu := []string{"BOM Menu", "Components Menu", "Exit"}
		for i, v := range menu {
			fmt.Printf("%v. %v\n", i+1, v)
		}
		input, err := getInput("option")
		if err != nil {
			fmt.Println(fmt.Errorf("error getting option: %v", err))
			continue run
		}
		switch input {
		case "1":
			bomMenu(bomReg, compReg)
			continue run
		case "2":
			componentMenu(compReg)
			continue run
		case "3":
			fmt.Println("Exiting Program")
			break run
		default:
			fmt.Println("Invalid Options")
			continue run
		}
	}
}
