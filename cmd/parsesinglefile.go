package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("../template/index2.gohtml")
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
		Coupon: "STARKINDUSTRIES",
		Amount: 5000,
	}
	err = tpl.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}