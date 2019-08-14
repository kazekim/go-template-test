package main

import (
	"os"
	"time"
	"text/template"
)

func main() {


	var fm = template.FuncMap{
		"fdateMDY": monthDayYear,
	}
	tpl := template.Must(template.New("function.tmpl").Funcs(fm).ParseFiles("../template_structure/function.gohtml"))

	//Since the templates created by ParseFiles are named by the base names of the argument files,
	// template should usually have the name of one of the (base) names of the files.
	err := tpl.ExecuteTemplate(os.Stdout, "function.gohtml", time.Now())
	if err != nil {
		panic(err)
	}
}

func monthDayYear(t time.Time) string {
	return t.Format("January 2, 2006")
}