// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeviceConfigurationMessage device configuration message
//
// swagger:model device-configuration-message
type DeviceConfigurationMessage struct {

	// configuration
	Configuration *DeviceConfiguration `json:"configuration,omitempty"`

	// Device identifier
	DeviceID string `json:"device_id,omitempty"`

	// version
	Version string `json:"version,omitempty"`

	// List of workloads deployed to the device
	Workloads WorkloadList `json:"workloads,omitempty"`

	// Defines the interval in seconds between the attempts to evaluate the workloads status and restart those that failed
	// Minimum: > 0
	WorkloadsMonitoringInterval int64 `json:"workloads_monitoring_interval,omitempty"`
}

// Validate validates this device configuration message
func (m *DeviceConfigurationMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloads(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkloadsMonitoringInterval(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeviceConfigurationMessage) validateConfiguration(formats strfmt.Registry) error {

	if swag.IsZero(m.Configuration) { // not required
		return nil
	}

	if m.Configuration != nil {
		if err := m.Configuration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("configuration")
			}
			return err
		}
	}

	return nil
}

func (m *DeviceConfigurationMessage) validateWorkloads(formats strfmt.Registry) error {

	if swag.IsZero(m.Workloads) { // not required
		return nil
	}

	if err := m.Workloads.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("workloads")
		}
		return err
	}

	return nil
}

func (m *DeviceConfigurationMessage) validateWorkloadsMonitoringInterval(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkloadsMonitoringInterval) { // not required
		return nil
	}

	if err := validate.MinimumInt("workloads_monitoring_interval", "body", int64(m.WorkloadsMonitoringInterval), 0, true); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeviceConfigurationMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceConfigurationMessage) UnmarshalBinary(b []byte) error {
	var res DeviceConfigurationMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}