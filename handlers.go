package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func HandlerHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are a real chef!")
}

func HandlerListRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := RStore.All()
	if err != nil {
		log.Printf("Could not get recipes: %v", err)
		w.WriteHeader(500)
		return
	}

	t, err := template.ParseFiles("tmpl/recipes.html")
	if err != nil {
		log.Printf("Could not parse template: %v", err)
		w.WriteHeader(500)
		return
	}

	t.Execute(w, recipes)
}
