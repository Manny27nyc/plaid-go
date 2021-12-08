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

// ProcessorAuthGetResponse ProcessorAuthGetResponse defines the response schema for `/processor/auth/get`
type ProcessorAuthGetResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	Numbers ProcessorNumber `json:"numbers"`
	Account AccountBase `json:"account"`
	AdditionalProperties map[string]interface{}
}

type _ProcessorAuthGetResponse ProcessorAuthGetResponse

// NewProcessorAuthGetResponse instantiates a new ProcessorAuthGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessorAuthGetResponse(requestId string, numbers ProcessorNumber, account AccountBase) *ProcessorAuthGetResponse {
	this := ProcessorAuthGetResponse{}
	this.RequestId = requestId
	this.Numbers = numbers
	this.Account = account
	return &this
}

// NewProcessorAuthGetResponseWithDefaults instantiates a new ProcessorAuthGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessorAuthGetResponseWithDefaults() *ProcessorAuthGetResponse {
	this := ProcessorAuthGetResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *ProcessorAuthGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ProcessorAuthGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ProcessorAuthGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

// GetNumbers returns the Numbers field value
func (o *ProcessorAuthGetResponse) GetNumbers() ProcessorNumber {
	if o == nil {
		var ret ProcessorNumber
		return ret
	}

	return o.Numbers
}

// GetNumbersOk returns a tuple with the Numbers field value
// and a boolean to check if the value has been set.
func (o *ProcessorAuthGetResponse) GetNumbersOk() (*ProcessorNumber, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Numbers, true
}

// SetNumbers sets field value
func (o *ProcessorAuthGetResponse) SetNumbers(v ProcessorNumber) {
	o.Numbers = v
}

// GetAccount returns the Account field value
func (o *ProcessorAuthGetResponse) GetAccount() AccountBase {
	if o == nil {
		var ret AccountBase
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *ProcessorAuthGetResponse) GetAccountOk() (*AccountBase, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *ProcessorAuthGetResponse) SetAccount(v AccountBase) {
	o.Account = v
}

func (o ProcessorAuthGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}
	if true {
		toSerialize["numbers"] = o.Numbers
	}
	if true {
		toSerialize["account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ProcessorAuthGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varProcessorAuthGetResponse := _ProcessorAuthGetResponse{}

	if err = json.Unmarshal(bytes, &varProcessorAuthGetResponse); err == nil {
		*o = ProcessorAuthGetResponse(varProcessorAuthGetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		delete(additionalProperties, "numbers")
		delete(additionalProperties, "account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProcessorAuthGetResponse struct {
	value *ProcessorAuthGetResponse
	isSet bool
}

func (v NullableProcessorAuthGetResponse) Get() *ProcessorAuthGetResponse {
	return v.value
}

func (v *NullableProcessorAuthGetResponse) Set(val *ProcessorAuthGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessorAuthGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessorAuthGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessorAuthGetResponse(val *ProcessorAuthGetResponse) *NullableProcessorAuthGetResponse {
	return &NullableProcessorAuthGetResponse{value: val, isSet: true}
}

func (v NullableProcessorAuthGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessorAuthGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


