package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// APIResponse represents the standard DooTask API response structure
type APIResponse[T any] struct {
	Ret  int    `json:"ret"` // 1: success, 0: error
	Msg  string `json:"msg"`
	Data T      `json:"data,omitempty"`
}

// ParseResponse parses the HTTP response and decodes the JSON into the provided interface.
func ParseResponse(resp *http.Response, result interface{}) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		return HandleError(resp.StatusCode, body)
	}

	return json.Unmarshal(body, result)
}

// ParseAPIResponse parses a standard DooTask API response
func ParseAPIResponse[T any](resp *http.Response, result *T) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		return HandleError(resp.StatusCode, body)
	}

	// Parse the response structure
	var apiResp struct {
		Ret  int             `json:"ret"`
		Msg  string          `json:"msg"`
		Data json.RawMessage `json:"data"`
	}

	if err := json.Unmarshal(body, &apiResp); err != nil {
		return err
	}

	if apiResp.Ret != 1 {
		return APIError{Ret: apiResp.Ret, Msg: apiResp.Msg}
	}

	// Try to unmarshal data directly
	if err := json.Unmarshal(apiResp.Data, result); err != nil {
		return err
	}

	return nil
}

// ParseAPIResponseToStruct parses a standard DooTask API response directly to a struct
func ParseAPIResponseToStruct[T any](resp *http.Response) (T, error) {
	var zero T

	var apiResp APIResponse[T]
	if err := ParseResponse(resp, &apiResp); err != nil {
		return zero, err
	}

	if apiResp.Ret != 1 {
		return zero, APIError{Ret: apiResp.Ret, Msg: apiResp.Msg}
	}

	return apiResp.Data, nil
}

// ParseAPIResponseFlexible parses a standard DooTask API response with flexible data type
// It handles cases where data might be an array or object
func ParseAPIResponseFlexible[T any](resp *http.Response, result *T) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		return HandleError(resp.StatusCode, body)
	}

	// First, try to parse as standard APIResponse
	var apiResp APIResponse[json.RawMessage]
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return err
	}

	if apiResp.Ret != 1 {
		return APIError{Ret: apiResp.Ret, Msg: apiResp.Msg}
	}

	// Try to unmarshal data directly to result
	if err := json.Unmarshal(apiResp.Data, result); err != nil {
		return err
	}

	return nil
}
