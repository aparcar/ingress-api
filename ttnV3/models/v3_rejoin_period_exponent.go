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

// V3RejoinPeriodExponent v3 rejoin period exponent
//
// swagger:model v3RejoinPeriodExponent
type V3RejoinPeriodExponent string

const (

	// V3RejoinPeriodExponentREJOINPERIOD0 captures enum value "REJOIN_PERIOD_0"
	V3RejoinPeriodExponentREJOINPERIOD0 V3RejoinPeriodExponent = "REJOIN_PERIOD_0"

	// V3RejoinPeriodExponentREJOINPERIOD1 captures enum value "REJOIN_PERIOD_1"
	V3RejoinPeriodExponentREJOINPERIOD1 V3RejoinPeriodExponent = "REJOIN_PERIOD_1"

	// V3RejoinPeriodExponentREJOINPERIOD2 captures enum value "REJOIN_PERIOD_2"
	V3RejoinPeriodExponentREJOINPERIOD2 V3RejoinPeriodExponent = "REJOIN_PERIOD_2"

	// V3RejoinPeriodExponentREJOINPERIOD3 captures enum value "REJOIN_PERIOD_3"
	V3RejoinPeriodExponentREJOINPERIOD3 V3RejoinPeriodExponent = "REJOIN_PERIOD_3"

	// V3RejoinPeriodExponentREJOINPERIOD4 captures enum value "REJOIN_PERIOD_4"
	V3RejoinPeriodExponentREJOINPERIOD4 V3RejoinPeriodExponent = "REJOIN_PERIOD_4"

	// V3RejoinPeriodExponentREJOINPERIOD5 captures enum value "REJOIN_PERIOD_5"
	V3RejoinPeriodExponentREJOINPERIOD5 V3RejoinPeriodExponent = "REJOIN_PERIOD_5"

	// V3RejoinPeriodExponentREJOINPERIOD6 captures enum value "REJOIN_PERIOD_6"
	V3RejoinPeriodExponentREJOINPERIOD6 V3RejoinPeriodExponent = "REJOIN_PERIOD_6"

	// V3RejoinPeriodExponentREJOINPERIOD7 captures enum value "REJOIN_PERIOD_7"
	V3RejoinPeriodExponentREJOINPERIOD7 V3RejoinPeriodExponent = "REJOIN_PERIOD_7"
)

// for schema
var v3RejoinPeriodExponentEnum []interface{}

func init() {
	var res []V3RejoinPeriodExponent
	if err := json.Unmarshal([]byte(`["REJOIN_PERIOD_0","REJOIN_PERIOD_1","REJOIN_PERIOD_2","REJOIN_PERIOD_3","REJOIN_PERIOD_4","REJOIN_PERIOD_5","REJOIN_PERIOD_6","REJOIN_PERIOD_7"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v3RejoinPeriodExponentEnum = append(v3RejoinPeriodExponentEnum, v)
	}
}

func (m V3RejoinPeriodExponent) validateV3RejoinPeriodExponentEnum(path, location string, value V3RejoinPeriodExponent) error {
	if err := validate.Enum(path, location, value, v3RejoinPeriodExponentEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v3 rejoin period exponent
func (m V3RejoinPeriodExponent) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV3RejoinPeriodExponentEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
