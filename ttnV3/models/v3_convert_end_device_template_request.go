// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3ConvertEndDeviceTemplateRequest v3 convert end device template request
//
// swagger:model v3ConvertEndDeviceTemplateRequest
type V3ConvertEndDeviceTemplateRequest struct {

	// Data to convert.
	// Format: byte
	Data strfmt.Base64 `json:"data,omitempty"`

	// ID of the format.
	FormatID string `json:"format_id,omitempty"`
}

// Validate validates this v3 convert end device template request
func (m *V3ConvertEndDeviceTemplateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V3ConvertEndDeviceTemplateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3ConvertEndDeviceTemplateRequest) UnmarshalBinary(b []byte) error {
	var res V3ConvertEndDeviceTemplateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
