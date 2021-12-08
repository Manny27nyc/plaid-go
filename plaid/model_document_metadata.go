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

// DocumentMetadata An object representing metadata from the end user's uploaded document.
type DocumentMetadata struct {
	// The name of the document.
	Name *string `json:"name,omitempty"`
	// The processing status of the document.
	Status *string `json:"status,omitempty"`
	// An identifier of the document that is also present in the paystub response.
	DocId *string `json:"doc_id,omitempty"`
	DocType *DocType `json:"doc_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DocumentMetadata DocumentMetadata

// NewDocumentMetadata instantiates a new DocumentMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentMetadata() *DocumentMetadata {
	this := DocumentMetadata{}
	return &this
}

// NewDocumentMetadataWithDefaults instantiates a new DocumentMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentMetadataWithDefaults() *DocumentMetadata {
	this := DocumentMetadata{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DocumentMetadata) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DocumentMetadata) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DocumentMetadata) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DocumentMetadata) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DocumentMetadata) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DocumentMetadata) SetStatus(v string) {
	o.Status = &v
}

// GetDocId returns the DocId field value if set, zero value otherwise.
func (o *DocumentMetadata) GetDocId() string {
	if o == nil || o.DocId == nil {
		var ret string
		return ret
	}
	return *o.DocId
}

// GetDocIdOk returns a tuple with the DocId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetDocIdOk() (*string, bool) {
	if o == nil || o.DocId == nil {
		return nil, false
	}
	return o.DocId, true
}

// HasDocId returns a boolean if a field has been set.
func (o *DocumentMetadata) HasDocId() bool {
	if o != nil && o.DocId != nil {
		return true
	}

	return false
}

// SetDocId gets a reference to the given string and assigns it to the DocId field.
func (o *DocumentMetadata) SetDocId(v string) {
	o.DocId = &v
}

// GetDocType returns the DocType field value if set, zero value otherwise.
func (o *DocumentMetadata) GetDocType() DocType {
	if o == nil || o.DocType == nil {
		var ret DocType
		return ret
	}
	return *o.DocType
}

// GetDocTypeOk returns a tuple with the DocType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocumentMetadata) GetDocTypeOk() (*DocType, bool) {
	if o == nil || o.DocType == nil {
		return nil, false
	}
	return o.DocType, true
}

// HasDocType returns a boolean if a field has been set.
func (o *DocumentMetadata) HasDocType() bool {
	if o != nil && o.DocType != nil {
		return true
	}

	return false
}

// SetDocType gets a reference to the given DocType and assigns it to the DocType field.
func (o *DocumentMetadata) SetDocType(v DocType) {
	o.DocType = &v
}

func (o DocumentMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.DocId != nil {
		toSerialize["doc_id"] = o.DocId
	}
	if o.DocType != nil {
		toSerialize["doc_type"] = o.DocType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *DocumentMetadata) UnmarshalJSON(bytes []byte) (err error) {
	varDocumentMetadata := _DocumentMetadata{}

	if err = json.Unmarshal(bytes, &varDocumentMetadata); err == nil {
		*o = DocumentMetadata(varDocumentMetadata)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "doc_id")
		delete(additionalProperties, "doc_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDocumentMetadata struct {
	value *DocumentMetadata
	isSet bool
}

func (v NullableDocumentMetadata) Get() *DocumentMetadata {
	return v.value
}

func (v *NullableDocumentMetadata) Set(val *DocumentMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentMetadata(val *DocumentMetadata) *NullableDocumentMetadata {
	return &NullableDocumentMetadata{value: val, isSet: true}
}

func (v NullableDocumentMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


