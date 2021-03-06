// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sendinblue/APIv3-go-library/models"
)

// GetListReader is a Reader for the GetList structure.
type GetListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetListNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetListOK creates a GetListOK with default headers values
func NewGetListOK() *GetListOK {
	return &GetListOK{}
}

/*GetListOK handles this case with default header values.

List informations
*/
type GetListOK struct {
	Payload *models.GetExtendedList
}

func (o *GetListOK) Error() string {
	return fmt.Sprintf("[GET /contacts/lists/{listId}][%d] getListOK  %+v", 200, o.Payload)
}

func (o *GetListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetExtendedList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListBadRequest creates a GetListBadRequest with default headers values
func NewGetListBadRequest() *GetListBadRequest {
	return &GetListBadRequest{}
}

/*GetListBadRequest handles this case with default header values.

bad request
*/
type GetListBadRequest struct {
	Payload *models.ErrorModel
}

func (o *GetListBadRequest) Error() string {
	return fmt.Sprintf("[GET /contacts/lists/{listId}][%d] getListBadRequest  %+v", 400, o.Payload)
}

func (o *GetListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetListNotFound creates a GetListNotFound with default headers values
func NewGetListNotFound() *GetListNotFound {
	return &GetListNotFound{}
}

/*GetListNotFound handles this case with default header values.

List ID not found
*/
type GetListNotFound struct {
	Payload *models.ErrorModel
}

func (o *GetListNotFound) Error() string {
	return fmt.Sprintf("[GET /contacts/lists/{listId}][%d] getListNotFound  %+v", 404, o.Payload)
}

func (o *GetListNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
