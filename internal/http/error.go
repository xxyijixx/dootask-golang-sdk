package http

import (
	"encoding/json"
	"fmt"
	"strings"
)

// APIError represents a DooTask API error
type APIError struct {
	Ret int    `json:"ret"`
	Msg string `json:"msg"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("API error [%d]: %s", e.Ret, e.Msg)
}

// HandleError processes the HTTP response and returns an error if the status code indicates a failure.
func HandleError(statusCode int, body []byte) error {
	var apiError APIError
	if err := json.Unmarshal(body, &apiError); err != nil {
		// If we can't parse the error response, return a generic error
		return fmt.Errorf("HTTP %d: %s", statusCode, string(body))
	}
	// If the API error doesn't have ret field set, use the HTTP status code
	if apiError.Ret == 0 {
		apiError.Ret = statusCode
	}
	return apiError
}

// ValidateRequired validates that required fields are not empty
func ValidateRequired(fields map[string]interface{}) error {
	var missing []string
	for name, value := range fields {
		if value == nil {
			missing = append(missing, name)
			continue
		}

		switch v := value.(type) {
		case string:
			if strings.TrimSpace(v) == "" {
				missing = append(missing, name)
			}
		case *string:
			if v == nil || strings.TrimSpace(*v) == "" {
				missing = append(missing, name)
			}
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing required fields: %s", strings.Join(missing, ", "))
	}

	return nil
}
