package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type Book struct {
   ID int
   ISBN string
   Author string
   PublishedYear string
}

func main() {
	f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%q", r.UserAgent())
		book := Book{ID: 123, ISBN: "0-201-03801-3", Author: "Donald Knuth", PublishedYear: "1968"}
		jsonData, _ := json.Marshal(book)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})
	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}