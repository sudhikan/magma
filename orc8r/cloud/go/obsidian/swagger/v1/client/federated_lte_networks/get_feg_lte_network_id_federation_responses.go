// Code generated by go-swagger; DO NOT EDIT.

package federated_lte_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// GetFegLTENetworkIDFederationReader is a Reader for the GetFegLTENetworkIDFederation structure.
type GetFegLTENetworkIDFederationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFegLTENetworkIDFederationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFegLTENetworkIDFederationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFegLTENetworkIDFederationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFegLTENetworkIDFederationOK creates a GetFegLTENetworkIDFederationOK with default headers values
func NewGetFegLTENetworkIDFederationOK() *GetFegLTENetworkIDFederationOK {
	return &GetFegLTENetworkIDFederationOK{}
}

/*GetFegLTENetworkIDFederationOK handles this case with default header values.

Retrieved Network FeG Configs
*/
type GetFegLTENetworkIDFederationOK struct {
	Payload *models.FederatedNetworkConfigs
}

func (o *GetFegLTENetworkIDFederationOK) Error() string {
	return fmt.Sprintf("[GET /feg_lte/{network_id}/federation][%d] getFegLteNetworkIdFederationOK  %+v", 200, o.Payload)
}

func (o *GetFegLTENetworkIDFederationOK) GetPayload() *models.FederatedNetworkConfigs {
	return o.Payload
}

func (o *GetFegLTENetworkIDFederationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FederatedNetworkConfigs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFegLTENetworkIDFederationDefault creates a GetFegLTENetworkIDFederationDefault with default headers values
func NewGetFegLTENetworkIDFederationDefault(code int) *GetFegLTENetworkIDFederationDefault {
	return &GetFegLTENetworkIDFederationDefault{
		_statusCode: code,
	}
}

/*GetFegLTENetworkIDFederationDefault handles this case with default header values.

Unexpected Error
*/
type GetFegLTENetworkIDFederationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get feg LTE network ID federation default response
func (o *GetFegLTENetworkIDFederationDefault) Code() int {
	return o._statusCode
}

func (o *GetFegLTENetworkIDFederationDefault) Error() string {
	return fmt.Sprintf("[GET /feg_lte/{network_id}/federation][%d] GetFegLTENetworkIDFederation default  %+v", o._statusCode, o.Payload)
}

func (o *GetFegLTENetworkIDFederationDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetFegLTENetworkIDFederationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
