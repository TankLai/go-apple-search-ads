package searchads

import (
	"encoding/json"
	"fmt"
)

// Status type to represent enum of Status (ENABLED/PAUSED)
type Status byte

// ENABLED and PAUSED enum values
const (
	ENABLED Status = iota + 1
	PAUSED
	VALID
	INVALID
)

var (
	_StatusNameToValue = map[string]Status{
		"ENABLED": ENABLED,
		"PAUSED":  PAUSED,
		"VALID":   VALID,
		"INVALID": INVALID,
	}

	_StatusValueToName = map[Status]string{
		ENABLED: "ENABLED",
		PAUSED:  "PAUSED",
		VALID:   "VALID",
		INVALID: "INVALID",
	}
)

func init() {
	var v Status
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_StatusNameToValue = map[string]Status{
			interface{}(ENABLED).(fmt.Stringer).String(): ENABLED,
			interface{}(PAUSED).(fmt.Stringer).String():  PAUSED,
			interface{}(VALID).(fmt.Stringer).String():   VALID,
			interface{}(INVALID).(fmt.Stringer).String(): INVALID,
		}
	}
}

// MarshalJSON is generated so Status satisfies json.Marshaler.
func (r Status) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _StatusValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid Status: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so Status satisfies json.Unmarshaler.
func (r *Status) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Status should be a string, got %s", data)
	}
	v, ok := _StatusNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid Status %q", s)
	}
	*r = v
	return nil
}

// ParseStatus to turn a String into the Status
func ParseStatus(name string) (Status, error) {
	v, ok := _StatusNameToValue[name]
	if ok {
		return v, nil
	}
	return Status(0), fmt.Errorf("invalid Status: %s", name)
}

// String to return the String of a Status
func (r Status) String() (string, error) {
	s, ok := _StatusValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid Status: %d", r)
	}
	return s, nil
}
