// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetChildInfoAllOf1Statistics Statistics about your child account activity
// swagger:model getChildInfoAllOf1Statistics
type GetChildInfoAllOf1Statistics struct {

	// Overall emails sent for current month
	CurrentMonthTotalSent int64 `json:"currentMonthTotalSent,omitempty"`

	// Overall emails sent for the previous month
	PreviousMonthTotalSent int64 `json:"previousMonthTotalSent,omitempty"`

	// Overall emails sent for since the account exists
	TotalSent int64 `json:"totalSent,omitempty"`
}

// Validate validates this get child info all of1 statistics
func (m *GetChildInfoAllOf1Statistics) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetChildInfoAllOf1Statistics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetChildInfoAllOf1Statistics) UnmarshalBinary(b []byte) error {
	var res GetChildInfoAllOf1Statistics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
