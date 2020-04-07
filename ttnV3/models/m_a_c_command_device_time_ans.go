// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MACCommandDeviceTimeAns m a c command device time ans
//
// swagger:model MACCommandDeviceTimeAns
type MACCommandDeviceTimeAns struct {

	// time
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`
}

// Validate validates this m a c command device time ans
func (m *MACCommandDeviceTimeAns) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MACCommandDeviceTimeAns) validateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MACCommandDeviceTimeAns) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MACCommandDeviceTimeAns) UnmarshalBinary(b []byte) error {
	var res MACCommandDeviceTimeAns
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
