// Code generated by go-swagger; DO NOT EDIT.

package app

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"AppHub/models"
)

// GetAppsOKCode is the HTTP code returned for type GetAppsOK
const GetAppsOKCode int = 200

/*GetAppsOK A list of apps

swagger:response getAppsOK
*/
type GetAppsOK struct {

	/*
	  In: Body
	*/
	Payload *models.GetAppsOKBody `json:"body,omitempty"`
}

// NewGetAppsOK creates GetAppsOK with default headers values
func NewGetAppsOK() *GetAppsOK {
	return &GetAppsOK{}
}

// WithPayload adds the payload to the get apps o k response
func (o *GetAppsOK) WithPayload(payload *models.GetAppsOKBody) *GetAppsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get apps o k response
func (o *GetAppsOK) SetPayload(payload *models.GetAppsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAppsInternalServerErrorCode is the HTTP code returned for type GetAppsInternalServerError
const GetAppsInternalServerErrorCode int = 500

/*GetAppsInternalServerError An unexpected error occured.

swagger:response getAppsInternalServerError
*/
type GetAppsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAppsInternalServerError creates GetAppsInternalServerError with default headers values
func NewGetAppsInternalServerError() *GetAppsInternalServerError {
	return &GetAppsInternalServerError{}
}

// WithPayload adds the payload to the get apps internal server error response
func (o *GetAppsInternalServerError) WithPayload(payload *models.Error) *GetAppsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get apps internal server error response
func (o *GetAppsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAppsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}