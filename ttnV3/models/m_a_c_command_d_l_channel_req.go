// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MACCommandDLChannelReq m a c command d l channel req
//
// swagger:model MACCommandDLChannelReq
type MACCommandDLChannelReq struct {

	// channel index
	ChannelIndex int64 `json:"channel_index,omitempty"`

	// frequency
	Frequency string `json:"frequency,omitempty"`
}

// Validate validates this m a c command d l channel req
func (m *MACCommandDLChannelReq) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MACCommandDLChannelReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MACCommandDLChannelReq) UnmarshalBinary(b []byte) error {
	var res MACCommandDLChannelReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
