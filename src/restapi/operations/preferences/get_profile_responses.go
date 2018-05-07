// Code generated by go-swagger; DO NOT EDIT.

package preferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"models"
)

// GetProfileOKCode is the HTTP code returned for type GetProfileOK
const GetProfileOKCode int = 200

/*GetProfileOK successful operation

swagger:response getProfileOK
*/
type GetProfileOK struct {

	/*
	  In: Body
	*/
	Payload *models.Profile `json:"body,omitempty"`
}

// NewGetProfileOK creates GetProfileOK with default headers values
func NewGetProfileOK() *GetProfileOK {
	return &GetProfileOK{}
}

// WithPayload adds the payload to the get profile o k response
func (o *GetProfileOK) WithPayload(payload *models.Profile) *GetProfileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get profile o k response
func (o *GetProfileOK) SetPayload(payload *models.Profile) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProfileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProfileBadRequestCode is the HTTP code returned for type GetProfileBadRequest
const GetProfileBadRequestCode int = 400

/*GetProfileBadRequest Invalid ID supplied

swagger:response getProfileBadRequest
*/
type GetProfileBadRequest struct {
}

// NewGetProfileBadRequest creates GetProfileBadRequest with default headers values
func NewGetProfileBadRequest() *GetProfileBadRequest {
	return &GetProfileBadRequest{}
}

// WriteResponse to the client
func (o *GetProfileBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetProfileNotFoundCode is the HTTP code returned for type GetProfileNotFound
const GetProfileNotFoundCode int = 404

/*GetProfileNotFound Not found

swagger:response getProfileNotFound
*/
type GetProfileNotFound struct {
}

// NewGetProfileNotFound creates GetProfileNotFound with default headers values
func NewGetProfileNotFound() *GetProfileNotFound {
	return &GetProfileNotFound{}
}

// WriteResponse to the client
func (o *GetProfileNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetProfileInternalServerErrorCode is the HTTP code returned for type GetProfileInternalServerError
const GetProfileInternalServerErrorCode int = 500

/*GetProfileInternalServerError Internal Error

swagger:response getProfileInternalServerError
*/
type GetProfileInternalServerError struct {
}

// NewGetProfileInternalServerError creates GetProfileInternalServerError with default headers values
func NewGetProfileInternalServerError() *GetProfileInternalServerError {
	return &GetProfileInternalServerError{}
}

// WriteResponse to the client
func (o *GetProfileInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
