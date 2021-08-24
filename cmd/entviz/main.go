package main

import (
	"log"
	"os"

	"entgo.io/ent/entc/gen"
	"github.com/hedwigz/entviz"
)

func main() {
	path := "./ent/schema"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	b, err := entviz.GeneratePage(path, &gen.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	if err := os.WriteFile("schema-viz.html", b, 0644); err != nil {
		log.Fatal(err)
	}
}
