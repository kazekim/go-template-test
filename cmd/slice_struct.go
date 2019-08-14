package main

import (
	"html/template"
	"log"
	"os"
)

type superhero struct {
Name  string
Motto string
}

func main() {
	im := superhero{
		Name:  "Iron man",
		Motto: "I am iron man",
	}
	ca := superhero{
		Name:  "Captain America",
		Motto: "Avengers assemble",
	}
	ds := superhero{
		Name:  "Doctor Strange",
		Motto: "I see things",
	}
	tpl := template.Must(template.ParseFiles("../template_structure/slice_struct.gohtml"))
	superheroes := []superhero{im, ca, ds}
	err := tpl.Execute(os.Stdout, superheroes)
	if err != nil {
		log.Fatalln(err)
	}
}