package main

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/joshwi/go-utils/graphdb"
	"github.com/joshwi/go-utils/parser"
	"github.com/joshwi/go-utils/utils"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

var regexp_1 = regexp.MustCompile(`\s+`)
var regexp_2 = regexp.MustCompile(`[^a-zA-Z\d]+`)
var regexp_3 = regexp.MustCompile(`\_{2,}`)

func InitJob(query map[string]string, urls []string, keys []string) []string {
	output := []string{}

	for _, url := range urls {
		for _, key := range keys {
			re, _ := regexp.Compile(fmt.Sprintf("{%v}", key))
			url = re.ReplaceAllString(url, query[key])
		}
		output = append(output, url)
	}

	return output
}

func RunJob(query map[string]string, urls []string, config parser.Config) (string, string, parser.Output) {

	text := ``

	for _, url := range urls {

		response := utils.Get(url)
		if response.Status == 200 {
			text = response.Data
			break
		}
	}

	label := ``
	bucket := config.Name

	output := parser.Collect(text, config.Parser)

	//Sort keys alphabetically
	keys := make([]string, 0, len(query))
	for k := range query {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Add query parameters to output Tags
	for n := range keys {
		if !strings.Contains(keys[n], "url") {
			if len(label) == 0 {
				label += query[keys[n]]
			} else {
				label += `_` + query[keys[n]]
			}
		}
		output.Tags = append(output.Tags, parser.Tag{Name: keys[n], Value: query[keys[n]]})
	}

	label = regexp_1.ReplaceAllString(label, "_")
	label = regexp_2.ReplaceAllString(label, "_")
	label = regexp_3.ReplaceAllString(label, "_")

	return label, bucket, output
}

func StoreDB(driver neo4j.Session, label string, bucket string, data parser.Output) {
	for _, item := range data.Collections {
		for n, entry := range item.Value {
			properties := []parser.Tag{}
			properties = append(properties, data.Tags...)
			properties = append(properties, entry...)
			new_bucket := bucket + "_" + item.Name
			new_label := label + "_" + strconv.Itoa(n+1)
			graphdb.PutNode(driver, new_bucket, new_label, properties)
		}
	}
}

func main() {

	var query string
	var name string

	// flags declaration using flag package
	flag.StringVar(&query, `q`, ``, `Specify config. Default: <empty>`)
	flag.StringVar(&name, `c`, `pfr_team_season`, `Specify config. Default: pfr_team_season`)
	flag.Parse()

	username := utils.Env("NEO4J_USERNAME")
	password := utils.Env("NEO4J_PASSWORD")
	host := utils.Env("NEO4J_SERVICE_HOST")
	port := utils.Env("NEO4J_SERVICE_PORT")
	uri := "bolt://" + host + ":" + port
	driver := graphdb.Connect(uri, username, password)
	sessionConfig := neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}
	session, err := driver.NewSession(sessionConfig)
	if err != nil {
		log.Println(err)
	}

	config := parser.Config{Name: "", Urls: []string{}, Keys: []string{}, Parser: []parser.Parser{}}

	for _, item := range parser.CONFIG_LIST {
		if name == item.Name {
			config = item
		}
	}

	config.Parser = parser.Compile(config.Parser)

	inputs := [][]parser.Tag{{
		parser.Tag{Name: "name", Value: config.Name},
	}}

	if len(query) > 0 {
		inputs = graphdb.RunCypher(session, query)
	}

	for _, entry := range inputs {

		params := map[string]string{}

		for _, item := range entry {
			params[item.Name] = item.Value
		}

		urls := InitJob(params, config.Urls, config.Keys)

		_, _, data := RunJob(params, urls, config)

		// log.Println(data)

		label, bucket, data := RunJob(params, urls, config)

		StoreDB(session, label, bucket, data)

	}

}
