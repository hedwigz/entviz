// +build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/hedwigz/entviz"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{
		Hooks: []gen.Hook{
			entviz.VisualizeSchema(),
		},
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
