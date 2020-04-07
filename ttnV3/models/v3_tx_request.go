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

// V3TxRequest TxRequest is a request for transmission.
// If sent to a roaming partner, this request is used to generate the DLMetadata Object (see Backend Interfaces 1.0, Table 22).
// If the gateway has a scheduler, this request is sent to the gateway, in the order of gateway_ids.
// Otherwise, the Gateway Server attempts to schedule the request and creates the TxSettings.
//
// swagger:model v3TxRequest
type V3TxRequest struct {

	// Time when the downlink message should be transmitted.
	// This value is only valid for class C downlink; class A downlink uses uplink tokens and class B downlink is scheduled on ping slots.
	// This requires the gateway to have GPS time sychronization.
	// If the absolute time is not set, the first available time will be used that does not conflict or violate regional limitations.
	// Format: date-time
	AbsoluteTime strfmt.DateTime `json:"absolute_time,omitempty"`

	// Advanced metadata fields
	// - can be used for advanced information or experimental features that are not yet formally defined in the API
	// - field names are written in snake_case
	Advanced interface{} `json:"advanced,omitempty"`

	// class
	Class V3Class `json:"class,omitempty"`

	// Downlink paths used to select a gateway for downlink.
	// In class A, the downlink paths are required to only contain uplink tokens.
	// In class B and C, the downlink paths may contain uplink tokens and fixed gateways antenna identifiers.
	DownlinkPaths []*V3DownlinkPath `json:"downlink_paths"`

	// Frequency plan ID from which the frequencies in this message are retrieved.
	FrequencyPlanID string `json:"frequency_plan_id,omitempty"`

	// Priority for scheduling.
	// Requests with a higher priority are allocated more channel time than messages with a lower priority, in duty-cycle limited regions.
	// A priority of HIGH or higher sets the HiPriorityFlag in the DLMetadata Object.
	Priority V3TxSchedulePriority `json:"priority,omitempty"`

	// LoRaWAN data rate index for Rx1.
	Rx1DataRateIndex V3DataRateIndex `json:"rx1_data_rate_index,omitempty"`

	// Rx1 delay (Rx2 delay is Rx1 delay + 1 second).
	Rx1Delay V3RxDelay `json:"rx1_delay,omitempty"`

	// Frequency (Hz) for Rx1.
	Rx1Frequency string `json:"rx1_frequency,omitempty"`

	// LoRaWAN data rate index for Rx2.
	Rx2DataRateIndex V3DataRateIndex `json:"rx2_data_rate_index,omitempty"`

	// Frequency (Hz) for Rx2.
	Rx2Frequency string `json:"rx2_frequency,omitempty"`
}

// Validate validates this v3 tx request
func (m *V3TxRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbsoluteTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClass(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDownlinkPaths(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRx1DataRateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRx1Delay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRx2DataRateIndex(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3TxRequest) validateAbsoluteTime(formats strfmt.Registry) error {

	if swag.IsZero(m.AbsoluteTime) { // not required
		return nil
	}

	if err := validate.FormatOf("absolute_time", "body", "date-time", m.AbsoluteTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V3TxRequest) validateClass(formats strfmt.Registry) error {

	if swag.IsZero(m.Class) { // not required
		return nil
	}

	if err := m.Class.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("class")
		}
		return err
	}

	return nil
}

func (m *V3TxRequest) validateDownlinkPaths(formats strfmt.Registry) error {

	if swag.IsZero(m.DownlinkPaths) { // not required
		return nil
	}

	for i := 0; i < len(m.DownlinkPaths); i++ {
		if swag.IsZero(m.DownlinkPaths[i]) { // not required
			continue
		}

		if m.DownlinkPaths[i] != nil {
			if err := m.DownlinkPaths[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("downlink_paths" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V3TxRequest) validatePriority(formats strfmt.Registry) error {

	if swag.IsZero(m.Priority) { // not required
		return nil
	}

	if err := m.Priority.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("priority")
		}
		return err
	}

	return nil
}

func (m *V3TxRequest) validateRx1DataRateIndex(formats strfmt.Registry) error {

	if swag.IsZero(m.Rx1DataRateIndex) { // not required
		return nil
	}

	if err := m.Rx1DataRateIndex.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rx1_data_rate_index")
		}
		return err
	}

	return nil
}

func (m *V3TxRequest) validateRx1Delay(formats strfmt.Registry) error {

	if swag.IsZero(m.Rx1Delay) { // not required
		return nil
	}

	if err := m.Rx1Delay.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rx1_delay")
		}
		return err
	}

	return nil
}

func (m *V3TxRequest) validateRx2DataRateIndex(formats strfmt.Registry) error {

	if swag.IsZero(m.Rx2DataRateIndex) { // not required
		return nil
	}

	if err := m.Rx2DataRateIndex.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("rx2_data_rate_index")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3TxRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3TxRequest) UnmarshalBinary(b []byte) error {
	var res V3TxRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
