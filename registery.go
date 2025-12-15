package main

import "fmt"

// Registry is a generic type that stores items of type T indexed by string IDs
type Registry[T any] struct {
	items map[string]T
}

// NewRegistry creates and initializes a new Registry
func NewRegistry[T any]() *Registry[T] {
	return &Registry[T]{
		items: make(map[string]T),
	}
}

// Add adds an item with the given ID to the registry
// Returns error if ID already exists
func (r *Registry[T]) Add(id string, item T) error {
	if _, exists := r.items[id]; exists {
		return fmt.Errorf("item with id %q already exists", id)
	}
	r.items[id] = item
	return nil
}

// Get retrieves an item by ID
// Returns the item and true if found, zero value and false otherwise
func (r *Registry[T]) Get(id string) (T, bool) {
	item, ok := r.items[id]
	return item, ok
}

// Delete removes an item by ID
func (r *Registry[T]) Delete(id string) {
	delete(r.items, id)
}

// ListAll returns all items in the registry
func (r *Registry[T]) ListAll() []T {
	result := make([]T, 0, len(r.items))
	for _, item := range r.items {
		result = append(result, item)
	}
	return result
}

// Len returns the number of items in the registry
func (r *Registry[T]) Len() int {
	return len(r.items)
}
