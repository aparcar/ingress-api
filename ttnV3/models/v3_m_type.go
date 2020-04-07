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

// V3MType v3 m type
//
// swagger:model v3MType
type V3MType string

const (

	// V3MTypeJOINREQUEST captures enum value "JOIN_REQUEST"
	V3MTypeJOINREQUEST V3MType = "JOIN_REQUEST"

	// V3MTypeJOINACCEPT captures enum value "JOIN_ACCEPT"
	V3MTypeJOINACCEPT V3MType = "JOIN_ACCEPT"

	// V3MTypeUNCONFIRMEDUP captures enum value "UNCONFIRMED_UP"
	V3MTypeUNCONFIRMEDUP V3MType = "UNCONFIRMED_UP"

	// V3MTypeUNCONFIRMEDDOWN captures enum value "UNCONFIRMED_DOWN"
	V3MTypeUNCONFIRMEDDOWN V3MType = "UNCONFIRMED_DOWN"

	// V3MTypeCONFIRMEDUP captures enum value "CONFIRMED_UP"
	V3MTypeCONFIRMEDUP V3MType = "CONFIRMED_UP"

	// V3MTypeCONFIRMEDDOWN captures enum value "CONFIRMED_DOWN"
	V3MTypeCONFIRMEDDOWN V3MType = "CONFIRMED_DOWN"

	// V3MTypeREJOINREQUEST captures enum value "REJOIN_REQUEST"
	V3MTypeREJOINREQUEST V3MType = "REJOIN_REQUEST"

	// V3MTypePROPRIETARY captures enum value "PROPRIETARY"
	V3MTypePROPRIETARY V3MType = "PROPRIETARY"
)

// for schema
var v3MTypeEnum []interface{}

func init() {
	var res []V3MType
	if err := json.Unmarshal([]byte(`["JOIN_REQUEST","JOIN_ACCEPT","UNCONFIRMED_UP","UNCONFIRMED_DOWN","CONFIRMED_UP","CONFIRMED_DOWN","REJOIN_REQUEST","PROPRIETARY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v3MTypeEnum = append(v3MTypeEnum, v)
	}
}

func (m V3MType) validateV3MTypeEnum(path, location string, value V3MType) error {
	if err := validate.Enum(path, location, value, v3MTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v3 m type
func (m V3MType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV3MTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
