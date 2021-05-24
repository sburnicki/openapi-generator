/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"encoding/json"
)

// BigCat struct for BigCat
type BigCat struct {
	Cat
	Kind *string `json:"kind,omitempty"`
}

// NewBigCat instantiates a new BigCat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBigCat(className string) *BigCat {
	this := BigCat{}
	this.ClassName = className
	var color string = "red"
	this.Color = &color
	return &this
}

// NewBigCatWithDefaults instantiates a new BigCat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBigCatWithDefaults() *BigCat {
	this := BigCat{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *BigCat) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BigCat) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *BigCat) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *BigCat) SetKind(v string) {
	o.Kind = &v
}

func (o BigCat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCat, errCat := json.Marshal(o.Cat)
	if errCat != nil {
		return []byte{}, errCat
	}
	errCat = json.Unmarshal([]byte(serializedCat), &toSerialize)
	if errCat != nil {
		return []byte{}, errCat
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	return json.Marshal(toSerialize)
}

type NullableBigCat struct {
	value *BigCat
	isSet bool
}

func (v NullableBigCat) Get() *BigCat {
	return v.value
}

func (v *NullableBigCat) Set(val *BigCat) {
	v.value = val
	v.isSet = true
}

func (v NullableBigCat) IsSet() bool {
	return v.isSet
}

func (v *NullableBigCat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBigCat(val *BigCat) *NullableBigCat {
	return &NullableBigCat{value: val, isSet: true}
}

func (v NullableBigCat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBigCat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


