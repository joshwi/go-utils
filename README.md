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

### Installation and Use

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
NEO4J_URL=<neo4j_url>
NEO4J_USERNAME=<neo4j_username>
NEO4J_PASSWORD=<neo4j_password>
```
5. Build main.go into executable
```
go build main.go
```
6. Open cronjob editor
```
crontab -e
```
7. Create job for collection script
```
*/1 * * * * /path/to/repo/main > /path/to/repo/output.log
```