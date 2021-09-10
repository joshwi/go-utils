package main

import (
	"log"
	"regexp"

	utils "./utils"
)

type Tag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type RegexTag struct {
	Name  string
	Value regexp.Regexp
}

type Collection struct {
	Name  string
	Value [][]Tag
}

type Output struct {
	Tags        []Tag
	Collections []Collection
}

type Config struct {
	Label string `json:"label"`
	Tags  []Tag  `json:"tags"`
	Regex []Tag  `json:"regex"`
}

type Parser struct {
	Label string
	Tags  []Tag
	Regex []RegexTag
}

var WIKI_MOVIE = []Config{
	{
		Label: "producer",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<tr><th[^>]+>Produced by<\\/th><td[^>]+>.*?<\\/td><\\/tr>"},
			{Name: "", Value: "<a.href=\"(?P<url>(.*?))\"[^>]+>(?P<producer>(.*?))<\\/a>"},
		},
	}, {
		Label: "director",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<tr><th[^>]+>Directed by<\\/th><td[^>]+>.*?<\\/td><\\/tr>"},
			{Name: "", Value: "<a.href=\"(?P<url>(.*?))\"[^>]+>(?P<director>(.*?))<\\/a>"},
		},
	}, {
		Label: "screenplay",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<tr><th[^>]+>Screenplay by<\\/th><td[^>]+>.*?<\\/td><\\/tr>"},
			{Name: "", Value: "<a.href=\"(?P<url>(.*?))\"[^>]+>(?P<writer>(.*?))<\\/a>"},
		},
	}, {
		Label: "cast",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<tr><th[^>]+>Starring<\\/th><td[^>]+>.*?<\\/td><\\/tr>"},
			{Name: "", Value: "<a.href=\"(?P<url>(.*?))\"[^>]+>(?P<actor>(.*?))<\\/a>"},
		},
	}, {
		Label: "score",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<tr><th[^>]+>Music by<\\/th><td[^>]+>.*?<\\/td><\\/tr>"},
			{Name: "", Value: "<a.href=\"(?P<url>(.*?))\"[^>]+>(?P<artist>(.*?))<\\/a>"},
		},
	}, {
		Label: "releaseDate",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<tr><th[^>]+>.*?Release date.*?<\\/th>.*?<\\/tr>"},
			{Name: "", Value: "<span[^>]+>(?P<releaseDate>(\\d{4}-\\d{2}-\\d{2}))<\\/span>"},
		},
	},
	{
		Label: "runtime",
		Tags: []Tag{
			{Name: "test", Value: "test"},
		},
		Regex: []Tag{
			{Name: "", Value: "(?ms)<tr><th[^>]+>.*?Running time.*?<\\/th>.*?<\\/tr>"},
			{Name: "", Value: "<td[^>]+>(?P<length>(\\d+)).minutes.*?<\\/td>"},
		},
	},
}

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

func main() {

	// log.Println([]schema.Config{})

	// schema.Test()

	// file, err := ioutil.ReadFile("input.txt")

	// if err != nil {
	// 	log.Println(err)
	// }

	// text := string(file)

	// parser := Compile(WIKI_MOVIE)

	// output := Collect(text, parser)

	// log.Println(output)

	output := utils.Scan("/Users/josh/Desktop/go-utils")

	log.Println(output)

}
