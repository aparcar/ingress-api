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

// TxAcknowledgmentResult tx acknowledgment result
//
// swagger:model TxAcknowledgmentResult
type TxAcknowledgmentResult string

const (

	// TxAcknowledgmentResultSUCCESS captures enum value "SUCCESS"
	TxAcknowledgmentResultSUCCESS TxAcknowledgmentResult = "SUCCESS"

	// TxAcknowledgmentResultUNKNOWNERROR captures enum value "UNKNOWN_ERROR"
	TxAcknowledgmentResultUNKNOWNERROR TxAcknowledgmentResult = "UNKNOWN_ERROR"

	// TxAcknowledgmentResultTOOLATE captures enum value "TOO_LATE"
	TxAcknowledgmentResultTOOLATE TxAcknowledgmentResult = "TOO_LATE"

	// TxAcknowledgmentResultTOOEARLY captures enum value "TOO_EARLY"
	TxAcknowledgmentResultTOOEARLY TxAcknowledgmentResult = "TOO_EARLY"

	// TxAcknowledgmentResultCOLLISIONPACKET captures enum value "COLLISION_PACKET"
	TxAcknowledgmentResultCOLLISIONPACKET TxAcknowledgmentResult = "COLLISION_PACKET"

	// TxAcknowledgmentResultCOLLISIONBEACON captures enum value "COLLISION_BEACON"
	TxAcknowledgmentResultCOLLISIONBEACON TxAcknowledgmentResult = "COLLISION_BEACON"

	// TxAcknowledgmentResultTXFREQ captures enum value "TX_FREQ"
	TxAcknowledgmentResultTXFREQ TxAcknowledgmentResult = "TX_FREQ"

	// TxAcknowledgmentResultTXPOWER captures enum value "TX_POWER"
	TxAcknowledgmentResultTXPOWER TxAcknowledgmentResult = "TX_POWER"

	// TxAcknowledgmentResultGPSUNLOCKED captures enum value "GPS_UNLOCKED"
	TxAcknowledgmentResultGPSUNLOCKED TxAcknowledgmentResult = "GPS_UNLOCKED"
)

// for schema
var txAcknowledgmentResultEnum []interface{}

func init() {
	var res []TxAcknowledgmentResult
	if err := json.Unmarshal([]byte(`["SUCCESS","UNKNOWN_ERROR","TOO_LATE","TOO_EARLY","COLLISION_PACKET","COLLISION_BEACON","TX_FREQ","TX_POWER","GPS_UNLOCKED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		txAcknowledgmentResultEnum = append(txAcknowledgmentResultEnum, v)
	}
}

func (m TxAcknowledgmentResult) validateTxAcknowledgmentResultEnum(path, location string, value TxAcknowledgmentResult) error {
	if err := validate.Enum(path, location, value, txAcknowledgmentResultEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this tx acknowledgment result
func (m TxAcknowledgmentResult) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTxAcknowledgmentResultEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
