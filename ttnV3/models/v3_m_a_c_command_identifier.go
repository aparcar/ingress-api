// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V3MACCommandIdentifier v3 m a c command identifier
//
// swagger:model v3MACCommandIdentifier
type V3MACCommandIdentifier string

const (

	// V3MACCommandIdentifierCIDRFU0 captures enum value "CID_RFU_0"
	V3MACCommandIdentifierCIDRFU0 V3MACCommandIdentifier = "CID_RFU_0"

	// V3MACCommandIdentifierCIDRESET captures enum value "CID_RESET"
	V3MACCommandIdentifierCIDRESET V3MACCommandIdentifier = "CID_RESET"

	// V3MACCommandIdentifierCIDLINKCHECK captures enum value "CID_LINK_CHECK"
	V3MACCommandIdentifierCIDLINKCHECK V3MACCommandIdentifier = "CID_LINK_CHECK"

	// V3MACCommandIdentifierCIDLINKADR captures enum value "CID_LINK_ADR"
	V3MACCommandIdentifierCIDLINKADR V3MACCommandIdentifier = "CID_LINK_ADR"

	// V3MACCommandIdentifierCIDDUTYCYCLE captures enum value "CID_DUTY_CYCLE"
	V3MACCommandIdentifierCIDDUTYCYCLE V3MACCommandIdentifier = "CID_DUTY_CYCLE"

	// V3MACCommandIdentifierCIDRXPARAMSETUP captures enum value "CID_RX_PARAM_SETUP"
	V3MACCommandIdentifierCIDRXPARAMSETUP V3MACCommandIdentifier = "CID_RX_PARAM_SETUP"

	// V3MACCommandIdentifierCIDDEVSTATUS captures enum value "CID_DEV_STATUS"
	V3MACCommandIdentifierCIDDEVSTATUS V3MACCommandIdentifier = "CID_DEV_STATUS"

	// V3MACCommandIdentifierCIDNEWCHANNEL captures enum value "CID_NEW_CHANNEL"
	V3MACCommandIdentifierCIDNEWCHANNEL V3MACCommandIdentifier = "CID_NEW_CHANNEL"

	// V3MACCommandIdentifierCIDRXTIMINGSETUP captures enum value "CID_RX_TIMING_SETUP"
	V3MACCommandIdentifierCIDRXTIMINGSETUP V3MACCommandIdentifier = "CID_RX_TIMING_SETUP"

	// V3MACCommandIdentifierCIDTXPARAMSETUP captures enum value "CID_TX_PARAM_SETUP"
	V3MACCommandIdentifierCIDTXPARAMSETUP V3MACCommandIdentifier = "CID_TX_PARAM_SETUP"

	// V3MACCommandIdentifierCIDDLCHANNEL captures enum value "CID_DL_CHANNEL"
	V3MACCommandIdentifierCIDDLCHANNEL V3MACCommandIdentifier = "CID_DL_CHANNEL"

	// V3MACCommandIdentifierCIDREKEY captures enum value "CID_REKEY"
	V3MACCommandIdentifierCIDREKEY V3MACCommandIdentifier = "CID_REKEY"

	// V3MACCommandIdentifierCIDADRPARAMSETUP captures enum value "CID_ADR_PARAM_SETUP"
	V3MACCommandIdentifierCIDADRPARAMSETUP V3MACCommandIdentifier = "CID_ADR_PARAM_SETUP"

	// V3MACCommandIdentifierCIDDEVICETIME captures enum value "CID_DEVICE_TIME"
	V3MACCommandIdentifierCIDDEVICETIME V3MACCommandIdentifier = "CID_DEVICE_TIME"

	// V3MACCommandIdentifierCIDFORCEREJOIN captures enum value "CID_FORCE_REJOIN"
	V3MACCommandIdentifierCIDFORCEREJOIN V3MACCommandIdentifier = "CID_FORCE_REJOIN"

	// V3MACCommandIdentifierCIDREJOINPARAMSETUP captures enum value "CID_REJOIN_PARAM_SETUP"
	V3MACCommandIdentifierCIDREJOINPARAMSETUP V3MACCommandIdentifier = "CID_REJOIN_PARAM_SETUP"

	// V3MACCommandIdentifierCIDPINGSLOTINFO captures enum value "CID_PING_SLOT_INFO"
	V3MACCommandIdentifierCIDPINGSLOTINFO V3MACCommandIdentifier = "CID_PING_SLOT_INFO"

	// V3MACCommandIdentifierCIDPINGSLOTCHANNEL captures enum value "CID_PING_SLOT_CHANNEL"
	V3MACCommandIdentifierCIDPINGSLOTCHANNEL V3MACCommandIdentifier = "CID_PING_SLOT_CHANNEL"

	// V3MACCommandIdentifierCIDBEACONTIMING captures enum value "CID_BEACON_TIMING"
	V3MACCommandIdentifierCIDBEACONTIMING V3MACCommandIdentifier = "CID_BEACON_TIMING"

	// V3MACCommandIdentifierCIDBEACONFREQ captures enum value "CID_BEACON_FREQ"
	V3MACCommandIdentifierCIDBEACONFREQ V3MACCommandIdentifier = "CID_BEACON_FREQ"

	// V3MACCommandIdentifierCIDDEVICEMODE captures enum value "CID_DEVICE_MODE"
	V3MACCommandIdentifierCIDDEVICEMODE V3MACCommandIdentifier = "CID_DEVICE_MODE"
)

// for schema
var v3MACCommandIdentifierEnum []interface{}

func init() {
	var res []V3MACCommandIdentifier
	if err := json.Unmarshal([]byte(`["CID_RFU_0","CID_RESET","CID_LINK_CHECK","CID_LINK_ADR","CID_DUTY_CYCLE","CID_RX_PARAM_SETUP","CID_DEV_STATUS","CID_NEW_CHANNEL","CID_RX_TIMING_SETUP","CID_TX_PARAM_SETUP","CID_DL_CHANNEL","CID_REKEY","CID_ADR_PARAM_SETUP","CID_DEVICE_TIME","CID_FORCE_REJOIN","CID_REJOIN_PARAM_SETUP","CID_PING_SLOT_INFO","CID_PING_SLOT_CHANNEL","CID_BEACON_TIMING","CID_BEACON_FREQ","CID_DEVICE_MODE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v3MACCommandIdentifierEnum = append(v3MACCommandIdentifierEnum, v)
	}
}

func (m V3MACCommandIdentifier) validateV3MACCommandIdentifierEnum(path, location string, value V3MACCommandIdentifier) error {
	if err := validate.Enum(path, location, value, v3MACCommandIdentifierEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v3 m a c command identifier
func (m V3MACCommandIdentifier) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV3MACCommandIdentifierEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
