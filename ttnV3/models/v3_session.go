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

// V3Session v3 session
//
// swagger:model v3Session
type V3Session struct {

	// Device Address, issued by the Network Server or chosen by device manufacturer in case of testing range (beginning with 00-03).
	// Known by Network Server, Application Server and Join Server. Owned by Network Server.
	// Format: byte
	DevAddr strfmt.Base64 `json:"dev_addr,omitempty"`

	// keys
	Keys *V3SessionKeys `json:"keys,omitempty"`

	// Last application downlink frame counter value used. Application Server only.
	LastafCntDown int64 `json:"last_a_f_cnt_down,omitempty"`

	// Frame counter of the last confirmed downlink message sent. Network Server only.
	LastConffCntDown int64 `json:"last_conf_f_cnt_down,omitempty"`

	// Last uplink frame counter value used. Network Server only. Application Server assumes the Network Server checked it.
	LastfCntUp int64 `json:"last_f_cnt_up,omitempty"`

	// Last network downlink frame counter value used. Network Server only.
	LastnfCntDown int64 `json:"last_n_f_cnt_down,omitempty"`

	// Queued Application downlink messages. Stored in Application Server and Network Server.
	QueuedApplicationDownlinks []*V3ApplicationDownlink `json:"queued_application_downlinks"`

	// Time when the session started. Network Server only.
	// Format: date-time
	StartedAt strfmt.DateTime `json:"started_at,omitempty"`
}

// Validate validates this v3 session
func (m *V3Session) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeys(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueuedApplicationDownlinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3Session) validateKeys(formats strfmt.Registry) error {

	if swag.IsZero(m.Keys) { // not required
		return nil
	}

	if m.Keys != nil {
		if err := m.Keys.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keys")
			}
			return err
		}
	}

	return nil
}

func (m *V3Session) validateQueuedApplicationDownlinks(formats strfmt.Registry) error {

	if swag.IsZero(m.QueuedApplicationDownlinks) { // not required
		return nil
	}

	for i := 0; i < len(m.QueuedApplicationDownlinks); i++ {
		if swag.IsZero(m.QueuedApplicationDownlinks[i]) { // not required
			continue
		}

		if m.QueuedApplicationDownlinks[i] != nil {
			if err := m.QueuedApplicationDownlinks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("queued_application_downlinks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V3Session) validateStartedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("started_at", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3Session) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3Session) UnmarshalBinary(b []byte) error {
	var res V3Session
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
