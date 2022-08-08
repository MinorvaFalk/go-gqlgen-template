package utils

import "errors"

const (
	RESTDataSource = "https://jsonplaceholder.typicode.com"
	ServerStatus   = "server is running"
)

// Errors
var (
	ErrEnvFileNotFound = errors.New(".env file not found")
)
