// Code generated by go-swagger; DO NOT EDIT.

package change

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cvbarros/go-teamcity-sdk/models"
)

// ServeChangeReader is a Reader for the ServeChange structure.
type ServeChangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServeChangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServeChangeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServeChangeOK creates a ServeChangeOK with default headers values
func NewServeChangeOK() *ServeChangeOK {
	return &ServeChangeOK{}
}

/*ServeChangeOK handles this case with default header values.

successful operation
*/
type ServeChangeOK struct {
	Payload *models.Change
}

func (o *ServeChangeOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/changes/{changeLocator}][%d] serveChangeOK  %+v", 200, o.Payload)
}

func (o *ServeChangeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Change)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}