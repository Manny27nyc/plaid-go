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

// DepositSwitchCreateRequestOptions Options to configure the `/deposit_switch/create` request. If provided, cannot be `null`.
type DepositSwitchCreateRequestOptions struct {
	// The URL registered to receive webhooks when the status of a deposit switch request has changed. 
	Webhook NullableString `json:"webhook,omitempty"`
	// An array of access tokens corresponding to transaction items to use when attempting to match the user to their Payroll Provider. These tokens must be created by the same client id as the one creating the switch, and have access to the transactions product.
	TransactionItemAccessTokens *[]string `json:"transaction_item_access_tokens,omitempty"`
}

// NewDepositSwitchCreateRequestOptions instantiates a new DepositSwitchCreateRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositSwitchCreateRequestOptions() *DepositSwitchCreateRequestOptions {
	this := DepositSwitchCreateRequestOptions{}
	return &this
}

// NewDepositSwitchCreateRequestOptionsWithDefaults instantiates a new DepositSwitchCreateRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositSwitchCreateRequestOptionsWithDefaults() *DepositSwitchCreateRequestOptions {
	this := DepositSwitchCreateRequestOptions{}
	return &this
}

// GetWebhook returns the Webhook field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DepositSwitchCreateRequestOptions) GetWebhook() string {
	if o == nil || o.Webhook.Get() == nil {
		var ret string
		return ret
	}
	return *o.Webhook.Get()
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositSwitchCreateRequestOptions) GetWebhookOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Webhook.Get(), o.Webhook.IsSet()
}

// HasWebhook returns a boolean if a field has been set.
func (o *DepositSwitchCreateRequestOptions) HasWebhook() bool {
	if o != nil && o.Webhook.IsSet() {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given NullableString and assigns it to the Webhook field.
func (o *DepositSwitchCreateRequestOptions) SetWebhook(v string) {
	o.Webhook.Set(&v)
}
// SetWebhookNil sets the value for Webhook to be an explicit nil
func (o *DepositSwitchCreateRequestOptions) SetWebhookNil() {
	o.Webhook.Set(nil)
}

// UnsetWebhook ensures that no value is present for Webhook, not even an explicit nil
func (o *DepositSwitchCreateRequestOptions) UnsetWebhook() {
	o.Webhook.Unset()
}

// GetTransactionItemAccessTokens returns the TransactionItemAccessTokens field value if set, zero value otherwise.
func (o *DepositSwitchCreateRequestOptions) GetTransactionItemAccessTokens() []string {
	if o == nil || o.TransactionItemAccessTokens == nil {
		var ret []string
		return ret
	}
	return *o.TransactionItemAccessTokens
}

// GetTransactionItemAccessTokensOk returns a tuple with the TransactionItemAccessTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositSwitchCreateRequestOptions) GetTransactionItemAccessTokensOk() (*[]string, bool) {
	if o == nil || o.TransactionItemAccessTokens == nil {
		return nil, false
	}
	return o.TransactionItemAccessTokens, true
}

// HasTransactionItemAccessTokens returns a boolean if a field has been set.
func (o *DepositSwitchCreateRequestOptions) HasTransactionItemAccessTokens() bool {
	if o != nil && o.TransactionItemAccessTokens != nil {
		return true
	}

	return false
}

// SetTransactionItemAccessTokens gets a reference to the given []string and assigns it to the TransactionItemAccessTokens field.
func (o *DepositSwitchCreateRequestOptions) SetTransactionItemAccessTokens(v []string) {
	o.TransactionItemAccessTokens = &v
}

func (o DepositSwitchCreateRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Webhook.IsSet() {
		toSerialize["webhook"] = o.Webhook.Get()
	}
	if o.TransactionItemAccessTokens != nil {
		toSerialize["transaction_item_access_tokens"] = o.TransactionItemAccessTokens
	}
	return json.Marshal(toSerialize)
}

type NullableDepositSwitchCreateRequestOptions struct {
	value *DepositSwitchCreateRequestOptions
	isSet bool
}

func (v NullableDepositSwitchCreateRequestOptions) Get() *DepositSwitchCreateRequestOptions {
	return v.value
}

func (v *NullableDepositSwitchCreateRequestOptions) Set(val *DepositSwitchCreateRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositSwitchCreateRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositSwitchCreateRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositSwitchCreateRequestOptions(val *DepositSwitchCreateRequestOptions) *NullableDepositSwitchCreateRequestOptions {
	return &NullableDepositSwitchCreateRequestOptions{value: val, isSet: true}
}

func (v NullableDepositSwitchCreateRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositSwitchCreateRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


