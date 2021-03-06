// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3DLSettings v3 d l settings
//
// swagger:model v3DLSettings
type V3DLSettings struct {

	// OptNeg is set if Network Server implements LoRaWAN 1.1 or greater.
	OptNeg bool `json:"opt_neg,omitempty"`

	// rx1 dr offset
	Rx1DrOffset int64 `json:"rx1_dr_offset,omitempty"`

	// rx2 dr
	Rx2Dr V3DataRateIndex `json:"rx2_dr,omitempty"`
}

// Validate validates this v3 d l settings
func (m *V3DLSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRx2Dr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3DLSettings) validateRx2Dr(formats strfmt.Registry) error {

	if swag.IsZero(m.Rx2Dr) { // not required
		return nil
	}

	if err := m.Rx2Dr.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rx2_dr")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3DLSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3DLSettings) UnmarshalBinary(b []byte) error {
	var res V3DLSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
