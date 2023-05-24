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
	"fmt"
)

// ActionType the model 'ActionType'
type ActionType int32

// List of ActionType
const (
	ACTIONTYPE__10 ActionType = 10
	ACTIONTYPE__20 ActionType = 20
	ACTIONTYPE__30 ActionType = 30
	ACTIONTYPE__32 ActionType = 32
	ACTIONTYPE__34 ActionType = 34
	ACTIONTYPE__36 ActionType = 36
	ACTIONTYPE__38 ActionType = 38
	ACTIONTYPE__40 ActionType = 40
	ACTIONTYPE__42 ActionType = 42
	ACTIONTYPE__44 ActionType = 44
	ACTIONTYPE__50 ActionType = 50
	ACTIONTYPE__60 ActionType = 60
	ACTIONTYPE__70 ActionType = 70
	ACTIONTYPE__80 ActionType = 80
	ACTIONTYPE__90 ActionType = 90
	ACTIONTYPE__100 ActionType = 100
	ACTIONTYPE__105 ActionType = 105
	ACTIONTYPE__106 ActionType = 106
	ACTIONTYPE__107 ActionType = 107
	ACTIONTYPE__108 ActionType = 108
	ACTIONTYPE__109 ActionType = 109
	ACTIONTYPE__110 ActionType = 110
	ACTIONTYPE__120 ActionType = 120
	ACTIONTYPE__130 ActionType = 130
	ACTIONTYPE__140 ActionType = 140
	ACTIONTYPE__150 ActionType = 150
	ACTIONTYPE__155 ActionType = 155
	ACTIONTYPE__156 ActionType = 156
	ACTIONTYPE__160 ActionType = 160
	ACTIONTYPE__165 ActionType = 165
	ACTIONTYPE__170 ActionType = 170
	ACTIONTYPE__175 ActionType = 175
	ACTIONTYPE__176 ActionType = 176
	ACTIONTYPE__177 ActionType = 177
	ACTIONTYPE__180 ActionType = 180
	ACTIONTYPE__190 ActionType = 190
	ACTIONTYPE__200 ActionType = 200
	ACTIONTYPE__210 ActionType = 210
	ACTIONTYPE__220 ActionType = 220
	ACTIONTYPE__230 ActionType = 230
	ACTIONTYPE__240 ActionType = 240
	ACTIONTYPE__250 ActionType = 250
	ACTIONTYPE__260 ActionType = 260
	ACTIONTYPE__270 ActionType = 270
	ACTIONTYPE__280 ActionType = 280
	ACTIONTYPE__290 ActionType = 290
	ACTIONTYPE__300 ActionType = 300
	ACTIONTYPE__310 ActionType = 310
	ACTIONTYPE__320 ActionType = 320
	ACTIONTYPE__330 ActionType = 330
	ACTIONTYPE__340 ActionType = 340
	ACTIONTYPE__350 ActionType = 350
	ACTIONTYPE__360 ActionType = 360
	ACTIONTYPE__370 ActionType = 370
	ACTIONTYPE__380 ActionType = 380
	ACTIONTYPE__390 ActionType = 390
	ACTIONTYPE__400 ActionType = 400
	ACTIONTYPE__410 ActionType = 410
)

// All allowed values of ActionType enum
var AllowedActionTypeEnumValues = []ActionType{
	10,
	20,
	30,
	32,
	34,
	36,
	38,
	40,
	42,
	44,
	50,
	60,
	70,
	80,
	90,
	100,
	105,
	106,
	107,
	108,
	109,
	110,
	120,
	130,
	140,
	150,
	155,
	156,
	160,
	165,
	170,
	175,
	176,
	177,
	180,
	190,
	200,
	210,
	220,
	230,
	240,
	250,
	260,
	270,
	280,
	290,
	300,
	310,
	320,
	330,
	340,
	350,
	360,
	370,
	380,
	390,
	400,
	410,
}

func (v *ActionType) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ActionType(value)
	for _, existing := range AllowedActionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ActionType", value)
}

// NewActionTypeFromValue returns a pointer to a valid ActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewActionTypeFromValue(v int32) (*ActionType, error) {
	ev := ActionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ActionType: valid values are %v", v, AllowedActionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ActionType) IsValid() bool {
	for _, existing := range AllowedActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ActionType value
func (v ActionType) Ptr() *ActionType {
	return &v
}

type NullableActionType struct {
	value *ActionType
	isSet bool
}

func (v NullableActionType) Get() *ActionType {
	return v.value
}

func (v *NullableActionType) Set(val *ActionType) {
	v.value = val
	v.isSet = true
}

func (v NullableActionType) IsSet() bool {
	return v.isSet
}

func (v *NullableActionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionType(val *ActionType) *NullableActionType {
	return &NullableActionType{value: val, isSet: true}
}

func (v NullableActionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
