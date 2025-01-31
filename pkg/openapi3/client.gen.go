// Package openapi3 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package openapi3

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// GetTasks request
	GetTasks(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostTasksWithBody request with any body
	PostTasksWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostTasks(ctx context.Context, body PostTasksJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteTasksTaskId request
	DeleteTasksTaskId(ctx context.Context, taskId TaskId, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetTasksTaskId request
	GetTasksTaskId(ctx context.Context, taskId TaskId, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PutTasksTaskIdWithBody request with any body
	PutTasksTaskIdWithBody(ctx context.Context, taskId TaskId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PutTasksTaskId(ctx context.Context, taskId TaskId, body PutTasksTaskIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostTokenWithBody request with any body
	PostTokenWithBody(ctx context.Context, params *PostTokenParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostToken(ctx context.Context, params *PostTokenParams, body PostTokenJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetTasks(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetTasksRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostTasksWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostTasksRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostTasks(ctx context.Context, body PostTasksJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostTasksRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteTasksTaskId(ctx context.Context, taskId TaskId, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteTasksTaskIdRequest(c.Server, taskId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetTasksTaskId(ctx context.Context, taskId TaskId, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetTasksTaskIdRequest(c.Server, taskId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PutTasksTaskIdWithBody(ctx context.Context, taskId TaskId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutTasksTaskIdRequestWithBody(c.Server, taskId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PutTasksTaskId(ctx context.Context, taskId TaskId, body PutTasksTaskIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPutTasksTaskIdRequest(c.Server, taskId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostTokenWithBody(ctx context.Context, params *PostTokenParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostTokenRequestWithBody(c.Server, params, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostToken(ctx context.Context, params *PostTokenParams, body PostTokenJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostTokenRequest(c.Server, params, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetTasksRequest generates requests for GetTasks
func NewGetTasksRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/tasks")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPostTasksRequest calls the generic PostTasks builder with application/json body
func NewPostTasksRequest(server string, body PostTasksJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostTasksRequestWithBody(server, "application/json", bodyReader)
}

// NewPostTasksRequestWithBody generates requests for PostTasks with any type of body
func NewPostTasksRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/tasks")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeleteTasksTaskIdRequest generates requests for DeleteTasksTaskId
func NewDeleteTasksTaskIdRequest(server string, taskId TaskId) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "taskId", runtime.ParamLocationPath, taskId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/tasks/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetTasksTaskIdRequest generates requests for GetTasksTaskId
func NewGetTasksTaskIdRequest(server string, taskId TaskId) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "taskId", runtime.ParamLocationPath, taskId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/tasks/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewPutTasksTaskIdRequest calls the generic PutTasksTaskId builder with application/json body
func NewPutTasksTaskIdRequest(server string, taskId TaskId, body PutTasksTaskIdJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPutTasksTaskIdRequestWithBody(server, taskId, "application/json", bodyReader)
}

// NewPutTasksTaskIdRequestWithBody generates requests for PutTasksTaskId with any type of body
func NewPutTasksTaskIdRequestWithBody(server string, taskId TaskId, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "taskId", runtime.ParamLocationPath, taskId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/tasks/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewPostTokenRequest calls the generic PostToken builder with application/json body
func NewPostTokenRequest(server string, params *PostTokenParams, body PostTokenJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostTokenRequestWithBody(server, params, "application/json", bodyReader)
}

// NewPostTokenRequestWithBody generates requests for PostToken with any type of body
func NewPostTokenRequestWithBody(server string, params *PostTokenParams, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/token")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	if params != nil {

		var headerParam0 string

		headerParam0, err = runtime.StyleParamWithLocation("simple", false, "Authorization", runtime.ParamLocationHeader, params.Authorization)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", headerParam0)

	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// GetTasksWithResponse request
	GetTasksWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetTasksResponse, error)

	// PostTasksWithBodyWithResponse request with any body
	PostTasksWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostTasksResponse, error)

	PostTasksWithResponse(ctx context.Context, body PostTasksJSONRequestBody, reqEditors ...RequestEditorFn) (*PostTasksResponse, error)

	// DeleteTasksTaskIdWithResponse request
	DeleteTasksTaskIdWithResponse(ctx context.Context, taskId TaskId, reqEditors ...RequestEditorFn) (*DeleteTasksTaskIdResponse, error)

	// GetTasksTaskIdWithResponse request
	GetTasksTaskIdWithResponse(ctx context.Context, taskId TaskId, reqEditors ...RequestEditorFn) (*GetTasksTaskIdResponse, error)

	// PutTasksTaskIdWithBodyWithResponse request with any body
	PutTasksTaskIdWithBodyWithResponse(ctx context.Context, taskId TaskId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutTasksTaskIdResponse, error)

	PutTasksTaskIdWithResponse(ctx context.Context, taskId TaskId, body PutTasksTaskIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutTasksTaskIdResponse, error)

	// PostTokenWithBodyWithResponse request with any body
	PostTokenWithBodyWithResponse(ctx context.Context, params *PostTokenParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostTokenResponse, error)

	PostTokenWithResponse(ctx context.Context, params *PostTokenParams, body PostTokenJSONRequestBody, reqEditors ...RequestEditorFn) (*PostTokenResponse, error)
}

type GetTasksResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Task
}

// Status returns HTTPResponse.Status
func (r GetTasksResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetTasksResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostTasksResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *Task
	JSON400      *Error
}

// Status returns HTTPResponse.Status
func (r PostTasksResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostTasksResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteTasksTaskIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON400      *Error
	JSON404      *Error
}

// Status returns HTTPResponse.Status
func (r DeleteTasksTaskIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteTasksTaskIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetTasksTaskIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Task
	JSON400      *Error
	JSON404      *Error
}

// Status returns HTTPResponse.Status
func (r GetTasksTaskIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetTasksTaskIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PutTasksTaskIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Task
	JSON400      *Error
	JSON404      *Error
}

// Status returns HTTPResponse.Status
func (r PutTasksTaskIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PutTasksTaskIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostTokenResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		// AccessToken A newly issued access token for the flow
		AccessToken string `json:"access_token"`

		// ExpiresIn The time to live of the access token in seconds
		ExpiresIn int `json:"expires_in"`

		// Scope Space separated string of issued scopes. If not present, the requested scopes were issued. If present, the issued scopes may differ from the requested scopes.
		Scope     *string   `json:"scope,omitempty"`
		TokenType TokenType `json:"token_type"`
	}
	JSON400 *Error
	JSON401 *Error
}

// Status returns HTTPResponse.Status
func (r PostTokenResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostTokenResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetTasksWithResponse request returning *GetTasksResponse
func (c *ClientWithResponses) GetTasksWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetTasksResponse, error) {
	rsp, err := c.GetTasks(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetTasksResponse(rsp)
}

// PostTasksWithBodyWithResponse request with arbitrary body returning *PostTasksResponse
func (c *ClientWithResponses) PostTasksWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostTasksResponse, error) {
	rsp, err := c.PostTasksWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostTasksResponse(rsp)
}

func (c *ClientWithResponses) PostTasksWithResponse(ctx context.Context, body PostTasksJSONRequestBody, reqEditors ...RequestEditorFn) (*PostTasksResponse, error) {
	rsp, err := c.PostTasks(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostTasksResponse(rsp)
}

// DeleteTasksTaskIdWithResponse request returning *DeleteTasksTaskIdResponse
func (c *ClientWithResponses) DeleteTasksTaskIdWithResponse(ctx context.Context, taskId TaskId, reqEditors ...RequestEditorFn) (*DeleteTasksTaskIdResponse, error) {
	rsp, err := c.DeleteTasksTaskId(ctx, taskId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteTasksTaskIdResponse(rsp)
}

// GetTasksTaskIdWithResponse request returning *GetTasksTaskIdResponse
func (c *ClientWithResponses) GetTasksTaskIdWithResponse(ctx context.Context, taskId TaskId, reqEditors ...RequestEditorFn) (*GetTasksTaskIdResponse, error) {
	rsp, err := c.GetTasksTaskId(ctx, taskId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetTasksTaskIdResponse(rsp)
}

// PutTasksTaskIdWithBodyWithResponse request with arbitrary body returning *PutTasksTaskIdResponse
func (c *ClientWithResponses) PutTasksTaskIdWithBodyWithResponse(ctx context.Context, taskId TaskId, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PutTasksTaskIdResponse, error) {
	rsp, err := c.PutTasksTaskIdWithBody(ctx, taskId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutTasksTaskIdResponse(rsp)
}

func (c *ClientWithResponses) PutTasksTaskIdWithResponse(ctx context.Context, taskId TaskId, body PutTasksTaskIdJSONRequestBody, reqEditors ...RequestEditorFn) (*PutTasksTaskIdResponse, error) {
	rsp, err := c.PutTasksTaskId(ctx, taskId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePutTasksTaskIdResponse(rsp)
}

// PostTokenWithBodyWithResponse request with arbitrary body returning *PostTokenResponse
func (c *ClientWithResponses) PostTokenWithBodyWithResponse(ctx context.Context, params *PostTokenParams, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostTokenResponse, error) {
	rsp, err := c.PostTokenWithBody(ctx, params, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostTokenResponse(rsp)
}

func (c *ClientWithResponses) PostTokenWithResponse(ctx context.Context, params *PostTokenParams, body PostTokenJSONRequestBody, reqEditors ...RequestEditorFn) (*PostTokenResponse, error) {
	rsp, err := c.PostToken(ctx, params, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostTokenResponse(rsp)
}

// ParseGetTasksResponse parses an HTTP response from a GetTasksWithResponse call
func ParseGetTasksResponse(rsp *http.Response) (*GetTasksResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetTasksResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Task
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParsePostTasksResponse parses an HTTP response from a PostTasksWithResponse call
func ParsePostTasksResponse(rsp *http.Response) (*PostTasksResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostTasksResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest Task
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	}

	return response, nil
}

// ParseDeleteTasksTaskIdResponse parses an HTTP response from a DeleteTasksTaskIdWithResponse call
func ParseDeleteTasksTaskIdResponse(rsp *http.Response) (*DeleteTasksTaskIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteTasksTaskIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParseGetTasksTaskIdResponse parses an HTTP response from a GetTasksTaskIdWithResponse call
func ParseGetTasksTaskIdResponse(rsp *http.Response) (*GetTasksTaskIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetTasksTaskIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Task
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParsePutTasksTaskIdResponse parses an HTTP response from a PutTasksTaskIdWithResponse call
func ParsePutTasksTaskIdResponse(rsp *http.Response) (*PutTasksTaskIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PutTasksTaskIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Task
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 404:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON404 = &dest

	}

	return response, nil
}

// ParsePostTokenResponse parses an HTTP response from a PostTokenWithResponse call
func ParsePostTokenResponse(rsp *http.Response) (*PostTokenResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &PostTokenResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			// AccessToken A newly issued access token for the flow
			AccessToken string `json:"access_token"`

			// ExpiresIn The time to live of the access token in seconds
			ExpiresIn int `json:"expires_in"`

			// Scope Space separated string of issued scopes. If not present, the requested scopes were issued. If present, the issued scopes may differ from the requested scopes.
			Scope     *string   `json:"scope,omitempty"`
			TokenType TokenType `json:"token_type"`
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 400:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON400 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 401:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON401 = &dest

	}

	return response, nil
}
