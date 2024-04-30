/*
 * Svix API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AppPortalAccessIn struct for AppPortalAccessIn
type AppPortalAccessIn struct {
	// How long the token will be valid for, in seconds.  Valid values are between 1 hour and 7 days. The default is 7 days.
	Expiry NullableInt32 `json:"expiry,omitempty"`
	// The set of feature flags the created token will have access to.
	FeatureFlags *[]string `json:"featureFlags,omitempty"`
	// Whether the app portal should be in read-only mode.
	ReadOnly NullableBool `json:"readOnly,omitempty"`
}

// NewAppPortalAccessIn instantiates a new AppPortalAccessIn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPortalAccessIn() *AppPortalAccessIn {
	this := AppPortalAccessIn{}
	var expiry int32 = 604800
	this.Expiry = *NewNullableInt32(&expiry)
	return &this
}

// NewAppPortalAccessInWithDefaults instantiates a new AppPortalAccessIn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPortalAccessInWithDefaults() *AppPortalAccessIn {
	this := AppPortalAccessIn{}
	var expiry int32 = 604800
	this.Expiry = *NewNullableInt32(&expiry)
	return &this
}

// GetExpiry returns the Expiry field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppPortalAccessIn) GetExpiry() int32 {
	if o == nil || o.Expiry.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Expiry.Get()
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppPortalAccessIn) GetExpiryOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Expiry.Get(), o.Expiry.IsSet()
}

// HasExpiry returns a boolean if a field has been set.
func (o *AppPortalAccessIn) HasExpiry() bool {
	if o != nil && o.Expiry.IsSet() {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given NullableInt32 and assigns it to the Expiry field.
func (o *AppPortalAccessIn) SetExpiry(v int32) {
	o.Expiry.Set(&v)
}
// SetExpiryNil sets the value for Expiry to be an explicit nil
func (o *AppPortalAccessIn) SetExpiryNil() {
	o.Expiry.Set(nil)
}

// UnsetExpiry ensures that no value is present for Expiry, not even an explicit nil
func (o *AppPortalAccessIn) UnsetExpiry() {
	o.Expiry.Unset()
}

// GetFeatureFlags returns the FeatureFlags field value if set, zero value otherwise.
func (o *AppPortalAccessIn) GetFeatureFlags() []string {
	if o == nil || o.FeatureFlags == nil {
		var ret []string
		return ret
	}
	return *o.FeatureFlags
}

// GetFeatureFlagsOk returns a tuple with the FeatureFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPortalAccessIn) GetFeatureFlagsOk() (*[]string, bool) {
	if o == nil || o.FeatureFlags == nil {
		return nil, false
	}
	return o.FeatureFlags, true
}

// HasFeatureFlags returns a boolean if a field has been set.
func (o *AppPortalAccessIn) HasFeatureFlags() bool {
	if o != nil && o.FeatureFlags != nil {
		return true
	}

	return false
}

// SetFeatureFlags gets a reference to the given []string and assigns it to the FeatureFlags field.
func (o *AppPortalAccessIn) SetFeatureFlags(v []string) {
	o.FeatureFlags = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppPortalAccessIn) GetReadOnly() bool {
	if o == nil || o.ReadOnly.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly.Get()
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppPortalAccessIn) GetReadOnlyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReadOnly.Get(), o.ReadOnly.IsSet()
}

// HasReadOnly returns a boolean if a field has been set.
func (o *AppPortalAccessIn) HasReadOnly() bool {
	if o != nil && o.ReadOnly.IsSet() {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given NullableBool and assigns it to the ReadOnly field.
func (o *AppPortalAccessIn) SetReadOnly(v bool) {
	o.ReadOnly.Set(&v)
}
// SetReadOnlyNil sets the value for ReadOnly to be an explicit nil
func (o *AppPortalAccessIn) SetReadOnlyNil() {
	o.ReadOnly.Set(nil)
}

// UnsetReadOnly ensures that no value is present for ReadOnly, not even an explicit nil
func (o *AppPortalAccessIn) UnsetReadOnly() {
	o.ReadOnly.Unset()
}

func (o AppPortalAccessIn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Expiry.IsSet() {
		toSerialize["expiry"] = o.Expiry.Get()
	}
	if o.FeatureFlags != nil {
		toSerialize["featureFlags"] = o.FeatureFlags
	}
	if o.ReadOnly.IsSet() {
		toSerialize["readOnly"] = o.ReadOnly.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAppPortalAccessIn struct {
	value *AppPortalAccessIn
	isSet bool
}

func (v NullableAppPortalAccessIn) Get() *AppPortalAccessIn {
	return v.value
}

func (v *NullableAppPortalAccessIn) Set(val *AppPortalAccessIn) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPortalAccessIn) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPortalAccessIn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPortalAccessIn(val *AppPortalAccessIn) *NullableAppPortalAccessIn {
	return &NullableAppPortalAccessIn{value: val, isSet: true}
}

func (v NullableAppPortalAccessIn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPortalAccessIn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


