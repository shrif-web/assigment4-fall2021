package go_server

import (
	"encoding/json"
	"time"
)

type Fraud struct {
	VehicleCode uint32
	Type        FraudType
	DateTime    time.Time
}

type FraudType string

const (
	Time        FraudType = "time"
	Temperature FraudType = "temp"
)

type VehicleLog struct {
	Code uint64 `json:"code"`
	Speed uint32 `json:"speed"`
	DateTime time.Time `json:"date_time"`
	Latitude float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Temperature float32 `json:"temperature"`
	IsOff bool `json:"is_off"`
}

func (s *VehicleLog) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, s)
}

func (s VehicleLog) MarshalBinary() (data []byte, err error) {
	return json.Marshal(s)
}
