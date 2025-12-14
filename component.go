// Provides Components,
package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	minUnitCost     = 0.0
	minLeadTimeDays = 0
)

func validateID(id string) error {
	id = strings.TrimSpace(id)
	if id == "" {
		return fmt.Errorf("id cannot be empty")
	}
	if strings.ContainsAny(id, " \t\n") {
		return fmt.Errorf("id cannot contain whitespace")
	}
	return nil
}

type Component struct {
	ID            string
	Name          string
	Description   string
	UnitOfMeasure string
	UnitCost      float64
	LeadTimeDays  int
}

func NewComponent(
	id string,
	name string,
	description string,
	unitOfMeasure string,
	unitCost float64,
	leadTimeDays int,
) (*Component, error) {
	if err := validateID(id); err != nil {
		return nil, err
	}
	if name == "" {
		return nil, fmt.Errorf("name can't be empty")
	}
	if unitCost < minUnitCost {
		return nil, fmt.Errorf("unit cost must be at least %.2f", minUnitCost)
	}
	if leadTimeDays < minLeadTimeDays {
		return nil, fmt.Errorf("lead time days must be at least %d", minLeadTimeDays)
	}
	return &Component{
		ID:            id,
		Name:          name,
		Description:   description,
		UnitOfMeasure: unitOfMeasure,
		UnitCost:      unitCost,
		LeadTimeDays:  leadTimeDays,
	}, nil
}

func (c *Component) String() string {
	unitCostStr := fmt.Sprintf("$%.2f per %s", c.UnitCost, c.UnitOfMeasure)
	leadTimeStr := fmt.Sprintf("%d days", c.LeadTimeDays)

	return fmt.Sprintf(
		"┌─ Component ──────────────────────────┐\n"+
			"│ ID:          %-23s │\n"+
			"│ Name:        %-23s │\n"+
			"│ Description: %-23s │\n"+
			"│ Unit Cost:   %-23s │\n"+
			"│ Lead Time:   %-23s │\n"+
			"└──────────────────────────────────────┘",
		c.ID,
		c.Name,
		c.Description,
		unitCostStr,
		leadTimeStr,
	)
}

type Components map[string]*Component

type ComponentRegistry struct {
	Components Components
}

func NewComponentRegistry() *ComponentRegistry {
	return &ComponentRegistry{
		Components: make(map[string]*Component),
	}
}

func (cr *ComponentRegistry) AddComponent(c *Component) error {
	_, ok := cr.Components[c.ID]
	if ok {
		return fmt.Errorf("component %v already exists", c.ID)
	}
	cr.Components[c.ID] = c
	return nil
}

func (cr *ComponentRegistry) GetComponent(id string) (*Component, bool) {
	c, ok := cr.Components[id]
	if !ok {
		return nil, false
	}
	return c, true
}

func (cr *ComponentRegistry) ListAll() []*Component {
	c := make([]*Component, 0, len(cr.Components))
	for _, v := range cr.Components {
		c = append(c, v)
	}
	return c
}

func (cr *ComponentRegistry) DeleteComponent(id string) {
	delete(cr.Components, id)
}

func getComponentData() (*Component, error) {
	id, err := getInput("id")
	if err != nil {
		return nil, fmt.Errorf("getting id: %v", err)
	}

	name, err := getInput("name")
	if err != nil {
		return nil, fmt.Errorf("getting name: %v", err)
	}

	description, err := getInput("descreption")
	if err != nil {
		return nil, fmt.Errorf("getting descreption: %v", err)
	}

	unitOfMeasure, err := getInput("unit of measure")
	if err != nil {
		return nil, fmt.Errorf("getting unitOfMeasure input: %v", err)
	}

	uc, err := getInput("unit cost")
	if err != nil {
		return nil, fmt.Errorf("getting unit cost: %v", err)
	}
	unitCost, err := strconv.ParseFloat(uc, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid unit cost '%s': must be a number", uc)
	}

	lTD, err := getInput("lead time days")
	if err != nil {
		return nil, fmt.Errorf("getting lead time days: %v", err)
	}
	leadTimeDays, err := strconv.ParseInt(lTD, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid lead time days '%s': must be a number", lTD)
	}

	return NewComponent(id, name, description, unitOfMeasure, unitCost, int(leadTimeDays))
}

func componentMenu(cr *ComponentRegistry) {
	const (
		AddComponent      = "1"
		ViewComponent     = "2"
		ListAllComponents = "3"
		DeleteComponent   = "4"
		Exit              = "5"
	)
	menu := []string{"Add Component", "View Component", "List All Components", "Delete Component", "Exit"}
start:
	for {
		fmt.Println("===Components Menu!===")
		for i, v := range menu {
			fmt.Printf("%v. %v\n", i+1, v)
		}

		input, err := getInput("option")
		if err != nil {
			fmt.Println(fmt.Errorf("error getting option: %v", err))
			continue start
		}

		switch input {
		case AddComponent:
			fmt.Println("===Adding Component===")
			c, err := getComponentData()
			if err != nil {
				fmt.Printf("error creating component: %v\n", err)
				continue start
			}
			err = cr.AddComponent(c)
			if err != nil {
				fmt.Printf("error creating component: %v\n", err)
				continue start
			}
			fmt.Println("=======Success========")
			continue start

		case ViewComponent:
			id, err := getInput("id")
			if err != nil {
				fmt.Printf("error getting component: %v\n", err)
				continue start
			}
			c, ok := cr.GetComponent(id)
			if !ok {
				fmt.Println("no component with that id")
				continue start
			}
			fmt.Println(c.String())
			continue start

		case ListAllComponents:
			components := cr.ListAll()
			if len(components) == 0 {
				fmt.Println("Registry is empty.")
				continue start
			}
			for _, v := range components {
				fmt.Println(v.String())
			}
			continue start

		case DeleteComponent:
			id, err := getInput("id")
			if err != nil {
				fmt.Printf("error deleting component: %v\n", err)
				continue start
			}

			// Get component first
			c, ok := cr.GetComponent(id)
			if !ok {
				fmt.Println("no component found with that id")
				continue start
			}

			// Show what will be deleted
			fmt.Println("\nComponent to delete:")
			fmt.Println(c.String())

			// Ask confirmation
			confirm, err := getInput("type 'yes' to confirm deletion")
			if err != nil || confirm != "yes" {
				fmt.Println("deletion cancelled")
				continue start
			}

			cr.DeleteComponent(id)
			fmt.Println("component deleted successfully")
			continue start

		case Exit:
			fmt.Println("Closing Component Menu")
			break start

		default:
			fmt.Println("invalid options")
			continue start
		}
	}
	fmt.Println("======================")
}
