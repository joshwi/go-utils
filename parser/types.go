package parser

import (
	"regexp"
)

// REGEX PARSER CONFIGS

// Config struct contains precompiled regexp parsing template
type Config struct {
	Label string `json:"label"`
	Tags  []Tag  `json:"tags"`
	Regex []Tag  `json:"regex"`
}

// Parser structure contains post compiled regexp parsing template
type Parser struct {
	Label string
	Tags  []Tag
	Regex []RegexTag
}

// Tag struct containing post compiled regexp
type RegexTag struct {
	Name  string
	Value regexp.Regexp
}

// DATA OUTPUT STURCTURES

// Tag strut for key value data storage
type Tag struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Collection struct contains array of key value pairs
type Collection struct {
	Name  string
	Value [][]Tag
}

// Final output contains a slice of key value tags and a collection of key value tags
type Output struct {
	Tags        []Tag
	Collections []Collection
}
