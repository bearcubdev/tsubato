package tsubato

import (
	"encoding/json"
	"fmt"
)

// CatalogPage represents a 4chan board catalog page
type CatalogPage struct {
	Page    int      `json:"page"`
	Threads []Thread `json:"threads"`
}

// Catalog.Pages => []CatalogPage
type Catalog struct {
	Pages []CatalogPage
}

// GetCatalog returns the catalog for the given board name
func GetCatalog(boardname string) (Catalog, error) {
	JSON, err := get(fmt.Sprintf("/%s/catalog.json", boardname))
	if err != nil {
		return Catalog{}, err
	}

	cat := Catalog{}
	err = json.Unmarshal(JSON, &cat.Pages)
	if err != nil {
		return Catalog{}, err
	}

	return cat, nil
}
