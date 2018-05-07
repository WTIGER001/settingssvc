// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PreferenceOwner preference owner
// swagger:model PreferenceOwner
type PreferenceOwner struct {

	// active
	Active string `json:"active,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// owner type
	OwnerType string `json:"owner-type,omitempty"`

	// profile ids
	ProfileIds []string `json:"profile-ids"`
}

// Validate validates this preference owner
func (m *PreferenceOwner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProfileIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PreferenceOwner) validateProfileIds(formats strfmt.Registry) error {

	if swag.IsZero(m.ProfileIds) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PreferenceOwner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PreferenceOwner) UnmarshalBinary(b []byte) error {
	var res PreferenceOwner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
