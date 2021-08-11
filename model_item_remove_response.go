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

// ItemRemoveResponse ItemRemoveResponse defines the response schema for `/item/remove`
type ItemRemoveResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _ItemRemoveResponse ItemRemoveResponse

// NewItemRemoveResponse instantiates a new ItemRemoveResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemRemoveResponse(requestId string) *ItemRemoveResponse {
	this := ItemRemoveResponse{}
	this.RequestId = requestId
	return &this
}

// NewItemRemoveResponseWithDefaults instantiates a new ItemRemoveResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemRemoveResponseWithDefaults() *ItemRemoveResponse {
	this := ItemRemoveResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *ItemRemoveResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *ItemRemoveResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *ItemRemoveResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o ItemRemoveResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ItemRemoveResponse) UnmarshalJSON(bytes []byte) (err error) {
	varItemRemoveResponse := _ItemRemoveResponse{}

	if err = json.Unmarshal(bytes, &varItemRemoveResponse); err == nil {
		*o = ItemRemoveResponse(varItemRemoveResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableItemRemoveResponse struct {
	value *ItemRemoveResponse
	isSet bool
}

func (v NullableItemRemoveResponse) Get() *ItemRemoveResponse {
	return v.value
}

func (v *NullableItemRemoveResponse) Set(val *ItemRemoveResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableItemRemoveResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableItemRemoveResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemRemoveResponse(val *ItemRemoveResponse) *NullableItemRemoveResponse {
	return &NullableItemRemoveResponse{value: val, isSet: true}
}

func (v NullableItemRemoveResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemRemoveResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


