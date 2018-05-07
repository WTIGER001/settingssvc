// Code generated by go-swagger; DO NOT EDIT.

package preferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"models"
)

// GetProfilesOKCode is the HTTP code returned for type GetProfilesOK
const GetProfilesOKCode int = 200

/*GetProfilesOK successful operation

swagger:response getProfilesOK
*/
type GetProfilesOK struct {

	/*
	  In: Body
	*/
	Payload models.ProfileArray `json:"body,omitempty"`
}

// NewGetProfilesOK creates GetProfilesOK with default headers values
func NewGetProfilesOK() *GetProfilesOK {
	return &GetProfilesOK{}
}

// WithPayload adds the payload to the get profiles o k response
func (o *GetProfilesOK) WithPayload(payload models.ProfileArray) *GetProfilesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get profiles o k response
func (o *GetProfilesOK) SetPayload(payload models.ProfileArray) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProfilesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.ProfileArray, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetProfilesBadRequestCode is the HTTP code returned for type GetProfilesBadRequest
const GetProfilesBadRequestCode int = 400

/*GetProfilesBadRequest Invalid ID supplied

swagger:response getProfilesBadRequest
*/
type GetProfilesBadRequest struct {
}

// NewGetProfilesBadRequest creates GetProfilesBadRequest with default headers values
func NewGetProfilesBadRequest() *GetProfilesBadRequest {
	return &GetProfilesBadRequest{}
}

// WriteResponse to the client
func (o *GetProfilesBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// GetProfilesNotFoundCode is the HTTP code returned for type GetProfilesNotFound
const GetProfilesNotFoundCode int = 404

/*GetProfilesNotFound Not found

swagger:response getProfilesNotFound
*/
type GetProfilesNotFound struct {
}

// NewGetProfilesNotFound creates GetProfilesNotFound with default headers values
func NewGetProfilesNotFound() *GetProfilesNotFound {
	return &GetProfilesNotFound{}
}

// WriteResponse to the client
func (o *GetProfilesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetProfilesInternalServerErrorCode is the HTTP code returned for type GetProfilesInternalServerError
const GetProfilesInternalServerErrorCode int = 500

/*GetProfilesInternalServerError Internal Error

swagger:response getProfilesInternalServerError
*/
type GetProfilesInternalServerError struct {
}

// NewGetProfilesInternalServerError creates GetProfilesInternalServerError with default headers values
func NewGetProfilesInternalServerError() *GetProfilesInternalServerError {
	return &GetProfilesInternalServerError{}
}

// WriteResponse to the client
func (o *GetProfilesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}