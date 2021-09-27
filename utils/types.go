package utils

// HTTP REST STRUCTS

// Response struct for HTTP Requests
type Response struct {
	Url    string
	Type   string
	Status int
	Data   string
	Error  string
}
