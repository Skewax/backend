// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Skewax/backend/pkg/swagger/server/models"
)

// GetTokenLoginOKCode is the HTTP code returned for type GetTokenLoginOK
const GetTokenLoginOKCode int = 200

/*
GetTokenLoginOK OK

swagger:response getTokenLoginOK
*/
type GetTokenLoginOK struct {

	/*
	  In: Body
	*/
	Payload *models.LoginResponse `json:"body,omitempty"`
}

// NewGetTokenLoginOK creates GetTokenLoginOK with default headers values
func NewGetTokenLoginOK() *GetTokenLoginOK {

	return &GetTokenLoginOK{}
}

// WithPayload adds the payload to the get token login o k response
func (o *GetTokenLoginOK) WithPayload(payload *models.LoginResponse) *GetTokenLoginOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get token login o k response
func (o *GetTokenLoginOK) SetPayload(payload *models.LoginResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTokenLoginOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTokenLoginBadRequestCode is the HTTP code returned for type GetTokenLoginBadRequest
const GetTokenLoginBadRequestCode int = 400

/*
GetTokenLoginBadRequest an incorrect/incorrectly formatted ID

swagger:response getTokenLoginBadRequest
*/
type GetTokenLoginBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.BasicResponse `json:"body,omitempty"`
}

// NewGetTokenLoginBadRequest creates GetTokenLoginBadRequest with default headers values
func NewGetTokenLoginBadRequest() *GetTokenLoginBadRequest {

	return &GetTokenLoginBadRequest{}
}

// WithPayload adds the payload to the get token login bad request response
func (o *GetTokenLoginBadRequest) WithPayload(payload *models.BasicResponse) *GetTokenLoginBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get token login bad request response
func (o *GetTokenLoginBadRequest) SetPayload(payload *models.BasicResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTokenLoginBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTokenLoginInternalServerErrorCode is the HTTP code returned for type GetTokenLoginInternalServerError
const GetTokenLoginInternalServerErrorCode int = 500

/*
GetTokenLoginInternalServerError generic server error

swagger:response getTokenLoginInternalServerError
*/
type GetTokenLoginInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.BasicResponse `json:"body,omitempty"`
}

// NewGetTokenLoginInternalServerError creates GetTokenLoginInternalServerError with default headers values
func NewGetTokenLoginInternalServerError() *GetTokenLoginInternalServerError {

	return &GetTokenLoginInternalServerError{}
}

// WithPayload adds the payload to the get token login internal server error response
func (o *GetTokenLoginInternalServerError) WithPayload(payload *models.BasicResponse) *GetTokenLoginInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get token login internal server error response
func (o *GetTokenLoginInternalServerError) SetPayload(payload *models.BasicResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTokenLoginInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
