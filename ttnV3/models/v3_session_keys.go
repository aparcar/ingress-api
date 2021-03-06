// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3SessionKeys Session keys for a LoRaWAN session.
// Only the components for which the keys were meant, will have the key-encryption-key (KEK) to decrypt the individual keys.
//
// swagger:model v3SessionKeys
type V3SessionKeys struct {

	// The (encrypted) Application Session Key.
	// This key is stored by the Application Server.
	AppsKey *V3KeyEnvelope `json:"app_s_key,omitempty"`

	// The (encrypted) Forwarding Network Session Integrity Key (or Network Session Key in 1.0 compatibility mode).
	// This key is stored by the (forwarding) Network Server.
	FNwksIntKey *V3KeyEnvelope `json:"f_nwk_s_int_key,omitempty"`

	// The (encrypted) Network Session Encryption Key.
	// This key is stored by the (serving) Network Server.
	NwksEncKey *V3KeyEnvelope `json:"nwk_s_enc_key,omitempty"`

	// The (encrypted) Serving Network Session Integrity Key.
	// This key is stored by the (serving) Network Server.
	SNwksIntKey *V3KeyEnvelope `json:"s_nwk_s_int_key,omitempty"`

	// Join Server issued identifier for the session keys.
	// This ID can be used to request the keys from the Join Server in case the are lost.
	// Format: byte
	SessionKeyID strfmt.Base64 `json:"session_key_id,omitempty"`
}

// Validate validates this v3 session keys
func (m *V3SessionKeys) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppsKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFNwksIntKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNwksEncKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSNwksIntKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3SessionKeys) validateAppsKey(formats strfmt.Registry) error {

	if swag.IsZero(m.AppsKey) { // not required
		return nil
	}

	if m.AppsKey != nil {
		if err := m.AppsKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("app_s_key")
			}
			return err
		}
	}

	return nil
}

func (m *V3SessionKeys) validateFNwksIntKey(formats strfmt.Registry) error {

	if swag.IsZero(m.FNwksIntKey) { // not required
		return nil
	}

	if m.FNwksIntKey != nil {
		if err := m.FNwksIntKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("f_nwk_s_int_key")
			}
			return err
		}
	}

	return nil
}

func (m *V3SessionKeys) validateNwksEncKey(formats strfmt.Registry) error {

	if swag.IsZero(m.NwksEncKey) { // not required
		return nil
	}

	if m.NwksEncKey != nil {
		if err := m.NwksEncKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nwk_s_enc_key")
			}
			return err
		}
	}

	return nil
}

func (m *V3SessionKeys) validateSNwksIntKey(formats strfmt.Registry) error {

	if swag.IsZero(m.SNwksIntKey) { // not required
		return nil
	}

	if m.SNwksIntKey != nil {
		if err := m.SNwksIntKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s_nwk_s_int_key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3SessionKeys) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3SessionKeys) UnmarshalBinary(b []byte) error {
	var res V3SessionKeys
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
