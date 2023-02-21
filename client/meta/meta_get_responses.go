//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2023 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

// Code generated by go-swagger; DO NOT EDIT.

package meta

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/weaviate/weaviate/entities/models"
)

// MetaGetReader is a Reader for the MetaGet structure.
type MetaGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetaGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMetaGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewMetaGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMetaGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewMetaGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMetaGetOK creates a MetaGetOK with default headers values
func NewMetaGetOK() *MetaGetOK {
	return &MetaGetOK{}
}

/*
MetaGetOK describes a response with status code 200, with default header values.

Successful response.
*/
type MetaGetOK struct {
	Payload *models.Meta
}

// IsSuccess returns true when this meta get o k response has a 2xx status code
func (o *MetaGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this meta get o k response has a 3xx status code
func (o *MetaGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this meta get o k response has a 4xx status code
func (o *MetaGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this meta get o k response has a 5xx status code
func (o *MetaGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this meta get o k response a status code equal to that given
func (o *MetaGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the meta get o k response
func (o *MetaGetOK) Code() int {
	return 200
}

func (o *MetaGetOK) Error() string {
	return fmt.Sprintf("[GET /meta][%d] metaGetOK  %+v", 200, o.Payload)
}

func (o *MetaGetOK) String() string {
	return fmt.Sprintf("[GET /meta][%d] metaGetOK  %+v", 200, o.Payload)
}

func (o *MetaGetOK) GetPayload() *models.Meta {
	return o.Payload
}

func (o *MetaGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Meta)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetaGetUnauthorized creates a MetaGetUnauthorized with default headers values
func NewMetaGetUnauthorized() *MetaGetUnauthorized {
	return &MetaGetUnauthorized{}
}

/*
MetaGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized or invalid credentials.
*/
type MetaGetUnauthorized struct {
}

// IsSuccess returns true when this meta get unauthorized response has a 2xx status code
func (o *MetaGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this meta get unauthorized response has a 3xx status code
func (o *MetaGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this meta get unauthorized response has a 4xx status code
func (o *MetaGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this meta get unauthorized response has a 5xx status code
func (o *MetaGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this meta get unauthorized response a status code equal to that given
func (o *MetaGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the meta get unauthorized response
func (o *MetaGetUnauthorized) Code() int {
	return 401
}

func (o *MetaGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /meta][%d] metaGetUnauthorized ", 401)
}

func (o *MetaGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /meta][%d] metaGetUnauthorized ", 401)
}

func (o *MetaGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMetaGetForbidden creates a MetaGetForbidden with default headers values
func NewMetaGetForbidden() *MetaGetForbidden {
	return &MetaGetForbidden{}
}

/*
MetaGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type MetaGetForbidden struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this meta get forbidden response has a 2xx status code
func (o *MetaGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this meta get forbidden response has a 3xx status code
func (o *MetaGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this meta get forbidden response has a 4xx status code
func (o *MetaGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this meta get forbidden response has a 5xx status code
func (o *MetaGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this meta get forbidden response a status code equal to that given
func (o *MetaGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the meta get forbidden response
func (o *MetaGetForbidden) Code() int {
	return 403
}

func (o *MetaGetForbidden) Error() string {
	return fmt.Sprintf("[GET /meta][%d] metaGetForbidden  %+v", 403, o.Payload)
}

func (o *MetaGetForbidden) String() string {
	return fmt.Sprintf("[GET /meta][%d] metaGetForbidden  %+v", 403, o.Payload)
}

func (o *MetaGetForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MetaGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetaGetInternalServerError creates a MetaGetInternalServerError with default headers values
func NewMetaGetInternalServerError() *MetaGetInternalServerError {
	return &MetaGetInternalServerError{}
}

/*
MetaGetInternalServerError describes a response with status code 500, with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type MetaGetInternalServerError struct {
	Payload *models.ErrorResponse
}

// IsSuccess returns true when this meta get internal server error response has a 2xx status code
func (o *MetaGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this meta get internal server error response has a 3xx status code
func (o *MetaGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this meta get internal server error response has a 4xx status code
func (o *MetaGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this meta get internal server error response has a 5xx status code
func (o *MetaGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this meta get internal server error response a status code equal to that given
func (o *MetaGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the meta get internal server error response
func (o *MetaGetInternalServerError) Code() int {
	return 500
}

func (o *MetaGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /meta][%d] metaGetInternalServerError  %+v", 500, o.Payload)
}

func (o *MetaGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /meta][%d] metaGetInternalServerError  %+v", 500, o.Payload)
}

func (o *MetaGetInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *MetaGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
