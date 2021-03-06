// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3LoRaDataRate v3 lo ra data rate
//
// swagger:model v3LoRaDataRate
type V3LoRaDataRate struct {

	// Bandwidth (Hz).
	Bandwidth int64 `json:"bandwidth,omitempty"`

	// spreading factor
	SpreadingFactor int64 `json:"spreading_factor,omitempty"`
}

// Validate validates this v3 lo ra data rate
func (m *V3LoRaDataRate) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V3LoRaDataRate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3LoRaDataRate) UnmarshalBinary(b []byte) error {
	var res V3LoRaDataRate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
