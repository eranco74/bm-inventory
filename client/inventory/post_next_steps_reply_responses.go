// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/filanov/bm-inventory/models"
)

// PostNextStepsReplyReader is a Reader for the PostNextStepsReply structure.
type PostNextStepsReplyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostNextStepsReplyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostNextStepsReplyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPostNextStepsReplyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostNextStepsReplyOK creates a PostNextStepsReplyOK with default headers values
func NewPostNextStepsReplyOK() *PostNextStepsReplyOK {
	return &PostNextStepsReplyOK{}
}

/*PostNextStepsReplyOK handles this case with default header values.

Steps reply
*/
type PostNextStepsReplyOK struct {
	Payload models.StepsReply
}

func (o *PostNextStepsReplyOK) Error() string {
	return fmt.Sprintf("[POST /nodes/{node_id}/next-steps/reply][%d] postNextStepsReplyOK  %+v", 200, o.Payload)
}

func (o *PostNextStepsReplyOK) GetPayload() models.StepsReply {
	return o.Payload
}

func (o *PostNextStepsReplyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostNextStepsReplyNotFound creates a PostNextStepsReplyNotFound with default headers values
func NewPostNextStepsReplyNotFound() *PostNextStepsReplyNotFound {
	return &PostNextStepsReplyNotFound{}
}

/*PostNextStepsReplyNotFound handles this case with default header values.

Node not found
*/
type PostNextStepsReplyNotFound struct {
}

func (o *PostNextStepsReplyNotFound) Error() string {
	return fmt.Sprintf("[POST /nodes/{node_id}/next-steps/reply][%d] postNextStepsReplyNotFound ", 404)
}

func (o *PostNextStepsReplyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
