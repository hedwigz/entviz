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
func toJsGraph(g *gen.Graph) jsGraph {
	graph := jsGraph{}
	for _, n := range g.Nodes {
		node := jsNode{ID: n.Name}
		for _, f := range n.Fields {
			node.Fields = append(node.Fields, jsField{
				Name: f.Name,
				Type: f.Type.String(),
			})
		}
		graph.Nodes = append(graph.Nodes, node)
		for _, e := range n.Edges {
			if e.IsInverse() {
				continue
			}
			graph.Edges = append(graph.Edges, jsEdge{
				From:  n.Name,
				To:    e.Type.Name,
				Label: e.Name,
			})
		}

	}
	return graph
}

var (
	//go:embed viz.tmpl
	tmplhtml string
	//go:embed entviz.go.tmpl
	tmplfile   string
	viztmpl    = template.Must(template.New("viz").Parse(tmplhtml))
	entviztmpl = template.Must(template.New("entviz").Parse(tmplfile))
)

func generateHTML(g *gen.Graph) ([]byte, error) {
	graph := toJsGraph(g)
	buf, err := json.Marshal(&graph)
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer
	if err := viztmpl.Execute(&b, string(buf)); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func generateGoFile() ([]byte, error) {
	var b bytes.Buffer
	if err := entviztmpl.Execute(&b, nil); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// VisualizeSchema is an ent hook that generates a static html page that visualizes the schema graph.
func VisualizeSchema() gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			buf, err := generateHTML(g)
			if err != nil {
				return err
			}
			gbuf, err := generateGoFile()
			if err != nil {
				return err
			}

			if err := ioutil.WriteFile("schema-viz.html", buf, 0644); err != nil {
				return err
			}
			if err = ioutil.WriteFile("entviz.go", gbuf, 0644); err != nil {
				return err
			}
			return nil
		})
	}
}
