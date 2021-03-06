// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MACCommandDeviceModeConf m a c command device mode conf
//
// swagger:model MACCommandDeviceModeConf
type MACCommandDeviceModeConf struct {

	// class
	Class V3Class `json:"class,omitempty"`
}

// Validate validates this m a c command device mode conf
func (m *MACCommandDeviceModeConf) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClass(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MACCommandDeviceModeConf) validateClass(formats strfmt.Registry) error {

	if swag.IsZero(m.Class) { // not required
		return nil
	}

	if err := m.Class.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("class")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MACCommandDeviceModeConf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MACCommandDeviceModeConf) UnmarshalBinary(b []byte) error {
	var res MACCommandDeviceModeConf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
