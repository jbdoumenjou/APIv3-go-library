// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetChildrenListChildrenItems get children list children items
// swagger:model getChildrenListChildrenItems
type GetChildrenListChildrenItems struct {
	GetChildInfo

	GetChildrenListChildrenItemsAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *GetChildrenListChildrenItems) UnmarshalJSON(raw []byte) error {

	var aO0 GetChildInfo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.GetChildInfo = aO0

	var aO1 GetChildrenListChildrenItemsAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.GetChildrenListChildrenItemsAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m GetChildrenListChildrenItems) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	aO0, err := swag.WriteJSON(m.GetChildInfo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.GetChildrenListChildrenItemsAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get children list children items
func (m *GetChildrenListChildrenItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.GetChildInfo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.GetChildrenListChildrenItemsAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *GetChildrenListChildrenItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetChildrenListChildrenItems) UnmarshalBinary(b []byte) error {
	var res GetChildrenListChildrenItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
