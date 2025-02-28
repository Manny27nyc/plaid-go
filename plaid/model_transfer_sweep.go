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
	"time"
)

// TransferSweep Describes a sweep of funds to / from the sweep account.  A sweep is associated with many sweep events (events of type `swept` or `reverse_swept`) which can be retrieved by invoking the `/transfer/event/list` endpoint with the corresponding `sweep_id`.  `swept` events occur when the transfer amount is credited or debited from your sweep account, depending on the `type` of the transfer. `reverse_swept` events occur when a transfer is reversed and Plaid undoes the credit or debit.  The total sum of the `swept` and `reverse_swept` events is equal to the `amount` of the sweep Plaid creates and matches the amount of the entry on your sweep account ledger.
type TransferSweep struct {
	// Identifier of the sweep.
	Id string `json:"id"`
	// The datetime when the sweep occurred, in RFC 3339 format.
	CreatedAt time.Time `json:"created_at"`
	// Signed decimal amount of the sweep as it appears on your sweep account ledger (e.g. \"-10.00\")  If amount is not present, the sweep was net-settled to zero and outstanding debits and credits between the sweep account and Plaid are balanced.
	Amount *string `json:"amount,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransferSweep TransferSweep

// NewTransferSweep instantiates a new TransferSweep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferSweep(id string, createdAt time.Time) *TransferSweep {
	this := TransferSweep{}
	this.Id = id
	this.CreatedAt = createdAt
	return &this
}

// NewTransferSweepWithDefaults instantiates a new TransferSweep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferSweepWithDefaults() *TransferSweep {
	this := TransferSweep{}
	return &this
}

// GetId returns the Id field value
func (o *TransferSweep) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TransferSweep) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TransferSweep) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TransferSweep) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TransferSweep) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TransferSweep) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *TransferSweep) GetAmount() string {
	if o == nil || o.Amount == nil {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferSweep) GetAmountOk() (*string, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *TransferSweep) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *TransferSweep) SetAmount(v string) {
	o.Amount = &v
}

func (o TransferSweep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferSweep) UnmarshalJSON(bytes []byte) (err error) {
	varTransferSweep := _TransferSweep{}

	if err = json.Unmarshal(bytes, &varTransferSweep); err == nil {
		*o = TransferSweep(varTransferSweep)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "amount")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferSweep struct {
	value *TransferSweep
	isSet bool
}

func (v NullableTransferSweep) Get() *TransferSweep {
	return v.value
}

func (v *NullableTransferSweep) Set(val *TransferSweep) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferSweep) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferSweep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferSweep(val *TransferSweep) *NullableTransferSweep {
	return &NullableTransferSweep{value: val, isSet: true}
}

func (v NullableTransferSweep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferSweep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


