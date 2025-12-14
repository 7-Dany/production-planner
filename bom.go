package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	minQuantity = 0.0
)

type BOMItem struct {
	Component *Component
	Quantity  float64
}

func NewBOMItem(c *Component, q float64) *BOMItem {
	return &BOMItem{c, q}
}

type BOM struct {
	ProductID   string
	ProductName string
	Items       []*BOMItem
}

func NewBOM(productID, productName string) (*BOM, error) {
	if productID == "" {
		return nil, fmt.Errorf("product id can't be empty")
	}
	if productName == "" {
		return nil, fmt.Errorf("product name can't be empty")
	}
	return &BOM{
		ProductID:   productID,
		ProductName: productName,
		Items:       []*BOMItem{},
	}, nil
}

func (bom *BOM) AddItem(componentID string, quantity float64, Registry *ComponentRegistry) error {
	c, ok := Registry.GetComponent(componentID)
	if !ok {
		return fmt.Errorf("component with id %q not found", componentID)
	}
	if quantity <= minQuantity {
		return fmt.Errorf("quantity must be greater than %.2f", minQuantity)
	}
	bom.Items = append(bom.Items, NewBOMItem(c, quantity))
	return nil
}

func (bom *BOM) TotalCost() float64 {
	var cost float64 = 0
	for _, item := range bom.Items {
		cost += item.Quantity * item.Component.UnitCost
	}
	return cost
}

func (bom *BOM) String() string {
	var b strings.Builder

	// Header
	fmt.Fprintf(&b, "╔═══════════════════════════════════════════════════════╗\n")
	fmt.Fprintf(&b, "║ Product: %-44s ║\n", bom.ProductName)
	fmt.Fprintf(&b, "║ ID: %-49s ║\n", bom.ProductID)
	fmt.Fprintf(&b, "╠═══════════════════════════════════════════════════════╣\n")

	// Items
	if len(bom.Items) == 0 {
		fmt.Fprintf(&b, "║ No components added yet                               ║\n")
	} else {
		fmt.Fprintf(&b, "║ Components:                                           ║\n")
		for i, item := range bom.Items {
			lineTotal := item.Quantity * item.Component.UnitCost
			fmt.Fprintf(&b, "║ %2d. %-14s %6.2f %-4s × $%7.2f = $%8.2f ║\n",
				i+1,
				item.Component.Name,
				item.Quantity,
				item.Component.UnitOfMeasure,
				item.Component.UnitCost,
				lineTotal,
			)
		}
	}

	// Total
	total := fmt.Sprintf("Total: $%8.2f", bom.TotalCost())
	fmt.Fprintf(&b, "╠═══════════════════════════════════════════════════════╣\n")
	fmt.Fprintf(&b, "║ %-53s ║\n", total)
	fmt.Fprintf(&b, "╚═══════════════════════════════════════════════════════╝\n")

	return b.String()
}

type BOMRegistry struct {
	BOMS map[string]*BOM
}

func NewBOMRegistry() *BOMRegistry {
	return &BOMRegistry{make(map[string]*BOM)}
}

func (br *BOMRegistry) ListAll() {
	if len(br.BOMS) == 0 {
		fmt.Println("BOM Registry is Empty")
		return
	}
	fmt.Println("===All BOM===")
	for _, v := range br.BOMS {
		fmt.Print(v.String())
	}
}

func (br *BOMRegistry) CreateBOM(b *BOM) error {
	_, ok := br.BOMS[b.ProductID]
	if ok {
		return fmt.Errorf("product id %q already exists", b.ProductID)
	}
	br.BOMS[b.ProductID] = b
	return nil
}

func (br *BOMRegistry) GetBOM(id string) (*BOM, bool) {
	b, ok := br.BOMS[id]
	if !ok {
		return nil, false
	}
	return b, true
}

func (br *BOMRegistry) DeleteBOM(id string) {
	delete(br.BOMS, id)
}

