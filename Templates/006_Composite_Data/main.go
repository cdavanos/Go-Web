package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	Wisdom    []sage
	Transport []car
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {

	// slice
	//sages := []string{"yo", "sup", "whatup"}

	// map
	// sages := map[string]string{
	// 	"India":    "Ghandi",
	// 	"America":  "MLK",
	// 	"Meditate": "Buddha",
	// 	"Love":     "Jesus",
	// 	"Prophet":  "Muhammad"}

	// struct
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	ghandi := sage{
		Name:  "Ghandi",
		Motto: "Be the change",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	// slice of structs
	sages := []sage{buddha, ghandi}
	cars := []car{f, c}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
