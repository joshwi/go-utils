# Go Utils

## Table of contents
* [Packages](#packages)
* [Setup](#setup)

## Packages

### Utils

- HTTP REST functions: GET, POST, PUT, DELETE

- File storage functions: Scan, Read, Write

### Parser

- Text parsing functions: Compile, Collect

### Neo4j

- NEO4J DB functions: Connect, RunCypher

## Setup

### Build Executable

1. Use go package manager to get go-utils: 
```
git clone https://github.com/joshwi/go-utils.git
```

2. Change directory into repo
```
cd go-utils
```
3. Create .env
```
nano .env
```
4. Paste environment variables in file
```
NEO4J_USERNAME=neo4j
NEO4J_PASSWORD=password123
NEO4J_SERVICE_HOST=localhost
NEO4J_SERVICE_PORT=7687
```
5. Build go code into executable
```
go build -o mycollector
```
6. Run the collector

Example: Get all games from 2020 NFL Season
```
./mycollector -c='pfr_map_team'
./mycollector -c='pfr_map_season'
./mycollector -c='pfr_team_season' -q='MATCH (n:pfr_map_team_teams),(m:pfr_map_season_years) WHERE m.year="2020" RETURN DISTINCT n.tag as tag, m.year as year'
```