// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProvisionEndDevicesRequestIdentifiersRange provision end devices request identifiers range
//
// swagger:model ProvisionEndDevicesRequestIdentifiersRange
type ProvisionEndDevicesRequestIdentifiersRange struct {

	// join eui
	// Format: byte
	JoinEui strfmt.Base64 `json:"join_eui,omitempty"`

	// DevEUI to start issuing from.
	// Format: byte
	StartDevEui strfmt.Base64 `json:"start_dev_eui,omitempty"`
}

// Validate validates this provision end devices request identifiers range
func (m *ProvisionEndDevicesRequestIdentifiersRange) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProvisionEndDevicesRequestIdentifiersRange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProvisionEndDevicesRequestIdentifiersRange) UnmarshalBinary(b []byte) error {
	var res ProvisionEndDevicesRequestIdentifiersRange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
