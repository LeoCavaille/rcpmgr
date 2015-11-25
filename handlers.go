package main

import (
	"fmt"
	"net/http"
)

func HandlerHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are a real chef!")
}
