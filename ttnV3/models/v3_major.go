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

// V3Major v3 major
//
// swagger:model v3Major
type V3Major string

const (

	// V3MajorLORAWANR1 captures enum value "LORAWAN_R1"
	V3MajorLORAWANR1 V3Major = "LORAWAN_R1"
)

// for schema
var v3MajorEnum []interface{}

func init() {
	var res []V3Major
	if err := json.Unmarshal([]byte(`["LORAWAN_R1"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v3MajorEnum = append(v3MajorEnum, v)
	}
}

func (m V3Major) validateV3MajorEnum(path, location string, value V3Major) error {
	if err := validate.Enum(path, location, value, v3MajorEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v3 major
func (m V3Major) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV3MajorEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
