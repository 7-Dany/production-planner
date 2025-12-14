package main

import (
	"fmt"
	"strconv"
	"strings"
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
		return fmt.Errorf("add item failed, no component with that id")
	}
	if quantity <= 0 {
		return fmt.Errorf("add item failed, quantity must be greater than 0")
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
	fmt.Fprintf(&b, "Product ID: %v\n", bom.ProductID)
	fmt.Fprintf(&b, "Product Name: %v\n", bom.ProductName)
	for _, item := range bom.Items {
		if item.Quantity > 0 {
			fmt.Fprint(&b, item.Component.String())
		}
	}
	fmt.Fprintf(&b, "Cost: %v\n", bom.TotalCost())
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
		fmt.Println("=============")
	}
}

func (br *BOMRegistry) CreateBOM(b *BOM) error {
	_, ok := br.BOMS[b.ProductID]
	if ok {
		return fmt.Errorf("Product id already exists")
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
		return fmt.Errorf("error getting id: %v", err)
	}
	q, err := getInput("quantity")
	if err != nil {
		return fmt.Errorf("error getting quantity input: %v", err)
	}
	quantity, err := strconv.ParseFloat(q, 64)
	if err != nil {
		return fmt.Errorf("error parsing quantity to int: %v", err)
	}
	err = b.AddItem(id, quantity, cr)
	if err != nil {
		return err
	}
	return nil
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
			fmt.Println(fmt.Errorf("error getting option: %v", err))
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
				fmt.Printf("error adding BOM item: %v\n", err)
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
		return nil, fmt.Errorf("error getting product ID: %v", err)
	}
	pName, err := getInput("product name")
	if err != nil {
		return nil, fmt.Errorf("error getting product name: %v", err)
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
			fmt.Println(fmt.Errorf("error getting option: %v", err))
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
				fmt.Printf("error creating BOM: %v", err)
				continue start
			}
			err = br.CreateBOM(b)
			if err != nil {
				fmt.Printf("error creating BOM: %v", err)
				continue start
			}
			fmt.Println("======Success======")
			continue start
		case SelectBOM:
			id, err := getInput("BOM id")
			if err != nil {
				fmt.Printf("error viewing BOM: %v", err)
				continue start
			}
			b, ok := br.GetBOM(id)
			if !ok {
				fmt.Printf("BOM does not exist!")
				continue start
			}
			bomItemMenu(b, cr)
			continue start
		case ViewBOM:
			id, err := getInput("BOM id")
			if err != nil {
				fmt.Printf("error viewing BOM: %v", err)
				continue start
			}
			b, ok := br.GetBOM(id)
			if !ok {
				fmt.Printf("BOM not exist!")
				continue start
			}
			fmt.Print(b.String())
			continue start
		case DeleteBOM:
			id, err := getInput("BOM id")
			if err != nil {
				fmt.Printf("error viewing BOM: %v", err)
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
