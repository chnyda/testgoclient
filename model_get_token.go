/*
Taikun - WebApi

This Api will be responsible for overall data distribution and authorization.

API version: v1
Contact: noreply@itera.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the GetToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetToken{}

// GetToken struct for GetToken
type GetToken struct {
	Token NullableString `json:"token,omitempty"`
	RefreshToken NullableString `json:"refreshToken,omitempty"`
	RefreshTokenExpireTime *time.Time `json:"refreshTokenExpireTime,omitempty"`
}

// NewGetToken instantiates a new GetToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetToken() *GetToken {
	this := GetToken{}
	return &this
}

// NewGetTokenWithDefaults instantiates a new GetToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTokenWithDefaults() *GetToken {
	this := GetToken{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetToken) GetToken() string {
	if o == nil || IsNil(o.Token.Get()) {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetToken) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *GetToken) HasToken() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *GetToken) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *GetToken) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *GetToken) UnsetToken() {
	o.Token.Unset()
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetToken) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken.Get()) {
		var ret string
		return ret
	}
	return *o.RefreshToken.Get()
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetToken) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshToken.Get(), o.RefreshToken.IsSet()
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *GetToken) HasRefreshToken() bool {
	if o != nil && o.RefreshToken.IsSet() {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given NullableString and assigns it to the RefreshToken field.
func (o *GetToken) SetRefreshToken(v string) {
	o.RefreshToken.Set(&v)
}
// SetRefreshTokenNil sets the value for RefreshToken to be an explicit nil
func (o *GetToken) SetRefreshTokenNil() {
	o.RefreshToken.Set(nil)
}

// UnsetRefreshToken ensures that no value is present for RefreshToken, not even an explicit nil
func (o *GetToken) UnsetRefreshToken() {
	o.RefreshToken.Unset()
}

// GetRefreshTokenExpireTime returns the RefreshTokenExpireTime field value if set, zero value otherwise.
func (o *GetToken) GetRefreshTokenExpireTime() time.Time {
	if o == nil || IsNil(o.RefreshTokenExpireTime) {
		var ret time.Time
		return ret
	}
	return *o.RefreshTokenExpireTime
}

// GetRefreshTokenExpireTimeOk returns a tuple with the RefreshTokenExpireTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetToken) GetRefreshTokenExpireTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RefreshTokenExpireTime) {
		return nil, false
	}
	return o.RefreshTokenExpireTime, true
}

// HasRefreshTokenExpireTime returns a boolean if a field has been set.
func (o *GetToken) HasRefreshTokenExpireTime() bool {
	if o != nil && !IsNil(o.RefreshTokenExpireTime) {
		return true
	}

	return false
}

// SetRefreshTokenExpireTime gets a reference to the given time.Time and assigns it to the RefreshTokenExpireTime field.
func (o *GetToken) SetRefreshTokenExpireTime(v time.Time) {
	o.RefreshTokenExpireTime = &v
}

func (o GetToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	if o.RefreshToken.IsSet() {
		toSerialize["refreshToken"] = o.RefreshToken.Get()
	}
	if !IsNil(o.RefreshTokenExpireTime) {
		toSerialize["refreshTokenExpireTime"] = o.RefreshTokenExpireTime
	}
	return toSerialize, nil
}

type NullableGetToken struct {
	value *GetToken
	isSet bool
}

func (v NullableGetToken) Get() *GetToken {
	return v.value
}

func (v *NullableGetToken) Set(val *GetToken) {
	v.value = val
	v.isSet = true
}

func (v NullableGetToken) IsSet() bool {
	return v.isSet
}

func (v *NullableGetToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetToken(val *GetToken) *NullableGetToken {
	return &NullableGetToken{value: val, isSet: true}
}

func (v NullableGetToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

