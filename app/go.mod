module github.com/joshwi/go-utils/app

go 1.16

replace github.com/joshwi/go-utils/app/graphdb => ./app/graphdb

replace github.com/joshwi/go-utils/app/parser => ./app/parser

replace github.com/joshwi/go-utils/app/utils => ./app/utils

require (
	github.com/joho/godotenv v1.3.0
	github.com/neo4j/neo4j-go-driver v1.8.3
)
