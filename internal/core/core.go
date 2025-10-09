package core

import "net/http"

// HTTPDoer abstracts the subset of client behavior needed by services
type HTTPDoer interface {
	DoRequest(method, endpoint string, body interface{}) (*http.Response, error)
}
