package main

import (
	"net/http"

	"entgo.io/ent/entc/gen"
	"github.com/hedwigz/entviz"
)

func main() {
	h, err := entviz.Serve("./ent/schema", &gen.Config{})
	if err != nil {
		panic(err)
	}
	http.ListenAndServe("localhost:3002", h)
}
