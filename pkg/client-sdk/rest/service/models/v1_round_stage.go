// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// V1RoundStage v1 round stage
// swagger:model v1RoundStage
type V1RoundStage string

const (

	// V1RoundStageROUNDSTAGEUNSPECIFIED captures enum value "ROUND_STAGE_UNSPECIFIED"
	V1RoundStageROUNDSTAGEUNSPECIFIED V1RoundStage = "ROUND_STAGE_UNSPECIFIED"

	// V1RoundStageROUNDSTAGEREGISTRATION captures enum value "ROUND_STAGE_REGISTRATION"
	V1RoundStageROUNDSTAGEREGISTRATION V1RoundStage = "ROUND_STAGE_REGISTRATION"

	// V1RoundStageROUNDSTAGEFINALIZATION captures enum value "ROUND_STAGE_FINALIZATION"
	V1RoundStageROUNDSTAGEFINALIZATION V1RoundStage = "ROUND_STAGE_FINALIZATION"

	// V1RoundStageROUNDSTAGEFINALIZED captures enum value "ROUND_STAGE_FINALIZED"
	V1RoundStageROUNDSTAGEFINALIZED V1RoundStage = "ROUND_STAGE_FINALIZED"

	// V1RoundStageROUNDSTAGEFAILED captures enum value "ROUND_STAGE_FAILED"
	V1RoundStageROUNDSTAGEFAILED V1RoundStage = "ROUND_STAGE_FAILED"
)

// for schema
var v1RoundStageEnum []interface{}

func init() {
	var res []V1RoundStage
	if err := json.Unmarshal([]byte(`["ROUND_STAGE_UNSPECIFIED","ROUND_STAGE_REGISTRATION","ROUND_STAGE_FINALIZATION","ROUND_STAGE_FINALIZED","ROUND_STAGE_FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1RoundStageEnum = append(v1RoundStageEnum, v)
	}
}

func (m V1RoundStage) validateV1RoundStageEnum(path, location string, value V1RoundStage) error {
	if err := validate.Enum(path, location, value, v1RoundStageEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 round stage
func (m V1RoundStage) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1RoundStageEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
