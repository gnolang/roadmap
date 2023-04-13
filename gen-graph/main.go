package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"strings"
	"time"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"moul.io/godev"
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
	graph.SetLabel(fmt.Sprintf(`\n\nhttps://github.com/gnolang/roadmap - generated on %s`, time.Now().Format("2006/01/02 15:04")))
	defer func() {
		checkErr(graph.Close())
		g.Close()
	}()

	nodes := make(map[string]*cgraph.Node)

	for _, task := range roadmap {
		node, err := graph.CreateNode(task.ID)
		checkErr(err)
		if task.ID == "https://github.com/gnolang/roadmap/issues/1" {
			println(godev.PrettyJSON(task))
		}

		// default
		// node.SetLabel(fmt.Sprintf("#%s - %s", task.ShortID(), task.Title))
		node.SetLabel(task.Title)
		node.SetHref(task.ID)
		node.SetShape("box")
		node.SetStyle("rounded")

		// exceptions
		if task.LabelExists("focus") {
			node.SetStyle("filled,bold,rounded")
			node.SetFillColor("#ffeeee")
		}
		if task.LabelExists("vision") {
			node.SetStyle("filled,rounded")
			node.SetFillColor("#eeeeee")
		}

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

	checkErr(g.RenderFilename(graph, "dot", "output/roadmap.dot"))
}

func (t task) LabelExists(label string) bool {
	for _, l := range t.HasLabel {
		short := l[strings.LastIndex(l, "/")+1:]
		if short == label {
			return true
		}
	}
	return false
}

func (t task) ShortID() string {
	return t.ID[strings.LastIndex(t.ID, "/")+1:]
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
