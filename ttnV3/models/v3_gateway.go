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

// V3Gateway Gateway is the message that defines a gateway on the network.
//
// swagger:model v3Gateway
type V3Gateway struct {

	// antennas
	Antennas []*V3GatewayAntenna `json:"antennas"`

	// attributes
	Attributes map[string]string `json:"attributes,omitempty"`

	// auto update
	AutoUpdate bool `json:"auto_update,omitempty"`

	// contact info
	ContactInfo []*V3ContactInfo `json:"contact_info"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"created_at,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// downlink path constraint
	DownlinkPathConstraint V3DownlinkPathConstraint `json:"downlink_path_constraint,omitempty"`

	// Enforcing gateway duty cycle is recommended for all gateways to respect spectrum regulations. Disable enforcing the
	// duty cycle only in controlled research and development environments.
	EnforceDutyCycle bool `json:"enforce_duty_cycle,omitempty"`

	// Frequency plan ID of the gateway.
	// This equals the first element of the frequency_plan_ids field.
	FrequencyPlanID string `json:"frequency_plan_id,omitempty"`

	// Frequency plan IDs of the gateway.
	// The first element equals the frequency_plan_id field.
	FrequencyPlanIds []string `json:"frequency_plan_ids"`

	// The address of the Gateway Server to connect to.
	// The typical format of the address is "host:port". If the port is omitted,
	// the normal port inference (with DNS lookup, otherwise defaults) is used.
	// The connection shall be established with transport layer security (TLS).
	// Custom certificate authorities may be configured out-of-band.
	GatewayServerAddress string `json:"gateway_server_address,omitempty"`

	// ids
	Ids *V3GatewayIdentifiers `json:"ids,omitempty"`

	// The location of this gateway may be publicly displayed.
	LocationPublic bool `json:"location_public,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Adjust the time that GS schedules class C messages in advance. This is useful for gateways that have a known high latency backhaul, like 3G and satellite.
	ScheduleAnytimeDelay string `json:"schedule_anytime_delay,omitempty"`

	// Enable server-side buffering of downlink messages. This is recommended for gateways using the Semtech UDP Packet
	// Forwarder v2.x or older, as it does not feature a just-in-time queue. If enabled, the Gateway Server schedules the
	// downlink message late to the gateway so that it does not overwrite previously scheduled downlink messages that have
	// not been transmitted yet.
	ScheduleDownlinkLate bool `json:"schedule_downlink_late,omitempty"`

	// The status of this gateway may be publicly displayed.
	StatusPublic bool `json:"status_public,omitempty"`

	// update channel
	UpdateChannel string `json:"update_channel,omitempty"`

	// update the location of this gateway from status messages
	UpdateLocationFromStatus bool `json:"update_location_from_status,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updated_at,omitempty"`

	// version ids
	VersionIds *V3GatewayVersionIdentifiers `json:"version_ids,omitempty"`
}

// Validate validates this v3 gateway
func (m *V3Gateway) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAntennas(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContactInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDownlinkPathConstraint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V3Gateway) validateAntennas(formats strfmt.Registry) error {

	if swag.IsZero(m.Antennas) { // not required
		return nil
	}

	for i := 0; i < len(m.Antennas); i++ {
		if swag.IsZero(m.Antennas[i]) { // not required
			continue
		}

		if m.Antennas[i] != nil {
			if err := m.Antennas[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("antennas" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V3Gateway) validateContactInfo(formats strfmt.Registry) error {

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

func (m *V3Gateway) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("created_at", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V3Gateway) validateDownlinkPathConstraint(formats strfmt.Registry) error {

	if swag.IsZero(m.DownlinkPathConstraint) { // not required
		return nil
	}

	if err := m.DownlinkPathConstraint.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("downlink_path_constraint")
		}
		return err
	}

	return nil
}

func (m *V3Gateway) validateIds(formats strfmt.Registry) error {

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

func (m *V3Gateway) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updated_at", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V3Gateway) validateVersionIds(formats strfmt.Registry) error {

	if swag.IsZero(m.VersionIds) { // not required
		return nil
	}

	if m.VersionIds != nil {
		if err := m.VersionIds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version_ids")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V3Gateway) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V3Gateway) UnmarshalBinary(b []byte) error {
	var res V3Gateway
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