func addBomItem(b *BOM, cr *ComponentRegistry) error {
	id, err := getInput("id")
	if err != nil {
		return fmt.Errorf("getting id: %v", err)
	}
	q, err := getInput("quantity")
	if err != nil {
		return fmt.Errorf("getting quantity input: %v", err)
	}
	quantity, err := strconv.ParseFloat(q, 64)
	if err != nil {
		return fmt.Errorf("parsing quantity to float: %v", err)
	}
	return b.AddItem(id, quantity, cr)
}

func bomItemMenu(b *BOM, cr *ComponentRegistry) {
	const (
		ViewDetails = "1"
		AddItem     = "2"
		Exit        = "3"
	)
	menu := []string{"View Details", "Add Item", "Exit"}
start:
	for {
		fmt.Printf("===BOM %v Menu!===\n", b.ProductID)
		for i, v := range menu {
			fmt.Printf("%v. %v\n", i+1, v)
		}
		input, err := getInput("option")
		if err != nil {
			fmt.Println(fmt.Errorf("getting option: %v", err))
			continue start
		}
		switch input {
		case ViewDetails:
			fmt.Println("==Bom Details!===")
			fmt.Print(b.String())
			continue start
		case AddItem:
			err := addBomItem(b, cr)
			if err != nil {
				fmt.Printf("adding BOM item: %v\n", err)
			}
			continue start
		case Exit:
			fmt.Printf("Closing BOM %v Menu\n", b.ProductID)
			break start
		}
	}
}

func getBOMData() (*BOM, error) {
	pID, err := getInput("product ID")
	if err != nil {
		return nil, fmt.Errorf("getting product ID: %v", err)
	}
	pName, err := getInput("product name")
	if err != nil {
		return nil, fmt.Errorf("getting product name: %v", err)
	}
	return NewBOM(pID, pName)
}

func bomMenu(br *BOMRegistry, cr *ComponentRegistry) {
	const (
		ListAll   = "1"
		CreateBOM = "2"
		SelectBOM = "3"
		ViewBOM   = "4"
		DeleteBOM = "5"
		Exit      = "6"
	)
	menu := []string{"List All", "Create BOM", "Select BOM", "View BOM", "Delete BOM", "Exit"}
start:
	for {
		fmt.Println("===BOM Menu!===")
		for i, v := range menu {
			fmt.Printf("%v. %v\n", i+1, v)
		}
		input, err := getInput("option")
		if err != nil {
			fmt.Println(fmt.Errorf("getting option: %v", err))
			continue start
		}
		switch input {
		case ListAll:
			br.ListAll()
			continue start
		case CreateBOM:
			fmt.Println("===Creating BOM!===")
			b, err := getBOMData()
			if err != nil {
				fmt.Printf("creating BOM: %v\n", err)
				continue start
			}
			err = br.CreateBOM(b)
			if err != nil {
				fmt.Printf("creating BOM: %v\n", err)
				continue start
			}
			fmt.Println("======Success======")
			continue start
		case SelectBOM:
			id, err := getInput("BOM id")
			if err != nil {
				fmt.Printf("viewing BOM: %v\n", err)
				continue start
			}
			b, ok := br.GetBOM(id)
			if !ok {
				fmt.Printf("no BOM found with id: %s\n", id)
				continue start
			}
			bomItemMenu(b, cr)
			continue start
		case ViewBOM:
			id, err := getInput("BOM id")
			if err != nil {
				fmt.Printf("viewing BOM: %v\n", err)
				continue start
			}
			b, ok := br.GetBOM(id)
			if !ok {
				fmt.Printf("no BOM found with id: %s\n", id)
				continue start
			}
			fmt.Print(b.String())
			continue start
		case DeleteBOM:
			id, err := getInput("BOM id")
			if err != nil {
				fmt.Printf("viewing BOM: %v\n", err)
				continue start
			}
			br.DeleteBOM(id)
			continue start
		case Exit:
			fmt.Print("Closing BOM Menu\n")
			break start
		default:
			fmt.Println("Invalid Options")
			continue start
		}
	}
	fmt.Println("===============")
}
