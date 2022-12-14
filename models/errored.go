// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Errored errored
//
// swagger:model Errored
type Errored struct {

	// Error message
	Message string `json:"message,omitempty"`
}

// Validate validates this errored
func (m *Errored) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this errored based on context it is used
func (m *Errored) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Errored) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Errored) UnmarshalBinary(b []byte) error {
	var res Errored
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
