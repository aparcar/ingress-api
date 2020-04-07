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

// V3JoinEUIPrefixes v3 join e UI prefixes
//
// swagger:model v3JoinEUIPrefixes
type V3JoinEUIPrefixes struct {

	// prefixes
	Prefixes []*V3JoinEUIPrefix `json:"prefixes"`
}

// Validate validates this v3 join e UI prefixes
func (m *V3JoinEUIPrefixes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrefixes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3JoinEUIPrefixes) validatePrefixes(formats strfmt.Registry) error {

	if swag.IsZero(m.Prefixes) { // not required
		return nil
	}

	for i := 0; i < len(m.Prefixes); i++ {
		if swag.IsZero(m.Prefixes[i]) { // not required
			continue
		}

		if m.Prefixes[i] != nil {
			if err := m.Prefixes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prefixes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3JoinEUIPrefixes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3JoinEUIPrefixes) UnmarshalBinary(b []byte) error {
	var res V3JoinEUIPrefixes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
