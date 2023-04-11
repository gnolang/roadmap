package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
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
	file, err := ioutil.ReadFile("roadmap.json")
	checkErr(err)
	var tasks []task
	err = json.Unmarshal(file, &tasks)
	checkErr(err)

	roadmap := make(map[string]task)
	for _, t := range tasks {
		roadmap[t.ID] = t
	}

	fmt.Println(roadmap)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
