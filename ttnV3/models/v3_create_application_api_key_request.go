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

// V3CreateApplicationAPIKeyRequest v3 create application API key request
//
// swagger:model v3CreateApplicationAPIKeyRequest
type V3CreateApplicationAPIKeyRequest struct {

	// application ids
	ApplicationIds *V3ApplicationIdentifiers `json:"application_ids,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// rights
	Rights []V3Right `json:"rights"`
}

// Validate validates this v3 create application API key request
func (m *V3CreateApplicationAPIKeyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRights(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3CreateApplicationAPIKeyRequest) validateApplicationIds(formats strfmt.Registry) error {

	if swag.IsZero(m.ApplicationIds) { // not required
		return nil
	}

	if m.ApplicationIds != nil {
		if err := m.ApplicationIds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application_ids")
			}
			return err
		}
	}

	return nil
}

func (m *V3CreateApplicationAPIKeyRequest) validateRights(formats strfmt.Registry) error {

	if swag.IsZero(m.Rights) { // not required
		return nil
	}

	for i := 0; i < len(m.Rights); i++ {

		if err := m.Rights[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rights" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3CreateApplicationAPIKeyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3CreateApplicationAPIKeyRequest) UnmarshalBinary(b []byte) error {
	var res V3CreateApplicationAPIKeyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
