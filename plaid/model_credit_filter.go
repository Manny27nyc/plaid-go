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

// CreditFilter A filter to apply to `credit`-type accounts
type CreditFilter struct {
	// An array of account subtypes to display in Link. If not specified, all account subtypes will be shown. For a full list of valid types and subtypes, see the [Account schema](https://plaid.com/docs/api/accounts#accounts-schema). 
	AccountSubtypes []AccountSubtype `json:"account_subtypes"`
	AdditionalProperties map[string]interface{}
}

type _CreditFilter CreditFilter

// NewCreditFilter instantiates a new CreditFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreditFilter(accountSubtypes []AccountSubtype) *CreditFilter {
	this := CreditFilter{}
	this.AccountSubtypes = accountSubtypes
	return &this
}

// NewCreditFilterWithDefaults instantiates a new CreditFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreditFilterWithDefaults() *CreditFilter {
	this := CreditFilter{}
	return &this
}

// GetAccountSubtypes returns the AccountSubtypes field value
func (o *CreditFilter) GetAccountSubtypes() []AccountSubtype {
	if o == nil {
		var ret []AccountSubtype
		return ret
	}

	return o.AccountSubtypes
}

// GetAccountSubtypesOk returns a tuple with the AccountSubtypes field value
// and a boolean to check if the value has been set.
func (o *CreditFilter) GetAccountSubtypesOk() (*[]AccountSubtype, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountSubtypes, true
}

// SetAccountSubtypes sets field value
func (o *CreditFilter) SetAccountSubtypes(v []AccountSubtype) {
	o.AccountSubtypes = v
}

func (o CreditFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["account_subtypes"] = o.AccountSubtypes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreditFilter) UnmarshalJSON(bytes []byte) (err error) {
	varCreditFilter := _CreditFilter{}

	if err = json.Unmarshal(bytes, &varCreditFilter); err == nil {
		*o = CreditFilter(varCreditFilter)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "account_subtypes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreditFilter struct {
	value *CreditFilter
	isSet bool
}

func (v NullableCreditFilter) Get() *CreditFilter {
	return v.value
}

func (v *NullableCreditFilter) Set(val *CreditFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableCreditFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableCreditFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreditFilter(val *CreditFilter) *NullableCreditFilter {
	return &NullableCreditFilter{value: val, isSet: true}
}

func (v NullableCreditFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreditFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


