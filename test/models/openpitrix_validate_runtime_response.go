// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixValidateRuntimeResponse openpitrix validate runtime response
// swagger:model openpitrixValidateRuntimeResponse
type OpenpitrixValidateRuntimeResponse struct {

	// validate ok or not
	Ok bool `json:"ok,omitempty"`
}

// Validate validates this openpitrix validate runtime response
func (m *OpenpitrixValidateRuntimeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixValidateRuntimeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixValidateRuntimeResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixValidateRuntimeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}