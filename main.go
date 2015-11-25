package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/LeoCavaille/rcpmgr/store"
)

var options struct {
	initMode bool
}

func init() {
	flag.BoolVar(&options.initMode, "init", false, "Initialize rcpmgr (first run).")
}

func main() {
	init()

	if options.initMode {
		err := store.CreateRecipeIndex()
		if err != nil {
			panic(err)
		}

		log.Info("Initialized rcpmgr, now run me!")
		return
	}

	http.HandleFunc("/", HandlerHomePage)
	http.ListenAndServe(":3456", nil)
}
