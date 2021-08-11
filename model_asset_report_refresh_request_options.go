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

// AssetReportRefreshRequestOptions An optional object to filter `/asset_report/refresh` results. If provided, cannot be `null`. If not specified, the `options` from the original call to `/asset_report/create` will be used.
type AssetReportRefreshRequestOptions struct {
	// Client-generated identifier, which can be used by lenders to track loan applications.
	ClientReportId *string `json:"client_report_id,omitempty"`
	// URL to which Plaid will send Assets webhooks, for example when the requested Asset Report is ready.
	Webhook *string `json:"webhook,omitempty"`
	User *AssetReportUser `json:"user,omitempty"`
}

// NewAssetReportRefreshRequestOptions instantiates a new AssetReportRefreshRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportRefreshRequestOptions() *AssetReportRefreshRequestOptions {
	this := AssetReportRefreshRequestOptions{}
	return &this
}

// NewAssetReportRefreshRequestOptionsWithDefaults instantiates a new AssetReportRefreshRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportRefreshRequestOptionsWithDefaults() *AssetReportRefreshRequestOptions {
	this := AssetReportRefreshRequestOptions{}
	return &this
}

// GetClientReportId returns the ClientReportId field value if set, zero value otherwise.
func (o *AssetReportRefreshRequestOptions) GetClientReportId() string {
	if o == nil || o.ClientReportId == nil {
		var ret string
		return ret
	}
	return *o.ClientReportId
}

// GetClientReportIdOk returns a tuple with the ClientReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportRefreshRequestOptions) GetClientReportIdOk() (*string, bool) {
	if o == nil || o.ClientReportId == nil {
		return nil, false
	}
	return o.ClientReportId, true
}

// HasClientReportId returns a boolean if a field has been set.
func (o *AssetReportRefreshRequestOptions) HasClientReportId() bool {
	if o != nil && o.ClientReportId != nil {
		return true
	}

	return false
}

// SetClientReportId gets a reference to the given string and assigns it to the ClientReportId field.
func (o *AssetReportRefreshRequestOptions) SetClientReportId(v string) {
	o.ClientReportId = &v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *AssetReportRefreshRequestOptions) GetWebhook() string {
	if o == nil || o.Webhook == nil {
		var ret string
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportRefreshRequestOptions) GetWebhookOk() (*string, bool) {
	if o == nil || o.Webhook == nil {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *AssetReportRefreshRequestOptions) HasWebhook() bool {
	if o != nil && o.Webhook != nil {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given string and assigns it to the Webhook field.
func (o *AssetReportRefreshRequestOptions) SetWebhook(v string) {
	o.Webhook = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *AssetReportRefreshRequestOptions) GetUser() AssetReportUser {
	if o == nil || o.User == nil {
		var ret AssetReportUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportRefreshRequestOptions) GetUserOk() (*AssetReportUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *AssetReportRefreshRequestOptions) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given AssetReportUser and assigns it to the User field.
func (o *AssetReportRefreshRequestOptions) SetUser(v AssetReportUser) {
	o.User = &v
}

func (o AssetReportRefreshRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientReportId != nil {
		toSerialize["client_report_id"] = o.ClientReportId
	}
	if o.Webhook != nil {
		toSerialize["webhook"] = o.Webhook
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableAssetReportRefreshRequestOptions struct {
	value *AssetReportRefreshRequestOptions
	isSet bool
}

func (v NullableAssetReportRefreshRequestOptions) Get() *AssetReportRefreshRequestOptions {
	return v.value
}

func (v *NullableAssetReportRefreshRequestOptions) Set(val *AssetReportRefreshRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportRefreshRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportRefreshRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportRefreshRequestOptions(val *AssetReportRefreshRequestOptions) *NullableAssetReportRefreshRequestOptions {
	return &NullableAssetReportRefreshRequestOptions{value: val, isSet: true}
}

func (v NullableAssetReportRefreshRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportRefreshRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


