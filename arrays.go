package flag

import (
	goflag "flag"
	"strconv"
)

// String array
type stringArrayValue_t []string

// NewStringArrayValue returns a new go.Value from the specifed array.
// If the specified array is nil, a new []string is created
func NewStringArrayValue(a *[]string) goflag.Value {
	if *a == nil {
		*a = make([]string, 0)
	}
	return (*stringArrayValue_t)(a)
}

func (a *stringArrayValue_t) Set(s string) error {
	*a = append(*a, s)
	return nil
}

func (a *stringArrayValue_t) String() string { return "" }

// Int array
type intArrayValue_t []int

// NewIntArrayValue returns a new go.Value from the specifed array.
// If the specified array is nil, a new []int is created
func NewIntArrayValue(a *[]int) goflag.Value {
	if *a == nil {
		*a = make([]int, 0)
	}
	return (*intArrayValue_t)(a)
}

func (a *intArrayValue_t) Set(s string) error {
	v, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return err
	}
	*a = append(*a, int(v))
	return nil
}

func (a *intArrayValue_t) String() string { return "" }
