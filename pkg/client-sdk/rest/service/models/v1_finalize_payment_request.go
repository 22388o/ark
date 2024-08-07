// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V1FinalizePaymentRequest v1 finalize payment request
// swagger:model v1FinalizePaymentRequest
type V1FinalizePaymentRequest struct {

	// Forfeit txs signed by the user.
	SignedForfeitTxs []string `json:"signedForfeitTxs"`
}

// Validate validates this v1 finalize payment request
func (m *V1FinalizePaymentRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1FinalizePaymentRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1FinalizePaymentRequest) UnmarshalBinary(b []byte) error {
	var res V1FinalizePaymentRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
