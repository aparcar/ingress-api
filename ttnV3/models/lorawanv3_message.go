// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Lorawanv3Message lorawanv3 message
//
// swagger:model lorawanv3Message
type Lorawanv3Message struct {

	// join accept payload
	JoinAcceptPayload *V3JoinAcceptPayload `json:"join_accept_payload,omitempty"`

	// join request payload
	JoinRequestPayload *V3JoinRequestPayload `json:"join_request_payload,omitempty"`

	// m hdr
	MHdr *V3MHDR `json:"m_hdr,omitempty"`

	// mac payload
	MacPayload *V3MACPayload `json:"mac_payload,omitempty"`

	// mic
	// Format: byte
	Mic strfmt.Base64 `json:"mic,omitempty"`

	// rejoin request payload
	RejoinRequestPayload *V3RejoinRequestPayload `json:"rejoin_request_payload,omitempty"`
}

// Validate validates this lorawanv3 message
func (m *Lorawanv3Message) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJoinAcceptPayload(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJoinRequestPayload(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMHdr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMacPayload(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRejoinRequestPayload(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Lorawanv3Message) validateJoinAcceptPayload(formats strfmt.Registry) error {

	if swag.IsZero(m.JoinAcceptPayload) { // not required
		return nil
	}

	if m.JoinAcceptPayload != nil {
		if err := m.JoinAcceptPayload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("join_accept_payload")
			}
			return err
		}
	}

	return nil
}

func (m *Lorawanv3Message) validateJoinRequestPayload(formats strfmt.Registry) error {

	if swag.IsZero(m.JoinRequestPayload) { // not required
		return nil
	}

	if m.JoinRequestPayload != nil {
		if err := m.JoinRequestPayload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("join_request_payload")
			}
			return err
		}
	}

	return nil
}

func (m *Lorawanv3Message) validateMHdr(formats strfmt.Registry) error {

	if swag.IsZero(m.MHdr) { // not required
		return nil
	}

	if m.MHdr != nil {
		if err := m.MHdr.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("m_hdr")
			}
			return err
		}
	}

	return nil
}

func (m *Lorawanv3Message) validateMacPayload(formats strfmt.Registry) error {

	if swag.IsZero(m.MacPayload) { // not required
		return nil
	}

	if m.MacPayload != nil {
		if err := m.MacPayload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mac_payload")
			}
			return err
		}
	}

	return nil
}

func (m *Lorawanv3Message) validateRejoinRequestPayload(formats strfmt.Registry) error {

	if swag.IsZero(m.RejoinRequestPayload) { // not required
		return nil
	}

	if m.RejoinRequestPayload != nil {
		if err := m.RejoinRequestPayload.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rejoin_request_payload")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Lorawanv3Message) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Lorawanv3Message) UnmarshalBinary(b []byte) error {
	var res Lorawanv3Message
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
