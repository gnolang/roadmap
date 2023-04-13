package main

import (
	//"encoding/json"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"github.com/gogo/protobuf/jsonpb"
	"moul.io/depviz/v3/pkg/dvmodel"
	"moul.io/godev"
)

type task dvmodel.Task

type roadmap struct {
	Tasks []task `json:"tasks"`
}

func main() {
	file, err := ioutil.ReadFile("output/roadmap.json")
	checkErr(err)
	var roadmapFile dvmodel.Batch
	var u jsonpb.Unmarshaler
	err = u.Unmarshal(bytes.NewReader(file), &roadmapFile)
	//err = json.Unmarshal(file, &roadmapFile)
	checkErr(err)

	roadmap := make(map[string]*dvmodel.Task)
	for _, t := range roadmapFile.Tasks {
		if t.Kind != 1 {
			continue
		}
		roadmap[t.ID.String()] = t
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
		node, err := graph.CreateNode(task.ID.String())
		checkErr(err)
		if task.ID == "https://github.com/gnolang/roadmap/issues/1" {
			println(godev.PrettyJSON(task))
		}

		// default
		// node.SetLabel(fmt.Sprintf("#%s - %s", task.ShortID(), task.Title))
		node.SetLabel(task.Title)
		node.SetHref(task.ID.String())
		node.SetShape("box")
		node.SetStyle("rounded")

		// exceptions
		if taskLabelExists(task, "focus") {
			node.SetStyle("filled,bold,rounded")
			node.SetFillColor("#ffeeee")
		}
		if taskLabelExists(task, "vision") {
			node.SetStyle("filled,rounded")
			node.SetFillColor("#eeeeee")
		}

		nodes[task.ID.String()] = node
	}

	for _, task := range roadmap {
		for _, dependentID := range task.IsBlocking {
			dependent := roadmap[dependentID.String()]
			name := task.ID.String() + dependent.ID.String()
			edge, err := graph.CreateEdge(name, nodes[task.ID.String()], nodes[dependent.ID.String()])
			checkErr(err)
			_ = edge
			// edge.SetLabel("blocking")
		}
		for _, dependingID := range task.IsDependingOn {
			depending := roadmap[dependingID.String()]
			name := depending.ID.String() + task.ID.String()
			edge, err := graph.CreateEdge(name, nodes[depending.ID.String()], nodes[task.ID.String()])
			checkErr(err)
			_ = edge
			// edge.SetLabel("depending")
		}
	}

	checkErr(g.RenderFilename(graph, graphviz.XDOT, "output/roadmap.dot"))
}

func taskLabelExists(t *dvmodel.Task, label string) bool {
	for _, l := range t.HasLabel {
		short := l.String()[strings.LastIndex(l.String(), "/")+1:]
		if short == label {
			return true
		}
	}
	return false
}

func (t task) ShortID() string {
	return t.ID.String()[strings.LastIndex(t.ID.String(), "/")+1:]
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
