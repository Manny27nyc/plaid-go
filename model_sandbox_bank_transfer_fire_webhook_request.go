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

// SandboxBankTransferFireWebhookRequest SandboxBankTransferFireWebhookRequest defines the request schema for `/sandbox/bank_transfer/fire_webhook`
type SandboxBankTransferFireWebhookRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The URL to which the webhook should be sent.
	Webhook string `json:"webhook"`
}

// NewSandboxBankTransferFireWebhookRequest instantiates a new SandboxBankTransferFireWebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSandboxBankTransferFireWebhookRequest(webhook string) *SandboxBankTransferFireWebhookRequest {
	this := SandboxBankTransferFireWebhookRequest{}
	this.Webhook = webhook
	return &this
}

// NewSandboxBankTransferFireWebhookRequestWithDefaults instantiates a new SandboxBankTransferFireWebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSandboxBankTransferFireWebhookRequestWithDefaults() *SandboxBankTransferFireWebhookRequest {
	this := SandboxBankTransferFireWebhookRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *SandboxBankTransferFireWebhookRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxBankTransferFireWebhookRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *SandboxBankTransferFireWebhookRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *SandboxBankTransferFireWebhookRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *SandboxBankTransferFireWebhookRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SandboxBankTransferFireWebhookRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *SandboxBankTransferFireWebhookRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *SandboxBankTransferFireWebhookRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetWebhook returns the Webhook field value
func (o *SandboxBankTransferFireWebhookRequest) GetWebhook() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value
// and a boolean to check if the value has been set.
func (o *SandboxBankTransferFireWebhookRequest) GetWebhookOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Webhook, true
}

// SetWebhook sets field value
func (o *SandboxBankTransferFireWebhookRequest) SetWebhook(v string) {
	o.Webhook = v
}

func (o SandboxBankTransferFireWebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["webhook"] = o.Webhook
	}
	return json.Marshal(toSerialize)
}

type NullableSandboxBankTransferFireWebhookRequest struct {
	value *SandboxBankTransferFireWebhookRequest
	isSet bool
}

func (v NullableSandboxBankTransferFireWebhookRequest) Get() *SandboxBankTransferFireWebhookRequest {
	return v.value
}

func (v *NullableSandboxBankTransferFireWebhookRequest) Set(val *SandboxBankTransferFireWebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSandboxBankTransferFireWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSandboxBankTransferFireWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSandboxBankTransferFireWebhookRequest(val *SandboxBankTransferFireWebhookRequest) *NullableSandboxBankTransferFireWebhookRequest {
	return &NullableSandboxBankTransferFireWebhookRequest{value: val, isSet: true}
}

func (v NullableSandboxBankTransferFireWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSandboxBankTransferFireWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


