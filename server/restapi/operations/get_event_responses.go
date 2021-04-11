// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetEventAcceptedCode is the HTTP code returned for type GetEventAccepted
const GetEventAcceptedCode int = 202

/*GetEventAccepted Accepted

swagger:response getEventAccepted
*/
type GetEventAccepted struct {

	/*
	  In: Body
	*/
	Payload *GetEventAcceptedBody `json:"body,omitempty"`
}

// NewGetEventAccepted creates GetEventAccepted with default headers values
func NewGetEventAccepted() *GetEventAccepted {

	return &GetEventAccepted{}
}

// WithPayload adds the payload to the get event accepted response
func (o *GetEventAccepted) WithPayload(payload *GetEventAcceptedBody) *GetEventAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get event accepted response
func (o *GetEventAccepted) SetPayload(payload *GetEventAcceptedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEventAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
