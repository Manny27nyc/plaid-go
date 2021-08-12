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
	"fmt"
)

// PaymentScheduleInterval The frequency interval of the payment.
type PaymentScheduleInterval string

// List of PaymentScheduleInterval
const (
	PAYMENTSCHEDULEINTERVAL_WEEKLY PaymentScheduleInterval = "WEEKLY"
	PAYMENTSCHEDULEINTERVAL_MONTHLY PaymentScheduleInterval = "MONTHLY"
)

var allowedPaymentScheduleIntervalEnumValues = []PaymentScheduleInterval{
	"WEEKLY",
	"MONTHLY",
}

func (v *PaymentScheduleInterval) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentScheduleInterval(value)
	for _, existing := range allowedPaymentScheduleIntervalEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentScheduleInterval", value)
}

// NewPaymentScheduleIntervalFromValue returns a pointer to a valid PaymentScheduleInterval
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentScheduleIntervalFromValue(v string) (*PaymentScheduleInterval, error) {
	ev := PaymentScheduleInterval(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentScheduleInterval: valid values are %v", v, allowedPaymentScheduleIntervalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentScheduleInterval) IsValid() bool {
	for _, existing := range allowedPaymentScheduleIntervalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentScheduleInterval value
func (v PaymentScheduleInterval) Ptr() *PaymentScheduleInterval {
	return &v
}

type NullablePaymentScheduleInterval struct {
	value *PaymentScheduleInterval
	isSet bool
}

func (v NullablePaymentScheduleInterval) Get() *PaymentScheduleInterval {
	return v.value
}

func (v *NullablePaymentScheduleInterval) Set(val *PaymentScheduleInterval) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentScheduleInterval) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentScheduleInterval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentScheduleInterval(val *PaymentScheduleInterval) *NullablePaymentScheduleInterval {
	return &NullablePaymentScheduleInterval{value: val, isSet: true}
}

func (v NullablePaymentScheduleInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentScheduleInterval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

