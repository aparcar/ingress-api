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

// V3ApplicationInvalidatedDownlinks v3 application invalidated downlinks
//
// swagger:model v3ApplicationInvalidatedDownlinks
type V3ApplicationInvalidatedDownlinks struct {

	// downlinks
	Downlinks []*V3ApplicationDownlink `json:"downlinks"`

	// last f cnt down
	LastfCntDown int64 `json:"last_f_cnt_down,omitempty"`
}

// Validate validates this v3 application invalidated downlinks
func (m *V3ApplicationInvalidatedDownlinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDownlinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3ApplicationInvalidatedDownlinks) validateDownlinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Downlinks) { // not required
		return nil
	}

	for i := 0; i < len(m.Downlinks); i++ {
		if swag.IsZero(m.Downlinks[i]) { // not required
			continue
		}

		if m.Downlinks[i] != nil {
			if err := m.Downlinks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("downlinks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3ApplicationInvalidatedDownlinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3ApplicationInvalidatedDownlinks) UnmarshalBinary(b []byte) error {
	var res V3ApplicationInvalidatedDownlinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}