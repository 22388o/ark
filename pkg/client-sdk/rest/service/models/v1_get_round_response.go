// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1GetRoundResponse v1 get round response
// swagger:model v1GetRoundResponse
type V1GetRoundResponse struct {

	// round
	Round *V1Round `json:"round,omitempty"`
}

// Validate validates this v1 get round response
func (m *V1GetRoundResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRound(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1GetRoundResponse) validateRound(formats strfmt.Registry) error {

	if swag.IsZero(m.Round) { // not required
		return nil
	}

	if m.Round != nil {
		if err := m.Round.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("round")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1GetRoundResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1GetRoundResponse) UnmarshalBinary(b []byte) error {
	var res V1GetRoundResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}