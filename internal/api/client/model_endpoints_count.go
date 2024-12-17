/*
Ghost API

The Ghost API

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EndpointsCount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointsCount{}

// EndpointsCount struct for EndpointsCount
type EndpointsCount struct {
	Count int32 `json:"count"`
}

type _EndpointsCount EndpointsCount

// NewEndpointsCount instantiates a new EndpointsCount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointsCount(count int32) *EndpointsCount {
	this := EndpointsCount{}
	this.Count = count
	return &this
}

// NewEndpointsCountWithDefaults instantiates a new EndpointsCount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointsCountWithDefaults() *EndpointsCount {
	this := EndpointsCount{}
	return &this
}

// GetCount returns the Count field value
func (o *EndpointsCount) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *EndpointsCount) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *EndpointsCount) SetCount(v int32) {
	o.Count = v
}

func (o EndpointsCount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointsCount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	return toSerialize, nil
}

func (o *EndpointsCount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEndpointsCount := _EndpointsCount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndpointsCount)

	if err != nil {
		return err
	}

	*o = EndpointsCount(varEndpointsCount)

	return err
}

type NullableEndpointsCount struct {
	value *EndpointsCount
	isSet bool
}

func (v NullableEndpointsCount) Get() *EndpointsCount {
	return v.value
}

func (v *NullableEndpointsCount) Set(val *EndpointsCount) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointsCount) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointsCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointsCount(val *EndpointsCount) *NullableEndpointsCount {
	return &NullableEndpointsCount{value: val, isSet: true}
}

func (v NullableEndpointsCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointsCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

