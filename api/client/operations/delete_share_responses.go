package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/opensds/opensds/api/models"
)

// DeleteShareReader is a Reader for the DeleteShare structure.
type DeleteShareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteShareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteShareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteShareOK creates a DeleteShareOK with default headers values
func NewDeleteShareOK() *DeleteShareOK {
	return &DeleteShareOK{}
}

/*DeleteShareOK handles this case with default header values.

OK
*/
type DeleteShareOK struct {
	Payload *models.DefaultResponse
}

func (o *DeleteShareOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/shares/{resourceType}/{id}][%d] deleteShareOK  %+v", 200, o.Payload)
}

func (o *DeleteShareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
