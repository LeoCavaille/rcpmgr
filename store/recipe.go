package store

import (
	"encoding/json"

	"github.com/LeoCavaille/rcpmgr/model"
	"github.com/blevesearch/bleve"
)

const RecipeIndexPath = "recipes.bleve"

func CreateRecipeIndex() error {
	mapping := bleve.NewIndexMapping()
	_, err := bleve.New(RecipeIndexPath, mapping)
	return err
}

func GetRecipeIndex() (bleve.Index, error) {
	return bleve.Open(RecipeIndexPath)
}

func GetRecipes() ([]model.Recipe, error) {
	recipes := []model.Recipe{}

	idx, err := GetRecipeIndex()
	if err != nil {
		return nil, err
	}

	q := bleve.NewMatchAllQuery()
	src := bleve.NewSearchRequest(q)
	res, err := idx.Search(src)

	for _, hit := range res.Hits {
		data, err := idx.GetInternal([]byte(hit.ID))
		if err != nil {
			panic(err)
		}

		recipe := model.Recipe{}
		json.Unmarshal(data, &recipe)
		recipes = append(recipes, recipe)
	}

	return recipes, err
}

func WriteRecipe(r model.Recipe) error {
	idx, err := GetRecipeIndex()
	if err != nil {
		return err
	}

	data, err := json.Marshal(r)
	idx.SetInternal([]byte(r.Title), data)

	err = idx.Index(r.Title, data)

	return err
}
