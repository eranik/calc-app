// Code generated by go-swagger; DO NOT EDIT.

package math

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// SubtractOKCode is the HTTP code returned for type SubtractOK
const SubtractOKCode int = 200

/*SubtractOK subtract of two integers.

swagger:response subtractOK
*/
type SubtractOK struct {

	/*
	  In: Body
	*/
	Payload *SubtractOKBody `json:"body,omitempty"`
}

// NewSubtractOK creates SubtractOK with default headers values
func NewSubtractOK() *SubtractOK {

	return &SubtractOK{}
}

// WithPayload adds the payload to the subtract o k response
func (o *SubtractOK) WithPayload(payload *SubtractOKBody) *SubtractOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the subtract o k response
func (o *SubtractOK) SetPayload(payload *SubtractOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SubtractOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SubtractUnauthorizedCode is the HTTP code returned for type SubtractUnauthorized
const SubtractUnauthorizedCode int = 401

/*SubtractUnauthorized Unauthorized

swagger:response subtractUnauthorized
*/
type SubtractUnauthorized struct {
}

// NewSubtractUnauthorized creates SubtractUnauthorized with default headers values
func NewSubtractUnauthorized() *SubtractUnauthorized {

	return &SubtractUnauthorized{}
}

// WriteResponse to the client
func (o *SubtractUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// SubtractForbiddenCode is the HTTP code returned for type SubtractForbidden
const SubtractForbiddenCode int = 403

/*SubtractForbidden Forbidden

swagger:response subtractForbidden
*/
type SubtractForbidden struct {
}

// NewSubtractForbidden creates SubtractForbidden with default headers values
func NewSubtractForbidden() *SubtractForbidden {

	return &SubtractForbidden{}
}

// WriteResponse to the client
func (o *SubtractForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// SubtractInternalServerErrorCode is the HTTP code returned for type SubtractInternalServerError
const SubtractInternalServerErrorCode int = 500

/*SubtractInternalServerError Unexpected error

swagger:response subtractInternalServerError
*/
type SubtractInternalServerError struct {
}

// NewSubtractInternalServerError creates SubtractInternalServerError with default headers values
func NewSubtractInternalServerError() *SubtractInternalServerError {

	return &SubtractInternalServerError{}
}

// WriteResponse to the client
func (o *SubtractInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
