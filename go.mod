module github.com/joshwi/go-utils

go 1.16

replace github.com/joshwi/go-utils/parser => ./parser

replace github.com/joshwi/go-utils/utils => ./utils

require (
	github.com/joho/godotenv v1.3.0
	github.com/neo4j/neo4j-go-driver v1.8.3
)
