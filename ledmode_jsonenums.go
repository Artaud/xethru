// generated by jsonenums -type=ledMode; DO NOT EDIT

package xethru

import (
	"encoding/json"
	"fmt"
)

var (
	_ledModeNameToValue = map[string]ledMode{
		"LEDOff":        LEDOff,
		"LEDSimple":     LEDSimple,
		"LEDFull":       LEDFull,
		"LEDInhalation": LEDInhalation,
	}

	_ledModeValueToName = map[ledMode]string{
		LEDOff:        "LEDOff",
		LEDSimple:     "LEDSimple",
		LEDFull:       "LEDFull",
		LEDInhalation: "LEDInhalation",
	}
)

func init() {
	var v ledMode
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_ledModeNameToValue = map[string]ledMode{
			interface{}(LEDOff).(fmt.Stringer).String():        LEDOff,
			interface{}(LEDSimple).(fmt.Stringer).String():     LEDSimple,
			interface{}(LEDFull).(fmt.Stringer).String():       LEDFull,
			interface{}(LEDInhalation).(fmt.Stringer).String(): LEDInhalation,
		}
	}
}

// MarshalJSON is generated so ledMode satisfies json.Marshaler.
func (r ledMode) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _ledModeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid ledMode: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so ledMode satisfies json.Unmarshaler.
func (r *ledMode) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ledMode should be a string, got %s", data)
	}
	v, ok := _ledModeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid ledMode %q", s)
	}
	*r = v
	return nil
}
