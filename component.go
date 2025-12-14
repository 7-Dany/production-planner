package main

import (
	"fmt"
	"strconv"
)

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
	if id == "" {
		return nil, fmt.Errorf("id can't be empty")
	}
	if name == "" {
		return nil, fmt.Errorf("name can't be empty")
	}
	if unitCost < 0 {
		return nil, fmt.Errorf("unit cost must be greater than 0.")
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
	return fmt.Sprintf("===Component===\nID: %v\nName: %v\nDescription: %v\nUnit of Measure: %v\nUnit Cost: %v\nLead Time Days: %v\n================\n", c.ID, c.Name, c.Description, c.UnitOfMeasure, c.UnitCost, c.LeadTimeDays)
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
		return nil, fmt.Errorf("error getting id: %v", err)
	}
	name, err := getInput("name")
	if err != nil {
		return nil, fmt.Errorf("error getting name: %v", err)
	}
	description, err := getInput("descreption")
	if err != nil {
		return nil, fmt.Errorf("error getting descreption: %v", err)
	}
	unitOfMeasure, err := getInput("unit of measure")
	if err != nil {
		return nil, fmt.Errorf("error getting unitOfMeasure input: %v", err)
	}
	uc, err := getInput("unit cost")
	if err != nil {
		return nil, fmt.Errorf("error getting unitCost input: %v", err)
	}
	unitCost, err := strconv.ParseFloat(uc, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing unitCost to float64: %v", err)
	}
	lTD, err := getInput("lead time days")
	if err != nil {
		return nil, fmt.Errorf("error getting leadTimeDays input: %v", err)
	}
	leadTimeDays, err := strconv.ParseInt(lTD, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing leadTimeDays to int: %v", err)
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
			cr.DeleteComponent(id)
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
