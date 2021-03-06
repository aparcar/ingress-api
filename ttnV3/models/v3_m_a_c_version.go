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

// V3MACVersion v3 m a c version
//
// swagger:model v3MACVersion
type V3MACVersion string

const (

	// V3MACVersionMACUNKNOWN captures enum value "MAC_UNKNOWN"
	V3MACVersionMACUNKNOWN V3MACVersion = "MAC_UNKNOWN"

	// V3MACVersionMACV10 captures enum value "MAC_V1_0"
	V3MACVersionMACV10 V3MACVersion = "MAC_V1_0"

	// V3MACVersionMACV101 captures enum value "MAC_V1_0_1"
	V3MACVersionMACV101 V3MACVersion = "MAC_V1_0_1"

	// V3MACVersionMACV102 captures enum value "MAC_V1_0_2"
	V3MACVersionMACV102 V3MACVersion = "MAC_V1_0_2"

	// V3MACVersionMACV11 captures enum value "MAC_V1_1"
	V3MACVersionMACV11 V3MACVersion = "MAC_V1_1"

	// V3MACVersionMACV103 captures enum value "MAC_V1_0_3"
	V3MACVersionMACV103 V3MACVersion = "MAC_V1_0_3"

	// V3MACVersionMACV104 captures enum value "MAC_V1_0_4"
	V3MACVersionMACV104 V3MACVersion = "MAC_V1_0_4"
)

// for schema
var v3MACVersionEnum []interface{}

func init() {
	var res []V3MACVersion
	if err := json.Unmarshal([]byte(`["MAC_UNKNOWN","MAC_V1_0","MAC_V1_0_1","MAC_V1_0_2","MAC_V1_1","MAC_V1_0_3","MAC_V1_0_4"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v3MACVersionEnum = append(v3MACVersionEnum, v)
	}
}

func (m V3MACVersion) validateV3MACVersionEnum(path, location string, value V3MACVersion) error {
	if err := validate.Enum(path, location, value, v3MACVersionEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v3 m a c version
func (m V3MACVersion) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV3MACVersionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
