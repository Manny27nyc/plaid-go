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

// LinkTokenCreateRequestPaymentInitiation Specifies options for initializing Link for use with the Payment Initiation (Europe) product. This field is required if `payment_initiation` is included in the `products` array.
type LinkTokenCreateRequestPaymentInitiation struct {
	// The `payment_id` provided by the `/payment_initiation/payment/create` endpoint.
	PaymentId string `json:"payment_id"`
}

// NewLinkTokenCreateRequestPaymentInitiation instantiates a new LinkTokenCreateRequestPaymentInitiation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkTokenCreateRequestPaymentInitiation(paymentId string) *LinkTokenCreateRequestPaymentInitiation {
	this := LinkTokenCreateRequestPaymentInitiation{}
	this.PaymentId = paymentId
	return &this
}

// NewLinkTokenCreateRequestPaymentInitiationWithDefaults instantiates a new LinkTokenCreateRequestPaymentInitiation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkTokenCreateRequestPaymentInitiationWithDefaults() *LinkTokenCreateRequestPaymentInitiation {
	this := LinkTokenCreateRequestPaymentInitiation{}
	return &this
}

// GetPaymentId returns the PaymentId field value
func (o *LinkTokenCreateRequestPaymentInitiation) GetPaymentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value
// and a boolean to check if the value has been set.
func (o *LinkTokenCreateRequestPaymentInitiation) GetPaymentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentId, true
}

// SetPaymentId sets field value
func (o *LinkTokenCreateRequestPaymentInitiation) SetPaymentId(v string) {
	o.PaymentId = v
}

func (o LinkTokenCreateRequestPaymentInitiation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["payment_id"] = o.PaymentId
	}
	return json.Marshal(toSerialize)
}

type NullableLinkTokenCreateRequestPaymentInitiation struct {
	value *LinkTokenCreateRequestPaymentInitiation
	isSet bool
}

func (v NullableLinkTokenCreateRequestPaymentInitiation) Get() *LinkTokenCreateRequestPaymentInitiation {
	return v.value
}

func (v *NullableLinkTokenCreateRequestPaymentInitiation) Set(val *LinkTokenCreateRequestPaymentInitiation) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkTokenCreateRequestPaymentInitiation) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkTokenCreateRequestPaymentInitiation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkTokenCreateRequestPaymentInitiation(val *LinkTokenCreateRequestPaymentInitiation) *NullableLinkTokenCreateRequestPaymentInitiation {
	return &NullableLinkTokenCreateRequestPaymentInitiation{value: val, isSet: true}
}

func (v NullableLinkTokenCreateRequestPaymentInitiation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkTokenCreateRequestPaymentInitiation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


