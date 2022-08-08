package utils

import "errors"

const (
	RESTDataSource = "https://jsonplaceholder.typicode.com"
)

// Errors
var (
	ErrEnvFileNotFound = errors.New(".env file not found")
)
