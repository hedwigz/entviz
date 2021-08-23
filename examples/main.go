package main

import (
	"net/http"

	"github.com/hedwigz/entviz/examples/ent"
)

func main() {
	http.ListenAndServe("localhost:3002", ent.ServeEntviz())
}
