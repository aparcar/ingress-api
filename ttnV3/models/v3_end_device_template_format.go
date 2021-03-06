// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3EndDeviceTemplateFormat v3 end device template format
//
// swagger:model v3EndDeviceTemplateFormat
type V3EndDeviceTemplateFormat struct {

	// description
	Description string `json:"description,omitempty"`

	// file extensions
	FileExtensions []string `json:"file_extensions"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this v3 end device template format
func (m *V3EndDeviceTemplateFormat) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V3EndDeviceTemplateFormat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3EndDeviceTemplateFormat) UnmarshalBinary(b []byte) error {
	var res V3EndDeviceTemplateFormat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
