// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V3OAuthAccessToken v3 o auth access token
//
// swagger:model v3OAuthAccessToken
type V3OAuthAccessToken struct {

	// access token
	AccessToken string `json:"access_token,omitempty"`

	// client ids
	ClientIds *V3ClientIdentifiers `json:"client_ids,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// expires at
	// Format: date-time
	ExpiresAt strfmt.DateTime `json:"expires_at,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// refresh token
	RefreshToken string `json:"refresh_token,omitempty"`

	// rights
	Rights []V3Right `json:"rights"`

	// user ids
	UserIds *V3UserIdentifiers `json:"user_ids,omitempty"`

	// user session id
	UserSessionID string `json:"user_session_id,omitempty"`
}

// Validate validates this v3 o auth access token
func (m *V3OAuthAccessToken) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRights(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3OAuthAccessToken) validateClientIds(formats strfmt.Registry) error {

	if swag.IsZero(m.ClientIds) { // not required
		return nil
	}

	if m.ClientIds != nil {
		if err := m.ClientIds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client_ids")
			}
			return err
		}
	}

	return nil
}

func (m *V3OAuthAccessToken) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V3OAuthAccessToken) validateExpiresAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpiresAt) { // not required
		return nil
	}

	if err := validate.FormatOf("expires_at", "body", "date-time", m.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V3OAuthAccessToken) validateRights(formats strfmt.Registry) error {

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

func (m *V3OAuthAccessToken) validateUserIds(formats strfmt.Registry) error {

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
func (m *V3OAuthAccessToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3OAuthAccessToken) UnmarshalBinary(b []byte) error {
	var res V3OAuthAccessToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
