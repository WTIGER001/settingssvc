// Code generated by go-swagger; DO NOT EDIT.

package preferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetProfilesParams creates a new GetProfilesParams object
// with the default values initialized.
func NewGetProfilesParams() GetProfilesParams {
	var ()
	return GetProfilesParams{}
}

// GetProfilesParams contains all the bound params for the get profiles operation
// typically these are obtained from a http.Request
//
// swagger:parameters getProfiles
type GetProfilesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*ID of type to return
	  In: query
	*/
	ID []string
	/*ID of owner
	  In: query
	*/
	Ownerid *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *GetProfilesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qID, qhkID, _ := qs.GetOK("id")
	if err := o.bindID(qID, qhkID, route.Formats); err != nil {
		res = append(res, err)
	}

	qOwnerid, qhkOwnerid, _ := qs.GetOK("ownerid")
	if err := o.bindOwnerid(qOwnerid, qhkOwnerid, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetProfilesParams) bindID(rawData []string, hasKey bool, formats strfmt.Registry) error {

	var qvID string
	if len(rawData) > 0 {
		qvID = rawData[len(rawData)-1]
	}

	iDIC := swag.SplitByFormat(qvID, "")

	if len(iDIC) == 0 {
		return nil
	}

	var iDIR []string
	for _, iDIV := range iDIC {
		iDI := iDIV

		iDIR = append(iDIR, iDI)
	}

	o.ID = iDIR

	return nil
}

func (o *GetProfilesParams) bindOwnerid(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Ownerid = &raw

	return nil
}