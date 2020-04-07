// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3CreateEndDeviceRequest v3 create end device request
//
// swagger:model v3CreateEndDeviceRequest
type V3CreateEndDeviceRequest struct {

	// end device
	EndDevice *V3EndDevice `json:"end_device,omitempty"`
}

// Validate validates this v3 create end device request
func (m *V3CreateEndDeviceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndDevice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3CreateEndDeviceRequest) validateEndDevice(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *V3CreateEndDeviceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3CreateEndDeviceRequest) UnmarshalBinary(b []byte) error {
	var res V3CreateEndDeviceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
