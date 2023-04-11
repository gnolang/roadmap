package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type task struct {
	ID                string    `json:"id,omitempty"`
	CreatedAt         time.Time `json:"created_at,omitempty"`
	UpdatedAt         time.Time `json:"updated_at,omitempty"`
	LocalID           string    `json:"local_id,omitempty"`
	Kind              int       `json:"kind,omitempty"`
	Title             string    `json:"title,omitempty"`
	Description       string    `json:"description,omitempty"`
	Driver            int       `json:"driver,omitempty"`
	State             int       `json:"state,omitempty"`
	EstimatedDuration string    `json:"estimated_duration,omitempty"`
	HasAuthor         string    `json:"has_author,omitempty"`
	HasOwner          string    `json:"has_owner,omitempty"`
	HasAssignee       []string  `json:"has_assignee,omitempty"`
	HasLabel          []string  `json:"has_label,omitempty"`
	IsDependingOn     []string  `json:"is_depending_on,omitempty"`
	IsBlocking        []string  `json:"is_blocking,omitempty"`
	IsRelatedWith     []string  `json:"is_related_with,omitempty"`
}

func main() {
	file, err := ioutil.ReadFile("output/roadmap.json")
	checkErr(err)
	var tasks []task
	err = json.Unmarshal(file, &tasks)
	checkErr(err)

	roadmap := make(map[string]task)
	for _, t := range tasks {
		if t.Kind != 1 {
			continue
		}
		roadmap[t.ID] = t
	}

	g := graphviz.New()
	graph, err := g.Graph()
	checkErr(err)
	graph.SetRankDir("LR")
	defer func() {
		checkErr(graph.Close())
		g.Close()
	}()

	nodes := make(map[string]*cgraph.Node)

	for _, task := range roadmap {
		node, err := graph.CreateNode(task.ID)
		checkErr(err)
		node.SetLabel(task.Title)
		node.SetShape("box")
		node.SetStyle("rounded")
		nodes[task.ID] = node
	}

	for _, task := range roadmap {
		for _, dependentID := range task.IsBlocking {
			dependent := roadmap[dependentID]
			name := task.ID + dependent.ID
			edge, err := graph.CreateEdge(name, nodes[task.ID], nodes[dependent.ID])
			checkErr(err)
			_ = edge
			// edge.SetLabel("blocking")
		}
		for _, dependingID := range task.IsDependingOn {
			depending := roadmap[dependingID]
			name := depending.ID + task.ID
			edge, err := graph.CreateEdge(name, nodes[depending.ID], nodes[task.ID])
			checkErr(err)
			_ = edge
			// edge.SetLabel("depending")
		}
	}

	var buf bytes.Buffer
	checkErr(g.Render(graph, "dot", &buf))
	fmt.Println(buf.String())
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
