package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()
	// Mapping to methods is possible with HttpRouter
	router.ServeFiles("/static/*filepath", http.Dir("/Users/naren/static"))
	log.Fatal(http.ListenAndServe(":8000", router))
}
