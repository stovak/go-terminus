package collections

import (
	"github.com/stovak/go-terminus/config"
)

// CollectionInterface is the interface for all collections
type CollectionInterface interface {
	AddItem(item CollectionInterface)
	GetItems() []CollectionInterface
	GetItemByIndex(index int) CollectionInterface
	String() string
	GetPath() string
}

// Collection is the base struct for all collections
type Collection struct {
	tc    *config.TerminusConfig
	Items []CollectionInterface
}

// AddItem adds an item to the collection
func (c *Collection) AddItem(item CollectionInterface) {
	c.Items = append(c.Items, item)
}

// GetItems returns all items from the collection
func (c *Collection) GetItems() []CollectionInterface {
	return c.Items
}

// GetItemByIndex returns a single item from the collection
func (c *Collection) GetItemByIndex(index int) CollectionInterface {
	return c.Items[index]
}

// String returns the name of the collection
func (c *Collection) String() string {
	return "Collection"
}
