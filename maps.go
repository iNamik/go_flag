package flag

import (
	"errors"
	goflag "flag"
	"strconv"
	"strings"
)

// String map
type stringMapValue_t map[string]string

// NewStringMapValue returns a new go.Value from the specifed map.
// If the specified map is nil, a new map[string]string is created
func NewStringMapValue(m *map[string]string) goflag.Value {
	if *m == nil {
		*m = make(map[string]string)
	}
	return (*stringMapValue_t)(m)
}

func (m *stringMapValue_t) Set(s string) error {
	var k, v string
	if i := strings.IndexRune(s, '='); i > 0 {
		k = s[0:i]
		v = s[i+1:]
	}
	if k != "" && v != "" {
		(*m)[k] = v
	} else {
		// invalid value "{VALUE}" for flag -{FLAG}: {ERROR_MSG}
		return errors.New("Key and value must both be non-empty")
	}
	return nil
}

func (m *stringMapValue_t) String() string { return "" }

// Int map
type intMapValue_t map[string]int

// NewIntMapValue returns a new go.Value from the specifed map.
// If the specified map is nil, a new map[string]int is created
func NewIntMapValue(m *map[string]int) goflag.Value {
	if *m == nil {
		*m = make(map[string]int)
	}
	return (*intMapValue_t)(m)
}

func (m *intMapValue_t) Set(s string) error {
	var k, v string
	if i := strings.IndexRune(s, '='); i > 0 {
		k = s[0:i]
		v = s[i+1:]
	}
	if k != "" && v != "" {
		iv, err := strconv.ParseInt(v, 0, 64)
		if err != nil {
			return err
		}
		(*m)[k] = int(iv)
	} else {
		// invalid value "{VALUE}" for flag -{FLAG}: {ERROR_MSG}
		return errors.New("Key and value must both be non-empty")
	}
	return nil
}

func (m *intMapValue_t) String() string { return "" }
