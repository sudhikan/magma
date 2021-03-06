// Code generated by go-swagger; DO NOT EDIT.

package federation_networks

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

// NewDeleteFegNetworkIDFederationParams creates a new DeleteFegNetworkIDFederationParams object
// with the default values initialized.
func NewDeleteFegNetworkIDFederationParams() *DeleteFegNetworkIDFederationParams {
	var ()
	return &DeleteFegNetworkIDFederationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFegNetworkIDFederationParamsWithTimeout creates a new DeleteFegNetworkIDFederationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteFegNetworkIDFederationParamsWithTimeout(timeout time.Duration) *DeleteFegNetworkIDFederationParams {
	var ()
	return &DeleteFegNetworkIDFederationParams{

		timeout: timeout,
	}
}

// NewDeleteFegNetworkIDFederationParamsWithContext creates a new DeleteFegNetworkIDFederationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteFegNetworkIDFederationParamsWithContext(ctx context.Context) *DeleteFegNetworkIDFederationParams {
	var ()
	return &DeleteFegNetworkIDFederationParams{

		Context: ctx,
	}
}

// NewDeleteFegNetworkIDFederationParamsWithHTTPClient creates a new DeleteFegNetworkIDFederationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteFegNetworkIDFederationParamsWithHTTPClient(client *http.Client) *DeleteFegNetworkIDFederationParams {
	var ()
	return &DeleteFegNetworkIDFederationParams{
		HTTPClient: client,
	}
}

/*DeleteFegNetworkIDFederationParams contains all the parameters to send to the API endpoint
for the delete feg network ID federation operation typically these are written to a http.Request
*/
type DeleteFegNetworkIDFederationParams struct {

	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete feg network ID federation params
func (o *DeleteFegNetworkIDFederationParams) WithTimeout(timeout time.Duration) *DeleteFegNetworkIDFederationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete feg network ID federation params
func (o *DeleteFegNetworkIDFederationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete feg network ID federation params
func (o *DeleteFegNetworkIDFederationParams) WithContext(ctx context.Context) *DeleteFegNetworkIDFederationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete feg network ID federation params
func (o *DeleteFegNetworkIDFederationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete feg network ID federation params
func (o *DeleteFegNetworkIDFederationParams) WithHTTPClient(client *http.Client) *DeleteFegNetworkIDFederationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete feg network ID federation params
func (o *DeleteFegNetworkIDFederationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetworkID adds the networkID to the delete feg network ID federation params
func (o *DeleteFegNetworkIDFederationParams) WithNetworkID(networkID string) *DeleteFegNetworkIDFederationParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete feg network ID federation params
func (o *DeleteFegNetworkIDFederationParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFegNetworkIDFederationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
