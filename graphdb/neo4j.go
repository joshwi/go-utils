package graphdb

import (
	"fmt"
	"log"

	"github.com/joshwi/go-utils/parser"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func Connect(url string, username string, password string) neo4j.Driver {

	Neo4jConfig := func(conf *neo4j.Config) { conf.Encrypted = false }

	driver, err := neo4j.NewDriver(url, neo4j.BasicAuth(username, password, ""), Neo4jConfig)
	if err != nil {
		log.Println(err)
	}

	// handle driver lifetime based on your application lifetime requirements
	// driver's lifetime is usually bound by the application lifetime, which usually implies one driver instance per application
	// defer driver.Close()

	return driver
}

func RunCypher(driver neo4j.Driver, query string) [][]parser.Tag {

	output := [][]parser.Tag{}

	sessionConfig := neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}
	session, err := driver.NewSession(sessionConfig)
	if err != nil {
		log.Println(err)
	}
	// defer session.Close()

	result, err := session.Run(query, map[string]interface{}{})
	if err != nil {
		log.Println(err)
	}

	for result.Next() {
		entry := []parser.Tag{}
		keys := result.Record().Keys()
		for n := 0; n < len(keys); n++ {
			value := fmt.Sprintf("%v", result.Record().GetByIndex(n))
			input := parser.Tag{
				Name:  keys[n],
				Value: value,
			}
			entry = append(entry, input)
		}
		output = append(output, entry)
	}

	return output
}

// func GetNode(driver string, node string, query string) [][]parser.Tag {

// 	return [][]parser.Tag{}
// }

func PutNode(driver neo4j.Driver, node string, label string, properties []parser.Tag) {

	// log.Println(`[ Function: PutNode ] [ Start ]`)

	sessionConfig := neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite}
	session, err := driver.NewSession(sessionConfig)
	if err != nil {
		log.Println(err)
	}

	cypher := `MERGE (n: ` + node + ` { label: "` + label + `" })`

	for _, item := range properties {
		cypher += ` SET n.` + item.Name + ` = "` + item.Value + `"`
	}

	result, err := session.Run(cypher, map[string]interface{}{})

	log.Println(result)

	// summary, err := result.Summary()

	// counters := summary.Counters()

	// log.Println(fmt.Sprintf(`[ Function: PutNode ] [ Label: %v ] [ Node: %v ] [ Properties Set: %v ]`, label, node, counters.PropertiesSet()))

	// log.Println(`[ Function: PutNode ] [ Finish ]`)

	session.Close()

}
