package main

import (
	"log"
	"os"
	"text/template"
)

func main() {

	// First Template will be used as main template, the others are needed to be included in first file if you need to show it
	// See example of included file in index.gohtml
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

	err = tpl.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}


}