package store

import (
	"encoding/json"
	"fmt"

	"github.com/LeoCavaille/rcpmgr/model"
	"github.com/blevesearch/bleve"
)

// The RecipeStore is the interface used to interact with our recipes data.
type RecipeStore interface {
	// All returns all the recipes available
	All() ([]model.Recipe, error)
	// Write adds a new recipe to the store
	Write(model.Recipe) error
}

// BleveRecipeStore is a RecipeStore implementation backed up by bleve (no deps, yay!)
type BleveRecipeStore struct {
	idx bleve.Index
}

func NewBleveRecipeStore(path string) RecipeStore {
	idx, err := bleve.Open(path)
	if err == bleve.ErrorIndexPathDoesNotExist {
		mapping := bleve.NewIndexMapping()
		idx, err = bleve.New(path, mapping)
	}
	if err != nil {
		panic(fmt.Errorf("could not open/create bleve idx at %s: %v", path, err))
	}

	return &BleveRecipeStore{idx: idx}
}

func (brs *BleveRecipeStore) All() ([]model.Recipe, error) {
	recipes := []model.Recipe{}

	q := bleve.NewMatchAllQuery()
	src := bleve.NewSearchRequest(q)
	res, err := brs.idx.Search(src)

	for _, hit := range res.Hits {
		data, err := brs.idx.GetInternal([]byte(hit.ID))
		if err != nil {
			panic(err)
		}

		recipe := model.Recipe{}
		json.Unmarshal(data, &recipe)
		recipes = append(recipes, recipe)
	}

	return recipes, err
}

func (brs *BleveRecipeStore) Write(r model.Recipe) error {
	data, err := json.Marshal(r)
	brs.idx.SetInternal([]byte(r.Title), data)

	err = brs.idx.Index(r.Title, data)

	return err
}
