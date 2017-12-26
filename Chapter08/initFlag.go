package main

import (
	"flag"
	"log"
	)

var name string
var age int

func init() {
	flag.StringVar(&name, "name", "stranger", "your wonderful name")
	flag.IntVar(&age, "age", 0, "your graceful age")
}

func main(){
	flag.Parse()
	log.Printf("Hello %s (%d years), Welcome to the command line world", name, age)
}