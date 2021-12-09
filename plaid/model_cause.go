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

// Cause An error object and associated `item_id` used to identify a specific Item and error when a batch operation operating on multiple Items has encountered an error in one of the Items.
type Cause struct {
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	Error NullableError `json:"error"`
	AdditionalProperties map[string]interface{}
}

type _Cause Cause

// NewCause instantiates a new Cause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCause(itemId string, error_ NullableError) *Cause {
	this := Cause{}
	this.ItemId = itemId
	this.Error = error_
	return &this
}

// NewCauseWithDefaults instantiates a new Cause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCauseWithDefaults() *Cause {
	this := Cause{}
	return &this
}

// GetItemId returns the ItemId field value
func (o *Cause) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *Cause) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *Cause) SetItemId(v string) {
	o.ItemId = v
}

// GetError returns the Error field value
// If the value is explicit nil, the zero value for Error will be returned
func (o *Cause) GetError() Error {
	if o == nil || o.Error.Get() == nil {
		var ret Error
		return ret
	}

	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Cause) GetErrorOk() (*Error, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// SetError sets field value
func (o *Cause) SetError(v Error) {
	o.Error.Set(&v)
}

func (o Cause) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["item_id"] = o.ItemId
	}
	if true {
		toSerialize["error"] = o.Error.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Cause) UnmarshalJSON(bytes []byte) (err error) {
	varCause := _Cause{}

	if err = json.Unmarshal(bytes, &varCause); err == nil {
		*o = Cause(varCause)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCause struct {
	value *Cause
	isSet bool
}

func (v NullableCause) Get() *Cause {
	return v.value
}

func (v *NullableCause) Set(val *Cause) {
	v.value = val
	v.isSet = true
}

func (v NullableCause) IsSet() bool {
	return v.isSet
}

func (v *NullableCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCause(val *Cause) *NullableCause {
	return &NullableCause{value: val, isSet: true}
}

func (v NullableCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


