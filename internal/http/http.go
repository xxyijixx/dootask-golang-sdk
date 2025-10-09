package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// LogRequest logs HTTP request details for debugging
func LogRequest(method, url string, body interface{}) {
	log.Printf("[REQUEST] %s %s", method, url)
	if body != nil {
		if bodyStr, err := json.MarshalIndent(body, "", "  "); err == nil {
			log.Printf("[REQUEST BODY] %s", string(bodyStr))
		}
	}
}

// LogResponse logs HTTP response details for debugging
func LogResponse(resp *http.Response) {
	log.Printf("[RESPONSE] %s %d", resp.Request.Method, resp.StatusCode)
}

// BuildQueryString builds a query string from parameters
func BuildQueryString(params map[string]interface{}) string {
	if len(params) == 0 {
		return ""
	}

	var parts []string
	for key, value := range params {
		if value != nil {
			parts = append(parts, fmt.Sprintf("%s=%v", key, value))
		}
	}

	if len(parts) > 0 {
		return "?" + strings.Join(parts, "&")
	}
	return ""
}
