// Code generated by go-swagger; DO NOT EDIT.

package app_manager

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

	"openpitrix.io/openpitrix/test/models"
)

// NewKsAdminRejectAppVersionParams creates a new KsAdminRejectAppVersionParams object
// with the default values initialized.
func NewKsAdminRejectAppVersionParams() *KsAdminRejectAppVersionParams {
	var ()
	return &KsAdminRejectAppVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewKsAdminRejectAppVersionParamsWithTimeout creates a new KsAdminRejectAppVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewKsAdminRejectAppVersionParamsWithTimeout(timeout time.Duration) *KsAdminRejectAppVersionParams {
	var ()
	return &KsAdminRejectAppVersionParams{

		timeout: timeout,
	}
}

// NewKsAdminRejectAppVersionParamsWithContext creates a new KsAdminRejectAppVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewKsAdminRejectAppVersionParamsWithContext(ctx context.Context) *KsAdminRejectAppVersionParams {
	var ()
	return &KsAdminRejectAppVersionParams{

		Context: ctx,
	}
}

// NewKsAdminRejectAppVersionParamsWithHTTPClient creates a new KsAdminRejectAppVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewKsAdminRejectAppVersionParamsWithHTTPClient(client *http.Client) *KsAdminRejectAppVersionParams {
	var ()
	return &KsAdminRejectAppVersionParams{
		HTTPClient: client,
	}
}

/*KsAdminRejectAppVersionParams contains all the parameters to send to the API endpoint
for the ks admin reject app version operation typically these are written to a http.Request
*/
type KsAdminRejectAppVersionParams struct {

	/*Body*/
	Body *models.OpenpitrixRejectAppVersionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ks admin reject app version params
func (o *KsAdminRejectAppVersionParams) WithTimeout(timeout time.Duration) *KsAdminRejectAppVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ks admin reject app version params
func (o *KsAdminRejectAppVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ks admin reject app version params
func (o *KsAdminRejectAppVersionParams) WithContext(ctx context.Context) *KsAdminRejectAppVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ks admin reject app version params
func (o *KsAdminRejectAppVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ks admin reject app version params
func (o *KsAdminRejectAppVersionParams) WithHTTPClient(client *http.Client) *KsAdminRejectAppVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ks admin reject app version params
func (o *KsAdminRejectAppVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the ks admin reject app version params
func (o *KsAdminRejectAppVersionParams) WithBody(body *models.OpenpitrixRejectAppVersionRequest) *KsAdminRejectAppVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the ks admin reject app version params
func (o *KsAdminRejectAppVersionParams) SetBody(body *models.OpenpitrixRejectAppVersionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *KsAdminRejectAppVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}