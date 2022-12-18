// Code generated by go-swagger; DO NOT EDIT.

package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Skewax/backend/pkg/swagger/server/models"
)

// PostCreateFileOKCode is the HTTP code returned for type PostCreateFileOK
const PostCreateFileOKCode int = 200

/*
PostCreateFileOK OK: updated file

swagger:response postCreateFileOK
*/
type PostCreateFileOK struct {

	/*
	  In: Body
	*/
	Payload *PostCreateFileOKBody `json:"body,omitempty"`
}

// NewPostCreateFileOK creates PostCreateFileOK with default headers values
func NewPostCreateFileOK() *PostCreateFileOK {

	return &PostCreateFileOK{}
}

// WithPayload adds the payload to the post create file o k response
func (o *PostCreateFileOK) WithPayload(payload *PostCreateFileOKBody) *PostCreateFileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post create file o k response
func (o *PostCreateFileOK) SetPayload(payload *PostCreateFileOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCreateFileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostCreateFileBadRequestCode is the HTTP code returned for type PostCreateFileBadRequest
const PostCreateFileBadRequestCode int = 400

/*
PostCreateFileBadRequest malformed/invalid session token

swagger:response postCreateFileBadRequest
*/
type PostCreateFileBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.BasicResponse `json:"body,omitempty"`
}

// NewPostCreateFileBadRequest creates PostCreateFileBadRequest with default headers values
func NewPostCreateFileBadRequest() *PostCreateFileBadRequest {

	return &PostCreateFileBadRequest{}
}

// WithPayload adds the payload to the post create file bad request response
func (o *PostCreateFileBadRequest) WithPayload(payload *models.BasicResponse) *PostCreateFileBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post create file bad request response
func (o *PostCreateFileBadRequest) SetPayload(payload *models.BasicResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCreateFileBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostCreateFileInternalServerErrorCode is the HTTP code returned for type PostCreateFileInternalServerError
const PostCreateFileInternalServerErrorCode int = 500

/*
PostCreateFileInternalServerError generic server error

swagger:response postCreateFileInternalServerError
*/
type PostCreateFileInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.BasicResponse `json:"body,omitempty"`
}

// NewPostCreateFileInternalServerError creates PostCreateFileInternalServerError with default headers values
func NewPostCreateFileInternalServerError() *PostCreateFileInternalServerError {

	return &PostCreateFileInternalServerError{}
}

// WithPayload adds the payload to the post create file internal server error response
func (o *PostCreateFileInternalServerError) WithPayload(payload *models.BasicResponse) *PostCreateFileInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post create file internal server error response
func (o *PostCreateFileInternalServerError) SetPayload(payload *models.BasicResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostCreateFileInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}