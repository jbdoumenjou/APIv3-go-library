// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateChildDomain update child domain
// swagger:model updateChildDomain
type UpdateChildDomain struct {

	// Value for the sender domain that will replace the existing domain
	Domain string `json:"domain,omitempty"`
}

// Validate validates this update child domain
func (m *UpdateChildDomain) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpdateChildDomain) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateChildDomain) UnmarshalBinary(b []byte) error {
	var res UpdateChildDomain
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}