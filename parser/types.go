package parser

import (
	"regexp"
)

// PARSER INPUT STRUCTURES

// Config structure containing parser and metadata
type Config struct {
	Name   string
	Urls   []string
	Keys   []string
	Parser []Parser
}

// Parser structure contains post compiled regexp parsing template
type Parser struct {
	Label string  `json:"label"`
	Tags  []Tag   `json:"tags"`
	Regex []Regex `json:"regex"`
}

// Regex struct containing pre & post compiled regexp
type Regex struct {
	Name  string
	Value regexp.Regexp
}

// DATA OUTPUT STRUCTURES

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
