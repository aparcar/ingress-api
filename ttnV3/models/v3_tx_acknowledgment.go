// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V3TxAcknowledgment v3 tx acknowledgment
//
// swagger:model v3TxAcknowledgment
type V3TxAcknowledgment struct {

	// correlation ids
	CorrelationIds []string `json:"correlation_ids"`

	// result
	Result TxAcknowledgmentResult `json:"result,omitempty"`
}

// Validate validates this v3 tx acknowledgment
func (m *V3TxAcknowledgment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3TxAcknowledgment) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if err := m.Result.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("result")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3TxAcknowledgment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3TxAcknowledgment) UnmarshalBinary(b []byte) error {
	var res V3TxAcknowledgment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
