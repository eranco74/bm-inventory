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

// ListNodesReader is a Reader for the ListNodes structure.
type ListNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListNodesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListNodesOK creates a ListNodesOK with default headers values
func NewListNodesOK() *ListNodesOK {
	return &ListNodesOK{}
}

/*ListNodesOK handles this case with default header values.

Node list
*/
type ListNodesOK struct {
	Payload models.NodeList
}

func (o *ListNodesOK) Error() string {
	return fmt.Sprintf("[GET /nodes][%d] listNodesOK  %+v", 200, o.Payload)
}

func (o *ListNodesOK) GetPayload() models.NodeList {
	return o.Payload
}

func (o *ListNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListNodesInternalServerError creates a ListNodesInternalServerError with default headers values
func NewListNodesInternalServerError() *ListNodesInternalServerError {
	return &ListNodesInternalServerError{}
}

/*ListNodesInternalServerError handles this case with default header values.

Internal server error
*/
type ListNodesInternalServerError struct {
}

func (o *ListNodesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /nodes][%d] listNodesInternalServerError ", 500)
}

func (o *ListNodesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
