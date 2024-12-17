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

// checks if the PaginatedEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedEndpoint{}

// PaginatedEndpoint struct for PaginatedEndpoint
type PaginatedEndpoint struct {
	Items []Endpoint `json:"items"`
	Page int32 `json:"page"`
	Pages int32 `json:"pages"`
	Size int32 `json:"size"`
	Total int32 `json:"total"`
}

type _PaginatedEndpoint PaginatedEndpoint

// NewPaginatedEndpoint instantiates a new PaginatedEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedEndpoint(items []Endpoint, page int32, pages int32, size int32, total int32) *PaginatedEndpoint {
	this := PaginatedEndpoint{}
	this.Items = items
	this.Page = page
	this.Pages = pages
	this.Size = size
	this.Total = total
	return &this
}

// NewPaginatedEndpointWithDefaults instantiates a new PaginatedEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedEndpointWithDefaults() *PaginatedEndpoint {
	this := PaginatedEndpoint{}
	return &this
}

// GetItems returns the Items field value
func (o *PaginatedEndpoint) GetItems() []Endpoint {
	if o == nil {
		var ret []Endpoint
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *PaginatedEndpoint) GetItemsOk() ([]Endpoint, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *PaginatedEndpoint) SetItems(v []Endpoint) {
	o.Items = v
}

// GetPage returns the Page field value
func (o *PaginatedEndpoint) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *PaginatedEndpoint) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *PaginatedEndpoint) SetPage(v int32) {
	o.Page = v
}

// GetPages returns the Pages field value
func (o *PaginatedEndpoint) GetPages() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Pages
}

// GetPagesOk returns a tuple with the Pages field value
// and a boolean to check if the value has been set.
func (o *PaginatedEndpoint) GetPagesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pages, true
}

// SetPages sets field value
func (o *PaginatedEndpoint) SetPages(v int32) {
	o.Pages = v
}

// GetSize returns the Size field value
func (o *PaginatedEndpoint) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *PaginatedEndpoint) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *PaginatedEndpoint) SetSize(v int32) {
	o.Size = v
}

// GetTotal returns the Total field value
func (o *PaginatedEndpoint) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *PaginatedEndpoint) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *PaginatedEndpoint) SetTotal(v int32) {
	o.Total = v
}

func (o PaginatedEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["page"] = o.Page
	toSerialize["pages"] = o.Pages
	toSerialize["size"] = o.Size
	toSerialize["total"] = o.Total
	return toSerialize, nil
}

func (o *PaginatedEndpoint) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
		"page",
		"pages",
		"size",
		"total",
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

	varPaginatedEndpoint := _PaginatedEndpoint{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPaginatedEndpoint)

	if err != nil {
		return err
	}

	*o = PaginatedEndpoint(varPaginatedEndpoint)

	return err
}

type NullablePaginatedEndpoint struct {
	value *PaginatedEndpoint
	isSet bool
}

func (v NullablePaginatedEndpoint) Get() *PaginatedEndpoint {
	return v.value
}

func (v *NullablePaginatedEndpoint) Set(val *PaginatedEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedEndpoint(val *PaginatedEndpoint) *NullablePaginatedEndpoint {
	return &NullablePaginatedEndpoint{value: val, isSet: true}
}

func (v NullablePaginatedEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


