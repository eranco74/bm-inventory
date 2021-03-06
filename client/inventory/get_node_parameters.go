// Code generated by go-swagger; DO NOT EDIT.

package inventory

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

// NewGetNodeParams creates a new GetNodeParams object
// with the default values initialized.
func NewGetNodeParams() *GetNodeParams {
	var ()
	return &GetNodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNodeParamsWithTimeout creates a new GetNodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNodeParamsWithTimeout(timeout time.Duration) *GetNodeParams {
	var ()
	return &GetNodeParams{

		timeout: timeout,
	}
}

// NewGetNodeParamsWithContext creates a new GetNodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNodeParamsWithContext(ctx context.Context) *GetNodeParams {
	var ()
	return &GetNodeParams{

		Context: ctx,
	}
}

// NewGetNodeParamsWithHTTPClient creates a new GetNodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNodeParamsWithHTTPClient(client *http.Client) *GetNodeParams {
	var ()
	return &GetNodeParams{
		HTTPClient: client,
	}
}

/*GetNodeParams contains all the parameters to send to the API endpoint
for the get node operation typically these are written to a http.Request
*/
type GetNodeParams struct {

	/*NodeID
	  The ID of the node to retrieve

	*/
	NodeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get node params
func (o *GetNodeParams) WithTimeout(timeout time.Duration) *GetNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get node params
func (o *GetNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get node params
func (o *GetNodeParams) WithContext(ctx context.Context) *GetNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get node params
func (o *GetNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get node params
func (o *GetNodeParams) WithHTTPClient(client *http.Client) *GetNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get node params
func (o *GetNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNodeID adds the nodeID to the get node params
func (o *GetNodeParams) WithNodeID(nodeID string) *GetNodeParams {
	o.SetNodeID(nodeID)
	return o
}

// SetNodeID adds the nodeId to the get node params
func (o *GetNodeParams) SetNodeID(nodeID string) {
	o.NodeID = nodeID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param node_id
	if err := r.SetPathParam("node_id", o.NodeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
