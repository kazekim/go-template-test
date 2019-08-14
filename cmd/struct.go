package main

import (
	"html/template"
	"log"
	"os"
)


func main() {
	superhero := struct {
		Name  string
		Position string
	}{
		Name:  "Jirawat Harnsiriwatanakit",
		Position: "Full Stack Developer",
	}

	tpl := template.Must(template.ParseFiles("../template_structure/struct.gohtml"))
	err := tpl.Execute(os.Stdout, superhero)
	if err != nil {
		log.Fatalln(err)
	}
}