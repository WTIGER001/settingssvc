// Code generated by go-swagger; DO NOT EDIT.

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddCategoryHandlerFunc turns a function with the right signature into a add category handler
type AddCategoryHandlerFunc func(AddCategoryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddCategoryHandlerFunc) Handle(params AddCategoryParams) middleware.Responder {
	return fn(params)
}

// AddCategoryHandler interface for that can handle valid add category params
type AddCategoryHandler interface {
	Handle(AddCategoryParams) middleware.Responder
}

// NewAddCategory creates a new http.Handler for the add category operation
func NewAddCategory(ctx *middleware.Context, handler AddCategoryHandler) *AddCategory {
	return &AddCategory{Context: ctx, Handler: handler}
}

/*AddCategory swagger:route POST /category Configuration addCategory

Save categories

*/
type AddCategory struct {
	Context *middleware.Context
	Handler AddCategoryHandler
}

func (o *AddCategory) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddCategoryParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
