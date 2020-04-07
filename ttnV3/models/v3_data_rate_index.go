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

// V3DataRateIndex v3 data rate index
//
// swagger:model v3DataRateIndex
type V3DataRateIndex string

const (

	// V3DataRateIndexDATARATE0 captures enum value "DATA_RATE_0"
	V3DataRateIndexDATARATE0 V3DataRateIndex = "DATA_RATE_0"

	// V3DataRateIndexDATARATE1 captures enum value "DATA_RATE_1"
	V3DataRateIndexDATARATE1 V3DataRateIndex = "DATA_RATE_1"

	// V3DataRateIndexDATARATE2 captures enum value "DATA_RATE_2"
	V3DataRateIndexDATARATE2 V3DataRateIndex = "DATA_RATE_2"

	// V3DataRateIndexDATARATE3 captures enum value "DATA_RATE_3"
	V3DataRateIndexDATARATE3 V3DataRateIndex = "DATA_RATE_3"

	// V3DataRateIndexDATARATE4 captures enum value "DATA_RATE_4"
	V3DataRateIndexDATARATE4 V3DataRateIndex = "DATA_RATE_4"

	// V3DataRateIndexDATARATE5 captures enum value "DATA_RATE_5"
	V3DataRateIndexDATARATE5 V3DataRateIndex = "DATA_RATE_5"

	// V3DataRateIndexDATARATE6 captures enum value "DATA_RATE_6"
	V3DataRateIndexDATARATE6 V3DataRateIndex = "DATA_RATE_6"

	// V3DataRateIndexDATARATE7 captures enum value "DATA_RATE_7"
	V3DataRateIndexDATARATE7 V3DataRateIndex = "DATA_RATE_7"

	// V3DataRateIndexDATARATE8 captures enum value "DATA_RATE_8"
	V3DataRateIndexDATARATE8 V3DataRateIndex = "DATA_RATE_8"

	// V3DataRateIndexDATARATE9 captures enum value "DATA_RATE_9"
	V3DataRateIndexDATARATE9 V3DataRateIndex = "DATA_RATE_9"

	// V3DataRateIndexDATARATE10 captures enum value "DATA_RATE_10"
	V3DataRateIndexDATARATE10 V3DataRateIndex = "DATA_RATE_10"

	// V3DataRateIndexDATARATE11 captures enum value "DATA_RATE_11"
	V3DataRateIndexDATARATE11 V3DataRateIndex = "DATA_RATE_11"

	// V3DataRateIndexDATARATE12 captures enum value "DATA_RATE_12"
	V3DataRateIndexDATARATE12 V3DataRateIndex = "DATA_RATE_12"

	// V3DataRateIndexDATARATE13 captures enum value "DATA_RATE_13"
	V3DataRateIndexDATARATE13 V3DataRateIndex = "DATA_RATE_13"

	// V3DataRateIndexDATARATE14 captures enum value "DATA_RATE_14"
	V3DataRateIndexDATARATE14 V3DataRateIndex = "DATA_RATE_14"

	// V3DataRateIndexDATARATE15 captures enum value "DATA_RATE_15"
	V3DataRateIndexDATARATE15 V3DataRateIndex = "DATA_RATE_15"
)

// for schema
var v3DataRateIndexEnum []interface{}

func init() {
	var res []V3DataRateIndex
	if err := json.Unmarshal([]byte(`["DATA_RATE_0","DATA_RATE_1","DATA_RATE_2","DATA_RATE_3","DATA_RATE_4","DATA_RATE_5","DATA_RATE_6","DATA_RATE_7","DATA_RATE_8","DATA_RATE_9","DATA_RATE_10","DATA_RATE_11","DATA_RATE_12","DATA_RATE_13","DATA_RATE_14","DATA_RATE_15"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v3DataRateIndexEnum = append(v3DataRateIndexEnum, v)
	}
}

func (m V3DataRateIndex) validateV3DataRateIndexEnum(path, location string, value V3DataRateIndex) error {
	if err := validate.Enum(path, location, value, v3DataRateIndexEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v3 data rate index
func (m V3DataRateIndex) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV3DataRateIndexEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
