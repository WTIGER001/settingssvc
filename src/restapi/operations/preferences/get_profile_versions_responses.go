// Code generated by go-swagger; DO NOT EDIT.

package preferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"models"
)

// GetProfileVersionsOKCode is the HTTP code returned for type GetProfileVersionsOK
const GetProfileVersionsOKCode int = 200

/*GetProfileVersionsOK successful operation

swagger:response getProfileVersionsOK
*/
type GetProfileVersionsOK struct {

	/*
	  In: Body
	*/
	Payload *models.ProfileVersions `json:"body,omitempty"`
}

// NewGetProfileVersionsOK creates GetProfileVersionsOK with default headers values
func NewGetProfileVersionsOK() *GetProfileVersionsOK {
	return &GetProfileVersionsOK{}
}

// WithPayload adds the payload to the get profile versions o k response
func (o *GetProfileVersionsOK) WithPayload(payload *models.ProfileVersions) *GetProfileVersionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get profile versions o k response
func (o *GetProfileVersionsOK) SetPayload(payload *models.ProfileVersions) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProfileVersionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProfileVersionsBadRequestCode is the HTTP code returned for type GetProfileVersionsBadRequest
const GetProfileVersionsBadRequestCode int = 400

/*GetProfileVersionsBadRequest Invalid ID supplied

swagger:response getProfileVersionsBadRequest
*/
type GetProfileVersionsBadRequest struct {
}

// NewGetProfileVersionsBadRequest creates GetProfileVersionsBadRequest with default headers values
func NewGetProfileVersionsBadRequest() *GetProfileVersionsBadRequest {
	return &GetProfileVersionsBadRequest{}
}

// WriteResponse to the client
func (o *GetProfileVersionsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetProfileVersionsNotFoundCode is the HTTP code returned for type GetProfileVersionsNotFound
const GetProfileVersionsNotFoundCode int = 404

/*GetProfileVersionsNotFound Not found

swagger:response getProfileVersionsNotFound
*/
type GetProfileVersionsNotFound struct {
}

// NewGetProfileVersionsNotFound creates GetProfileVersionsNotFound with default headers values
func NewGetProfileVersionsNotFound() *GetProfileVersionsNotFound {
	return &GetProfileVersionsNotFound{}
}

// WriteResponse to the client
func (o *GetProfileVersionsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetProfileVersionsInternalServerErrorCode is the HTTP code returned for type GetProfileVersionsInternalServerError
const GetProfileVersionsInternalServerErrorCode int = 500

/*GetProfileVersionsInternalServerError Internal Error

swagger:response getProfileVersionsInternalServerError
*/
type GetProfileVersionsInternalServerError struct {
}

// NewGetProfileVersionsInternalServerError creates GetProfileVersionsInternalServerError with default headers values
func NewGetProfileVersionsInternalServerError() *GetProfileVersionsInternalServerError {
	return &GetProfileVersionsInternalServerError{}
}

// WriteResponse to the client
func (o *GetProfileVersionsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
