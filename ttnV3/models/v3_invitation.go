// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V3Invitation v3 invitation
//
// swagger:model v3Invitation
type V3Invitation struct {

	// accepted at
	// Format: date-time
	AcceptedAt strfmt.DateTime `json:"accepted_at,omitempty"`

	// accepted by
	AcceptedBy *V3UserIdentifiers `json:"accepted_by,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// email
	Email string `json:"email,omitempty"`

	// expires at
	// Format: date-time
	ExpiresAt strfmt.DateTime `json:"expires_at,omitempty"`

	// token
	Token string `json:"token,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this v3 invitation
func (m *V3Invitation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAcceptedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3Invitation) validateAcceptedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.AcceptedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("accepted_at", "body", "date-time", m.AcceptedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V3Invitation) validateAcceptedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.AcceptedBy) { // not required
		return nil
	}

	if m.AcceptedBy != nil {
		if err := m.AcceptedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accepted_by")
			}
			return err
		}
	}

	return nil
}

func (m *V3Invitation) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V3Invitation) validateExpiresAt(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpiresAt) { // not required
		return nil
	}

	if err := validate.FormatOf("expires_at", "body", "date-time", m.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V3Invitation) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3Invitation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3Invitation) UnmarshalBinary(b []byte) error {
	var res V3Invitation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
