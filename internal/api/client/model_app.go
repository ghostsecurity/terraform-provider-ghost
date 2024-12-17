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

// checks if the App type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &App{}

// App struct for App
type App struct {
	ApiCounts *ApiCounts `json:"api_counts,omitempty"`
	AssetCounts AssetCounts `json:"asset_counts"`
	CreatedAt string `json:"created_at"`
	EndpointCounts EndpointCounts `json:"endpoint_counts"`
	Environment string `json:"environment"`
	Host Host `json:"host"`
	Id string `json:"id"`
	IsPublic bool `json:"is_public"`
	LastScannedAt string `json:"last_scanned_at"`
	Name string `json:"name"`
	ScreenshotUrl string `json:"screenshot_url"`
	UpdatedAt string `json:"updated_at"`
	Vulnerabilities *VulnerabilityCount `json:"vulnerabilities,omitempty"`
}

type _App App

// NewApp instantiates a new App object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApp(assetCounts AssetCounts, createdAt string, endpointCounts EndpointCounts, environment string, host Host, id string, isPublic bool, lastScannedAt string, name string, screenshotUrl string, updatedAt string) *App {
	this := App{}
	this.AssetCounts = assetCounts
	this.CreatedAt = createdAt
	this.EndpointCounts = endpointCounts
	this.Environment = environment
	this.Host = host
	this.Id = id
	this.IsPublic = isPublic
	this.LastScannedAt = lastScannedAt
	this.Name = name
	this.ScreenshotUrl = screenshotUrl
	this.UpdatedAt = updatedAt
	return &this
}

// NewAppWithDefaults instantiates a new App object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppWithDefaults() *App {
	this := App{}
	return &this
}

// GetApiCounts returns the ApiCounts field value if set, zero value otherwise.
func (o *App) GetApiCounts() ApiCounts {
	if o == nil || IsNil(o.ApiCounts) {
		var ret ApiCounts
		return ret
	}
	return *o.ApiCounts
}

// GetApiCountsOk returns a tuple with the ApiCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetApiCountsOk() (*ApiCounts, bool) {
	if o == nil || IsNil(o.ApiCounts) {
		return nil, false
	}
	return o.ApiCounts, true
}

// HasApiCounts returns a boolean if a field has been set.
func (o *App) HasApiCounts() bool {
	if o != nil && !IsNil(o.ApiCounts) {
		return true
	}

	return false
}

// SetApiCounts gets a reference to the given ApiCounts and assigns it to the ApiCounts field.
func (o *App) SetApiCounts(v ApiCounts) {
	o.ApiCounts = &v
}

// GetAssetCounts returns the AssetCounts field value
func (o *App) GetAssetCounts() AssetCounts {
	if o == nil {
		var ret AssetCounts
		return ret
	}

	return o.AssetCounts
}

// GetAssetCountsOk returns a tuple with the AssetCounts field value
// and a boolean to check if the value has been set.
func (o *App) GetAssetCountsOk() (*AssetCounts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetCounts, true
}

// SetAssetCounts sets field value
func (o *App) SetAssetCounts(v AssetCounts) {
	o.AssetCounts = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *App) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *App) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *App) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetEndpointCounts returns the EndpointCounts field value
func (o *App) GetEndpointCounts() EndpointCounts {
	if o == nil {
		var ret EndpointCounts
		return ret
	}

	return o.EndpointCounts
}

// GetEndpointCountsOk returns a tuple with the EndpointCounts field value
// and a boolean to check if the value has been set.
func (o *App) GetEndpointCountsOk() (*EndpointCounts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointCounts, true
}

// SetEndpointCounts sets field value
func (o *App) SetEndpointCounts(v EndpointCounts) {
	o.EndpointCounts = v
}

// GetEnvironment returns the Environment field value
func (o *App) GetEnvironment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *App) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *App) SetEnvironment(v string) {
	o.Environment = v
}

