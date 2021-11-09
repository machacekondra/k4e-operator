// Code generated by go-swagger; DO NOT EDIT.

package yggdrasil

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

// NewGetControlMessageForDeviceParams creates a new GetControlMessageForDeviceParams object
// with the default values initialized.
func NewGetControlMessageForDeviceParams() *GetControlMessageForDeviceParams {
	var ()
	return &GetControlMessageForDeviceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetControlMessageForDeviceParamsWithTimeout creates a new GetControlMessageForDeviceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetControlMessageForDeviceParamsWithTimeout(timeout time.Duration) *GetControlMessageForDeviceParams {
	var ()
	return &GetControlMessageForDeviceParams{

		timeout: timeout,
	}
}

// NewGetControlMessageForDeviceParamsWithContext creates a new GetControlMessageForDeviceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetControlMessageForDeviceParamsWithContext(ctx context.Context) *GetControlMessageForDeviceParams {
	var ()
	return &GetControlMessageForDeviceParams{

		Context: ctx,
	}
}

// NewGetControlMessageForDeviceParamsWithHTTPClient creates a new GetControlMessageForDeviceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetControlMessageForDeviceParamsWithHTTPClient(client *http.Client) *GetControlMessageForDeviceParams {
	var ()
	return &GetControlMessageForDeviceParams{
		HTTPClient: client,
	}
}

/*GetControlMessageForDeviceParams contains all the parameters to send to the API endpoint
for the get control message for device operation typically these are written to a http.Request
*/
type GetControlMessageForDeviceParams struct {

	/*DeviceID
	  Device ID

	*/
	DeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get control message for device params
func (o *GetControlMessageForDeviceParams) WithTimeout(timeout time.Duration) *GetControlMessageForDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get control message for device params
func (o *GetControlMessageForDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get control message for device params
func (o *GetControlMessageForDeviceParams) WithContext(ctx context.Context) *GetControlMessageForDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get control message for device params
func (o *GetControlMessageForDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get control message for device params
func (o *GetControlMessageForDeviceParams) WithHTTPClient(client *http.Client) *GetControlMessageForDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get control message for device params
func (o *GetControlMessageForDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get control message for device params
func (o *GetControlMessageForDeviceParams) WithDeviceID(deviceID string) *GetControlMessageForDeviceParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get control message for device params
func (o *GetControlMessageForDeviceParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetControlMessageForDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param device_id
	if err := r.SetPathParam("device_id", o.DeviceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}