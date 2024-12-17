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

// checks if the Certificate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Certificate{}

// Certificate struct for Certificate
type Certificate struct {
	Issuer DistinguishedName `json:"issuer"`
	Metadata *CertificateMetadata `json:"metadata,omitempty"`
	Subject DistinguishedName `json:"subject"`
	SubjectAlternativeNames []string `json:"subject_alternative_names"`
	Validity CertificateValidity `json:"validity"`
}

type _Certificate Certificate

// NewCertificate instantiates a new Certificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificate(issuer DistinguishedName, subject DistinguishedName, subjectAlternativeNames []string, validity CertificateValidity) *Certificate {
	this := Certificate{}
	this.Issuer = issuer
	this.Subject = subject
	this.SubjectAlternativeNames = subjectAlternativeNames
	this.Validity = validity
	return &this
}

// NewCertificateWithDefaults instantiates a new Certificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateWithDefaults() *Certificate {
	this := Certificate{}
	return &this
}

// GetIssuer returns the Issuer field value
func (o *Certificate) GetIssuer() DistinguishedName {
	if o == nil {
		var ret DistinguishedName
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetIssuerOk() (*DistinguishedName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *Certificate) SetIssuer(v DistinguishedName) {
	o.Issuer = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Certificate) GetMetadata() CertificateMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret CertificateMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certificate) GetMetadataOk() (*CertificateMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Certificate) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given CertificateMetadata and assigns it to the Metadata field.
func (o *Certificate) SetMetadata(v CertificateMetadata) {
	o.Metadata = &v
}

// GetSubject returns the Subject field value
func (o *Certificate) GetSubject() DistinguishedName {
	if o == nil {
		var ret DistinguishedName
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetSubjectOk() (*DistinguishedName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *Certificate) SetSubject(v DistinguishedName) {
	o.Subject = v
}

// GetSubjectAlternativeNames returns the SubjectAlternativeNames field value
func (o *Certificate) GetSubjectAlternativeNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SubjectAlternativeNames
}

// GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetSubjectAlternativeNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubjectAlternativeNames, true
}

// SetSubjectAlternativeNames sets field value
func (o *Certificate) SetSubjectAlternativeNames(v []string) {
	o.SubjectAlternativeNames = v
}

// GetValidity returns the Validity field value
func (o *Certificate) GetValidity() CertificateValidity {
	if o == nil {
		var ret CertificateValidity
		return ret
	}

	return o.Validity
}

// GetValidityOk returns a tuple with the Validity field value
// and a boolean to check if the value has been set.
func (o *Certificate) GetValidityOk() (*CertificateValidity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Validity, true
}

// SetValidity sets field value
func (o *Certificate) SetValidity(v CertificateValidity) {
	o.Validity = v
}

func (o Certificate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Certificate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["issuer"] = o.Issuer
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["subject"] = o.Subject
	toSerialize["subject_alternative_names"] = o.SubjectAlternativeNames
	toSerialize["validity"] = o.Validity
	return toSerialize, nil
}

func (o *Certificate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"issuer",
		"subject",
		"subject_alternative_names",
		"validity",
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

	varCertificate := _Certificate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCertificate)

	if err != nil {
		return err
	}

	*o = Certificate(varCertificate)

	return err
}

type NullableCertificate struct {
	value *Certificate
	isSet bool
}

func (v NullableCertificate) Get() *Certificate {
	return v.value
}

func (v *NullableCertificate) Set(val *Certificate) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificate(val *Certificate) *NullableCertificate {
	return &NullableCertificate{value: val, isSet: true}
}

func (v NullableCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


