// Code generated by go-swagger; DO NOT EDIT.

package math

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SubtractReader is a Reader for the Subtract structure.
type SubtractReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubtractReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubtractOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSubtractUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSubtractForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSubtractInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSubtractOK creates a SubtractOK with default headers values
func NewSubtractOK() *SubtractOK {
	return &SubtractOK{}
}

/*SubtractOK handles this case with default header values.

subtract of two integers.
*/
type SubtractOK struct {
	Payload *SubtractOKBody
}

func (o *SubtractOK) Error() string {
	return fmt.Sprintf("[GET /subtract][%d] subtractOK  %+v", 200, o.Payload)
}

func (o *SubtractOK) GetPayload() *SubtractOKBody {
	return o.Payload
}

func (o *SubtractOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SubtractOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSubtractUnauthorized creates a SubtractUnauthorized with default headers values
func NewSubtractUnauthorized() *SubtractUnauthorized {
	return &SubtractUnauthorized{}
}

/*SubtractUnauthorized handles this case with default header values.

Unauthorized
*/
type SubtractUnauthorized struct {
}

func (o *SubtractUnauthorized) Error() string {
	return fmt.Sprintf("[GET /subtract][%d] subtractUnauthorized ", 401)
}

func (o *SubtractUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubtractForbidden creates a SubtractForbidden with default headers values
func NewSubtractForbidden() *SubtractForbidden {
	return &SubtractForbidden{}
}

/*SubtractForbidden handles this case with default header values.

Forbidden
*/
type SubtractForbidden struct {
}

func (o *SubtractForbidden) Error() string {
	return fmt.Sprintf("[GET /subtract][%d] subtractForbidden ", 403)
}

func (o *SubtractForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSubtractInternalServerError creates a SubtractInternalServerError with default headers values
func NewSubtractInternalServerError() *SubtractInternalServerError {
	return &SubtractInternalServerError{}
}

/*SubtractInternalServerError handles this case with default header values.

Unexpected error
*/
type SubtractInternalServerError struct {
}

func (o *SubtractInternalServerError) Error() string {
	return fmt.Sprintf("[GET /subtract][%d] subtractInternalServerError ", 500)
}

func (o *SubtractInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*SubtractOKBody subtract o k body
swagger:model SubtractOKBody
*/
type SubtractOKBody struct {

	// res
	Res int64 `json:"res"`
}

// Validate validates this subtract o k body
func (o *SubtractOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SubtractOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SubtractOKBody) UnmarshalBinary(b []byte) error {
	var res SubtractOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
