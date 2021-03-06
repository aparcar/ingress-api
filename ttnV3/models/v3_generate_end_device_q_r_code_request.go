// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3GenerateEndDeviceQRCodeRequest v3 generate end device q r code request
//
// swagger:model v3GenerateEndDeviceQRCodeRequest
type V3GenerateEndDeviceQRCodeRequest struct {

	// end device
	EndDevice *V3EndDevice `json:"end_device,omitempty"`

	// format id
	FormatID string `json:"format_id,omitempty"`

	// image
	Image *GenerateEndDeviceQRCodeRequestImage `json:"image,omitempty"`
}

// Validate validates this v3 generate end device q r code request
func (m *V3GenerateEndDeviceQRCodeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3GenerateEndDeviceQRCodeRequest) validateEndDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.EndDevice) { // not required
		return nil
	}

	if m.EndDevice != nil {
		if err := m.EndDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("end_device")
			}
			return err
		}
	}

	return nil
}

func (m *V3GenerateEndDeviceQRCodeRequest) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3GenerateEndDeviceQRCodeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3GenerateEndDeviceQRCodeRequest) UnmarshalBinary(b []byte) error {
	var res V3GenerateEndDeviceQRCodeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
