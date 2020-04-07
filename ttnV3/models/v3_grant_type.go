// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V3GrantType The OAuth2 flows an OAuth client can use to get an access token.
//
//  - GRANT_AUTHORIZATION_CODE: Grant type used to exchange an authorization code for an access token.
//  - GRANT_PASSWORD: Grant type used to exchange a user ID and password for an access token.
//  - GRANT_REFRESH_TOKEN: Grant type used to exchange a refresh token for an access token.
//
// swagger:model v3GrantType
type V3GrantType string

const (

	// V3GrantTypeGRANTAUTHORIZATIONCODE captures enum value "GRANT_AUTHORIZATION_CODE"
	V3GrantTypeGRANTAUTHORIZATIONCODE V3GrantType = "GRANT_AUTHORIZATION_CODE"

	// V3GrantTypeGRANTPASSWORD captures enum value "GRANT_PASSWORD"
	V3GrantTypeGRANTPASSWORD V3GrantType = "GRANT_PASSWORD"

	// V3GrantTypeGRANTREFRESHTOKEN captures enum value "GRANT_REFRESH_TOKEN"
	V3GrantTypeGRANTREFRESHTOKEN V3GrantType = "GRANT_REFRESH_TOKEN"
)

// for schema
var v3GrantTypeEnum []interface{}

func init() {
	var res []V3GrantType
	if err := json.Unmarshal([]byte(`["GRANT_AUTHORIZATION_CODE","GRANT_PASSWORD","GRANT_REFRESH_TOKEN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v3GrantTypeEnum = append(v3GrantTypeEnum, v)
	}
}

func (m V3GrantType) validateV3GrantTypeEnum(path, location string, value V3GrantType) error {
	if err := validate.Enum(path, location, value, v3GrantTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v3 grant type
func (m V3GrantType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV3GrantTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
