// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"

	"encoding/json"
)

// Profile profile
// swagger:model Profile
type Profile struct {

	// id
	ID string `json:"id,omitempty"`

	// preferences
	Preferences ProfilePreferences `json:"preferences"`

	// version
	Version int64 `json:"version,omitempty"`

	//HACK
	JsonData map[string]*json.RawMessage `json:"-"`
}

// Validate validates this profile
func (m *Profile) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Profile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Profile) UnmarshalBinary(b []byte) error {
	var res Profile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
