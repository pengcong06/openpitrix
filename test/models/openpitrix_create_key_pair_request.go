// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixCreateKeyPairRequest openpitrix create key pair request
// swagger:model openpitrixCreateKeyPairRequest
type OpenpitrixCreateKeyPairRequest struct {

	// keypair description
	Description string `json:"description,omitempty"`

	// keypair name
	Name string `json:"name,omitempty"`

	// public key
	PubKey string `json:"pub_key,omitempty"`
}

// Validate validates this openpitrix create key pair request
func (m *OpenpitrixCreateKeyPairRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixCreateKeyPairRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixCreateKeyPairRequest) UnmarshalBinary(b []byte) error {
	var res OpenpitrixCreateKeyPairRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
