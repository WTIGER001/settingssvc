// Code generated by go-swagger; DO NOT EDIT.

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AddOwnerTypeOKCode is the HTTP code returned for type AddOwnerTypeOK
const AddOwnerTypeOKCode int = 200

/*AddOwnerTypeOK success

swagger:response addOwnerTypeOK
*/
type AddOwnerTypeOK struct {
}

// NewAddOwnerTypeOK creates AddOwnerTypeOK with default headers values
func NewAddOwnerTypeOK() *AddOwnerTypeOK {
	return &AddOwnerTypeOK{}
}

// WriteResponse to the client
func (o *AddOwnerTypeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// AddOwnerTypeMethodNotAllowedCode is the HTTP code returned for type AddOwnerTypeMethodNotAllowed
const AddOwnerTypeMethodNotAllowedCode int = 405

/*AddOwnerTypeMethodNotAllowed Invalid input

swagger:response addOwnerTypeMethodNotAllowed
*/
type AddOwnerTypeMethodNotAllowed struct {
}

// NewAddOwnerTypeMethodNotAllowed creates AddOwnerTypeMethodNotAllowed with default headers values
func NewAddOwnerTypeMethodNotAllowed() *AddOwnerTypeMethodNotAllowed {
	return &AddOwnerTypeMethodNotAllowed{}
}

// WriteResponse to the client
func (o *AddOwnerTypeMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(405)
}

// AddOwnerTypeInternalServerErrorCode is the HTTP code returned for type AddOwnerTypeInternalServerError
const AddOwnerTypeInternalServerErrorCode int = 500

/*AddOwnerTypeInternalServerError Internal Error

swagger:response addOwnerTypeInternalServerError
*/
type AddOwnerTypeInternalServerError struct {
}

// NewAddOwnerTypeInternalServerError creates AddOwnerTypeInternalServerError with default headers values
func NewAddOwnerTypeInternalServerError() *AddOwnerTypeInternalServerError {
	return &AddOwnerTypeInternalServerError{}
}

// WriteResponse to the client
func (o *AddOwnerTypeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
