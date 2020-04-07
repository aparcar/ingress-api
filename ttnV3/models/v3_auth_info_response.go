// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3AuthInfoResponse v3 auth info response
//
// swagger:model v3AuthInfoResponse
type V3AuthInfoResponse struct {

	// api key
	APIKey *AuthInfoResponseAPIKeyAccess `json:"api_key,omitempty"`

	// is admin
	IsAdmin bool `json:"is_admin,omitempty"`

	// oauth access token
	OauthAccessToken *V3OAuthAccessToken `json:"oauth_access_token,omitempty"`

	// universal rights
	UniversalRights *V3Rights `json:"universal_rights,omitempty"`
}

// Validate validates this v3 auth info response
func (m *V3AuthInfoResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOauthAccessToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUniversalRights(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3AuthInfoResponse) validateAPIKey(formats strfmt.Registry) error {

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

func (m *V3AuthInfoResponse) validateOauthAccessToken(formats strfmt.Registry) error {

	if swag.IsZero(m.OauthAccessToken) { // not required
		return nil
	}

	if m.OauthAccessToken != nil {
		if err := m.OauthAccessToken.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oauth_access_token")
			}
			return err
		}
	}

	return nil
}

func (m *V3AuthInfoResponse) validateUniversalRights(formats strfmt.Registry) error {

	if swag.IsZero(m.UniversalRights) { // not required
		return nil
	}

	if m.UniversalRights != nil {
		if err := m.UniversalRights.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("universal_rights")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3AuthInfoResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3AuthInfoResponse) UnmarshalBinary(b []byte) error {
	var res V3AuthInfoResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
