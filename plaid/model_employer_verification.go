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

// EmployerVerification An object containing employer data.
type EmployerVerification struct {
	// Name of employer.
	Name NullableString `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmployerVerification EmployerVerification

// NewEmployerVerification instantiates a new EmployerVerification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployerVerification() *EmployerVerification {
	this := EmployerVerification{}
	return &this
}

// NewEmployerVerificationWithDefaults instantiates a new EmployerVerification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployerVerificationWithDefaults() *EmployerVerification {
	this := EmployerVerification{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmployerVerification) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmployerVerification) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *EmployerVerification) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *EmployerVerification) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *EmployerVerification) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *EmployerVerification) UnsetName() {
	o.Name.Unset()
}

func (o EmployerVerification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EmployerVerification) UnmarshalJSON(bytes []byte) (err error) {
	varEmployerVerification := _EmployerVerification{}

	if err = json.Unmarshal(bytes, &varEmployerVerification); err == nil {
		*o = EmployerVerification(varEmployerVerification)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmployerVerification struct {
	value *EmployerVerification
	isSet bool
}

func (v NullableEmployerVerification) Get() *EmployerVerification {
	return v.value
}

func (v *NullableEmployerVerification) Set(val *EmployerVerification) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployerVerification) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployerVerification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployerVerification(val *EmployerVerification) *NullableEmployerVerification {
	return &NullableEmployerVerification{value: val, isSet: true}
}

func (v NullableEmployerVerification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployerVerification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


