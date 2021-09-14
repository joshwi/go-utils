package main

import (
	"strconv"

	"github.com/joshwi/go-utils/graphdb"
	"github.com/joshwi/go-utils/parser"
	"github.com/joshwi/go-utils/utils"
)

func main() {

	username := utils.Env("NEO4J_USERNAME")
	password := utils.Env("NEO4J_PASSWORD")
	uri := utils.Env("NEO4J_URL")
	driver := graphdb.Connect(uri, username, password)

	query := map[string]string{"tag": "kan", "year": "2020"}
	text := utils.Get("https://www.pro-football-reference.com/teams/kan/2020.htm")
	search := parser.Compile(parser.PFR_TEAM_SEASON)
	output := parser.Collect(text.Data, search)

	label := ``
	bucket := "test"

	for key, value := range query {
		if len(label) == 0 {
			label += value
		} else {
			label += `_` + value
		}
		output.Tags = append(output.Tags, parser.Tag{Name: key, Value: value})
	}

	for _, item := range output.Collections {
		for n, entry := range item.Value {
			properties := []parser.Tag{}
			properties = append(properties, output.Tags...)
			properties = append(properties, entry...)
			new_bucket := bucket + "_" + item.Name
			new_label := label + "_" + strconv.Itoa(n)
			graphdb.PutNode(driver, new_bucket, new_label, properties)
		}
	}

}
