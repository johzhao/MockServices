package router

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	decoder2 "mock.services/common/router/decoder"
	encoder2 "mock.services/common/router/encoder"
)

type Endpoint func(ctx context.Context, request interface{}) (interface{}, error)

type Router interface {
	Method() string
	Path() string
	HandleRequest(ctx *gin.Context)
}

func NewJSONRouter(method string, path string, requestObject interface{}, endpoint Endpoint) Router {
	return NewCustomRouter(method, path, decoder2.NewJSONRequestDecoder(requestObject), endpoint, encoder2.NewJSONResponseEncoder())
}

func NewCustomRouter(method string, path string, decoder decoder2.RequestDecoder, endpoint Endpoint, encoder encoder2.ResponseEncoder) Router {
	return &router{
		method:   method,
		path:     path,
		decoder:  decoder,
		endpoint: endpoint,
		encoder:  encoder,
	}
}

type router struct {
	method   string
	path     string
	decoder  decoder2.RequestDecoder
	endpoint Endpoint
	encoder  encoder2.ResponseEncoder
}

func (r *router) Method() string {
	return r.method
}

func (r *router) Path() string {
	return r.path
}

func (r *router) HandleRequest(ctx *gin.Context) {
	req, err := r.decodeRequest(ctx)
	if err != nil {
		r.handleError(ctx, err)
		return
	}

	resp, err := r.endpoint(ctx, req)
	if err != nil {
		r.handleError(ctx, err)
		return
	}

	if err := r.encodeResponse(ctx, resp); err != nil {
		r.handleError(ctx, err)
		return
	}
}

func (r *router) decodeRequest(ctx *gin.Context) (interface{}, error) {
	if r.decoder == nil {
		return nil, nil
	}

	req, err := r.decoder.DecodeRequest(ctx)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (r *router) encodeResponse(ctx *gin.Context, data interface{}) error {
	if r.encoder == nil {
		return fmt.Errorf("missing response encoder")
	}

	return r.encoder.ResponseWithData(ctx, data)
}

func (r *router) handleError(ctx *gin.Context, err error) {
	if r.encoder != nil {
		r.encoder.ResponseWithError(ctx, err)
	}
}
