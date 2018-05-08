// Code generated by go-swagger; DO NOT EDIT.

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// DeleteCategoryOKCode is the HTTP code returned for type DeleteCategoryOK
const DeleteCategoryOKCode int = 200

/*DeleteCategoryOK Success

swagger:response deleteCategoryOK
*/
type DeleteCategoryOK struct {
}

// NewDeleteCategoryOK creates DeleteCategoryOK with default headers values
func NewDeleteCategoryOK() *DeleteCategoryOK {
	return &DeleteCategoryOK{}
}

// WriteResponse to the client
func (o *DeleteCategoryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// DeleteCategoryBadRequestCode is the HTTP code returned for type DeleteCategoryBadRequest
const DeleteCategoryBadRequestCode int = 400

/*DeleteCategoryBadRequest Invalid ID supplied

swagger:response deleteCategoryBadRequest
*/
type DeleteCategoryBadRequest struct {
}

// NewDeleteCategoryBadRequest creates DeleteCategoryBadRequest with default headers values
func NewDeleteCategoryBadRequest() *DeleteCategoryBadRequest {
	return &DeleteCategoryBadRequest{}
}

// WriteResponse to the client
func (o *DeleteCategoryBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// DeleteCategoryNotFoundCode is the HTTP code returned for type DeleteCategoryNotFound
const DeleteCategoryNotFoundCode int = 404

/*DeleteCategoryNotFound Category not found

swagger:response deleteCategoryNotFound
*/
type DeleteCategoryNotFound struct {
}

// NewDeleteCategoryNotFound creates DeleteCategoryNotFound with default headers values
func NewDeleteCategoryNotFound() *DeleteCategoryNotFound {
	return &DeleteCategoryNotFound{}
}

// WriteResponse to the client
func (o *DeleteCategoryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DeleteCategoryInternalServerErrorCode is the HTTP code returned for type DeleteCategoryInternalServerError
const DeleteCategoryInternalServerErrorCode int = 500

/*DeleteCategoryInternalServerError Internal Error

swagger:response deleteCategoryInternalServerError
*/
type DeleteCategoryInternalServerError struct {
}

// NewDeleteCategoryInternalServerError creates DeleteCategoryInternalServerError with default headers values
func NewDeleteCategoryInternalServerError() *DeleteCategoryInternalServerError {
	return &DeleteCategoryInternalServerError{}
}

// WriteResponse to the client
func (o *DeleteCategoryInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}