// GetHost returns the Host field value
func (o *App) GetHost() Host {
	if o == nil {
		var ret Host
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *App) GetHostOk() (*Host, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *App) SetHost(v Host) {
	o.Host = v
}

// GetId returns the Id field value
func (o *App) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *App) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *App) SetId(v string) {
	o.Id = v
}

// GetIsPublic returns the IsPublic field value
func (o *App) GetIsPublic() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value
// and a boolean to check if the value has been set.
func (o *App) GetIsPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPublic, true
}

// SetIsPublic sets field value
func (o *App) SetIsPublic(v bool) {
	o.IsPublic = v
}

// GetLastScannedAt returns the LastScannedAt field value
func (o *App) GetLastScannedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastScannedAt
}

// GetLastScannedAtOk returns a tuple with the LastScannedAt field value
// and a boolean to check if the value has been set.
func (o *App) GetLastScannedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastScannedAt, true
}

// SetLastScannedAt sets field value
func (o *App) SetLastScannedAt(v string) {
	o.LastScannedAt = v
}

// GetName returns the Name field value
func (o *App) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *App) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *App) SetName(v string) {
	o.Name = v
}

// GetScreenshotUrl returns the ScreenshotUrl field value
func (o *App) GetScreenshotUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScreenshotUrl
}

// GetScreenshotUrlOk returns a tuple with the ScreenshotUrl field value
// and a boolean to check if the value has been set.
func (o *App) GetScreenshotUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScreenshotUrl, true
}

// SetScreenshotUrl sets field value
func (o *App) SetScreenshotUrl(v string) {
	o.ScreenshotUrl = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *App) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *App) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *App) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetVulnerabilities returns the Vulnerabilities field value if set, zero value otherwise.
func (o *App) GetVulnerabilities() VulnerabilityCount {
	if o == nil || IsNil(o.Vulnerabilities) {
		var ret VulnerabilityCount
		return ret
	}
	return *o.Vulnerabilities
}

// GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *App) GetVulnerabilitiesOk() (*VulnerabilityCount, bool) {
	if o == nil || IsNil(o.Vulnerabilities) {
		return nil, false
	}
	return o.Vulnerabilities, true
}

// HasVulnerabilities returns a boolean if a field has been set.
func (o *App) HasVulnerabilities() bool {
	if o != nil && !IsNil(o.Vulnerabilities) {
		return true
	}

	return false
}

// SetVulnerabilities gets a reference to the given VulnerabilityCount and assigns it to the Vulnerabilities field.
func (o *App) SetVulnerabilities(v VulnerabilityCount) {
	o.Vulnerabilities = &v
}

func (o App) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o App) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiCounts) {
		toSerialize["api_counts"] = o.ApiCounts
	}
	toSerialize["asset_counts"] = o.AssetCounts
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["endpoint_counts"] = o.EndpointCounts
	toSerialize["environment"] = o.Environment
	toSerialize["host"] = o.Host
	toSerialize["id"] = o.Id
	toSerialize["is_public"] = o.IsPublic
	toSerialize["last_scanned_at"] = o.LastScannedAt
	toSerialize["name"] = o.Name
	toSerialize["screenshot_url"] = o.ScreenshotUrl
	toSerialize["updated_at"] = o.UpdatedAt
	if !IsNil(o.Vulnerabilities) {
		toSerialize["vulnerabilities"] = o.Vulnerabilities
	}
	return toSerialize, nil
}

func (o *App) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"asset_counts",
		"created_at",
		"endpoint_counts",
		"environment",
		"host",
		"id",
		"is_public",
		"last_scanned_at",
		"name",
		"screenshot_url",
		"updated_at",
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

	varApp := _App{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApp)

	if err != nil {
		return err
	}

	*o = App(varApp)

	return err
}

type NullableApp struct {
	value *App
	isSet bool
}

func (v NullableApp) Get() *App {
	return v.value
}

func (v *NullableApp) Set(val *App) {
	v.value = val
	v.isSet = true
}

func (v NullableApp) IsSet() bool {
	return v.isSet
}

func (v *NullableApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApp(val *App) *NullableApp {
	return &NullableApp{value: val, isSet: true}
}

func (v NullableApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

