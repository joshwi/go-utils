package parser

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/joshwi/go-utils/utils"
)

var regexp_1 = regexp.MustCompile(`\s+`)
var regexp_2 = regexp.MustCompile(`[^a-zA-Z\d]+`)
var regexp_3 = regexp.MustCompile(`\_{2,}`)

func AddParams(query map[string]string, urls []string, keys []string) []string {
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

func RunJob(query map[string]string, urls []string, config utils.Config) (string, string, utils.Output) {

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

	output := utils.Collect(text, config.Parser)

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
		output.Tags = append(output.Tags, Tag{Name: keys[n], Value: query[keys[n]]})
	}

	output.Tags = append(output.Tags, Tag{Name: "timeutc", Value: time.Now().Format(time.RFC3339)})

	label = regexp_1.ReplaceAllString(label, "_")
	label = regexp_2.ReplaceAllString(label, "_")
	label = regexp_3.ReplaceAllString(label, "_")

	return label, bucket, output
}
