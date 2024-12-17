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

// checks if the CertificateMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateMetadata{}

// CertificateMetadata struct for CertificateMetadata
type CertificateMetadata struct {
	InvalidCaReason *string `json:"invalid_ca_reason,omitempty"`
	InvalidHostnameReason *string `json:"invalid_hostname_reason,omitempty"`
	ValidCa bool `json:"valid_ca"`
	ValidHostname bool `json:"valid_hostname"`
}

type _CertificateMetadata CertificateMetadata

// NewCertificateMetadata instantiates a new CertificateMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateMetadata(validCa bool, validHostname bool) *CertificateMetadata {
	this := CertificateMetadata{}
	this.ValidCa = validCa
	this.ValidHostname = validHostname
	return &this
}

// NewCertificateMetadataWithDefaults instantiates a new CertificateMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateMetadataWithDefaults() *CertificateMetadata {
	this := CertificateMetadata{}
	return &this
}

// GetInvalidCaReason returns the InvalidCaReason field value if set, zero value otherwise.
func (o *CertificateMetadata) GetInvalidCaReason() string {
	if o == nil || IsNil(o.InvalidCaReason) {
		var ret string
		return ret
	}
	return *o.InvalidCaReason
}

// GetInvalidCaReasonOk returns a tuple with the InvalidCaReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateMetadata) GetInvalidCaReasonOk() (*string, bool) {
	if o == nil || IsNil(o.InvalidCaReason) {
		return nil, false
	}
	return o.InvalidCaReason, true
}

// HasInvalidCaReason returns a boolean if a field has been set.
func (o *CertificateMetadata) HasInvalidCaReason() bool {
	if o != nil && !IsNil(o.InvalidCaReason) {
		return true
	}

	return false
}

// SetInvalidCaReason gets a reference to the given string and assigns it to the InvalidCaReason field.
func (o *CertificateMetadata) SetInvalidCaReason(v string) {
	o.InvalidCaReason = &v
}

// GetInvalidHostnameReason returns the InvalidHostnameReason field value if set, zero value otherwise.
func (o *CertificateMetadata) GetInvalidHostnameReason() string {
	if o == nil || IsNil(o.InvalidHostnameReason) {
		var ret string
		return ret
	}
	return *o.InvalidHostnameReason
}

// GetInvalidHostnameReasonOk returns a tuple with the InvalidHostnameReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateMetadata) GetInvalidHostnameReasonOk() (*string, bool) {
	if o == nil || IsNil(o.InvalidHostnameReason) {
		return nil, false
	}
	return o.InvalidHostnameReason, true
}

// HasInvalidHostnameReason returns a boolean if a field has been set.
func (o *CertificateMetadata) HasInvalidHostnameReason() bool {
	if o != nil && !IsNil(o.InvalidHostnameReason) {
		return true
	}

	return false
}

// SetInvalidHostnameReason gets a reference to the given string and assigns it to the InvalidHostnameReason field.
func (o *CertificateMetadata) SetInvalidHostnameReason(v string) {
	o.InvalidHostnameReason = &v
}

// GetValidCa returns the ValidCa field value
func (o *CertificateMetadata) GetValidCa() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ValidCa
}

// GetValidCaOk returns a tuple with the ValidCa field value
// and a boolean to check if the value has been set.
func (o *CertificateMetadata) GetValidCaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidCa, true
}

// SetValidCa sets field value
func (o *CertificateMetadata) SetValidCa(v bool) {
	o.ValidCa = v
}

// GetValidHostname returns the ValidHostname field value
func (o *CertificateMetadata) GetValidHostname() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ValidHostname
}

// GetValidHostnameOk returns a tuple with the ValidHostname field value
// and a boolean to check if the value has been set.
func (o *CertificateMetadata) GetValidHostnameOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidHostname, true
}

// SetValidHostname sets field value
func (o *CertificateMetadata) SetValidHostname(v bool) {
	o.ValidHostname = v
}

func (o CertificateMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvalidCaReason) {
		toSerialize["invalid_ca_reason"] = o.InvalidCaReason
	}
	if !IsNil(o.InvalidHostnameReason) {
		toSerialize["invalid_hostname_reason"] = o.InvalidHostnameReason
	}
	toSerialize["valid_ca"] = o.ValidCa
	toSerialize["valid_hostname"] = o.ValidHostname
	return toSerialize, nil
}

func (o *CertificateMetadata) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"valid_ca",
		"valid_hostname",
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

	varCertificateMetadata := _CertificateMetadata{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCertificateMetadata)

	if err != nil {
		return err
	}

	*o = CertificateMetadata(varCertificateMetadata)

	return err
}

type NullableCertificateMetadata struct {
	value *CertificateMetadata
	isSet bool
}

func (v NullableCertificateMetadata) Get() *CertificateMetadata {
	return v.value
}

func (v *NullableCertificateMetadata) Set(val *CertificateMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateMetadata(val *CertificateMetadata) *NullableCertificateMetadata {
	return &NullableCertificateMetadata{value: val, isSet: true}
}

func (v NullableCertificateMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


