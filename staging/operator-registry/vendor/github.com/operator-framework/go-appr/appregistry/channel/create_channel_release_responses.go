// Code generated by go-swagger; DO NOT EDIT.

package channel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/operator-framework/go-appr/models"
)

// CreateChannelReleaseReader is a Reader for the CreateChannelRelease structure.
type CreateChannelReleaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateChannelReleaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateChannelReleaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCreateChannelReleaseUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateChannelReleaseNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateChannelReleaseOK creates a CreateChannelReleaseOK with default headers values
func NewCreateChannelReleaseOK() *CreateChannelReleaseOK {
	return &CreateChannelReleaseOK{}
}

/*CreateChannelReleaseOK handles this case with default header values.

successful operation
*/
type CreateChannelReleaseOK struct {
	Payload *models.Channel
}

func (o *CreateChannelReleaseOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/packages/{namespace}/{package}/channels/{channel}/{release}][%d] createChannelReleaseOK  %+v", 200, o.Payload)
}

func (o *CreateChannelReleaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Channel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateChannelReleaseUnauthorized creates a CreateChannelReleaseUnauthorized with default headers values
func NewCreateChannelReleaseUnauthorized() *CreateChannelReleaseUnauthorized {
	return &CreateChannelReleaseUnauthorized{}
}

/*CreateChannelReleaseUnauthorized handles this case with default header values.

Not authorized to read the package
*/
type CreateChannelReleaseUnauthorized struct {
	Payload *models.Error
}

func (o *CreateChannelReleaseUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/packages/{namespace}/{package}/channels/{channel}/{release}][%d] createChannelReleaseUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateChannelReleaseUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateChannelReleaseNotFound creates a CreateChannelReleaseNotFound with default headers values
func NewCreateChannelReleaseNotFound() *CreateChannelReleaseNotFound {
	return &CreateChannelReleaseNotFound{}
}

/*CreateChannelReleaseNotFound handles this case with default header values.

Package not found
*/
type CreateChannelReleaseNotFound struct {
	Payload *models.Error
}

func (o *CreateChannelReleaseNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/packages/{namespace}/{package}/channels/{channel}/{release}][%d] createChannelReleaseNotFound  %+v", 404, o.Payload)
}

func (o *CreateChannelReleaseNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
