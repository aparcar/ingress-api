// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3UpdateUserPasswordRequest v3 update user password request
//
// swagger:model v3UpdateUserPasswordRequest
type V3UpdateUserPasswordRequest struct {

	// new
	New string `json:"new,omitempty"`

	// old
	Old string `json:"old,omitempty"`

	// Revoke active sessions and access tokens of user if true. To be used if credentials are suspected to be compromised.
	RevokeAllAccess bool `json:"revoke_all_access,omitempty"`

	// user ids
	UserIds *V3UserIdentifiers `json:"user_ids,omitempty"`
}

// Validate validates this v3 update user password request
func (m *V3UpdateUserPasswordRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3UpdateUserPasswordRequest) validateUserIds(formats strfmt.Registry) error {

	if swag.IsZero(m.UserIds) { // not required
		return nil
	}

	if m.UserIds != nil {
		if err := m.UserIds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user_ids")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3UpdateUserPasswordRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3UpdateUserPasswordRequest) UnmarshalBinary(b []byte) error {
	var res V3UpdateUserPasswordRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
