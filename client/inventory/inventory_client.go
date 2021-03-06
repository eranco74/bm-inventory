// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the inventory client
type API interface {
	/*
	   CreateImage creates an open shift bare metal cluster assist installation image*/
	CreateImage(ctx context.Context, params *CreateImageParams) (*CreateImageCreated, error)
	/*
	   DeregisterCluster deregisters open shift bare metal cluster*/
	DeregisterCluster(ctx context.Context, params *DeregisterClusterParams) (*DeregisterClusterNoContent, error)
	/*
	   DeregisterNode deregisters open shift bare metal node*/
	DeregisterNode(ctx context.Context, params *DeregisterNodeParams) (*DeregisterNodeNoContent, error)
	/*
	   GetCluster retrieves open shift bare metal cluster information*/
	GetCluster(ctx context.Context, params *GetClusterParams) (*GetClusterOK, error)
	/*
	   GetImage retrieves installation image information*/
	GetImage(ctx context.Context, params *GetImageParams) (*GetImageOK, error)
	/*
	   GetNextSteps retrieves the next operations that the agent need to perform*/
	GetNextSteps(ctx context.Context, params *GetNextStepsParams) (*GetNextStepsOK, error)
	/*
	   GetNode retrieves open shift bare metal node information*/
	GetNode(ctx context.Context, params *GetNodeParams) (*GetNodeOK, error)
	/*
	   ListClusters lists open shift bare metal clusters*/
	ListClusters(ctx context.Context, params *ListClustersParams) (*ListClustersOK, error)
	/*
	   ListImages lists installation images*/
	ListImages(ctx context.Context, params *ListImagesParams) (*ListImagesOK, error)
	/*
	   ListNodes lists open shift bare metal nodes*/
	ListNodes(ctx context.Context, params *ListNodesParams) (*ListNodesOK, error)
	/*
	   PostNextStepsReply posts the result of the required operations from the server*/
	PostNextStepsReply(ctx context.Context, params *PostNextStepsReplyParams) (*PostNextStepsReplyOK, error)
	/*
	   RegisterCluster registers a new open shift bare metal cluster*/
	RegisterCluster(ctx context.Context, params *RegisterClusterParams) (*RegisterClusterCreated, error)
	/*
	   RegisterNode registers a new open shift bare metal node*/
	RegisterNode(ctx context.Context, params *RegisterNodeParams) (*RegisterNodeCreated, error)
}

// New creates a new inventory API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for inventory API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
CreateImage creates an open shift bare metal cluster assist installation image
*/
func (a *Client) CreateImage(ctx context.Context, params *CreateImageParams) (*CreateImageCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateImage",
		Method:             "POST",
		PathPattern:        "/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateImageReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateImageCreated), nil

}

/*
DeregisterCluster deregisters open shift bare metal cluster
*/
func (a *Client) DeregisterCluster(ctx context.Context, params *DeregisterClusterParams) (*DeregisterClusterNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeregisterCluster",
		Method:             "DELETE",
		PathPattern:        "/clusters/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeregisterClusterReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeregisterClusterNoContent), nil

}

/*
DeregisterNode deregisters open shift bare metal node
*/
func (a *Client) DeregisterNode(ctx context.Context, params *DeregisterNodeParams) (*DeregisterNodeNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeregisterNode",
		Method:             "DELETE",
		PathPattern:        "/nodes/{node_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeregisterNodeReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeregisterNodeNoContent), nil

}

/*
GetCluster retrieves open shift bare metal cluster information
*/
func (a *Client) GetCluster(ctx context.Context, params *GetClusterParams) (*GetClusterOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetCluster",
		Method:             "GET",
		PathPattern:        "/clusters/{cluster_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterOK), nil

}

/*
GetImage retrieves installation image information
*/
func (a *Client) GetImage(ctx context.Context, params *GetImageParams) (*GetImageOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetImage",
		Method:             "GET",
		PathPattern:        "/images/{image_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetImageReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetImageOK), nil

}

/*
GetNextSteps retrieves the next operations that the agent need to perform
*/
func (a *Client) GetNextSteps(ctx context.Context, params *GetNextStepsParams) (*GetNextStepsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNextSteps",
		Method:             "GET",
		PathPattern:        "/nodes/{node_id}/next-steps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetNextStepsReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNextStepsOK), nil

}

/*
GetNode retrieves open shift bare metal node information
*/
func (a *Client) GetNode(ctx context.Context, params *GetNodeParams) (*GetNodeOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetNode",
		Method:             "GET",
		PathPattern:        "/nodes/{node_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetNodeReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodeOK), nil

}

/*
ListClusters lists open shift bare metal clusters
*/
func (a *Client) ListClusters(ctx context.Context, params *ListClustersParams) (*ListClustersOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListClusters",
		Method:             "GET",
		PathPattern:        "/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListClustersReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListClustersOK), nil

}

/*
ListImages lists installation images
*/
func (a *Client) ListImages(ctx context.Context, params *ListImagesParams) (*ListImagesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListImages",
		Method:             "GET",
		PathPattern:        "/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListImagesReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListImagesOK), nil

}

/*
ListNodes lists open shift bare metal nodes
*/
func (a *Client) ListNodes(ctx context.Context, params *ListNodesParams) (*ListNodesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListNodes",
		Method:             "GET",
		PathPattern:        "/nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListNodesReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListNodesOK), nil

}

/*
PostNextStepsReply posts the result of the required operations from the server
*/
func (a *Client) PostNextStepsReply(ctx context.Context, params *PostNextStepsReplyParams) (*PostNextStepsReplyOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostNextStepsReply",
		Method:             "POST",
		PathPattern:        "/nodes/{node_id}/next-steps/reply",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostNextStepsReplyReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostNextStepsReplyOK), nil

}

/*
RegisterCluster registers a new open shift bare metal cluster
*/
func (a *Client) RegisterCluster(ctx context.Context, params *RegisterClusterParams) (*RegisterClusterCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RegisterCluster",
		Method:             "POST",
		PathPattern:        "/clusters",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RegisterClusterReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RegisterClusterCreated), nil

}

/*
RegisterNode registers a new open shift bare metal node
*/
func (a *Client) RegisterNode(ctx context.Context, params *RegisterNodeParams) (*RegisterNodeCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RegisterNode",
		Method:             "POST",
		PathPattern:        "/nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RegisterNodeReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RegisterNodeCreated), nil

}
