// Code generated by github.com/infraboard/mcube
// DO NOT EDIT

package pipeline

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseTRIGGER_MODEFromString Parse TRIGGER_MODE from string
func ParseTRIGGER_MODEFromString(str string) (TRIGGER_MODE, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := TRIGGER_MODE_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown TRIGGER_MODE: %s", str)
	}

	return TRIGGER_MODE(v), nil
}

// Equal type compare
func (t TRIGGER_MODE) Equal(target TRIGGER_MODE) bool {
	return t == target
}

// IsIn todo
func (t TRIGGER_MODE) IsIn(targets ...TRIGGER_MODE) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t TRIGGER_MODE) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *TRIGGER_MODE) UnmarshalJSON(b []byte) error {
	ins, err := ParseTRIGGER_MODEFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
