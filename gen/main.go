package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tmpl := template.Must(template.ParseFiles("./gen/types.tmpl"))

	f, err := os.Create("arbitrary.gen.go")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	sizes := make([]int, 64)

	for idx := range sizes {
		sizes[idx] = idx + 1
	}

	if err := tmpl.Execute(f, sizes); err != nil {
		log.Fatal(err)
	}
}
