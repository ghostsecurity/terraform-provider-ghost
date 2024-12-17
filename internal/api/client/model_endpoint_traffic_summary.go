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

// checks if the EndpointTrafficSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointTrafficSummary{}

// EndpointTrafficSummary struct for EndpointTrafficSummary
type EndpointTrafficSummary struct {
	AverageRequestsPerHour int32 `json:"average_requests_per_hour"`
	FirstTrafficAt string `json:"first_traffic_at"`
	LastTrafficAt string `json:"last_traffic_at"`
	RequestCount int32 `json:"request_count"`
}

type _EndpointTrafficSummary EndpointTrafficSummary

// NewEndpointTrafficSummary instantiates a new EndpointTrafficSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointTrafficSummary(averageRequestsPerHour int32, firstTrafficAt string, lastTrafficAt string, requestCount int32) *EndpointTrafficSummary {
	this := EndpointTrafficSummary{}
	this.AverageRequestsPerHour = averageRequestsPerHour
	this.FirstTrafficAt = firstTrafficAt
	this.LastTrafficAt = lastTrafficAt
	this.RequestCount = requestCount
	return &this
}

// NewEndpointTrafficSummaryWithDefaults instantiates a new EndpointTrafficSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointTrafficSummaryWithDefaults() *EndpointTrafficSummary {
	this := EndpointTrafficSummary{}
	return &this
}

// GetAverageRequestsPerHour returns the AverageRequestsPerHour field value
func (o *EndpointTrafficSummary) GetAverageRequestsPerHour() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AverageRequestsPerHour
}

// GetAverageRequestsPerHourOk returns a tuple with the AverageRequestsPerHour field value
// and a boolean to check if the value has been set.
func (o *EndpointTrafficSummary) GetAverageRequestsPerHourOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AverageRequestsPerHour, true
}

// SetAverageRequestsPerHour sets field value
func (o *EndpointTrafficSummary) SetAverageRequestsPerHour(v int32) {
	o.AverageRequestsPerHour = v
}

// GetFirstTrafficAt returns the FirstTrafficAt field value
func (o *EndpointTrafficSummary) GetFirstTrafficAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstTrafficAt
}

// GetFirstTrafficAtOk returns a tuple with the FirstTrafficAt field value
// and a boolean to check if the value has been set.
func (o *EndpointTrafficSummary) GetFirstTrafficAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstTrafficAt, true
}

// SetFirstTrafficAt sets field value
func (o *EndpointTrafficSummary) SetFirstTrafficAt(v string) {
	o.FirstTrafficAt = v
}

// GetLastTrafficAt returns the LastTrafficAt field value
func (o *EndpointTrafficSummary) GetLastTrafficAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastTrafficAt
}

// GetLastTrafficAtOk returns a tuple with the LastTrafficAt field value
// and a boolean to check if the value has been set.
func (o *EndpointTrafficSummary) GetLastTrafficAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastTrafficAt, true
}

// SetLastTrafficAt sets field value
func (o *EndpointTrafficSummary) SetLastTrafficAt(v string) {
	o.LastTrafficAt = v
}

// GetRequestCount returns the RequestCount field value
func (o *EndpointTrafficSummary) GetRequestCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RequestCount
}

// GetRequestCountOk returns a tuple with the RequestCount field value
// and a boolean to check if the value has been set.
func (o *EndpointTrafficSummary) GetRequestCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestCount, true
}

// SetRequestCount sets field value
func (o *EndpointTrafficSummary) SetRequestCount(v int32) {
	o.RequestCount = v
}

func (o EndpointTrafficSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointTrafficSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["average_requests_per_hour"] = o.AverageRequestsPerHour
	toSerialize["first_traffic_at"] = o.FirstTrafficAt
	toSerialize["last_traffic_at"] = o.LastTrafficAt
	toSerialize["request_count"] = o.RequestCount
	return toSerialize, nil
}

func (o *EndpointTrafficSummary) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"average_requests_per_hour",
		"first_traffic_at",
		"last_traffic_at",
		"request_count",
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

	varEndpointTrafficSummary := _EndpointTrafficSummary{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEndpointTrafficSummary)

	if err != nil {
		return err
	}

	*o = EndpointTrafficSummary(varEndpointTrafficSummary)

	return err
}

type NullableEndpointTrafficSummary struct {
	value *EndpointTrafficSummary
	isSet bool
}

func (v NullableEndpointTrafficSummary) Get() *EndpointTrafficSummary {
	return v.value
}

func (v *NullableEndpointTrafficSummary) Set(val *EndpointTrafficSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointTrafficSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointTrafficSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointTrafficSummary(val *EndpointTrafficSummary) *NullableEndpointTrafficSummary {
	return &NullableEndpointTrafficSummary{value: val, isSet: true}
}

func (v NullableEndpointTrafficSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointTrafficSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

