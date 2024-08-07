// Code generated by go-swagger; DO NOT EDIT.

package admin_service

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

// NewAdminServiceGetBalanceParams creates a new AdminServiceGetBalanceParams object
// with the default values initialized.
func NewAdminServiceGetBalanceParams() *AdminServiceGetBalanceParams {

	return &AdminServiceGetBalanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAdminServiceGetBalanceParamsWithTimeout creates a new AdminServiceGetBalanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAdminServiceGetBalanceParamsWithTimeout(timeout time.Duration) *AdminServiceGetBalanceParams {

	return &AdminServiceGetBalanceParams{

		timeout: timeout,
	}
}

// NewAdminServiceGetBalanceParamsWithContext creates a new AdminServiceGetBalanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewAdminServiceGetBalanceParamsWithContext(ctx context.Context) *AdminServiceGetBalanceParams {

	return &AdminServiceGetBalanceParams{

		Context: ctx,
	}
}

// NewAdminServiceGetBalanceParamsWithHTTPClient creates a new AdminServiceGetBalanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAdminServiceGetBalanceParamsWithHTTPClient(client *http.Client) *AdminServiceGetBalanceParams {

	return &AdminServiceGetBalanceParams{
		HTTPClient: client,
	}
}

/*AdminServiceGetBalanceParams contains all the parameters to send to the API endpoint
for the admin service get balance operation typically these are written to a http.Request
*/
type AdminServiceGetBalanceParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the admin service get balance params
func (o *AdminServiceGetBalanceParams) WithTimeout(timeout time.Duration) *AdminServiceGetBalanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin service get balance params
func (o *AdminServiceGetBalanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin service get balance params
func (o *AdminServiceGetBalanceParams) WithContext(ctx context.Context) *AdminServiceGetBalanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin service get balance params
func (o *AdminServiceGetBalanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin service get balance params
func (o *AdminServiceGetBalanceParams) WithHTTPClient(client *http.Client) *AdminServiceGetBalanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin service get balance params
func (o *AdminServiceGetBalanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AdminServiceGetBalanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
