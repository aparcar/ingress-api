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

// V3Application Application is the message that defines an Application in the network.
//
// swagger:model v3Application
type V3Application struct {

	// attributes
	Attributes map[string]string `json:"attributes,omitempty"`

	// contact info
	ContactInfo []*V3ContactInfo `json:"contact_info"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// ids
	Ids *V3ApplicationIdentifiers `json:"ids,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`
}

// Validate validates this v3 application
func (m *V3Application) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContactInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIds(formats); err != nil {
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

func (m *V3Application) validateContactInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.ContactInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.ContactInfo); i++ {
		if swag.IsZero(m.ContactInfo[i]) { // not required
			continue
		}

		if m.ContactInfo[i] != nil {
			if err := m.ContactInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contact_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V3Application) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V3Application) validateIds(formats strfmt.Registry) error {

	if swag.IsZero(m.Ids) { // not required
		return nil
	}

	if m.Ids != nil {
		if err := m.Ids.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ids")
			}
			return err
		}
	}

	return nil
}

func (m *V3Application) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3Application) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3Application) UnmarshalBinary(b []byte) error {
	var res V3Application
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
