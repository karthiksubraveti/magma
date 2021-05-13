// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// DeleteNetworksNetworkIDPoliciesRulesRuleIDReader is a Reader for the DeleteNetworksNetworkIDPoliciesRulesRuleID structure.
type DeleteNetworksNetworkIDPoliciesRulesRuleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworksNetworkIDPoliciesRulesRuleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteNetworksNetworkIDPoliciesRulesRuleIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteNetworksNetworkIDPoliciesRulesRuleIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNetworksNetworkIDPoliciesRulesRuleIDNoContent creates a DeleteNetworksNetworkIDPoliciesRulesRuleIDNoContent with default headers values
func NewDeleteNetworksNetworkIDPoliciesRulesRuleIDNoContent() *DeleteNetworksNetworkIDPoliciesRulesRuleIDNoContent {
	return &DeleteNetworksNetworkIDPoliciesRulesRuleIDNoContent{}
}

/*DeleteNetworksNetworkIDPoliciesRulesRuleIDNoContent handles this case with default header values.

Success
*/
type DeleteNetworksNetworkIDPoliciesRulesRuleIDNoContent struct {
}

func (o *DeleteNetworksNetworkIDPoliciesRulesRuleIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /networks/{network_id}/policies/rules/{rule_id}][%d] deleteNetworksNetworkIdPoliciesRulesRuleIdNoContent ", 204)
}

func (o *DeleteNetworksNetworkIDPoliciesRulesRuleIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNetworksNetworkIDPoliciesRulesRuleIDDefault creates a DeleteNetworksNetworkIDPoliciesRulesRuleIDDefault with default headers values
func NewDeleteNetworksNetworkIDPoliciesRulesRuleIDDefault(code int) *DeleteNetworksNetworkIDPoliciesRulesRuleIDDefault {
	return &DeleteNetworksNetworkIDPoliciesRulesRuleIDDefault{
		_statusCode: code,
	}
}

/*DeleteNetworksNetworkIDPoliciesRulesRuleIDDefault handles this case with default header values.

Unexpected Error
*/
type DeleteNetworksNetworkIDPoliciesRulesRuleIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the delete networks network ID policies rules rule ID default response
func (o *DeleteNetworksNetworkIDPoliciesRulesRuleIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNetworksNetworkIDPoliciesRulesRuleIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /networks/{network_id}/policies/rules/{rule_id}][%d] DeleteNetworksNetworkIDPoliciesRulesRuleID default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteNetworksNetworkIDPoliciesRulesRuleIDDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteNetworksNetworkIDPoliciesRulesRuleIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}