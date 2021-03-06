// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetContactCampaignStats Campaign Statistics for the contact
// swagger:model getContactCampaignStats
type GetContactCampaignStats struct {

	// clicked
	Clicked GetContactCampaignStatsClicked `json:"clicked"`

	// complaints
	Complaints GetContactCampaignStatsComplaints `json:"complaints"`

	// hard bounces
	HardBounces GetContactCampaignStatsHardBounces `json:"hardBounces"`

	// messages sent
	MessagesSent GetContactCampaignStatsMessagesSent `json:"messagesSent"`

	// opened
	Opened GetContactCampaignStatsOpened `json:"opened"`

	// soft bounces
	SoftBounces GetContactCampaignStatsSoftBounces `json:"softBounces"`

	// transac attributes
	TransacAttributes GetContactCampaignStatsTransacAttributes `json:"transacAttributes"`

	// unsubscriptions
	Unsubscriptions *GetContactCampaignStatsUnsubscriptions `json:"unsubscriptions,omitempty"`
}

// Validate validates this get contact campaign stats
func (m *GetContactCampaignStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUnsubscriptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetContactCampaignStats) validateUnsubscriptions(formats strfmt.Registry) error {

	if swag.IsZero(m.Unsubscriptions) { // not required
		return nil
	}

	if m.Unsubscriptions != nil {

		if err := m.Unsubscriptions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unsubscriptions")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetContactCampaignStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetContactCampaignStats) UnmarshalBinary(b []byte) error {
	var res GetContactCampaignStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
