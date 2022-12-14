// Code generated by go-swagger; DO NOT EDIT.

package schedule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteScheduleParams creates a new DeleteScheduleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteScheduleParams() *DeleteScheduleParams {
	return &DeleteScheduleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteScheduleParamsWithTimeout creates a new DeleteScheduleParams object
// with the ability to set a timeout on a request.
func NewDeleteScheduleParamsWithTimeout(timeout time.Duration) *DeleteScheduleParams {
	return &DeleteScheduleParams{
		timeout: timeout,
	}
}

// NewDeleteScheduleParamsWithContext creates a new DeleteScheduleParams object
// with the ability to set a context for a request.
func NewDeleteScheduleParamsWithContext(ctx context.Context) *DeleteScheduleParams {
	return &DeleteScheduleParams{
		Context: ctx,
	}
}

// NewDeleteScheduleParamsWithHTTPClient creates a new DeleteScheduleParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteScheduleParamsWithHTTPClient(client *http.Client) *DeleteScheduleParams {
	return &DeleteScheduleParams{
		HTTPClient: client,
	}
}

/*
DeleteScheduleParams contains all the parameters to send to the API endpoint

	for the delete schedule operation.

	Typically these are written to a http.Request.
*/
type DeleteScheduleParams struct {

	/* ID.

	   The unique ID of the schedule.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteScheduleParams) WithDefaults() *DeleteScheduleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteScheduleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete schedule params
func (o *DeleteScheduleParams) WithTimeout(timeout time.Duration) *DeleteScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete schedule params
func (o *DeleteScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete schedule params
func (o *DeleteScheduleParams) WithContext(ctx context.Context) *DeleteScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete schedule params
func (o *DeleteScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete schedule params
func (o *DeleteScheduleParams) WithHTTPClient(client *http.Client) *DeleteScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete schedule params
func (o *DeleteScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete schedule params
func (o *DeleteScheduleParams) WithID(id strfmt.UUID) *DeleteScheduleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete schedule params
func (o *DeleteScheduleParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
