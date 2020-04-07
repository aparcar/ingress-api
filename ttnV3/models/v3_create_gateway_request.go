// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3CreateGatewayRequest v3 create gateway request
//
// swagger:model v3CreateGatewayRequest
type V3CreateGatewayRequest struct {

	// Collaborator to grant all rights on the newly created gateway.
	Collaborator *V3OrganizationOrUserIdentifiers `json:"collaborator,omitempty"`

	// gateway
	Gateway *V3Gateway `json:"gateway,omitempty"`
}

// Validate validates this v3 create gateway request
func (m *V3CreateGatewayRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCollaborator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3CreateGatewayRequest) validateCollaborator(formats strfmt.Registry) error {

	if swag.IsZero(m.Collaborator) { // not required
		return nil
	}

	if m.Collaborator != nil {
		if err := m.Collaborator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("collaborator")
			}
			return err
		}
	}

	return nil
}

func (m *V3CreateGatewayRequest) validateGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.Gateway) { // not required
		return nil
	}

	if m.Gateway != nil {
		if err := m.Gateway.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gateway")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3CreateGatewayRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3CreateGatewayRequest) UnmarshalBinary(b []byte) error {
	var res V3CreateGatewayRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
