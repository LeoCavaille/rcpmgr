package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/LeoCavaille/rcpmgr/fixtures"
	"github.com/LeoCavaille/rcpmgr/store"
)

// TODO: clean this?
var RStore store.RecipeStore

var options struct {
	recipeIdxPath string
}

func init() {
	flag.StringVar(&options.recipeIdxPath, "recipe_index_path", "/tmp/prod.recipes.bleve", "Path of the bleve recipe index")
}

func main() {
	// FIXME: hack for tests
	os.RemoveAll("/tmp/prod.recipes.bleve")
	RStore = store.NewBleveRecipeStore(options.recipeIdxPath)
	RStore.Write(fixtures.GaletteDesRois)
	log.Println("Loaded some fixtures, w00t w00t!")

	http.HandleFunc("/recipes", HandlerListRecipes)
	http.HandleFunc("/", HandlerHomePage)
	http.ListenAndServe(":3456", nil)
}
