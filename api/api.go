package api

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"moul.io/http2curl"
)

// TODO:
// Callback function config.
//
// Request sends an HTTP request to the given URL.
func Request(ctx context.Context, payload io.Reader, method, url string, header map[string]string) Response {
	var curl string

	req, err := createRequest(ctx, url, method, payload)
	if err != nil {
		return responseError(err, req.RequestURI, req.Method, curl)
	}

	addCustomHeader(req, header)

	curl, err = Curl(req)
	if err != nil {
		return responseError(err, req.RequestURI, req.Method, curl)
	}

	response, err := sendRequest(req)
	if err != nil {
		return responseError(err, req.RequestURI, req.Method, curl)
	}

	responseBody, err := readResponse(response)
	if err != nil {
		return responseError(err, req.RequestURI, req.Method, curl)
	}

	return NewResponse(response.StatusCode, responseBody, response.Request.URL.String(), response.Request.Method, curl, nil)
}

func createRequest(ctx context.Context, url, method string, payload io.Reader) (*http.Request, error) {
	var req *http.Request
	var err error

	req, err = http.NewRequestWithContext(ctx, method, url, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	return req, err
}

func readResponse(res *http.Response) ([]byte, error) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return body, fmt.Errorf("failed to read response body: %w", err)
	}

	defer res.Body.Close()

	return body, nil
}

func addCustomHeader(r *http.Request, headers map[string]string) {
	for k, v := range headers {
		r.Header.Add(k, v)
	}
}

func sendRequest(r *http.Request) (*http.Response, error) {
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	return res, nil
}

// Curl returns the curl for the given http.Request.
func Curl(r *http.Request) (string, error) {
	curl, err := http2curl.GetCurlCommand(r)
	if err != nil {
		return "", fmt.Errorf("failed to get curl command: %w", err)
	}

	return curl.String(), nil
}
