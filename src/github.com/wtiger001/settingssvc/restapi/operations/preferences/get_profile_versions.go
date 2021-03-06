// Code generated by go-swagger; DO NOT EDIT.

package preferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetProfileVersionsHandlerFunc turns a function with the right signature into a get profile versions handler
type GetProfileVersionsHandlerFunc func(GetProfileVersionsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetProfileVersionsHandlerFunc) Handle(params GetProfileVersionsParams) middleware.Responder {
	return fn(params)
}

// GetProfileVersionsHandler interface for that can handle valid get profile versions params
type GetProfileVersionsHandler interface {
	Handle(GetProfileVersionsParams) middleware.Responder
}

// NewGetProfileVersions creates a new http.Handler for the get profile versions operation
func NewGetProfileVersions(ctx *middleware.Context, handler GetProfileVersionsHandler) *GetProfileVersions {
	return &GetProfileVersions{Context: ctx, Handler: handler}
}

/*GetProfileVersions swagger:route GET /profile/{id}/version Preferences getProfileVersions

Versions of a profile

Returns a list of all profile versions

*/
type GetProfileVersions struct {
	Context *middleware.Context
	Handler GetProfileVersionsHandler
}

func (o *GetProfileVersions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetProfileVersionsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
