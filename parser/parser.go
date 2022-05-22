package parser

import (
	"encoding/json"
	"regexp"

	"github.com/joshwi/go-utils/utils"
)

func Init(name string, file string) (utils.Config, error) {

	config := utils.Config{Name: "", Urls: []string{}, Params: []string{}, Parser: []utils.Parser{}}

	// Open file with parsing configurations
	fileBytes, err := utils.Read(file)
	if err != nil {
		return config, err
	}

	// Unmarshall file into []Config struct
	var configurations []utils.Config
	json.Unmarshal(fileBytes, &configurations)

	// Find parsing config given in script args
	for _, item := range configurations {
		if name == item.Name {
			config = item
		}
	}

	// Compile parser config into regexp
	config.Parser = Compile(config.Parser)

	return config, nil
}

func Compile(parser []utils.Parser) []utils.Parser {

	// log.Println(`[ Function: Compile ] [ Start ]`)

	output := []utils.Parser{}

	for _, entry := range parser {
		tags := []utils.Regex{}
		for _, n := range entry.Regex {
			r := regexp.MustCompile(n.Name)
			exp := utils.Regex{Name: n.Name, Value: *r}
			tags = append(tags, exp)
		}
		parser := utils.Parser{Label: entry.Label, Tags: entry.Tags, Regex: tags}
		output = append(output, parser)
	}

	// log.Println(`[ Function: Compile ] [ Finish ]`)

	return output

}

func Collect(text string, parsers []utils.Parser) utils.Output {

	// log.Println(`[ Function: Collect ] [ Start ]`)

	output := utils.Output{}

	for _, parser := range parsers {
		input := Parse(text, parser.Label, parser.Regex, 0)
		output.Tags = append(output.Tags, input.Tags...)
		output.Collections = append(output.Collections, input.Collections...)
	}

	// log.Println(`[ Function: Collect ] [ Finish ]`)

	return output

}

func Parse(text string, title string, regex []utils.Regex, num int) utils.Output {

	output := utils.Output{}

	r := regex[num].Value

	response := r.FindAllStringSubmatch(text, -1)

	if len(response) > 0 {

		// If there are one or more submatches in regexp
		if len(r.SubexpNames()) > 1 {
			values := []utils.Tag{}
			collection := utils.Collection{Name: title}
			for i := range response {
				tags := []utils.Tag{}
				// Create a utils.Tag for each submatch
				for j, name := range r.SubexpNames() {
					if name != "" {
						tag := utils.Tag{Name: name, Value: response[i][j]}
						tags = append(tags, tag)
					}
				}
				if len(tags) > 1 {
					// If there are multiple tags create a collection
					collection.Value = append(collection.Value, tags)
				} else if len(tags) == 1 {
					// If there is one tag append to Tags
					values = append(values, tags...)
				}
			}
			output.Tags = append(output.Tags, values...)
			if len(collection.Value) > 0 {
				output.Collections = append(output.Collections, collection)
			}
		} else if len(r.SubexpNames()) == 1 && len(regex) > num+1 {
			// If there is one match but no submatches in the regexp
			if len(response[0]) > 0 {
				// Run the matched text against the next regexp
				result := Parse(response[0][0], title, regex, num+1)
				output.Tags = append(output.Tags, result.Tags...)
				output.Collections = append(output.Collections, result.Collections...)
			}
		}

	}

	return output

}
