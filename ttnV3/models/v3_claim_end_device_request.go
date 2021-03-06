// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3ClaimEndDeviceRequest v3 claim end device request
//
// swagger:model v3ClaimEndDeviceRequest
type V3ClaimEndDeviceRequest struct {

	// authenticated identifiers
	AuthenticatedIdentifiers *ClaimEndDeviceRequestAuthenticatedIdentifiers `json:"authenticated_identifiers,omitempty"`

	// If set, invalidate the authentication code with which the device gets claimed. This prohibits subsequent claiming requests.
	InvalidateAuthenticationCode bool `json:"invalidate_authentication_code,omitempty"`

	// qr code
	// Format: byte
	QrCode strfmt.Base64 `json:"qr_code,omitempty"`

	// Application identifiers of the target end device.
	TargetApplicationIds *V3ApplicationIdentifiers `json:"target_application_ids,omitempty"`

	// The address of the Application Server where the device will be registered.
	// If set and if the source device is currently registered on an Application Server, settings will be transferred.
	// If not set, the device shall not be registered on an Application Server.
	TargetApplicationServerAddress string `json:"target_application_server_address,omitempty"`

	// The AS-ID of the Application Server to use.
	TargetApplicationServerID string `json:"target_application_server_id,omitempty"`

	// The KEK label of the Application Server to use for wrapping the application session key.
	TargetApplicationServerKekLabel string `json:"target_application_server_kek_label,omitempty"`

	// End device ID of the target end device. If empty, use the source device ID.
	TargetDeviceID string `json:"target_device_id,omitempty"`

	// Home NetID.
	// Format: byte
	TargetNetID strfmt.Base64 `json:"target_net_id,omitempty"`

	// The address of the Network Server where the device will be registered.
	// If set and if the source device is currently registered on a Network Server, settings will be transferred.
	// If not set, the device shall not be registered on a Network Server.
	TargetNetworkServerAddress string `json:"target_network_server_address,omitempty"`

	// The KEK label of the Network Server to use for wrapping network session keys.
	TargetNetworkServerKekLabel string `json:"target_network_server_kek_label,omitempty"`
}

// Validate validates this v3 claim end device request
func (m *V3ClaimEndDeviceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthenticatedIdentifiers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetApplicationIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3ClaimEndDeviceRequest) validateAuthenticatedIdentifiers(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthenticatedIdentifiers) { // not required
		return nil
	}

	if m.AuthenticatedIdentifiers != nil {
		if err := m.AuthenticatedIdentifiers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authenticated_identifiers")
			}
			return err
		}
	}

	return nil
}

func (m *V3ClaimEndDeviceRequest) validateTargetApplicationIds(formats strfmt.Registry) error {

	if swag.IsZero(m.TargetApplicationIds) { // not required
		return nil
	}

	if m.TargetApplicationIds != nil {
		if err := m.TargetApplicationIds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("target_application_ids")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3ClaimEndDeviceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3ClaimEndDeviceRequest) UnmarshalBinary(b []byte) error {
	var res V3ClaimEndDeviceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
