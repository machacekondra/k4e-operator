// Code generated by go-swagger; DO NOT EDIT.

package yggdrasil

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostControlMessageForDeviceReader is a Reader for the PostControlMessageForDevice structure.
type PostControlMessageForDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostControlMessageForDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostControlMessageForDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostControlMessageForDeviceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostControlMessageForDeviceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostControlMessageForDeviceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostControlMessageForDeviceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostControlMessageForDeviceOK creates a PostControlMessageForDeviceOK with default headers values
func NewPostControlMessageForDeviceOK() *PostControlMessageForDeviceOK {
	return &PostControlMessageForDeviceOK{}
}

/*PostControlMessageForDeviceOK handles this case with default header values.

Success
*/
type PostControlMessageForDeviceOK struct {
}

func (o *PostControlMessageForDeviceOK) Error() string {
	return fmt.Sprintf("[POST /control/{device_id}/out][%d] postControlMessageForDeviceOK ", 200)
}

func (o *PostControlMessageForDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostControlMessageForDeviceUnauthorized creates a PostControlMessageForDeviceUnauthorized with default headers values
func NewPostControlMessageForDeviceUnauthorized() *PostControlMessageForDeviceUnauthorized {
	return &PostControlMessageForDeviceUnauthorized{}
}

/*PostControlMessageForDeviceUnauthorized handles this case with default header values.

Unauthorized
*/
type PostControlMessageForDeviceUnauthorized struct {
}

func (o *PostControlMessageForDeviceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /control/{device_id}/out][%d] postControlMessageForDeviceUnauthorized ", 401)
}

func (o *PostControlMessageForDeviceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostControlMessageForDeviceForbidden creates a PostControlMessageForDeviceForbidden with default headers values
func NewPostControlMessageForDeviceForbidden() *PostControlMessageForDeviceForbidden {
	return &PostControlMessageForDeviceForbidden{}
}

/*PostControlMessageForDeviceForbidden handles this case with default header values.

Forbidden
*/
type PostControlMessageForDeviceForbidden struct {
}

func (o *PostControlMessageForDeviceForbidden) Error() string {
	return fmt.Sprintf("[POST /control/{device_id}/out][%d] postControlMessageForDeviceForbidden ", 403)
}

func (o *PostControlMessageForDeviceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostControlMessageForDeviceNotFound creates a PostControlMessageForDeviceNotFound with default headers values
func NewPostControlMessageForDeviceNotFound() *PostControlMessageForDeviceNotFound {
	return &PostControlMessageForDeviceNotFound{}
}

/*PostControlMessageForDeviceNotFound handles this case with default header values.

Error
*/
type PostControlMessageForDeviceNotFound struct {
}

func (o *PostControlMessageForDeviceNotFound) Error() string {
	return fmt.Sprintf("[POST /control/{device_id}/out][%d] postControlMessageForDeviceNotFound ", 404)
}

func (o *PostControlMessageForDeviceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostControlMessageForDeviceInternalServerError creates a PostControlMessageForDeviceInternalServerError with default headers values
func NewPostControlMessageForDeviceInternalServerError() *PostControlMessageForDeviceInternalServerError {
	return &PostControlMessageForDeviceInternalServerError{}
}

/*PostControlMessageForDeviceInternalServerError handles this case with default header values.

Error
*/
type PostControlMessageForDeviceInternalServerError struct {
}

func (o *PostControlMessageForDeviceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /control/{device_id}/out][%d] postControlMessageForDeviceInternalServerError ", 500)
}

func (o *PostControlMessageForDeviceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
