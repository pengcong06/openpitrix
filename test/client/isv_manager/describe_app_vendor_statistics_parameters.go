// Code generated by go-swagger; DO NOT EDIT.

package isv_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDescribeAppVendorStatisticsParams creates a new DescribeAppVendorStatisticsParams object
// with the default values initialized.
func NewDescribeAppVendorStatisticsParams() *DescribeAppVendorStatisticsParams {
	var ()
	return &DescribeAppVendorStatisticsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeAppVendorStatisticsParamsWithTimeout creates a new DescribeAppVendorStatisticsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeAppVendorStatisticsParamsWithTimeout(timeout time.Duration) *DescribeAppVendorStatisticsParams {
	var ()
	return &DescribeAppVendorStatisticsParams{

		timeout: timeout,
	}
}

// NewDescribeAppVendorStatisticsParamsWithContext creates a new DescribeAppVendorStatisticsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeAppVendorStatisticsParamsWithContext(ctx context.Context) *DescribeAppVendorStatisticsParams {
	var ()
	return &DescribeAppVendorStatisticsParams{

		Context: ctx,
	}
}

// NewDescribeAppVendorStatisticsParamsWithHTTPClient creates a new DescribeAppVendorStatisticsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeAppVendorStatisticsParamsWithHTTPClient(client *http.Client) *DescribeAppVendorStatisticsParams {
	var ()
	return &DescribeAppVendorStatisticsParams{
		HTTPClient: client,
	}
}

/*DescribeAppVendorStatisticsParams contains all the parameters to send to the API endpoint
for the describe app vendor statistics operation typically these are written to a http.Request
*/
type DescribeAppVendorStatisticsParams struct {

	/*DisplayColumns
	  select column to display.

	*/
	DisplayColumns []string
	/*Limit
	  data limit per page, default value 20, max value 200.

	*/
	Limit *int64
	/*Offset
	  data offset, default 0.

	*/
	Offset *int64
	/*Owner
	  owner.

	*/
	Owner []string
	/*Reverse
	  value = 0 sort ASC, value = 1 sort DESC.

	*/
	Reverse *bool
	/*SearchWord
	  query key, support these fields(user_id, status).

	*/
	SearchWord *string
	/*SortKey
	  sort key, order by sort_key, default create_time.

	*/
	SortKey *string
	/*Status
	  status eg.[draft|submitted|passed|rejected|suspended|in-review|new].

	*/
	Status []string
	/*UserID
	  user ids.

	*/
	UserID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) WithTimeout(timeout time.Duration) *DescribeAppVendorStatisticsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) WithContext(ctx context.Context) *DescribeAppVendorStatisticsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) WithHTTPClient(client *http.Client) *DescribeAppVendorStatisticsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisplayColumns adds the displayColumns to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) WithDisplayColumns(displayColumns []string) *DescribeAppVendorStatisticsParams {
	o.SetDisplayColumns(displayColumns)
	return o
}

// SetDisplayColumns adds the displayColumns to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) SetDisplayColumns(displayColumns []string) {
	o.DisplayColumns = displayColumns
}

// WithLimit adds the limit to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) WithLimit(limit *int64) *DescribeAppVendorStatisticsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) WithOffset(offset *int64) *DescribeAppVendorStatisticsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOwner adds the owner to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) WithOwner(owner []string) *DescribeAppVendorStatisticsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) SetOwner(owner []string) {
	o.Owner = owner
}

// WithReverse adds the reverse to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) WithReverse(reverse *bool) *DescribeAppVendorStatisticsParams {
	o.SetReverse(reverse)
	return o
}

// SetReverse adds the reverse to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) SetReverse(reverse *bool) {
	o.Reverse = reverse
}

// WithSearchWord adds the searchWord to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) WithSearchWord(searchWord *string) *DescribeAppVendorStatisticsParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WithSortKey adds the sortKey to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) WithSortKey(sortKey *string) *DescribeAppVendorStatisticsParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WithStatus adds the status to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) WithStatus(status []string) *DescribeAppVendorStatisticsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) SetStatus(status []string) {
	o.Status = status
}

// WithUserID adds the userID to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) WithUserID(userID []string) *DescribeAppVendorStatisticsParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the describe app vendor statistics params
func (o *DescribeAppVendorStatisticsParams) SetUserID(userID []string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeAppVendorStatisticsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesDisplayColumns := o.DisplayColumns

	joinedDisplayColumns := swag.JoinByFormat(valuesDisplayColumns, "multi")
	// query array param display_columns
	if err := r.SetQueryParam("display_columns", joinedDisplayColumns...); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	valuesOwner := o.Owner

	joinedOwner := swag.JoinByFormat(valuesOwner, "multi")
	// query array param owner
	if err := r.SetQueryParam("owner", joinedOwner...); err != nil {
		return err
	}

	if o.Reverse != nil {

		// query param reverse
		var qrReverse bool
		if o.Reverse != nil {
			qrReverse = *o.Reverse
		}
		qReverse := swag.FormatBool(qrReverse)
		if qReverse != "" {
			if err := r.SetQueryParam("reverse", qReverse); err != nil {
				return err
			}
		}

	}

	if o.SearchWord != nil {

		// query param search_word
		var qrSearchWord string
		if o.SearchWord != nil {
			qrSearchWord = *o.SearchWord
		}
		qSearchWord := qrSearchWord
		if qSearchWord != "" {
			if err := r.SetQueryParam("search_word", qSearchWord); err != nil {
				return err
			}
		}

	}

	if o.SortKey != nil {

		// query param sort_key
		var qrSortKey string
		if o.SortKey != nil {
			qrSortKey = *o.SortKey
		}
		qSortKey := qrSortKey
		if qSortKey != "" {
			if err := r.SetQueryParam("sort_key", qSortKey); err != nil {
				return err
			}
		}

	}

	valuesStatus := o.Status

	joinedStatus := swag.JoinByFormat(valuesStatus, "multi")
	// query array param status
	if err := r.SetQueryParam("status", joinedStatus...); err != nil {
		return err
	}

	valuesUserID := o.UserID

	joinedUserID := swag.JoinByFormat(valuesUserID, "multi")
	// query array param user_id
	if err := r.SetQueryParam("user_id", joinedUserID...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}