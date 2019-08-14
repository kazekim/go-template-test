package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	champions := map[string]string{
		"England":        "Manchester City",
		"Spain":    "Barcelona",
		"Italy": "Juventus",
		"France": "Paris SG",
		"Germany": "Bayern Munchen",
	}
	tpl := template.Must(template.ParseFiles("../template_structure/map.gohtml"))
	err := tpl.Execute(os.Stdout, champions)
	if err != nil {
		log.Fatalln(err)
	}
}
