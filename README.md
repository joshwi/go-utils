# Go Utils

## Table of contents
* [Package](#package)
* [Install](#install)

## Package

#### Functions

- HTTP REST: GET

- Archive: Unzip, Zip

- File System: 

    - Scan: Get list of files and subdirectories
    - Read & Write: TXT, JSON, CSV

- Regex Parser:

    - Compile: Function to compile regex in config structures
    - Collect: Runs regex parser on the string input

#### Structures

- Tag: Building block data structure for go-utils. Stores one key value pair

- Collection: Struct containing two properties. Name is a string value that represents the category of the collection. Value is a 2D nested array of Tags.

- Response: HTTP response structure with fields: Url, Method, Status, Data, Error

## Install

1. Use go package manager to get go-utils: 
```
go get github.com/joshwi/go-utils
```