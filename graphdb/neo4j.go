package graphdb

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/joshwi/go-utils/parser"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

var regexp_1 = regexp.MustCompile(`"`)

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

func RunCypher(session neo4j.Session, query string) [][]parser.Tag {

	output := [][]parser.Tag{}

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

func PostNode(session neo4j.Session, node string, label string, properties []parser.Tag) {

	cypher := `CREATE (n: ` + node + ` { label: "` + label + `" })`

	for _, item := range properties {
		cypher += ` SET n.` + item.Name + ` = "` + regexp_1.ReplaceAllString(item.Value, "\\'") + `"`
	}

	// cypher = regexp_1.ReplaceAllString(cypher, "'")

	result, err := session.Run(cypher, map[string]interface{}{})
	if err != nil {
		log.Println(err)
	}

	summary, err := result.Summary()

	counters := summary.Counters()

	log.Println(fmt.Sprintf(`[ Function: PostNode ] [ Label: %v ] [ Node: %v ] [ Properties Set: %v ]`, label, node, counters.PropertiesSet()))

}

func PutNode(session neo4j.Session, node string, label string, properties []parser.Tag) {

	cypher := `MERGE (n: ` + node + ` { label: "` + label + `" })`

	for _, item := range properties {
		cypher += ` SET n.` + item.Name + ` = "` + regexp_1.ReplaceAllString(item.Value, "\\'") + `"`
	}

	result, err := session.Run(cypher, map[string]interface{}{})
	if err != nil {
		log.Println(err)
	}

	summary, err := result.Summary()
	if err != nil {
		log.Println(err)
	}

	counters := summary.Counters()

	log.Println(fmt.Sprintf(`[ Function: PutNode ] [ Label: %v ] [ Node: %v ] [ Properties Set: %v ]`, label, node, counters.PropertiesSet()))

}

func StoreDB(session neo4j.Session, label string, bucket string, data parser.Output) {
	for _, item := range data.Collections {
		for n, entry := range item.Value {
			properties := []parser.Tag{}
			properties = append(properties, data.Tags...)
			properties = append(properties, entry...)
			new_bucket := bucket + "_" + item.Name
			new_label := label + "_" + strconv.Itoa(n+1)

			PutNode(session, new_bucket, new_label, properties)
		}
	}
}

// func DeleteNode(driver string, node string, label string) {
// }
