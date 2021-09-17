package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"

	"github.com/joshwi/go-utils/graphdb"
	"github.com/joshwi/go-utils/parser"
	"github.com/joshwi/go-utils/utils"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func ComputeUrl(query map[string]string, urls []string, keys []string) []string {
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
		if response.Status == "200 OK" {
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
		if len(label) == 0 {
			label += query[keys[n]]
		} else {
			label += `_` + query[keys[n]]
		}
		output.Tags = append(output.Tags, parser.Tag{Name: keys[n], Value: query[keys[n]]})
	}

	return label, bucket, output
}

func StoreResults(driver neo4j.Driver, label string, bucket string, data parser.Output) {
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

	username := "neo4j"
	password := "nico"
	uri := "bolt://neo4j"
	driver := graphdb.Connect(uri, username, password)

	log.Println(driver.VerifyConnectivity())

	name := "pfr_team_season"
	teams := []string{"atl", "buf", "car", "chi", "cin", "cle", "clt", "crd", "dal", "den", "det", "gnb", "htx", "jax", "kan", "mia", "min", "nor", "nwe", "nyg", "nyj", "oti", "phi", "pit", "rai", "ram", "rav", "sdg", "sea", "sfo", "tam", "was"}
	year := 2021

	config := parser.Config{Name: "", Urls: []string{}, Keys: []string{}, Parser: []parser.Parser{}}

	for _, item := range parser.CONFIG_LIST {
		if name == item.Name {
			config = item
		}
	}

	config.Parser = parser.Compile(config.Parser)

	for _, team := range teams {

		query := map[string]string{"tag": team, "year": strconv.Itoa(year)}

		urls := ComputeUrl(query, config.Urls, config.Keys)

		label, bucket, data := RunJob(query, urls, config)

		StoreResults(driver, label, bucket, data)

	}

}
