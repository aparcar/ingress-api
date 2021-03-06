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

// V3LocationSource  - SOURCE_UNKNOWN: The source of the location is not known or not set.
//  - SOURCE_GPS: The location is determined by GPS.
//  - SOURCE_REGISTRY: The location is set in and updated from a registry.
//  - SOURCE_IP_GEOLOCATION: The location is estimated with IP geolocation.
//  - SOURCE_WIFI_RSSI_GEOLOCATION: The location is estimated with WiFi RSSI geolocation.
//  - SOURCE_BT_RSSI_GEOLOCATION: The location is estimated with BT/BLE RSSI geolocation.
//  - SOURCE_LORA_RSSI_GEOLOCATION: The location is estimated with LoRa RSSI geolocation.
//  - SOURCE_LORA_TDOA_GEOLOCATION: The location is estimated with LoRa TDOA geolocation.
//  - SOURCE_COMBINED_GEOLOCATION: The location is estimated by a combination of geolocation sources.
//
// swagger:model v3LocationSource
type V3LocationSource string

const (

	// V3LocationSourceSOURCEUNKNOWN captures enum value "SOURCE_UNKNOWN"
	V3LocationSourceSOURCEUNKNOWN V3LocationSource = "SOURCE_UNKNOWN"

	// V3LocationSourceSOURCEGPS captures enum value "SOURCE_GPS"
	V3LocationSourceSOURCEGPS V3LocationSource = "SOURCE_GPS"

	// V3LocationSourceSOURCEREGISTRY captures enum value "SOURCE_REGISTRY"
	V3LocationSourceSOURCEREGISTRY V3LocationSource = "SOURCE_REGISTRY"

	// V3LocationSourceSOURCEIPGEOLOCATION captures enum value "SOURCE_IP_GEOLOCATION"
	V3LocationSourceSOURCEIPGEOLOCATION V3LocationSource = "SOURCE_IP_GEOLOCATION"

	// V3LocationSourceSOURCEWIFIRSSIGEOLOCATION captures enum value "SOURCE_WIFI_RSSI_GEOLOCATION"
	V3LocationSourceSOURCEWIFIRSSIGEOLOCATION V3LocationSource = "SOURCE_WIFI_RSSI_GEOLOCATION"

	// V3LocationSourceSOURCEBTRSSIGEOLOCATION captures enum value "SOURCE_BT_RSSI_GEOLOCATION"
	V3LocationSourceSOURCEBTRSSIGEOLOCATION V3LocationSource = "SOURCE_BT_RSSI_GEOLOCATION"

	// V3LocationSourceSOURCELORARSSIGEOLOCATION captures enum value "SOURCE_LORA_RSSI_GEOLOCATION"
	V3LocationSourceSOURCELORARSSIGEOLOCATION V3LocationSource = "SOURCE_LORA_RSSI_GEOLOCATION"

	// V3LocationSourceSOURCELORATDOAGEOLOCATION captures enum value "SOURCE_LORA_TDOA_GEOLOCATION"
	V3LocationSourceSOURCELORATDOAGEOLOCATION V3LocationSource = "SOURCE_LORA_TDOA_GEOLOCATION"

	// V3LocationSourceSOURCECOMBINEDGEOLOCATION captures enum value "SOURCE_COMBINED_GEOLOCATION"
	V3LocationSourceSOURCECOMBINEDGEOLOCATION V3LocationSource = "SOURCE_COMBINED_GEOLOCATION"
)

// for schema
var v3LocationSourceEnum []interface{}

func init() {
	var res []V3LocationSource
	if err := json.Unmarshal([]byte(`["SOURCE_UNKNOWN","SOURCE_GPS","SOURCE_REGISTRY","SOURCE_IP_GEOLOCATION","SOURCE_WIFI_RSSI_GEOLOCATION","SOURCE_BT_RSSI_GEOLOCATION","SOURCE_LORA_RSSI_GEOLOCATION","SOURCE_LORA_TDOA_GEOLOCATION","SOURCE_COMBINED_GEOLOCATION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v3LocationSourceEnum = append(v3LocationSourceEnum, v)
	}
}

func (m V3LocationSource) validateV3LocationSourceEnum(path, location string, value V3LocationSource) error {
	if err := validate.Enum(path, location, value, v3LocationSourceEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v3 location source
func (m V3LocationSource) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV3LocationSourceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
