// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V3ADRAckLimitExponent v3 a d r ack limit exponent
//
// swagger:model v3ADRAckLimitExponent
type V3ADRAckLimitExponent string

const (

	// V3ADRAckLimitExponentADRACKLIMIT1 captures enum value "ADR_ACK_LIMIT_1"
	V3ADRAckLimitExponentADRACKLIMIT1 V3ADRAckLimitExponent = "ADR_ACK_LIMIT_1"

	// V3ADRAckLimitExponentADRACKLIMIT2 captures enum value "ADR_ACK_LIMIT_2"
	V3ADRAckLimitExponentADRACKLIMIT2 V3ADRAckLimitExponent = "ADR_ACK_LIMIT_2"

	// V3ADRAckLimitExponentADRACKLIMIT4 captures enum value "ADR_ACK_LIMIT_4"
	V3ADRAckLimitExponentADRACKLIMIT4 V3ADRAckLimitExponent = "ADR_ACK_LIMIT_4"

	// V3ADRAckLimitExponentADRACKLIMIT8 captures enum value "ADR_ACK_LIMIT_8"
	V3ADRAckLimitExponentADRACKLIMIT8 V3ADRAckLimitExponent = "ADR_ACK_LIMIT_8"

	// V3ADRAckLimitExponentADRACKLIMIT16 captures enum value "ADR_ACK_LIMIT_16"
	V3ADRAckLimitExponentADRACKLIMIT16 V3ADRAckLimitExponent = "ADR_ACK_LIMIT_16"

	// V3ADRAckLimitExponentADRACKLIMIT32 captures enum value "ADR_ACK_LIMIT_32"
	V3ADRAckLimitExponentADRACKLIMIT32 V3ADRAckLimitExponent = "ADR_ACK_LIMIT_32"

	// V3ADRAckLimitExponentADRACKLIMIT64 captures enum value "ADR_ACK_LIMIT_64"
	V3ADRAckLimitExponentADRACKLIMIT64 V3ADRAckLimitExponent = "ADR_ACK_LIMIT_64"

	// V3ADRAckLimitExponentADRACKLIMIT128 captures enum value "ADR_ACK_LIMIT_128"
	V3ADRAckLimitExponentADRACKLIMIT128 V3ADRAckLimitExponent = "ADR_ACK_LIMIT_128"

	// V3ADRAckLimitExponentADRACKLIMIT256 captures enum value "ADR_ACK_LIMIT_256"
	V3ADRAckLimitExponentADRACKLIMIT256 V3ADRAckLimitExponent = "ADR_ACK_LIMIT_256"

	// V3ADRAckLimitExponentADRACKLIMIT512 captures enum value "ADR_ACK_LIMIT_512"
	V3ADRAckLimitExponentADRACKLIMIT512 V3ADRAckLimitExponent = "ADR_ACK_LIMIT_512"

	// V3ADRAckLimitExponentADRACKLIMIT1024 captures enum value "ADR_ACK_LIMIT_1024"
	V3ADRAckLimitExponentADRACKLIMIT1024 V3ADRAckLimitExponent = "ADR_ACK_LIMIT_1024"

	// V3ADRAckLimitExponentADRACKLIMIT2048 captures enum value "ADR_ACK_LIMIT_2048"
	V3ADRAckLimitExponentADRACKLIMIT2048 V3ADRAckLimitExponent = "ADR_ACK_LIMIT_2048"

	// V3ADRAckLimitExponentADRACKLIMIT4096 captures enum value "ADR_ACK_LIMIT_4096"
	V3ADRAckLimitExponentADRACKLIMIT4096 V3ADRAckLimitExponent = "ADR_ACK_LIMIT_4096"

	// V3ADRAckLimitExponentADRACKLIMIT8192 captures enum value "ADR_ACK_LIMIT_8192"
	V3ADRAckLimitExponentADRACKLIMIT8192 V3ADRAckLimitExponent = "ADR_ACK_LIMIT_8192"

	// V3ADRAckLimitExponentADRACKLIMIT16384 captures enum value "ADR_ACK_LIMIT_16384"
	V3ADRAckLimitExponentADRACKLIMIT16384 V3ADRAckLimitExponent = "ADR_ACK_LIMIT_16384"

	// V3ADRAckLimitExponentADRACKLIMIT32768 captures enum value "ADR_ACK_LIMIT_32768"
	V3ADRAckLimitExponentADRACKLIMIT32768 V3ADRAckLimitExponent = "ADR_ACK_LIMIT_32768"
)

// for schema
var v3ADRAckLimitExponentEnum []interface{}

func init() {
	var res []V3ADRAckLimitExponent
	if err := json.Unmarshal([]byte(`["ADR_ACK_LIMIT_1","ADR_ACK_LIMIT_2","ADR_ACK_LIMIT_4","ADR_ACK_LIMIT_8","ADR_ACK_LIMIT_16","ADR_ACK_LIMIT_32","ADR_ACK_LIMIT_64","ADR_ACK_LIMIT_128","ADR_ACK_LIMIT_256","ADR_ACK_LIMIT_512","ADR_ACK_LIMIT_1024","ADR_ACK_LIMIT_2048","ADR_ACK_LIMIT_4096","ADR_ACK_LIMIT_8192","ADR_ACK_LIMIT_16384","ADR_ACK_LIMIT_32768"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v3ADRAckLimitExponentEnum = append(v3ADRAckLimitExponentEnum, v)
	}
}

func (m V3ADRAckLimitExponent) validateV3ADRAckLimitExponentEnum(path, location string, value V3ADRAckLimitExponent) error {
	if err := validate.Enum(path, location, value, v3ADRAckLimitExponentEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v3 a d r ack limit exponent
func (m V3ADRAckLimitExponent) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV3ADRAckLimitExponentEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
