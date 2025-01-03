// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ListWordsOKCode is the HTTP code returned for type ListWordsOK
const ListWordsOKCode int = 200

/*ListWordsOK successful operation

swagger:response listWordsOK
*/
type ListWordsOK struct {

	/*
	  In: Body
	*/
	Payload []string `json:"body,omitempty"`
}

// NewListWordsOK creates ListWordsOK with default headers values
func NewListWordsOK() *ListWordsOK {

	return &ListWordsOK{}
}

// WithPayload adds the payload to the list words o k response
func (o *ListWordsOK) WithPayload(payload []string) *ListWordsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list words o k response
func (o *ListWordsOK) SetPayload(payload []string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListWordsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]string, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
