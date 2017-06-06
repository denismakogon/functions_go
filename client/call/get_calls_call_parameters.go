package call

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

// NewGetCallsCallParams creates a new GetCallsCallParams object
// with the default values initialized.
func NewGetCallsCallParams() *GetCallsCallParams {
	var ()
	return &GetCallsCallParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCallsCallParamsWithTimeout creates a new GetCallsCallParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCallsCallParamsWithTimeout(timeout time.Duration) *GetCallsCallParams {
	var ()
	return &GetCallsCallParams{

		timeout: timeout,
	}
}

// NewGetCallsCallParamsWithContext creates a new GetCallsCallParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCallsCallParamsWithContext(ctx context.Context) *GetCallsCallParams {
	var ()
	return &GetCallsCallParams{

		Context: ctx,
	}
}

// NewGetCallsCallParamsWithHTTPClient creates a new GetCallsCallParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCallsCallParamsWithHTTPClient(client *http.Client) *GetCallsCallParams {
	var ()
	return &GetCallsCallParams{
		HTTPClient: client,
	}
}

/*GetCallsCallParams contains all the parameters to send to the API endpoint
for the get calls call operation typically these are written to a http.Request
*/
type GetCallsCallParams struct {

	/*Call
	  Call ID.

	*/
	Call string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get calls call params
func (o *GetCallsCallParams) WithTimeout(timeout time.Duration) *GetCallsCallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get calls call params
func (o *GetCallsCallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get calls call params
func (o *GetCallsCallParams) WithContext(ctx context.Context) *GetCallsCallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get calls call params
func (o *GetCallsCallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get calls call params
func (o *GetCallsCallParams) WithHTTPClient(client *http.Client) *GetCallsCallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get calls call params
func (o *GetCallsCallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCall adds the call to the get calls call params
func (o *GetCallsCallParams) WithCall(call string) *GetCallsCallParams {
	o.SetCall(call)
	return o
}

// SetCall adds the call to the get calls call params
func (o *GetCallsCallParams) SetCall(call string) {
	o.Call = call
}

// WriteToRequest writes these params to a swagger request
func (o *GetCallsCallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param call
	if err := r.SetPathParam("call", o.Call); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
