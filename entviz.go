package entviz

import (
	"bytes"
	_ "embed"
	"html/template"
	"io/ioutil"
	"strings"

	"entgo.io/ent/entc/gen"
)

//go:embed index.html
var tmplhtml string

func VisualizeSchema() gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			var b bytes.Buffer
			if err := tmpl.Execute(&b, g); err != nil {
				return err
			}
			if err := ioutil.WriteFile("schema-viz.html", b.Bytes(), 0644); err != nil {
				return err
			}
			return nil
		})
	}
}

var tmpl = template.Must(template.New("viz").
	Funcs(template.FuncMap{
		"fmtType": func(s string) string {
			return strings.NewReplacer(
				".", "DOT",
				"*", "STAR",
				"[", "LBRACK",
				"]", "RBRACK",
			).Replace(s)
		},
	}).
	Parse(tmplhtml))
