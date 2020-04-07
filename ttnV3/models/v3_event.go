// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V3Event v3 event
//
// swagger:model v3Event
type V3Event struct {

	// context
	Context map[string]strfmt.Base64 `json:"context,omitempty"`

	// correlation ids
	CorrelationIds []string `json:"correlation_ids"`

	// data
	Data *ProtobufAny `json:"data,omitempty"`

	// identifiers
	Identifiers []*V3EntityIdentifiers `json:"identifiers"`

	// name
	Name string `json:"name,omitempty"`

	// origin
	Origin string `json:"origin,omitempty"`

	// time
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`

	// The event will be visible to a caller that has any of these rights.
	Visibility *V3Rights `json:"visibility,omitempty"`
}

// Validate validates this v3 event
func (m *V3Event) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentifiers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVisibility(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3Event) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

func (m *V3Event) validateIdentifiers(formats strfmt.Registry) error {

	if swag.IsZero(m.Identifiers) { // not required
		return nil
	}

	for i := 0; i < len(m.Identifiers); i++ {
		if swag.IsZero(m.Identifiers[i]) { // not required
			continue
		}

		if m.Identifiers[i] != nil {
			if err := m.Identifiers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("identifiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V3Event) validateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V3Event) validateVisibility(formats strfmt.Registry) error {

	if swag.IsZero(m.Visibility) { // not required
		return nil
	}

	if m.Visibility != nil {
		if err := m.Visibility.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("visibility")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3Event) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3Event) UnmarshalBinary(b []byte) error {
	var res V3Event
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
