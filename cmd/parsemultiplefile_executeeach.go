package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	// First Template will be used as main template, the others are needed to be included in first file if you need to show it
	// See example of incuded file in index.gohtml
	tpl, err := template.ParseFiles("../template/index.gohtml", "../template/index2.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	type User struct {
		Name   string
		Coupon string
		Amount int64
	}
	user := User{
		Name:   "Kim",
		Coupon: "STARKCORP",
		Amount: 5000,
	}

	// If you want to show each file, use this way (file is needed to be parse into template too.)
	err = tpl.ExecuteTemplate(os.Stdout, "index2.gohtml", user)
	if err != nil {
		panic(err)
	}


}