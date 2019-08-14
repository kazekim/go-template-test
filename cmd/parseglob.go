package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseGlob("../template2/*")
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
		Coupon: "STARTINDUSTRIES",
		Amount: 5000,
	}
	err = tpl.ExecuteTemplate(os.Stdout, "main", user)
	if err != nil {
		panic(err)
	}

}