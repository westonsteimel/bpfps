// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IPAMResponse IPAM configuration of an endpoint
// swagger:model IPAMResponse

type IPAMResponse struct {

	// address
	// Required: true
	Address *AddressPair `json:"address"`

	// host addressing
	// Required: true
	HostAddressing *NodeAddressing `json:"host-addressing"`
}

/* polymorph IPAMResponse address false */

/* polymorph IPAMResponse host-addressing false */

// Validate validates this IP a m response
func (m *IPAMResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHostAddressing(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPAMResponse) validateAddress(formats strfmt.Registry) error {

	if err := validate.Required("address", "body", m.Address); err != nil {
		return err
	}

	if m.Address != nil {

		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *IPAMResponse) validateHostAddressing(formats strfmt.Registry) error {

	if err := validate.Required("host-addressing", "body", m.HostAddressing); err != nil {
		return err
	}

	if m.HostAddressing != nil {

		if err := m.HostAddressing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host-addressing")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPAMResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPAMResponse) UnmarshalBinary(b []byte) error {
	var res IPAMResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
