package utils

import (
	"log"
	"regexp"
)

func Compile(config []Config) []Parser {

	log.Println(`[ Function: Compile ] [ Start ]`)

	output := []Parser{}

	for _, entry := range config {
		tags := []RegexTag{}
		for _, n := range entry.Regex {
			r := regexp.MustCompile(n.Value)
			exp := RegexTag{Name: n.Name, Value: *r}
			tags = append(tags, exp)
		}
		parser := Parser{Label: entry.Label, Tags: entry.Tags, Regex: tags}
		output = append(output, parser)
	}

	log.Println(`[ Function: Compile ] [ Finish ]`)

	return output

}

func Collect(text string, parsers []Parser) Output {

	log.Println(`[ Function: Collect ] [ Start ]`)

	output := Output{}

	for _, parser := range parsers {
		input := Parse(text, parser.Label, parser.Regex, 0)
		output.Tags = append(output.Tags, input.Tags...)
		output.Collections = append(output.Collections, input.Collections...)
	}

	log.Println(`[ Function: Collect ] [ Finish ]`)

	return output

}

func Parse(text string, title string, regex []RegexTag, num int) Output {

	output := Output{}

	r := regex[num].Value

	response := r.FindAllStringSubmatch(text, -1)

	if len(response) > 0 {
		if len(r.SubexpNames()) > 1 {
			values := []Tag{}
			collection := Collection{Name: title}
			for i := range response {
				tags := []Tag{}
				for j, name := range r.SubexpNames() {
					if name != "" {
						tag := Tag{Name: name, Value: response[i][j]}
						tags = append(tags, tag)
					}
				}
				if len(tags) > 1 {
					collection.Value = append(collection.Value, tags)
				} else if len(tags) == 1 {
					values = append(values, tags...)
				}

			}
			output.Tags = append(output.Tags, values...)
			if len(collection.Value) > 0 {
				output.Collections = append(output.Collections, collection)
			}
		} else if len(r.SubexpNames()) == 1 && len(regex) > num+1 {
			if len(response[0]) > 0 {
				result := Parse(response[0][0], title, regex, num+1)
				output.Tags = append(output.Tags, result.Tags...)
				output.Collections = append(output.Collections, result.Collections...)
			}
		}

	}

	return output

}
