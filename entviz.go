package entviz

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"html/template"

	"io/ioutil"

	"entgo.io/ent/entc/gen"
)

type (
	jsGraph struct {
		Nodes []jsNode `json:"nodes"`
		Edges []jsEdge `json:"edges"`
	}

	jsNode struct {
		ID     string    `json:"id"`
		Fields []jsField `json:"fields"`
	}

	jsEdge struct {
		From  string `json:"from"`
		To    string `json:"to"`
		Label string `json:"label"`
	}

	jsField struct {
		Name string `json:"name"`
		Type string `json:"type"`
	}
)

// toJsGraph converts ent's graph into json serializable struct
func toJsGraph(g *gen.Graph) *jsGraph {
	var jsGraph = &jsGraph{}
	for _, n := range g.Nodes {
		node := jsNode{ID: n.Name}

		for _, f := range n.Fields {
			node.Fields = append(node.Fields, jsField{
				Name: f.Name,
				Type: f.Type.String(),
			})
		}
		jsGraph.Nodes = append(jsGraph.Nodes, node)

		for _, e := range n.Edges {
			if e.IsInverse() {
				continue
			}
			jsGraph.Edges = append(jsGraph.Edges, jsEdge{
				From:  n.Name,
				To:    e.Type.Name,
				Label: e.Name,
			})
		}

	}
	return jsGraph
}

//go:embed viz.html
var tmplhtml string

// VisualizeSchema is an ent's hook that generates static html that visualizes the schema graph.
func VisualizeSchema() gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			jsGraph := toJsGraph(g)
			graph, err := json.Marshal(jsGraph)
			if err != nil {
				return err
			}
			var b bytes.Buffer
			if err := tmpl.Execute(&b, string(graph)); err != nil {
				return err
			}

			if err := ioutil.WriteFile("schema-viz.html", b.Bytes(), 0644); err != nil {
				return err
			}
			return nil
		})
	}
}

var tmpl = template.Must(template.New("viz").Parse(tmplhtml))
