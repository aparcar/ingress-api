// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MACCommandPingSlotChannelAns m a c command ping slot channel ans
//
// swagger:model MACCommandPingSlotChannelAns
type MACCommandPingSlotChannelAns struct {

	// data rate index ack
	DataRateIndexAck bool `json:"data_rate_index_ack,omitempty"`

	// frequency ack
	FrequencyAck bool `json:"frequency_ack,omitempty"`
}

// Validate validates this m a c command ping slot channel ans
func (m *MACCommandPingSlotChannelAns) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MACCommandPingSlotChannelAns) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MACCommandPingSlotChannelAns) UnmarshalBinary(b []byte) error {
	var res MACCommandPingSlotChannelAns
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
