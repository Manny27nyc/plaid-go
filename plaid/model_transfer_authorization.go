/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.58.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransferAuthorization TransferAuthorization contains the authorization decision for a proposed transfer
type TransferAuthorization struct {
	// Plaid’s unique identifier for a transfer authorization.
	Id string `json:"id"`
	// The datetime representing when the authorization was created, in the format \"2006-01-02T15:04:05Z\".
	Created string `json:"created"`
	//  A decision regarding the proposed transfer.  `approved` – The proposed transfer has received the end user's consent and has been approved for processing. Plaid has also reviewed the proposed transfer and has approved it for processing.   `permitted` – Plaid was unable to fetch the information required to approve or decline the proposed transfer. You may proceed with the transfer, but further review is recommended. Plaid is awaiting further instructions from the client.  `declined` – Plaid reviewed the proposed transfer and declined processing. Refer to the `code` field in the `decision_rationale` object for details.
	Decision string `json:"decision"`
	DecisionRationale NullableTransferAuthorizationDecisionRationale `json:"decision_rationale"`
	ProposedTransfer TransferAuthorizationProposedTransfer `json:"proposed_transfer"`
	AdditionalProperties map[string]interface{}
}

type _TransferAuthorization TransferAuthorization

// NewTransferAuthorization instantiates a new TransferAuthorization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferAuthorization(id string, created string, decision string, decisionRationale NullableTransferAuthorizationDecisionRationale, proposedTransfer TransferAuthorizationProposedTransfer) *TransferAuthorization {
	this := TransferAuthorization{}
	this.Id = id
	this.Created = created
	this.Decision = decision
	this.DecisionRationale = decisionRationale
	this.ProposedTransfer = proposedTransfer
	return &this
}

// NewTransferAuthorizationWithDefaults instantiates a new TransferAuthorization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferAuthorizationWithDefaults() *TransferAuthorization {
	this := TransferAuthorization{}
	return &this
}

// GetId returns the Id field value
func (o *TransferAuthorization) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorization) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TransferAuthorization) SetId(v string) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *TransferAuthorization) GetCreated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorization) GetCreatedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *TransferAuthorization) SetCreated(v string) {
	o.Created = v
}

// GetDecision returns the Decision field value
func (o *TransferAuthorization) GetDecision() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorization) GetDecisionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Decision, true
}

// SetDecision sets field value
func (o *TransferAuthorization) SetDecision(v string) {
	o.Decision = v
}

// GetDecisionRationale returns the DecisionRationale field value
// If the value is explicit nil, the zero value for TransferAuthorizationDecisionRationale will be returned
func (o *TransferAuthorization) GetDecisionRationale() TransferAuthorizationDecisionRationale {
	if o == nil || o.DecisionRationale.Get() == nil {
		var ret TransferAuthorizationDecisionRationale
		return ret
	}

	return *o.DecisionRationale.Get()
}

// GetDecisionRationaleOk returns a tuple with the DecisionRationale field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferAuthorization) GetDecisionRationaleOk() (*TransferAuthorizationDecisionRationale, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DecisionRationale.Get(), o.DecisionRationale.IsSet()
}

// SetDecisionRationale sets field value
func (o *TransferAuthorization) SetDecisionRationale(v TransferAuthorizationDecisionRationale) {
	o.DecisionRationale.Set(&v)
}

// GetProposedTransfer returns the ProposedTransfer field value
func (o *TransferAuthorization) GetProposedTransfer() TransferAuthorizationProposedTransfer {
	if o == nil {
		var ret TransferAuthorizationProposedTransfer
		return ret
	}

	return o.ProposedTransfer
}

// GetProposedTransferOk returns a tuple with the ProposedTransfer field value
// and a boolean to check if the value has been set.
func (o *TransferAuthorization) GetProposedTransferOk() (*TransferAuthorizationProposedTransfer, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProposedTransfer, true
}

// SetProposedTransfer sets field value
func (o *TransferAuthorization) SetProposedTransfer(v TransferAuthorizationProposedTransfer) {
	o.ProposedTransfer = v
}

func (o TransferAuthorization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["decision"] = o.Decision
	}
	if true {
		toSerialize["decision_rationale"] = o.DecisionRationale.Get()
	}
	if true {
		toSerialize["proposed_transfer"] = o.ProposedTransfer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferAuthorization) UnmarshalJSON(bytes []byte) (err error) {
	varTransferAuthorization := _TransferAuthorization{}

	if err = json.Unmarshal(bytes, &varTransferAuthorization); err == nil {
		*o = TransferAuthorization(varTransferAuthorization)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created")
		delete(additionalProperties, "decision")
		delete(additionalProperties, "decision_rationale")
		delete(additionalProperties, "proposed_transfer")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferAuthorization struct {
	value *TransferAuthorization
	isSet bool
}

func (v NullableTransferAuthorization) Get() *TransferAuthorization {
	return v.value
}

func (v *NullableTransferAuthorization) Set(val *TransferAuthorization) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferAuthorization) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferAuthorization(val *TransferAuthorization) *NullableTransferAuthorization {
	return &NullableTransferAuthorization{value: val, isSet: true}
}

func (v NullableTransferAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


