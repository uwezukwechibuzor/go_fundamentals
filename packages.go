package main

import (
	"log"

	"github.com/uwezukwechibuzor/go_fundamentals/helpers"
)

func main() {
	var myVar helpers.SomeType
	myVar.TypeName = "Some name"

	log.Println(myVar.TypeName)
}
