// Package ratings provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package ratings

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
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
	// GetRatings request
	GetRatings(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetTrainingRating request
	GetTrainingRating(ctx context.Context, trainingUUID string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// PostTrainingRating request with any body
	PostTrainingRatingWithBody(ctx context.Context, trainingUUID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	PostTrainingRating(ctx context.Context, trainingUUID string, body PostTrainingRatingJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) GetRatings(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetRatingsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetTrainingRating(ctx context.Context, trainingUUID string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetTrainingRatingRequest(c.Server, trainingUUID)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostTrainingRatingWithBody(ctx context.Context, trainingUUID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostTrainingRatingRequestWithBody(c.Server, trainingUUID, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) PostTrainingRating(ctx context.Context, trainingUUID string, body PostTrainingRatingJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewPostTrainingRatingRequest(c.Server, trainingUUID, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewGetRatingsRequest generates requests for GetRatings
func NewGetRatingsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/ratings")
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

// NewGetTrainingRatingRequest generates requests for GetTrainingRating
func NewGetTrainingRatingRequest(server string, trainingUUID string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "trainingUUID", runtime.ParamLocationPath, trainingUUID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/ratings/%s", pathParam0)
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

// NewPostTrainingRatingRequest calls the generic PostTrainingRating builder with application/json body
func NewPostTrainingRatingRequest(server string, trainingUUID string, body PostTrainingRatingJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewPostTrainingRatingRequestWithBody(server, trainingUUID, "application/json", bodyReader)
}

// NewPostTrainingRatingRequestWithBody generates requests for PostTrainingRating with any type of body
func NewPostTrainingRatingRequestWithBody(server string, trainingUUID string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "trainingUUID", runtime.ParamLocationPath, trainingUUID)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/ratings/%s", pathParam0)
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
	// GetRatings request
	GetRatingsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetRatingsResponse, error)

	// GetTrainingRating request
	GetTrainingRatingWithResponse(ctx context.Context, trainingUUID string, reqEditors ...RequestEditorFn) (*GetTrainingRatingResponse, error)

	// PostTrainingRating request with any body
	PostTrainingRatingWithBodyWithResponse(ctx context.Context, trainingUUID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostTrainingRatingResponse, error)

	PostTrainingRatingWithResponse(ctx context.Context, trainingUUID string, body PostTrainingRatingJSONRequestBody, reqEditors ...RequestEditorFn) (*PostTrainingRatingResponse, error)
}

type GetRatingsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Ratings
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r GetRatingsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetRatingsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetTrainingRatingResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Rating
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r GetTrainingRatingResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetTrainingRatingResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type PostTrainingRatingResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSONDefault  *Error
}

// Status returns HTTPResponse.Status
func (r PostTrainingRatingResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r PostTrainingRatingResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// GetRatingsWithResponse request returning *GetRatingsResponse
func (c *ClientWithResponses) GetRatingsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*GetRatingsResponse, error) {
	rsp, err := c.GetRatings(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetRatingsResponse(rsp)
}

// GetTrainingRatingWithResponse request returning *GetTrainingRatingResponse
func (c *ClientWithResponses) GetTrainingRatingWithResponse(ctx context.Context, trainingUUID string, reqEditors ...RequestEditorFn) (*GetTrainingRatingResponse, error) {
	rsp, err := c.GetTrainingRating(ctx, trainingUUID, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetTrainingRatingResponse(rsp)
}

// PostTrainingRatingWithBodyWithResponse request with arbitrary body returning *PostTrainingRatingResponse
func (c *ClientWithResponses) PostTrainingRatingWithBodyWithResponse(ctx context.Context, trainingUUID string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*PostTrainingRatingResponse, error) {
	rsp, err := c.PostTrainingRatingWithBody(ctx, trainingUUID, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostTrainingRatingResponse(rsp)
}

func (c *ClientWithResponses) PostTrainingRatingWithResponse(ctx context.Context, trainingUUID string, body PostTrainingRatingJSONRequestBody, reqEditors ...RequestEditorFn) (*PostTrainingRatingResponse, error) {
	rsp, err := c.PostTrainingRating(ctx, trainingUUID, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParsePostTrainingRatingResponse(rsp)
}

// ParseGetRatingsResponse parses an HTTP response from a GetRatingsWithResponse call
func ParseGetRatingsResponse(rsp *http.Response) (*GetRatingsResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetRatingsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Ratings
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParseGetTrainingRatingResponse parses an HTTP response from a GetTrainingRatingWithResponse call
func ParseGetTrainingRatingResponse(rsp *http.Response) (*GetTrainingRatingResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetTrainingRatingResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Rating
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}

// ParsePostTrainingRatingResponse parses an HTTP response from a PostTrainingRatingWithResponse call
func ParsePostTrainingRatingResponse(rsp *http.Response) (*PostTrainingRatingResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &PostTrainingRatingResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && true:
		var dest Error
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSONDefault = &dest

	}

	return response, nil
}
