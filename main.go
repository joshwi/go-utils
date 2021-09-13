package main

import (
	"log"

	"github.com/joshwi/go-utils/parser"
	"github.com/joshwi/go-utils/utils"
)

func main() {

	// file, err := ioutil.ReadFile("input.txt")
	// if err != nil {
	// 	log.Println(err)
	// }
	// text := string(file)
	// parser := utils.Compile(WIKI_MOVIE)
	// output := utils.Collect(text, parser)

	//TESTS

	// output := utils.Get("http://localhost:5000/api")
	// output := utils.Scan("/Users/josh/Desktop/go-utils")
	// output := utils.Read("/Users/josh/Desktop/go-utils/test.json")

	text := utils.Get("https://www.pro-football-reference.com/teams/kan/2020.htm")

	search := parser.Compile(parser.PFR_TEAM_SEASON)

	output := parser.Collect(text.Data, search)

	log.Println(output)

}
