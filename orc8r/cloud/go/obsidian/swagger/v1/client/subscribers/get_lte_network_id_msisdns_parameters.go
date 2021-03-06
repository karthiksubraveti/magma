// Code generated by go-swagger; DO NOT EDIT.

package subscribers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetLTENetworkIDMsisdnsParams creates a new GetLTENetworkIDMsisdnsParams object
// with the default values initialized.
func NewGetLTENetworkIDMsisdnsParams() *GetLTENetworkIDMsisdnsParams {
	var ()
	return &GetLTENetworkIDMsisdnsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLTENetworkIDMsisdnsParamsWithTimeout creates a new GetLTENetworkIDMsisdnsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLTENetworkIDMsisdnsParamsWithTimeout(timeout time.Duration) *GetLTENetworkIDMsisdnsParams {
	var ()
	return &GetLTENetworkIDMsisdnsParams{

		timeout: timeout,
	}
}

// NewGetLTENetworkIDMsisdnsParamsWithContext creates a new GetLTENetworkIDMsisdnsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLTENetworkIDMsisdnsParamsWithContext(ctx context.Context) *GetLTENetworkIDMsisdnsParams {
	var ()
	return &GetLTENetworkIDMsisdnsParams{

		Context: ctx,
	}
}

// NewGetLTENetworkIDMsisdnsParamsWithHTTPClient creates a new GetLTENetworkIDMsisdnsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLTENetworkIDMsisdnsParamsWithHTTPClient(client *http.Client) *GetLTENetworkIDMsisdnsParams {
	var ()
	return &GetLTENetworkIDMsisdnsParams{
		HTTPClient: client,
	}
}

/*GetLTENetworkIDMsisdnsParams contains all the parameters to send to the API endpoint
for the get LTE network ID msisdns operation typically these are written to a http.Request
*/
type GetLTENetworkIDMsisdnsParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get LTE network ID msisdns params
func (o *GetLTENetworkIDMsisdnsParams) WithTimeout(timeout time.Duration) *GetLTENetworkIDMsisdnsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get LTE network ID msisdns params
func (o *GetLTENetworkIDMsisdnsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get LTE network ID msisdns params
func (o *GetLTENetworkIDMsisdnsParams) WithContext(ctx context.Context) *GetLTENetworkIDMsisdnsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get LTE network ID msisdns params
func (o *GetLTENetworkIDMsisdnsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get LTE network ID msisdns params
func (o *GetLTENetworkIDMsisdnsParams) WithHTTPClient(client *http.Client) *GetLTENetworkIDMsisdnsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get LTE network ID msisdns params
func (o *GetLTENetworkIDMsisdnsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the get LTE network ID msisdns params
func (o *GetLTENetworkIDMsisdnsParams) WithNetworkID(networkID string) *GetLTENetworkIDMsisdnsParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get LTE network ID msisdns params
func (o *GetLTENetworkIDMsisdnsParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLTENetworkIDMsisdnsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
