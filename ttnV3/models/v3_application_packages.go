// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3ApplicationPackages v3 application packages
//
// swagger:model v3ApplicationPackages
type V3ApplicationPackages struct {

	// packages
	Packages []*V3ApplicationPackage `json:"packages"`
}

// Validate validates this v3 application packages
func (m *V3ApplicationPackages) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePackages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3ApplicationPackages) validatePackages(formats strfmt.Registry) error {

	if swag.IsZero(m.Packages) { // not required
		return nil
	}

	for i := 0; i < len(m.Packages); i++ {
		if swag.IsZero(m.Packages[i]) { // not required
			continue
		}

		if m.Packages[i] != nil {
			if err := m.Packages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("packages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3ApplicationPackages) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3ApplicationPackages) UnmarshalBinary(b []byte) error {
	var res V3ApplicationPackages
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
