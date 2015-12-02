package store

import (
	"fmt"
	"os"
	"testing"

	"github.com/LeoCavaille/rcpmgr/fixtures"
	"github.com/stretchr/testify/assert"
)

func NewTestBleveRecipeStore() (RecipeStore, string) {
	path := fmt.Sprintf("%s/recipes.bleve", os.TempDir())
	rs := NewBleveRecipeStore(path)
	return rs, path
}

func TestWriteRecipe(t *testing.T) {
	assert := assert.New(t)
	brs, temp := NewTestBleveRecipeStore()
	defer os.Remove(temp)

	err := brs.Write(fixtures.GaletteDesRois)
	assert.Nil(err)

	// now fetch it
	recipes, err := brs.All()
	assert.Len(recipes, 1)
	assert.Equal(recipes[0].Title, "Galette des Rois")
}
