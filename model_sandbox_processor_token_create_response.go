/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.20.6
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// SandboxProcessorTokenCreateResponse struct for SandboxProcessorTokenCreateResponse
type SandboxProcessorTokenCreateResponse struct {
	// A processor token that can be used to call the `/processor/` endpoints.
	ProcessorToken string `json:"processor_token"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _SandboxProcessorTokenCreateResponse SandboxProcessorTokenCreateResponse

// NewSandboxProcessorTokenCreateResponse instantiates a new SandboxProcessorTokenCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxProcessorTokenCreateResponse(processorToken string, requestId string) *SandboxProcessorTokenCreateResponse {
	this := SandboxProcessorTokenCreateResponse{}
	this.ProcessorToken = processorToken
	this.RequestId = requestId
	return &this
}

// NewSandboxProcessorTokenCreateResponseWithDefaults instantiates a new SandboxProcessorTokenCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxProcessorTokenCreateResponseWithDefaults() *SandboxProcessorTokenCreateResponse {
	this := SandboxProcessorTokenCreateResponse{}
	return &this
}

// GetProcessorToken returns the ProcessorToken field value
func (o *SandboxProcessorTokenCreateResponse) GetProcessorToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProcessorToken
}

// GetProcessorTokenOk returns a tuple with the ProcessorToken field value
// and a boolean to check if the value has been set.
func (o *SandboxProcessorTokenCreateResponse) GetProcessorTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ProcessorToken, true
}

// SetProcessorToken sets field value
func (o *SandboxProcessorTokenCreateResponse) SetProcessorToken(v string) {
	o.ProcessorToken = v
}

// GetRequestId returns the RequestId field value
func (o *SandboxProcessorTokenCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *SandboxProcessorTokenCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *SandboxProcessorTokenCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o SandboxProcessorTokenCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["processor_token"] = o.ProcessorToken
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SandboxProcessorTokenCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSandboxProcessorTokenCreateResponse := _SandboxProcessorTokenCreateResponse{}

	if err = json.Unmarshal(bytes, &varSandboxProcessorTokenCreateResponse); err == nil {
		*o = SandboxProcessorTokenCreateResponse(varSandboxProcessorTokenCreateResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "processor_token")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSandboxProcessorTokenCreateResponse struct {
	value *SandboxProcessorTokenCreateResponse
	isSet bool
}

func (v NullableSandboxProcessorTokenCreateResponse) Get() *SandboxProcessorTokenCreateResponse {
	return v.value
}

func (v *NullableSandboxProcessorTokenCreateResponse) Set(val *SandboxProcessorTokenCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxProcessorTokenCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxProcessorTokenCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxProcessorTokenCreateResponse(val *SandboxProcessorTokenCreateResponse) *NullableSandboxProcessorTokenCreateResponse {
	return &NullableSandboxProcessorTokenCreateResponse{value: val, isSet: true}
}

func (v NullableSandboxProcessorTokenCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxProcessorTokenCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


