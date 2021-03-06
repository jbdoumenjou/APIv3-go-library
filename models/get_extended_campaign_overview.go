// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetExtendedCampaignOverview get extended campaign overview
// swagger:model getExtendedCampaignOverview
type GetExtendedCampaignOverview struct {
	GetCampaignOverview

	GetExtendedCampaignOverviewAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetExtendedCampaignOverview) UnmarshalJSON(raw []byte) error {

	var aO0 GetCampaignOverview
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.GetCampaignOverview = aO0

	var aO1 GetExtendedCampaignOverviewAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.GetExtendedCampaignOverviewAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetExtendedCampaignOverview) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.GetCampaignOverview)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.GetExtendedCampaignOverviewAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get extended campaign overview
func (m *GetExtendedCampaignOverview) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.GetCampaignOverview.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.GetExtendedCampaignOverviewAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetExtendedCampaignOverview) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetExtendedCampaignOverview) UnmarshalBinary(b []byte) error {
	var res GetExtendedCampaignOverview
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
