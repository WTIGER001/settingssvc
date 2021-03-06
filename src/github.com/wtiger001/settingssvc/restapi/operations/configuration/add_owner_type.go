// Code generated by go-swagger; DO NOT EDIT.

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddOwnerTypeHandlerFunc turns a function with the right signature into a add owner type handler
type AddOwnerTypeHandlerFunc func(AddOwnerTypeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddOwnerTypeHandlerFunc) Handle(params AddOwnerTypeParams) middleware.Responder {
	return fn(params)
}

// AddOwnerTypeHandler interface for that can handle valid add owner type params
type AddOwnerTypeHandler interface {
	Handle(AddOwnerTypeParams) middleware.Responder
}

// NewAddOwnerType creates a new http.Handler for the add owner type operation
func NewAddOwnerType(ctx *middleware.Context, handler AddOwnerTypeHandler) *AddOwnerType {
	return &AddOwnerType{Context: ctx, Handler: handler}
}

/*AddOwnerType swagger:route POST /type Configuration addOwnerType

Add a new Owner type to the set of available types

*/
type AddOwnerType struct {
	Context *middleware.Context
	Handler AddOwnerTypeHandler
}

func (o *AddOwnerType) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddOwnerTypeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
