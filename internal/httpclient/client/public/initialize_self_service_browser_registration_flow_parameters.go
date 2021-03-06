// Code generated by go-swagger; DO NOT EDIT.

package public

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

// NewInitializeSelfServiceBrowserRegistrationFlowParams creates a new InitializeSelfServiceBrowserRegistrationFlowParams object
// with the default values initialized.
func NewInitializeSelfServiceBrowserRegistrationFlowParams() *InitializeSelfServiceBrowserRegistrationFlowParams {

	return &InitializeSelfServiceBrowserRegistrationFlowParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInitializeSelfServiceBrowserRegistrationFlowParamsWithTimeout creates a new InitializeSelfServiceBrowserRegistrationFlowParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInitializeSelfServiceBrowserRegistrationFlowParamsWithTimeout(timeout time.Duration) *InitializeSelfServiceBrowserRegistrationFlowParams {

	return &InitializeSelfServiceBrowserRegistrationFlowParams{

		timeout: timeout,
	}
}

// NewInitializeSelfServiceBrowserRegistrationFlowParamsWithContext creates a new InitializeSelfServiceBrowserRegistrationFlowParams object
// with the default values initialized, and the ability to set a context for a request
func NewInitializeSelfServiceBrowserRegistrationFlowParamsWithContext(ctx context.Context) *InitializeSelfServiceBrowserRegistrationFlowParams {

	return &InitializeSelfServiceBrowserRegistrationFlowParams{

		Context: ctx,
	}
}

// NewInitializeSelfServiceBrowserRegistrationFlowParamsWithHTTPClient creates a new InitializeSelfServiceBrowserRegistrationFlowParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInitializeSelfServiceBrowserRegistrationFlowParamsWithHTTPClient(client *http.Client) *InitializeSelfServiceBrowserRegistrationFlowParams {

	return &InitializeSelfServiceBrowserRegistrationFlowParams{
		HTTPClient: client,
	}
}

/*InitializeSelfServiceBrowserRegistrationFlowParams contains all the parameters to send to the API endpoint
for the initialize self service browser registration flow operation typically these are written to a http.Request
*/
type InitializeSelfServiceBrowserRegistrationFlowParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the initialize self service browser registration flow params
func (o *InitializeSelfServiceBrowserRegistrationFlowParams) WithTimeout(timeout time.Duration) *InitializeSelfServiceBrowserRegistrationFlowParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the initialize self service browser registration flow params
func (o *InitializeSelfServiceBrowserRegistrationFlowParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the initialize self service browser registration flow params
func (o *InitializeSelfServiceBrowserRegistrationFlowParams) WithContext(ctx context.Context) *InitializeSelfServiceBrowserRegistrationFlowParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the initialize self service browser registration flow params
func (o *InitializeSelfServiceBrowserRegistrationFlowParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the initialize self service browser registration flow params
func (o *InitializeSelfServiceBrowserRegistrationFlowParams) WithHTTPClient(client *http.Client) *InitializeSelfServiceBrowserRegistrationFlowParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the initialize self service browser registration flow params
func (o *InitializeSelfServiceBrowserRegistrationFlowParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *InitializeSelfServiceBrowserRegistrationFlowParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
