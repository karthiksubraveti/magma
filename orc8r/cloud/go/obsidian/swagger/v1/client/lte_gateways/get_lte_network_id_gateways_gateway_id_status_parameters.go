// Code generated by go-swagger; DO NOT EDIT.

package lte_gateways

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

// NewGetLTENetworkIDGatewaysGatewayIDStatusParams creates a new GetLTENetworkIDGatewaysGatewayIDStatusParams object
// with the default values initialized.
func NewGetLTENetworkIDGatewaysGatewayIDStatusParams() *GetLTENetworkIDGatewaysGatewayIDStatusParams {
	var ()
	return &GetLTENetworkIDGatewaysGatewayIDStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDStatusParamsWithTimeout creates a new GetLTENetworkIDGatewaysGatewayIDStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLTENetworkIDGatewaysGatewayIDStatusParamsWithTimeout(timeout time.Duration) *GetLTENetworkIDGatewaysGatewayIDStatusParams {
	var ()
	return &GetLTENetworkIDGatewaysGatewayIDStatusParams{

		timeout: timeout,
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDStatusParamsWithContext creates a new GetLTENetworkIDGatewaysGatewayIDStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLTENetworkIDGatewaysGatewayIDStatusParamsWithContext(ctx context.Context) *GetLTENetworkIDGatewaysGatewayIDStatusParams {
	var ()
	return &GetLTENetworkIDGatewaysGatewayIDStatusParams{

		Context: ctx,
	}
}

// NewGetLTENetworkIDGatewaysGatewayIDStatusParamsWithHTTPClient creates a new GetLTENetworkIDGatewaysGatewayIDStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLTENetworkIDGatewaysGatewayIDStatusParamsWithHTTPClient(client *http.Client) *GetLTENetworkIDGatewaysGatewayIDStatusParams {
	var ()
	return &GetLTENetworkIDGatewaysGatewayIDStatusParams{
		HTTPClient: client,
	}
}

/*GetLTENetworkIDGatewaysGatewayIDStatusParams contains all the parameters to send to the API endpoint
for the get LTE network ID gateways gateway ID status operation typically these are written to a http.Request
*/
type GetLTENetworkIDGatewaysGatewayIDStatusParams struct {

	/*GatewayID
	  Gateway ID

	*/
	GatewayID string
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get LTE network ID gateways gateway ID status params
func (o *GetLTENetworkIDGatewaysGatewayIDStatusParams) WithTimeout(timeout time.Duration) *GetLTENetworkIDGatewaysGatewayIDStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get LTE network ID gateways gateway ID status params
func (o *GetLTENetworkIDGatewaysGatewayIDStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get LTE network ID gateways gateway ID status params
func (o *GetLTENetworkIDGatewaysGatewayIDStatusParams) WithContext(ctx context.Context) *GetLTENetworkIDGatewaysGatewayIDStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get LTE network ID gateways gateway ID status params
func (o *GetLTENetworkIDGatewaysGatewayIDStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get LTE network ID gateways gateway ID status params
func (o *GetLTENetworkIDGatewaysGatewayIDStatusParams) WithHTTPClient(client *http.Client) *GetLTENetworkIDGatewaysGatewayIDStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get LTE network ID gateways gateway ID status params
func (o *GetLTENetworkIDGatewaysGatewayIDStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayID adds the gatewayID to the get LTE network ID gateways gateway ID status params
func (o *GetLTENetworkIDGatewaysGatewayIDStatusParams) WithGatewayID(gatewayID string) *GetLTENetworkIDGatewaysGatewayIDStatusParams {
	o.SetGatewayID(gatewayID)
	return o
}

// SetGatewayID adds the gatewayId to the get LTE network ID gateways gateway ID status params
func (o *GetLTENetworkIDGatewaysGatewayIDStatusParams) SetGatewayID(gatewayID string) {
	o.GatewayID = gatewayID
}

// WithNetworkID adds the networkID to the get LTE network ID gateways gateway ID status params
func (o *GetLTENetworkIDGatewaysGatewayIDStatusParams) WithNetworkID(networkID string) *GetLTENetworkIDGatewaysGatewayIDStatusParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get LTE network ID gateways gateway ID status params
func (o *GetLTENetworkIDGatewaysGatewayIDStatusParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLTENetworkIDGatewaysGatewayIDStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gateway_id
	if err := r.SetPathParam("gateway_id", o.GatewayID); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
