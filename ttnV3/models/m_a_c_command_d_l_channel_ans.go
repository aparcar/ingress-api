// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MACCommandDLChannelAns m a c command d l channel ans
//
// swagger:model MACCommandDLChannelAns
type MACCommandDLChannelAns struct {

	// channel index ack
	ChannelIndexAck bool `json:"channel_index_ack,omitempty"`

	// frequency ack
	FrequencyAck bool `json:"frequency_ack,omitempty"`
}

// Validate validates this m a c command d l channel ans
func (m *MACCommandDLChannelAns) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MACCommandDLChannelAns) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MACCommandDLChannelAns) UnmarshalBinary(b []byte) error {
	var res MACCommandDLChannelAns
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
