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
)

// checks if the BillingInfoDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingInfoDto{}

// BillingInfoDto struct for BillingInfoDto
type BillingInfoDto struct {
	Country NullableString `json:"country,omitempty"`
	VatNumber NullableString `json:"vatNumber,omitempty"`
	LegalName NullableString `json:"legalName,omitempty"`
	City NullableString `json:"city,omitempty"`
	BillingEmail NullableString `json:"billingEmail,omitempty"`
	Address NullableString `json:"address,omitempty"`
}

// NewBillingInfoDto instantiates a new BillingInfoDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInfoDto() *BillingInfoDto {
	this := BillingInfoDto{}
	return &this
}

// NewBillingInfoDtoWithDefaults instantiates a new BillingInfoDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInfoDtoWithDefaults() *BillingInfoDto {
	this := BillingInfoDto{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInfoDto) GetCountry() string {
	if o == nil || IsNil(o.Country.Get()) {
		var ret string
		return ret
	}
	return *o.Country.Get()
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfoDto) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Country.Get(), o.Country.IsSet()
}

// HasCountry returns a boolean if a field has been set.
func (o *BillingInfoDto) HasCountry() bool {
	if o != nil && o.Country.IsSet() {
		return true
	}

	return false
}

// SetCountry gets a reference to the given NullableString and assigns it to the Country field.
func (o *BillingInfoDto) SetCountry(v string) {
	o.Country.Set(&v)
}
// SetCountryNil sets the value for Country to be an explicit nil
func (o *BillingInfoDto) SetCountryNil() {
	o.Country.Set(nil)
}

// UnsetCountry ensures that no value is present for Country, not even an explicit nil
func (o *BillingInfoDto) UnsetCountry() {
	o.Country.Unset()
}

// GetVatNumber returns the VatNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInfoDto) GetVatNumber() string {
	if o == nil || IsNil(o.VatNumber.Get()) {
		var ret string
		return ret
	}
	return *o.VatNumber.Get()
}

// GetVatNumberOk returns a tuple with the VatNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfoDto) GetVatNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VatNumber.Get(), o.VatNumber.IsSet()
}

// HasVatNumber returns a boolean if a field has been set.
func (o *BillingInfoDto) HasVatNumber() bool {
	if o != nil && o.VatNumber.IsSet() {
		return true
	}

	return false
}

// SetVatNumber gets a reference to the given NullableString and assigns it to the VatNumber field.
func (o *BillingInfoDto) SetVatNumber(v string) {
	o.VatNumber.Set(&v)
}
// SetVatNumberNil sets the value for VatNumber to be an explicit nil
func (o *BillingInfoDto) SetVatNumberNil() {
	o.VatNumber.Set(nil)
}

// UnsetVatNumber ensures that no value is present for VatNumber, not even an explicit nil
func (o *BillingInfoDto) UnsetVatNumber() {
	o.VatNumber.Unset()
}

// GetLegalName returns the LegalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInfoDto) GetLegalName() string {
	if o == nil || IsNil(o.LegalName.Get()) {
		var ret string
		return ret
	}
	return *o.LegalName.Get()
}

// GetLegalNameOk returns a tuple with the LegalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfoDto) GetLegalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LegalName.Get(), o.LegalName.IsSet()
}

// HasLegalName returns a boolean if a field has been set.
func (o *BillingInfoDto) HasLegalName() bool {
	if o != nil && o.LegalName.IsSet() {
		return true
	}

	return false
}

// SetLegalName gets a reference to the given NullableString and assigns it to the LegalName field.
func (o *BillingInfoDto) SetLegalName(v string) {
	o.LegalName.Set(&v)
}
// SetLegalNameNil sets the value for LegalName to be an explicit nil
func (o *BillingInfoDto) SetLegalNameNil() {
	o.LegalName.Set(nil)
}

// UnsetLegalName ensures that no value is present for LegalName, not even an explicit nil
func (o *BillingInfoDto) UnsetLegalName() {
	o.LegalName.Unset()
}

// GetCity returns the City field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInfoDto) GetCity() string {
	if o == nil || IsNil(o.City.Get()) {
		var ret string
		return ret
	}
	return *o.City.Get()
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfoDto) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.City.Get(), o.City.IsSet()
}

// HasCity returns a boolean if a field has been set.
func (o *BillingInfoDto) HasCity() bool {
	if o != nil && o.City.IsSet() {
		return true
	}

	return false
}

// SetCity gets a reference to the given NullableString and assigns it to the City field.
func (o *BillingInfoDto) SetCity(v string) {
	o.City.Set(&v)
}
// SetCityNil sets the value for City to be an explicit nil
func (o *BillingInfoDto) SetCityNil() {
	o.City.Set(nil)
}

// UnsetCity ensures that no value is present for City, not even an explicit nil
func (o *BillingInfoDto) UnsetCity() {
	o.City.Unset()
}

// GetBillingEmail returns the BillingEmail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInfoDto) GetBillingEmail() string {
	if o == nil || IsNil(o.BillingEmail.Get()) {
		var ret string
		return ret
	}
	return *o.BillingEmail.Get()
}

// GetBillingEmailOk returns a tuple with the BillingEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfoDto) GetBillingEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillingEmail.Get(), o.BillingEmail.IsSet()
}

// HasBillingEmail returns a boolean if a field has been set.
func (o *BillingInfoDto) HasBillingEmail() bool {
	if o != nil && o.BillingEmail.IsSet() {
		return true
	}

	return false
}

// SetBillingEmail gets a reference to the given NullableString and assigns it to the BillingEmail field.
func (o *BillingInfoDto) SetBillingEmail(v string) {
	o.BillingEmail.Set(&v)
}
// SetBillingEmailNil sets the value for BillingEmail to be an explicit nil
func (o *BillingInfoDto) SetBillingEmailNil() {
	o.BillingEmail.Set(nil)
}

// UnsetBillingEmail ensures that no value is present for BillingEmail, not even an explicit nil
func (o *BillingInfoDto) UnsetBillingEmail() {
	o.BillingEmail.Unset()
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BillingInfoDto) GetAddress() string {
	if o == nil || IsNil(o.Address.Get()) {
		var ret string
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BillingInfoDto) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *BillingInfoDto) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullableString and assigns it to the Address field.
func (o *BillingInfoDto) SetAddress(v string) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *BillingInfoDto) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *BillingInfoDto) UnsetAddress() {
	o.Address.Unset()
}

func (o BillingInfoDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingInfoDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Country.IsSet() {
		toSerialize["country"] = o.Country.Get()
	}
	if o.VatNumber.IsSet() {
		toSerialize["vatNumber"] = o.VatNumber.Get()
	}
	if o.LegalName.IsSet() {
		toSerialize["legalName"] = o.LegalName.Get()
	}
	if o.City.IsSet() {
		toSerialize["city"] = o.City.Get()
	}
	if o.BillingEmail.IsSet() {
		toSerialize["billingEmail"] = o.BillingEmail.Get()
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	return toSerialize, nil
}

type NullableBillingInfoDto struct {
	value *BillingInfoDto
	isSet bool
}

func (v NullableBillingInfoDto) Get() *BillingInfoDto {
	return v.value
}

func (v *NullableBillingInfoDto) Set(val *BillingInfoDto) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingInfoDto) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingInfoDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingInfoDto(val *BillingInfoDto) *NullableBillingInfoDto {
	return &NullableBillingInfoDto{value: val, isSet: true}
}

func (v NullableBillingInfoDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingInfoDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


