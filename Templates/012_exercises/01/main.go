package main

import (
	"html/template"
	"log"
	"os"
)

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

type hotels []hotel

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	h := hotels{
		hotel{
			Name:    "yo",
			Address: "135",
			City:    "California",
			Zip:     "11211",
			Region:  "Northern",
		},
		hotel{
			Name:    "H",
			Address: "4",
			City:    "L",
			Zip:     "95612",
			Region:  "southern",
		},
	}

	err := tpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
