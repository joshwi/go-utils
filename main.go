package main

import (
	"flag"
	"log"

	"github.com/joshwi/go-utils/graphdb"
	"github.com/joshwi/go-utils/parser"
	"github.com/joshwi/go-utils/utils"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func RunScript(session neo4j.Session, entry []parser.Tag, config parser.Config) {

	// Convert params from struct [][]parser.Tag -> map[string]string
	params := map[string]string{}

	for _, item := range entry {
		params[item.Name] = item.Value
	}

	// Add params to search urls
	urls := parser.AddParams(params, config.Urls, config.Params)

	// Run GET request and parsing collection
	label, bucket, data := parser.RunJob(params, urls, config)

	// Send output data to Neo4j
	graphdb.StoreDB(session, label, bucket, data)

}

func main() {

	// Init flag values
	var query string
	var name string

	// Define flag arguments for the application
	flag.StringVar(&query, `q`, ``, `Specify config. Default: <empty>`)
	flag.StringVar(&name, `c`, `pfr_team_season`, `Specify config. Default: pfr_team_season`)
	flag.Parse()

	// Pull in env variables: username, password, uri
	username := utils.Env("NEO4J_USERNAME")
	password := utils.Env("NEO4J_PASSWORD")
	host := utils.Env("NEO4J_SERVICE_HOST")
	port := utils.Env("NEO4J_SERVICE_PORT")

	// Create application session with Neo4j
	uri := "bolt://" + host + ":" + port
	driver := graphdb.Connect(uri, username, password)
	sessionConfig := neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}
	session, err := driver.NewSession(sessionConfig)
	if err != nil {
		log.Println(err)
	}

	// Find parsing config requested by user
	config := parser.Config{Name: "", Urls: []string{}, Params: []string{}, Parser: []parser.Parser{}}

	for _, item := range parser.CONFIG_LIST {
		if name == item.Name {
			config = item
		}
	}

	// Compile parser config into regexp
	config.Parser = parser.Compile(config.Parser)

	// Grab input parameters from  Neo4j
	inputs := [][]parser.Tag{{parser.Tag{Name: "name", Value: config.Name}}}

	if len(query) > 0 {
		inputs = graphdb.RunCypher(session, query)
	}

	// var wg sync.WaitGroup

	for _, entry := range inputs {

		// wg.Add(1)
		RunScript(session, entry, config)
		// go RunScript(session, entry, config)

	}

	// wg.Wait()

}
