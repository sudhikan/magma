// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_networks

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

// NewPostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams creates a new PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams object
// with the default values initialized.
func NewPostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams() *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams {
	var ()
	return &PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParamsWithTimeout creates a new PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParamsWithTimeout(timeout time.Duration) *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams {
	var ()
	return &PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams{

		timeout: timeout,
	}
}

// NewPostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParamsWithContext creates a new PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParamsWithContext(ctx context.Context) *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams {
	var ()
	return &PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams{

		Context: ctx,
	}
}

// NewPostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParamsWithHTTPClient creates a new PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParamsWithHTTPClient(client *http.Client) *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams {
	var ()
	return &PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams{
		HTTPClient: client,
	}
}

/*PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams contains all the parameters to send to the API endpoint
for the post cwf network ID subscriber config base names base name operation typically these are written to a http.Request
*/
type PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams struct {

	/*BaseName
	  Charging Rule Base Name

	*/
	BaseName string
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post cwf network ID subscriber config base names base name params
func (o *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams) WithTimeout(timeout time.Duration) *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post cwf network ID subscriber config base names base name params
func (o *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post cwf network ID subscriber config base names base name params
func (o *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams) WithContext(ctx context.Context) *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post cwf network ID subscriber config base names base name params
func (o *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post cwf network ID subscriber config base names base name params
func (o *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams) WithHTTPClient(client *http.Client) *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post cwf network ID subscriber config base names base name params
func (o *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBaseName adds the baseName to the post cwf network ID subscriber config base names base name params
func (o *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams) WithBaseName(baseName string) *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams {
	o.SetBaseName(baseName)
	return o
}

// SetBaseName adds the baseName to the post cwf network ID subscriber config base names base name params
func (o *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams) SetBaseName(baseName string) {
	o.BaseName = baseName
}

// WithNetworkID adds the networkID to the post cwf network ID subscriber config base names base name params
func (o *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams) WithNetworkID(networkID string) *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the post cwf network ID subscriber config base names base name params
func (o *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *PostCwfNetworkIDSubscriberConfigBaseNamesBaseNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param base_name
	if err := r.SetPathParam("base_name", o.BaseName); err != nil {
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
