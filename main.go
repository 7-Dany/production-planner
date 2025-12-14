// Package main provides Bill of Materials (BOM) management functionality
// for production planning and cost calculation.
package main

import (
	"fmt"
)

func displayWelcome(compReg *ComponentRegistry, bomReg *BOMRegistry) {
	fmt.Println("\n╔════════════════════════════════════════════╗")
	fmt.Println("║  Production Planning Dashboard v1.0        ║")
	fmt.Println("╚════════════════════════════════════════════╝")
	fmt.Printf("\nSystem Status:\n")
	fmt.Printf("  • Components: %d registered\n", len(compReg.Components))
	fmt.Printf("  • BOMs: %d active\n\n", len(bomReg.BOMS))
}

func main() {
	compReg := NewComponentRegistry()
	bomReg := NewBOMRegistry()
	displayWelcome(compReg, bomReg)
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
