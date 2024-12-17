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

// checks if the Endpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Endpoint{}

// Endpoint struct for Endpoint
type Endpoint struct {
	CreatedAt string `json:"created_at"`
	Format string `json:"format"`
	Host Host `json:"host"`
	Id string `json:"id"`
	IsFirstParty bool `json:"is_first_party"`
	Kinds []string `json:"kinds"`
	Method string `json:"method"`
	PathTemplate string `json:"path_template"`
	Port int32 `json:"port"`
	TrafficSummary *EndpointTrafficSummary `json:"traffic_summary,omitempty"`
}

type _Endpoint Endpoint

// NewEndpoint instantiates a new Endpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpoint(createdAt string, format string, host Host, id string, isFirstParty bool, kinds []string, method string, pathTemplate string, port int32) *Endpoint {
	this := Endpoint{}
	this.CreatedAt = createdAt
	this.Format = format
	this.Host = host
	this.Id = id
	this.IsFirstParty = isFirstParty
	this.Kinds = kinds
	this.Method = method
	this.PathTemplate = pathTemplate
	this.Port = port
	return &this
}

// NewEndpointWithDefaults instantiates a new Endpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointWithDefaults() *Endpoint {
	this := Endpoint{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *Endpoint) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Endpoint) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetFormat returns the Format field value
func (o *Endpoint) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *Endpoint) SetFormat(v string) {
	o.Format = v
}

// GetHost returns the Host field value
func (o *Endpoint) GetHost() Host {
	if o == nil {
		var ret Host
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetHostOk() (*Host, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *Endpoint) SetHost(v Host) {
	o.Host = v
}

// GetId returns the Id field value
func (o *Endpoint) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Endpoint) SetId(v string) {
	o.Id = v
}

// GetIsFirstParty returns the IsFirstParty field value
func (o *Endpoint) GetIsFirstParty() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFirstParty
}

// GetIsFirstPartyOk returns a tuple with the IsFirstParty field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetIsFirstPartyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFirstParty, true
}

// SetIsFirstParty sets field value
func (o *Endpoint) SetIsFirstParty(v bool) {
	o.IsFirstParty = v
}

// GetKinds returns the Kinds field value
func (o *Endpoint) GetKinds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Kinds
}

// GetKindsOk returns a tuple with the Kinds field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetKindsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Kinds, true
}

// SetKinds sets field value
func (o *Endpoint) SetKinds(v []string) {
	o.Kinds = v
}

// GetMethod returns the Method field value
func (o *Endpoint) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *Endpoint) SetMethod(v string) {
	o.Method = v
}

// GetPathTemplate returns the PathTemplate field value
func (o *Endpoint) GetPathTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PathTemplate
}

// GetPathTemplateOk returns a tuple with the PathTemplate field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetPathTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PathTemplate, true
}

// SetPathTemplate sets field value
func (o *Endpoint) SetPathTemplate(v string) {
	o.PathTemplate = v
}

// GetPort returns the Port field value
func (o *Endpoint) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *Endpoint) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *Endpoint) SetPort(v int32) {
	o.Port = v
}

// GetTrafficSummary returns the TrafficSummary field value if set, zero value otherwise.
func (o *Endpoint) GetTrafficSummary() EndpointTrafficSummary {
	if o == nil || IsNil(o.TrafficSummary) {
		var ret EndpointTrafficSummary
		return ret
	}
	return *o.TrafficSummary
}

// GetTrafficSummaryOk returns a tuple with the TrafficSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Endpoint) GetTrafficSummaryOk() (*EndpointTrafficSummary, bool) {
	if o == nil || IsNil(o.TrafficSummary) {
		return nil, false
	}
	return o.TrafficSummary, true
}

// HasTrafficSummary returns a boolean if a field has been set.
func (o *Endpoint) HasTrafficSummary() bool {
	if o != nil && !IsNil(o.TrafficSummary) {
		return true
	}

	return false
}

// SetTrafficSummary gets a reference to the given EndpointTrafficSummary and assigns it to the TrafficSummary field.
func (o *Endpoint) SetTrafficSummary(v EndpointTrafficSummary) {
	o.TrafficSummary = &v
}

func (o Endpoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Endpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["format"] = o.Format
	toSerialize["host"] = o.Host
	toSerialize["id"] = o.Id
	toSerialize["is_first_party"] = o.IsFirstParty
	toSerialize["kinds"] = o.Kinds
	toSerialize["method"] = o.Method
	toSerialize["path_template"] = o.PathTemplate
	toSerialize["port"] = o.Port
	if !IsNil(o.TrafficSummary) {
		toSerialize["traffic_summary"] = o.TrafficSummary
	}
	return toSerialize, nil
}

func (o *Endpoint) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_at",
		"format",
		"host",
		"id",
		"is_first_party",
		"kinds",
		"method",
		"path_template",
		"port",
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

	varEndpoint := _Endpoint{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndpoint)

	if err != nil {
		return err
	}

	*o = Endpoint(varEndpoint)

	return err
}

type NullableEndpoint struct {
	value *Endpoint
	isSet bool
}

func (v NullableEndpoint) Get() *Endpoint {
	return v.value
}

func (v *NullableEndpoint) Set(val *Endpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpoint(val *Endpoint) *NullableEndpoint {
	return &NullableEndpoint{value: val, isSet: true}
}

func (v NullableEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

