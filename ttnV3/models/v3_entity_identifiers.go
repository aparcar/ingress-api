// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3EntityIdentifiers EntityIdentifiers contains one of the possible entity identifiers.
//
// swagger:model v3EntityIdentifiers
type V3EntityIdentifiers struct {

	// application ids
	ApplicationIds *V3ApplicationIdentifiers `json:"application_ids,omitempty"`

	// client ids
	ClientIds *V3ClientIdentifiers `json:"client_ids,omitempty"`

	// device ids
	DeviceIds *V3EndDeviceIdentifiers `json:"device_ids,omitempty"`

	// gateway ids
	GatewayIds *V3GatewayIdentifiers `json:"gateway_ids,omitempty"`

	// organization ids
	OrganizationIds *V3OrganizationIdentifiers `json:"organization_ids,omitempty"`

	// user ids
	UserIds *V3UserIdentifiers `json:"user_ids,omitempty"`
}

// Validate validates this v3 entity identifiers
func (m *V3EntityIdentifiers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplicationIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGatewayIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationIds(formats); err != nil {
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

func (m *V3EntityIdentifiers) validateApplicationIds(formats strfmt.Registry) error {

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

func (m *V3EntityIdentifiers) validateClientIds(formats strfmt.Registry) error {

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

func (m *V3EntityIdentifiers) validateDeviceIds(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceIds) { // not required
		return nil
	}

	if m.DeviceIds != nil {
		if err := m.DeviceIds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_ids")
			}
			return err
		}
	}

	return nil
}

func (m *V3EntityIdentifiers) validateGatewayIds(formats strfmt.Registry) error {

	if swag.IsZero(m.GatewayIds) { // not required
		return nil
	}

	if m.GatewayIds != nil {
		if err := m.GatewayIds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gateway_ids")
			}
			return err
		}
	}

	return nil
}

func (m *V3EntityIdentifiers) validateOrganizationIds(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganizationIds) { // not required
		return nil
	}

	if m.OrganizationIds != nil {
		if err := m.OrganizationIds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("organization_ids")
			}
			return err
		}
	}

	return nil
}

func (m *V3EntityIdentifiers) validateUserIds(formats strfmt.Registry) error {

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
func (m *V3EntityIdentifiers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3EntityIdentifiers) UnmarshalBinary(b []byte) error {
	var res V3EntityIdentifiers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
