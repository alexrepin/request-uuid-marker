package main

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

type registerer string

var headerName = "X-Request-Id"
var pluginName = "RequestUUIDMarker"
var HandlerRegisterer = registerer(pluginName)

func (r registerer) RegisterHandlers(f func(
	name string,
	handler func(context.Context, map[string]interface{}, http.Handler) (http.Handler, error),
)) {
	f(string(r), r.registerHandlers)
}

func (r registerer) registerHandlers(
	c context.Context,
	extra map[string]interface{},
	h http.Handler,
) (http.Handler, error) {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		requestID := uuid.New().String()
		r := req.WithContext(c)
		r.Header.Set(headerName, requestID)
		w.Header().Set(headerName, requestID)

		h.ServeHTTP(w, r)
	}), nil
}

func main() {}
