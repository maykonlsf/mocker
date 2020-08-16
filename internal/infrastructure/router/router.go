package router

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/maykonlf/mocker/internal/model/entities"
	"github.com/valyala/fasthttp"
)

func NewRouter(addr string) Router {
	return &router{
		configs: map[string]map[string]*entities.APIResponse{},
		addr:    addr,
	}
}

type router struct {
	addr    string
	configs map[string]map[string]*entities.APIResponse
	handler fasthttp.RequestHandler
}

func (r *router) Listen() {
	r.handler = r.rootHandler
	fmt.Println("serving mock API at", r.addr)
	panic(fasthttp.ListenAndServe(r.addr, r.handler))
}

func (r *router) Set(route, method string, response *entities.APIResponse) error {
	routeKey := r.configs[route]
	if routeKey == nil {
		r.configs[route] = map[string]*entities.APIResponse{}
	}

	routeMethodKey := r.configs[route][method]
	if routeMethodKey != nil {
		return errors.New("route-method conflicted")
	}

	r.configs[route][method] = response
	return nil
}

func (r *router) rootHandler(ctx *fasthttp.RequestCtx) {
	if ctx.IsOptions() {
		r.setOptionsResponse(ctx)
		return
	}

	route := r.configs[string(ctx.Path())]
	if route == nil {
		ctx.Error("not found", fasthttp.StatusNotFound)
		return
	}

	r.setResponse(ctx, route)
}

func (r *router) setResponse(ctx *fasthttp.RequestCtx, route map[string]*entities.APIResponse) {
	methodKey := strings.ToLower(string(ctx.Method()))
	response := route[methodKey]
	if response == nil {
		ctx.Error("method not allowed", fasthttp.StatusMethodNotAllowed)
		return
	}

	r.setConfiguredResponse(ctx, response)
}

func (r *router) setConfiguredResponse(ctx *fasthttp.RequestCtx, response *entities.APIResponse) {
	time.Sleep(response.Time)
	ctx.Response.SetStatusCode(response.Status)
	ctx.Response.SetBody([]byte(response.Body))
	for key, value := range response.Headers {
		ctx.Response.Header.Set(key, value)
	}
}

func (r *router) setOptionsResponse(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "*")
}
