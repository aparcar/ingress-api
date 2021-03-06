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

// MQTTProviderQoS m q t t provider qo s
//
// swagger:model MQTTProviderQoS
type MQTTProviderQoS string

const (

	// MQTTProviderQoSATMOSTONCE captures enum value "AT_MOST_ONCE"
	MQTTProviderQoSATMOSTONCE MQTTProviderQoS = "AT_MOST_ONCE"

	// MQTTProviderQoSATLEASTONCE captures enum value "AT_LEAST_ONCE"
	MQTTProviderQoSATLEASTONCE MQTTProviderQoS = "AT_LEAST_ONCE"

	// MQTTProviderQoSEXACTLYONCE captures enum value "EXACTLY_ONCE"
	MQTTProviderQoSEXACTLYONCE MQTTProviderQoS = "EXACTLY_ONCE"
)

// for schema
var mQTTProviderQoSEnum []interface{}

func init() {
	var res []MQTTProviderQoS
	if err := json.Unmarshal([]byte(`["AT_MOST_ONCE","AT_LEAST_ONCE","EXACTLY_ONCE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mQTTProviderQoSEnum = append(mQTTProviderQoSEnum, v)
	}
}

func (m MQTTProviderQoS) validateMQTTProviderQoSEnum(path, location string, value MQTTProviderQoS) error {
	if err := validate.Enum(path, location, value, mQTTProviderQoSEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this m q t t provider qo s
func (m MQTTProviderQoS) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateMQTTProviderQoSEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
