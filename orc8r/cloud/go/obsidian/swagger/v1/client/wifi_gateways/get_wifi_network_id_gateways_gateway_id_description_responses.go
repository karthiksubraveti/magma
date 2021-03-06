// Code generated by go-swagger; DO NOT EDIT.

package wifi_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// GetWifiNetworkIDGatewaysGatewayIDDescriptionReader is a Reader for the GetWifiNetworkIDGatewaysGatewayIDDescription structure.
type GetWifiNetworkIDGatewaysGatewayIDDescriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWifiNetworkIDGatewaysGatewayIDDescriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWifiNetworkIDGatewaysGatewayIDDescriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWifiNetworkIDGatewaysGatewayIDDescriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWifiNetworkIDGatewaysGatewayIDDescriptionOK creates a GetWifiNetworkIDGatewaysGatewayIDDescriptionOK with default headers values
func NewGetWifiNetworkIDGatewaysGatewayIDDescriptionOK() *GetWifiNetworkIDGatewaysGatewayIDDescriptionOK {
	return &GetWifiNetworkIDGatewaysGatewayIDDescriptionOK{}
}

/*GetWifiNetworkIDGatewaysGatewayIDDescriptionOK handles this case with default header values.

The description of the gateway
*/
type GetWifiNetworkIDGatewaysGatewayIDDescriptionOK struct {
	Payload models.GatewayDescription
}

func (o *GetWifiNetworkIDGatewaysGatewayIDDescriptionOK) Error() string {
	return fmt.Sprintf("[GET /wifi/{network_id}/gateways/{gateway_id}/description][%d] getWifiNetworkIdGatewaysGatewayIdDescriptionOK  %+v", 200, o.Payload)
}

func (o *GetWifiNetworkIDGatewaysGatewayIDDescriptionOK) GetPayload() models.GatewayDescription {
	return o.Payload
}

func (o *GetWifiNetworkIDGatewaysGatewayIDDescriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWifiNetworkIDGatewaysGatewayIDDescriptionDefault creates a GetWifiNetworkIDGatewaysGatewayIDDescriptionDefault with default headers values
func NewGetWifiNetworkIDGatewaysGatewayIDDescriptionDefault(code int) *GetWifiNetworkIDGatewaysGatewayIDDescriptionDefault {
	return &GetWifiNetworkIDGatewaysGatewayIDDescriptionDefault{
		_statusCode: code,
	}
}

/*GetWifiNetworkIDGatewaysGatewayIDDescriptionDefault handles this case with default header values.

Unexpected Error
*/
type GetWifiNetworkIDGatewaysGatewayIDDescriptionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get wifi network ID gateways gateway ID description default response
func (o *GetWifiNetworkIDGatewaysGatewayIDDescriptionDefault) Code() int {
	return o._statusCode
}

func (o *GetWifiNetworkIDGatewaysGatewayIDDescriptionDefault) Error() string {
	return fmt.Sprintf("[GET /wifi/{network_id}/gateways/{gateway_id}/description][%d] GetWifiNetworkIDGatewaysGatewayIDDescription default  %+v", o._statusCode, o.Payload)
}

func (o *GetWifiNetworkIDGatewaysGatewayIDDescriptionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetWifiNetworkIDGatewaysGatewayIDDescriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
