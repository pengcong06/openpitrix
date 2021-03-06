// Code generated by go-swagger; DO NOT EDIT.

package release_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDescribeReleaseParams creates a new DescribeReleaseParams object
// with the default values initialized.
func NewDescribeReleaseParams() *DescribeReleaseParams {
	var ()
	return &DescribeReleaseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeReleaseParamsWithTimeout creates a new DescribeReleaseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeReleaseParamsWithTimeout(timeout time.Duration) *DescribeReleaseParams {
	var ()
	return &DescribeReleaseParams{

		timeout: timeout,
	}
}

// NewDescribeReleaseParamsWithContext creates a new DescribeReleaseParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeReleaseParamsWithContext(ctx context.Context) *DescribeReleaseParams {
	var ()
	return &DescribeReleaseParams{

		Context: ctx,
	}
}

// NewDescribeReleaseParamsWithHTTPClient creates a new DescribeReleaseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeReleaseParamsWithHTTPClient(client *http.Client) *DescribeReleaseParams {
	var ()
	return &DescribeReleaseParams{
		HTTPClient: client,
	}
}

/*DescribeReleaseParams contains all the parameters to send to the API endpoint
for the describe release operation typically these are written to a http.Request
*/
type DescribeReleaseParams struct {

	/*ReleaseName*/
	ReleaseName *string
	/*RuntimeID*/
	RuntimeID *string
	/*Status*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe release params
func (o *DescribeReleaseParams) WithTimeout(timeout time.Duration) *DescribeReleaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe release params
func (o *DescribeReleaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe release params
func (o *DescribeReleaseParams) WithContext(ctx context.Context) *DescribeReleaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe release params
func (o *DescribeReleaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe release params
func (o *DescribeReleaseParams) WithHTTPClient(client *http.Client) *DescribeReleaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe release params
func (o *DescribeReleaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReleaseName adds the releaseName to the describe release params
func (o *DescribeReleaseParams) WithReleaseName(releaseName *string) *DescribeReleaseParams {
	o.SetReleaseName(releaseName)
	return o
}

// SetReleaseName adds the releaseName to the describe release params
func (o *DescribeReleaseParams) SetReleaseName(releaseName *string) {
	o.ReleaseName = releaseName
}

// WithRuntimeID adds the runtimeID to the describe release params
func (o *DescribeReleaseParams) WithRuntimeID(runtimeID *string) *DescribeReleaseParams {
	o.SetRuntimeID(runtimeID)
	return o
}

// SetRuntimeID adds the runtimeId to the describe release params
func (o *DescribeReleaseParams) SetRuntimeID(runtimeID *string) {
	o.RuntimeID = runtimeID
}

// WithStatus adds the status to the describe release params
func (o *DescribeReleaseParams) WithStatus(status *string) *DescribeReleaseParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe release params
func (o *DescribeReleaseParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeReleaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ReleaseName != nil {

		// query param release_name
		var qrReleaseName string
		if o.ReleaseName != nil {
			qrReleaseName = *o.ReleaseName
		}
		qReleaseName := qrReleaseName
		if qReleaseName != "" {
			if err := r.SetQueryParam("release_name", qReleaseName); err != nil {
				return err
			}
		}

	}

	if o.RuntimeID != nil {

		// query param runtime_id
		var qrRuntimeID string
		if o.RuntimeID != nil {
			qrRuntimeID = *o.RuntimeID
		}
		qRuntimeID := qrRuntimeID
		if qRuntimeID != "" {
			if err := r.SetQueryParam("runtime_id", qRuntimeID); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
