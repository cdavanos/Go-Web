package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

type user struct {
	Name  string
	Motto string
	Admin bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	// xs := []string{"zero", "one", "two", "three"}

	// data := struct {
	// 	Words []string
	// 	Lname string
	// }{
	// 	xs,
	// 	"McFly",
	// }

	u1 := user{
		Name:  "Buddha",
		Motto: "Something",
		Admin: false,
	}

	u2 := user{
		Name:  "Ghandi",
		Motto: "Else",
		Admin: true,
	}

	u3 := user{
		Name:  "",
		Motto: "Nobody",
		Admin: true,
	}

	users := []user{u1, u2, u3}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}
