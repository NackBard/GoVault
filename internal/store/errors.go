package store

import "fmt"

var ErrNotFound = fmt.Errorf("not found")
var ErrInternalServer = fmt.Errorf("internal server error")
var ErrInvalidId = fmt.Errorf("invalid id")
var ErrInvalidRequestBody = fmt.Errorf("invalid request body")
