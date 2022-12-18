// Code generated by go-swagger; DO NOT EDIT.

package files

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/Skewax/backend/pkg/swagger/server/models"
)

// PostCreateFileHandlerFunc turns a function with the right signature into a post create file handler
type PostCreateFileHandlerFunc func(PostCreateFileParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostCreateFileHandlerFunc) Handle(params PostCreateFileParams) middleware.Responder {
	return fn(params)
}

// PostCreateFileHandler interface for that can handle valid post create file params
type PostCreateFileHandler interface {
	Handle(PostCreateFileParams) middleware.Responder
}

// NewPostCreateFile creates a new http.Handler for the post create file operation
func NewPostCreateFile(ctx *middleware.Context, handler PostCreateFileHandler) *PostCreateFile {
	return &PostCreateFile{Context: ctx, Handler: handler}
}

/*
	PostCreateFile swagger:route POST /createFile Files postCreateFile

create file with name and text, returns basic file object
*/
type PostCreateFile struct {
	Context *middleware.Context
	Handler PostCreateFileHandler
}

func (o *PostCreateFile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostCreateFileParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostCreateFileBody post create file body
//
// swagger:model PostCreateFileBody
type PostCreateFileBody struct {

	// file name
	// Required: true
	FileName *string `json:"file_name"`

	// session token
	// Required: true
	SessionToken *string `json:"session_token"`

	// text
	// Required: true
	Text *string `json:"text"`

	// user id
	// Required: true
	UserID *string `json:"user_id"`
}

// Validate validates this post create file body
func (o *PostCreateFileBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFileName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSessionToken(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateText(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCreateFileBody) validateFileName(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"file_name", "body", o.FileName); err != nil {
		return err
	}

	return nil
}

func (o *PostCreateFileBody) validateSessionToken(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"session_token", "body", o.SessionToken); err != nil {
		return err
	}

	return nil
}

func (o *PostCreateFileBody) validateText(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"text", "body", o.Text); err != nil {
		return err
	}

	return nil
}

func (o *PostCreateFileBody) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"user_id", "body", o.UserID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post create file body based on context it is used
func (o *PostCreateFileBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostCreateFileBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCreateFileBody) UnmarshalBinary(b []byte) error {
	var res PostCreateFileBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostCreateFileOKBody post create file o k body
//
// swagger:model PostCreateFileOKBody
type PostCreateFileOKBody struct {

	// error
	Error string `json:"error,omitempty"`

	// file
	File *models.BasicFileObject `json:"file,omitempty"`
}

// Validate validates this post create file o k body
func (o *PostCreateFileOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCreateFileOKBody) validateFile(formats strfmt.Registry) error {
	if swag.IsZero(o.File) { // not required
		return nil
	}

	if o.File != nil {
		if err := o.File.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postCreateFileOK" + "." + "file")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postCreateFileOK" + "." + "file")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post create file o k body based on the context it is used
func (o *PostCreateFileOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostCreateFileOKBody) contextValidateFile(ctx context.Context, formats strfmt.Registry) error {

	if o.File != nil {
		if err := o.File.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postCreateFileOK" + "." + "file")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postCreateFileOK" + "." + "file")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostCreateFileOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCreateFileOKBody) UnmarshalBinary(b []byte) error {
	var res PostCreateFileOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}