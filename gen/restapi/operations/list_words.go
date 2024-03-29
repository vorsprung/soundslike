// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListWordsHandlerFunc turns a function with the right signature into a list words handler
type ListWordsHandlerFunc func(ListWordsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListWordsHandlerFunc) Handle(params ListWordsParams) middleware.Responder {
	return fn(params)
}

// ListWordsHandler interface for that can handle valid list words params
type ListWordsHandler interface {
	Handle(ListWordsParams) middleware.Responder
}

// NewListWords creates a new http.Handler for the list words operation
func NewListWords(ctx *middleware.Context, handler ListWordsHandler) *ListWords {
	return &ListWords{Context: ctx, Handler: handler}
}

/* ListWords swagger:route POST /phononym listWords

Find sounds like

Optional description

*/
type ListWords struct {
	Context *middleware.Context
	Handler ListWordsHandler
}

func (o *ListWords) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListWordsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
