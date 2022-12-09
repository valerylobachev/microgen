// Code generated by microgen 0.9.0. DO NOT EDIT.

// Please, do not change functions names!
package transporthttp

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	transport "github.com/valerylobachev/microgen/examples/generated/transport"
	mux "github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"path"
)

func CommonHTTPRequestEncoder(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func CommonHTTPResponseEncoder(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func _Decode_Uppercase_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.UppercaseRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func _Decode_Count_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var (
		_param string
	)
	var ok bool
	_vars := mux.Vars(r)
	_param, ok = _vars["text"]
	if !ok {
		return nil, errors.New("param text not found")
	}
	text := _param
	_param, ok = _vars["symbol"]
	if !ok {
		return nil, errors.New("param symbol not found")
	}
	symbol := _param
	return &transport.CountRequest{
		Symbol: string(symbol),
		Text:   string(text),
	}, nil
}

func _Decode_TestCase_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.TestCaseRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func _Decode_DummyMethod_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.DummyMethodRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return &req, err
}

func _Decode_Uppercase_Response(_ context.Context, r *http.Response) (interface{}, error) {
	var resp transport.UppercaseResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func _Decode_Count_Response(_ context.Context, r *http.Response) (interface{}, error) {
	var resp transport.CountResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func _Decode_TestCase_Response(_ context.Context, r *http.Response) (interface{}, error) {
	var resp transport.TestCaseResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func _Decode_DummyMethod_Response(_ context.Context, r *http.Response) (interface{}, error) {
	var resp transport.DummyMethodResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func _Encode_Uppercase_Request(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "uppercase")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_Count_Request(ctx context.Context, r *http.Request, request interface{}) error {
	req := request.(*transport.CountRequest)
	r.URL.Path = path.Join(r.URL.Path, "count",
		req.Text,
		req.Symbol,
	)
	return nil
}

func _Encode_TestCase_Request(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "test-case")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_DummyMethod_Request(ctx context.Context, r *http.Request, request interface{}) error {
	r.URL.Path = path.Join(r.URL.Path, "dummy-method")
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_Uppercase_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func _Encode_Count_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func _Encode_TestCase_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func _Encode_DummyMethod_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}
