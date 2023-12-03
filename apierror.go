package gvk

import "fmt"

// APIError represents an error returned by the Telegram API.
type APIError struct {
	//desc string
	code int
}

// ErrorCode returns the error code received from the Telegram API.
func (a *APIError) ErrorCode() int {
	return a.code
}

//// Description returns the error description received from the Telegram API.
//func (a *APIError) Description() string {
//	return a.desc
//}

// Error returns the error string.
func (a *APIError) Error() string {
	return fmt.Sprintf("Failed error: %d", a.code)
}
