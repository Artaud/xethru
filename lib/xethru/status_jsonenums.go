// generated by jsonenums -type=status; DO NOT EDIT

package xethru

import (
	"encoding/json"
	"fmt"
)

var (
	_statusNameToValue = map[string]status{
		"respApp":    respApp,
		"sleepApp":   sleepApp,
		"basebandAP": basebandAP,
		"basebandIQ": basebandIQ,
	}

	_statusValueToName = map[status]string{
		respApp:    "respApp",
		sleepApp:   "sleepApp",
		basebandAP: "basebandAP",
		basebandIQ: "basebandIQ",
	}
)

func init() {
	var v status
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_statusNameToValue = map[string]status{
			interface{}(respApp).(fmt.Stringer).String():    respApp,
			interface{}(sleepApp).(fmt.Stringer).String():   sleepApp,
			interface{}(basebandAP).(fmt.Stringer).String(): basebandAP,
			interface{}(basebandIQ).(fmt.Stringer).String(): basebandIQ,
		}
	}
}

// MarshalJSON is generated so status satisfies json.Marshaler.
func (r status) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _statusValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid status: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so status satisfies json.Unmarshaler.
func (r *status) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("status should be a string, got %s", data)
	}
	v, ok := _statusNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid status %q", s)
	}
	*r = v
	return nil
}