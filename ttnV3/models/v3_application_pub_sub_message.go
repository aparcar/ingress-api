// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3ApplicationPubSubMessage v3 application pub sub message
//
// swagger:model v3ApplicationPubSubMessage
type V3ApplicationPubSubMessage struct {

	// The topic on which the Application Server publishes or receives the messages.
	Topic string `json:"topic,omitempty"`
}

// Validate validates this v3 application pub sub message
func (m *V3ApplicationPubSubMessage) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V3ApplicationPubSubMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3ApplicationPubSubMessage) UnmarshalBinary(b []byte) error {
	var res V3ApplicationPubSubMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
