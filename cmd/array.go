package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl := template.Must(template.ParseFiles("../template_structure/array.gohtml"))

	footballers := []string{"Christiano Ronaldo", "Paul Pogba", "Mason Greenwood", "David De Gea", "Juan Mata"}
	err := tpl.Execute(os.Stdout, footballers)
	if err != nil {
		log.Fatalln(err)
	}
}