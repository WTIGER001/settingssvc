// Code generated by go-swagger; DO NOT EDIT.

package preferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateProfileHandlerFunc turns a function with the right signature into a update profile handler
type UpdateProfileHandlerFunc func(UpdateProfileParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateProfileHandlerFunc) Handle(params UpdateProfileParams) middleware.Responder {
	return fn(params)
}

// UpdateProfileHandler interface for that can handle valid update profile params
type UpdateProfileHandler interface {
	Handle(UpdateProfileParams) middleware.Responder
}

// NewUpdateProfile creates a new http.Handler for the update profile operation
func NewUpdateProfile(ctx *middleware.Context, handler UpdateProfileHandler) *UpdateProfile {
	return &UpdateProfile{Context: ctx, Handler: handler}
}

/*UpdateProfile swagger:route POST /profile/{id} Preferences updateProfile

Update an existing profile

*/
type UpdateProfile struct {
	Context *middleware.Context
	Handler UpdateProfileHandler
}

func (o *UpdateProfile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateProfileParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}