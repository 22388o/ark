// Code generated by go-swagger; DO NOT EDIT.

package ark_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ark-network/ark-sdk/rest/service/models"
)

// ArkServiceGetInfoReader is a Reader for the ArkServiceGetInfo structure.
type ArkServiceGetInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArkServiceGetInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewArkServiceGetInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewArkServiceGetInfoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewArkServiceGetInfoOK creates a ArkServiceGetInfoOK with default headers values
func NewArkServiceGetInfoOK() *ArkServiceGetInfoOK {
	return &ArkServiceGetInfoOK{}
}

/*ArkServiceGetInfoOK handles this case with default header values.

A successful response.
*/
type ArkServiceGetInfoOK struct {
	Payload *models.V1GetInfoResponse
}

func (o *ArkServiceGetInfoOK) Error() string {
	return fmt.Sprintf("[GET /v1/info][%d] arkServiceGetInfoOK  %+v", 200, o.Payload)
}

func (o *ArkServiceGetInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetInfoResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArkServiceGetInfoDefault creates a ArkServiceGetInfoDefault with default headers values
func NewArkServiceGetInfoDefault(code int) *ArkServiceGetInfoDefault {
	return &ArkServiceGetInfoDefault{
		_statusCode: code,
	}
}

/*ArkServiceGetInfoDefault handles this case with default header values.

An unexpected error response.
*/
type ArkServiceGetInfoDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// Code gets the status code for the ark service get info default response
func (o *ArkServiceGetInfoDefault) Code() int {
	return o._statusCode
}

func (o *ArkServiceGetInfoDefault) Error() string {
	return fmt.Sprintf("[GET /v1/info][%d] ArkService_GetInfo default  %+v", o._statusCode, o.Payload)
}

func (o *ArkServiceGetInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
