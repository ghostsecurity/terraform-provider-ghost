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

// checks if the EndpointApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointApp{}

// EndpointApp struct for EndpointApp
type EndpointApp struct {
	App App `json:"app"`
	LastScannedAt string `json:"last_scanned_at"`
}

type _EndpointApp EndpointApp

// NewEndpointApp instantiates a new EndpointApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointApp(app App, lastScannedAt string) *EndpointApp {
	this := EndpointApp{}
	this.App = app
	this.LastScannedAt = lastScannedAt
	return &this
}

// NewEndpointAppWithDefaults instantiates a new EndpointApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointAppWithDefaults() *EndpointApp {
	this := EndpointApp{}
	return &this
}

// GetApp returns the App field value
func (o *EndpointApp) GetApp() App {
	if o == nil {
		var ret App
		return ret
	}

	return o.App
}

// GetAppOk returns a tuple with the App field value
// and a boolean to check if the value has been set.
func (o *EndpointApp) GetAppOk() (*App, bool) {
	if o == nil {
		return nil, false
	}
	return &o.App, true
}

// SetApp sets field value
func (o *EndpointApp) SetApp(v App) {
	o.App = v
}

// GetLastScannedAt returns the LastScannedAt field value
func (o *EndpointApp) GetLastScannedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastScannedAt
}

// GetLastScannedAtOk returns a tuple with the LastScannedAt field value
// and a boolean to check if the value has been set.
func (o *EndpointApp) GetLastScannedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastScannedAt, true
}

// SetLastScannedAt sets field value
func (o *EndpointApp) SetLastScannedAt(v string) {
	o.LastScannedAt = v
}

func (o EndpointApp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["app"] = o.App
	toSerialize["last_scanned_at"] = o.LastScannedAt
	return toSerialize, nil
}

func (o *EndpointApp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"app",
		"last_scanned_at",
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

	varEndpointApp := _EndpointApp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndpointApp)

	if err != nil {
		return err
	}

	*o = EndpointApp(varEndpointApp)

	return err
}

type NullableEndpointApp struct {
	value *EndpointApp
	isSet bool
}

func (v NullableEndpointApp) Get() *EndpointApp {
	return v.value
}

func (v *NullableEndpointApp) Set(val *EndpointApp) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointApp) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointApp(val *EndpointApp) *NullableEndpointApp {
	return &NullableEndpointApp{value: val, isSet: true}
}

func (v NullableEndpointApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


