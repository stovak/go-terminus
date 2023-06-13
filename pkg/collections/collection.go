package collections

import (
	"github.com/stovak/go-terminus/config"
	"github.com/stovak/go-terminus/pkg/models"
	"net/http"
)

// CollectionInterface is the interface for all collections
type CollectionInterface interface {
	AddItem(item CollectionInterface)
	GetItems() []CollectionInterface
	GetItemByIndex(index int) CollectionInterface
	String() string
	GetPath() string
	GetCollectionRequest() *http.Request
}

// Collection is the base struct for all collections
type Collection struct {
	tc    *config.TerminusConfig
	Items []models.ModelInterface
}

// AddItem adds an item to the collection
func (c *Collection) AddItem(item models.ModelInterface) {
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
