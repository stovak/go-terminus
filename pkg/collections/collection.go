package collections

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/stovak/go-terminus/config"
	"github.com/stovak/go-terminus/pkg/models"
)

var (
	cli = http.Client{}
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
	Path  string
	Tc    *config.TerminusConfig
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

func (c *Collection) CreateCollectionRequest(m string) *http.Request {
	return c.Tc.CreateRequest(m, c.GetPath(), nil)
}

func (c *Collection) ProcessCollectionResponse(req *http.Request) error {
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("error getting site: %s", resp.Status)
	}
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &c.Items)
	if err != nil {
		return err
	}
	return nil
}

// GetPath returns the path for the collection
func (c *Collection) GetPath() string {
	fmt.Println("instance of Collection")
	return c.Path
}
