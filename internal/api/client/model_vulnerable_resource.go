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

// checks if the VulnerableResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VulnerableResource{}

// VulnerableResource struct for VulnerableResource
type VulnerableResource struct {
	Id string `json:"id"`
	Kind string `json:"kind"`
	Name string `json:"name"`
}

type _VulnerableResource VulnerableResource

// NewVulnerableResource instantiates a new VulnerableResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVulnerableResource(id string, kind string, name string) *VulnerableResource {
	this := VulnerableResource{}
	this.Id = id
	this.Kind = kind
	this.Name = name
	return &this
}

// NewVulnerableResourceWithDefaults instantiates a new VulnerableResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVulnerableResourceWithDefaults() *VulnerableResource {
	this := VulnerableResource{}
	return &this
}

// GetId returns the Id field value
func (o *VulnerableResource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VulnerableResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VulnerableResource) SetId(v string) {
	o.Id = v
}

// GetKind returns the Kind field value
func (o *VulnerableResource) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *VulnerableResource) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *VulnerableResource) SetKind(v string) {
	o.Kind = v
}

// GetName returns the Name field value
func (o *VulnerableResource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VulnerableResource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VulnerableResource) SetName(v string) {
	o.Name = v
}

func (o VulnerableResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VulnerableResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["kind"] = o.Kind
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *VulnerableResource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"kind",
		"name",
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

	varVulnerableResource := _VulnerableResource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVulnerableResource)

	if err != nil {
		return err
	}

	*o = VulnerableResource(varVulnerableResource)

	return err
}

type NullableVulnerableResource struct {
	value *VulnerableResource
	isSet bool
}

func (v NullableVulnerableResource) Get() *VulnerableResource {
	return v.value
}

func (v *NullableVulnerableResource) Set(val *VulnerableResource) {
	v.value = val
	v.isSet = true
}

func (v NullableVulnerableResource) IsSet() bool {
	return v.isSet
}

func (v *NullableVulnerableResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVulnerableResource(val *VulnerableResource) *NullableVulnerableResource {
	return &NullableVulnerableResource{value: val, isSet: true}
}

func (v NullableVulnerableResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVulnerableResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

