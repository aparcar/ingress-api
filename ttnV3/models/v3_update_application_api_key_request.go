// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3UpdateApplicationAPIKeyRequest v3 update application API key request
//
// swagger:model v3UpdateApplicationAPIKeyRequest
type V3UpdateApplicationAPIKeyRequest struct {

	// api key
	APIKey *V3APIKey `json:"api_key,omitempty"`

	// application ids
	ApplicationIds *V3ApplicationIdentifiers `json:"application_ids,omitempty"`
}

// Validate validates this v3 update application API key request
func (m *V3UpdateApplicationAPIKeyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplicationIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3UpdateApplicationAPIKeyRequest) validateAPIKey(formats strfmt.Registry) error {

	if swag.IsZero(m.APIKey) { // not required
		return nil
	}

	if m.APIKey != nil {
		if err := m.APIKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("api_key")
			}
			return err
		}
	}

	return nil
}

func (m *V3UpdateApplicationAPIKeyRequest) validateApplicationIds(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *V3UpdateApplicationAPIKeyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3UpdateApplicationAPIKeyRequest) UnmarshalBinary(b []byte) error {
	var res V3UpdateApplicationAPIKeyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
