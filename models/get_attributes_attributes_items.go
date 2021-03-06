// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetAttributesAttributesItems get attributes attributes items
// swagger:model getAttributesAttributesItems
type GetAttributesAttributesItems struct {

	// Calculated value formula
	CalculatedValue string `json:"calculatedValue,omitempty"`

	// Category of the attribute
	// Required: true
	Category *string `json:"category"`

	// enumeration
	Enumeration GetAttributesAttributesItemsEnumeration `json:"enumeration"`

	// Name of the attribute
	// Required: true
	Name *string `json:"name"`

	// Type of the attribute
	Type string `json:"type,omitempty"`
}

// Validate validates this get attributes attributes items
func (m *GetAttributesAttributesItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getAttributesAttributesItemsTypeCategoryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["normal","transactional","category","calculated","global"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getAttributesAttributesItemsTypeCategoryPropEnum = append(getAttributesAttributesItemsTypeCategoryPropEnum, v)
	}
}

const (
	// GetAttributesAttributesItemsCategoryNormal captures enum value "normal"
	GetAttributesAttributesItemsCategoryNormal string = "normal"
	// GetAttributesAttributesItemsCategoryTransactional captures enum value "transactional"
	GetAttributesAttributesItemsCategoryTransactional string = "transactional"
	// GetAttributesAttributesItemsCategoryCategory captures enum value "category"
	GetAttributesAttributesItemsCategoryCategory string = "category"
	// GetAttributesAttributesItemsCategoryCalculated captures enum value "calculated"
	GetAttributesAttributesItemsCategoryCalculated string = "calculated"
	// GetAttributesAttributesItemsCategoryGlobal captures enum value "global"
	GetAttributesAttributesItemsCategoryGlobal string = "global"
)

// prop value enum
func (m *GetAttributesAttributesItems) validateCategoryEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getAttributesAttributesItemsTypeCategoryPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetAttributesAttributesItems) validateCategory(formats strfmt.Registry) error {

	if err := validate.Required("category", "body", m.Category); err != nil {
		return err
	}

	// value enum
	if err := m.validateCategoryEnum("category", "body", *m.Category); err != nil {
		return err
	}

	return nil
}

func (m *GetAttributesAttributesItems) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var getAttributesAttributesItemsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["text","date","float","id"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getAttributesAttributesItemsTypeTypePropEnum = append(getAttributesAttributesItemsTypeTypePropEnum, v)
	}
}

const (
	// GetAttributesAttributesItemsTypeText captures enum value "text"
	GetAttributesAttributesItemsTypeText string = "text"
	// GetAttributesAttributesItemsTypeDate captures enum value "date"
	GetAttributesAttributesItemsTypeDate string = "date"
	// GetAttributesAttributesItemsTypeFloat captures enum value "float"
	GetAttributesAttributesItemsTypeFloat string = "float"
	// GetAttributesAttributesItemsTypeID captures enum value "id"
	GetAttributesAttributesItemsTypeID string = "id"
)

// prop value enum
func (m *GetAttributesAttributesItems) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getAttributesAttributesItemsTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetAttributesAttributesItems) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetAttributesAttributesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetAttributesAttributesItems) UnmarshalBinary(b []byte) error {
	var res GetAttributesAttributesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
