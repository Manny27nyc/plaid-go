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

// DepositSwitchCreateRequest DepositSwitchCreateRequest defines the request schema for `/deposit_switch/create`
type DepositSwitchCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// Access token for the target Item, typically provided in the Import Item response. 
	TargetAccessToken string `json:"target_access_token"`
	// Plaid Account ID that specifies the target bank account. This account will become the recipient for a user's direct deposit.
	TargetAccountId string `json:"target_account_id"`
	// ISO-3166-1 alpha-2 country code standard.
	CountryCode NullableString `json:"country_code,omitempty"`
	Options *DepositSwitchCreateRequestOptions `json:"options,omitempty"`
}

// NewDepositSwitchCreateRequest instantiates a new DepositSwitchCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDepositSwitchCreateRequest(targetAccessToken string, targetAccountId string) *DepositSwitchCreateRequest {
	this := DepositSwitchCreateRequest{}
	this.TargetAccessToken = targetAccessToken
	this.TargetAccountId = targetAccountId
	return &this
}

// NewDepositSwitchCreateRequestWithDefaults instantiates a new DepositSwitchCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositSwitchCreateRequestWithDefaults() *DepositSwitchCreateRequest {
	this := DepositSwitchCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *DepositSwitchCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositSwitchCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *DepositSwitchCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *DepositSwitchCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *DepositSwitchCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositSwitchCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *DepositSwitchCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *DepositSwitchCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetTargetAccessToken returns the TargetAccessToken field value
func (o *DepositSwitchCreateRequest) GetTargetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetAccessToken
}

// GetTargetAccessTokenOk returns a tuple with the TargetAccessToken field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchCreateRequest) GetTargetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetAccessToken, true
}

// SetTargetAccessToken sets field value
func (o *DepositSwitchCreateRequest) SetTargetAccessToken(v string) {
	o.TargetAccessToken = v
}

// GetTargetAccountId returns the TargetAccountId field value
func (o *DepositSwitchCreateRequest) GetTargetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetAccountId
}

// GetTargetAccountIdOk returns a tuple with the TargetAccountId field value
// and a boolean to check if the value has been set.
func (o *DepositSwitchCreateRequest) GetTargetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TargetAccountId, true
}

// SetTargetAccountId sets field value
func (o *DepositSwitchCreateRequest) SetTargetAccountId(v string) {
	o.TargetAccountId = v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DepositSwitchCreateRequest) GetCountryCode() string {
	if o == nil || o.CountryCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.CountryCode.Get()
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DepositSwitchCreateRequest) GetCountryCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CountryCode.Get(), o.CountryCode.IsSet()
}

// HasCountryCode returns a boolean if a field has been set.
func (o *DepositSwitchCreateRequest) HasCountryCode() bool {
	if o != nil && o.CountryCode.IsSet() {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given NullableString and assigns it to the CountryCode field.
func (o *DepositSwitchCreateRequest) SetCountryCode(v string) {
	o.CountryCode.Set(&v)
}
// SetCountryCodeNil sets the value for CountryCode to be an explicit nil
func (o *DepositSwitchCreateRequest) SetCountryCodeNil() {
	o.CountryCode.Set(nil)
}

// UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
func (o *DepositSwitchCreateRequest) UnsetCountryCode() {
	o.CountryCode.Unset()
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *DepositSwitchCreateRequest) GetOptions() DepositSwitchCreateRequestOptions {
	if o == nil || o.Options == nil {
		var ret DepositSwitchCreateRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DepositSwitchCreateRequest) GetOptionsOk() (*DepositSwitchCreateRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *DepositSwitchCreateRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given DepositSwitchCreateRequestOptions and assigns it to the Options field.
func (o *DepositSwitchCreateRequest) SetOptions(v DepositSwitchCreateRequestOptions) {
	o.Options = &v
}

func (o DepositSwitchCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["target_access_token"] = o.TargetAccessToken
	}
	if true {
		toSerialize["target_account_id"] = o.TargetAccountId
	}
	if o.CountryCode.IsSet() {
		toSerialize["country_code"] = o.CountryCode.Get()
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableDepositSwitchCreateRequest struct {
	value *DepositSwitchCreateRequest
	isSet bool
}

func (v NullableDepositSwitchCreateRequest) Get() *DepositSwitchCreateRequest {
	return v.value
}

func (v *NullableDepositSwitchCreateRequest) Set(val *DepositSwitchCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDepositSwitchCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDepositSwitchCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDepositSwitchCreateRequest(val *DepositSwitchCreateRequest) *NullableDepositSwitchCreateRequest {
	return &NullableDepositSwitchCreateRequest{value: val, isSet: true}
}

func (v NullableDepositSwitchCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDepositSwitchCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


