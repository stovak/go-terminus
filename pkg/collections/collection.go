package collections

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/stovak/go-terminus/config"
	"github.com/stovak/go-terminus/pkg/models"
)

// CollectionInterface is the interface for all collections
type CollectionInterface interface {
	AddItem(item models.ModelInterface)
	GetItems() []models.ModelInterface
	GetItemByIndex(index int) models.ModelInterface
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
func (c *Collection) GetItems() []models.ModelInterface {
	return c.Items
}

// GetItemByIndex returns a single item from the collection
func (c *Collection) GetItemByIndex(index int) models.ModelInterface {
	return c.Items[index]
}

// String returns the name of the collection
func (c *Collection) String() string {
	return "Collection"
}

func (c *Collection) CreateCollectionRequest() *http.Request {
	return c.tc.CreateRequest("GET", sitePath, nil)
}

func (c *Collection) ProcessCollectionResponse(req *http.Request) error {
	resp := c.tc.SendRequest(req)
	if resp.StatusCode != 200 {
		return fmt.Errorf("error getting site: %s", resp.Status)
	}
	body, _ := io.ReadAll(resp.Body)
	err := json.Unmarshal(body, &c)
	if err != nil {
		return err
	}
	return nil
}
