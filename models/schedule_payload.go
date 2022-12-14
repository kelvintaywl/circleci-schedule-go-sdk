// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SchedulePayload schedule payload
//
// swagger:model SchedulePayload
type SchedulePayload struct {
	ScheduleBaseData

	// The attribution-actor of the scheduled pipeline.
	// Enum: [current system]
	AttributionActor string `json:"attribution-actor,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SchedulePayload) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ScheduleBaseData
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ScheduleBaseData = aO0

	// AO1
	var dataAO1 struct {
		AttributionActor string `json:"attribution-actor,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.AttributionActor = dataAO1.AttributionActor

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SchedulePayload) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ScheduleBaseData)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		AttributionActor string `json:"attribution-actor,omitempty"`
	}

	dataAO1.AttributionActor = m.AttributionActor

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this schedule payload
func (m *SchedulePayload) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ScheduleBaseData
	if err := m.ScheduleBaseData.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttributionActor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var schedulePayloadTypeAttributionActorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["current","system"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		schedulePayloadTypeAttributionActorPropEnum = append(schedulePayloadTypeAttributionActorPropEnum, v)
	}
}

// property enum
func (m *SchedulePayload) validateAttributionActorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, schedulePayloadTypeAttributionActorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SchedulePayload) validateAttributionActor(formats strfmt.Registry) error {

	if swag.IsZero(m.AttributionActor) { // not required
		return nil
	}

	// value enum
	if err := m.validateAttributionActorEnum("attribution-actor", "body", m.AttributionActor); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this schedule payload based on the context it is used
func (m *SchedulePayload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ScheduleBaseData
	if err := m.ScheduleBaseData.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SchedulePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchedulePayload) UnmarshalBinary(b []byte) error {
	var res SchedulePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
