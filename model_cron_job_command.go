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

// checks if the CronJobCommand type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CronJobCommand{}

// CronJobCommand struct for CronJobCommand
type CronJobCommand struct {
	CronPeriod string `json:"cronPeriod"`
}

// NewCronJobCommand instantiates a new CronJobCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCronJobCommand(cronPeriod string) *CronJobCommand {
	this := CronJobCommand{}
	this.CronPeriod = cronPeriod
	return &this
}

// NewCronJobCommandWithDefaults instantiates a new CronJobCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCronJobCommandWithDefaults() *CronJobCommand {
	this := CronJobCommand{}
	return &this
}

// GetCronPeriod returns the CronPeriod field value
func (o *CronJobCommand) GetCronPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CronPeriod
}

// GetCronPeriodOk returns a tuple with the CronPeriod field value
// and a boolean to check if the value has been set.
func (o *CronJobCommand) GetCronPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CronPeriod, true
}

// SetCronPeriod sets field value
func (o *CronJobCommand) SetCronPeriod(v string) {
	o.CronPeriod = v
}

func (o CronJobCommand) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CronJobCommand) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cronPeriod"] = o.CronPeriod
	return toSerialize, nil
}

type NullableCronJobCommand struct {
	value *CronJobCommand
	isSet bool
}

func (v NullableCronJobCommand) Get() *CronJobCommand {
	return v.value
}

func (v *NullableCronJobCommand) Set(val *CronJobCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableCronJobCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableCronJobCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCronJobCommand(val *CronJobCommand) *NullableCronJobCommand {
	return &NullableCronJobCommand{value: val, isSet: true}
}

func (v NullableCronJobCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCronJobCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


