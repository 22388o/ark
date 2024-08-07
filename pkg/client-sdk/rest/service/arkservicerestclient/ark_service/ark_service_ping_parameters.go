// Code generated by go-swagger; DO NOT EDIT.

package ark_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewArkServicePingParams creates a new ArkServicePingParams object
// with the default values initialized.
func NewArkServicePingParams() *ArkServicePingParams {
	var ()
	return &ArkServicePingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewArkServicePingParamsWithTimeout creates a new ArkServicePingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewArkServicePingParamsWithTimeout(timeout time.Duration) *ArkServicePingParams {
	var ()
	return &ArkServicePingParams{

		timeout: timeout,
	}
}

// NewArkServicePingParamsWithContext creates a new ArkServicePingParams object
// with the default values initialized, and the ability to set a context for a request
func NewArkServicePingParamsWithContext(ctx context.Context) *ArkServicePingParams {
	var ()
	return &ArkServicePingParams{

		Context: ctx,
	}
}

// NewArkServicePingParamsWithHTTPClient creates a new ArkServicePingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewArkServicePingParamsWithHTTPClient(client *http.Client) *ArkServicePingParams {
	var ()
	return &ArkServicePingParams{
		HTTPClient: client,
	}
}

/*ArkServicePingParams contains all the parameters to send to the API endpoint
for the ark service ping operation typically these are written to a http.Request
*/
type ArkServicePingParams struct {

	/*PaymentID*/
	PaymentID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ark service ping params
func (o *ArkServicePingParams) WithTimeout(timeout time.Duration) *ArkServicePingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ark service ping params
func (o *ArkServicePingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ark service ping params
func (o *ArkServicePingParams) WithContext(ctx context.Context) *ArkServicePingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ark service ping params
func (o *ArkServicePingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ark service ping params
func (o *ArkServicePingParams) WithHTTPClient(client *http.Client) *ArkServicePingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ark service ping params
func (o *ArkServicePingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPaymentID adds the paymentID to the ark service ping params
func (o *ArkServicePingParams) WithPaymentID(paymentID string) *ArkServicePingParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the ark service ping params
func (o *ArkServicePingParams) SetPaymentID(paymentID string) {
	o.PaymentID = paymentID
}

// WriteToRequest writes these params to a swagger request
func (o *ArkServicePingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
