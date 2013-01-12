go_flag
=======

**Struct-Based Interface To The Standard Go Flag Package**


ABOUT
-----

Package flag implements command-line flag parsing by way of an alternate
interface to the standard Go flag package.

We define a struct to store the flag values, and use struct tags to define the
flag parameters.

Maps (map[string]string and map[string]int) and Arrays ([]string and []int) are
supported via flags that can be used multiple times
(i.e. -map key1=value1 -map key2=value2)

Custom types can be used as long as they implement the Go flag.Value interface.


EXAMPLE
-------

Below is an example that shows all of the currently supported flag types,
including support for Maps (map[string][string] and map[string]int),
Arrays ([]string and []int), as well as custom types that implement the
Go flag.Value interface.

	package main

	import (
		"errors"
		"fmt"
		"github.com/iNamik/go_flag"
		"strings"
		"time"
	)

	type flag_t struct {
		MyBool        bool              `name:"bool"        usage:"a bool, as defined by strconv.ParseBool"`
		MyInt         int               `name:"int"         usage:"an int"`
		MyInt64       int64             `name:"int64"       usage:"an int64"`
		MyUint        uint              `name:"uint"        usage:"an uint"`
		MyUint64      uint64            `name:"uint64"      usage:"an uint64"`
		MyFloat64     float64           `name:"float64"     usage:"a float64"`
		MyString      string            `name:"string"      usage:"a string"`
		MyDuration    time.Duration     `name:"duration"    usage:"a duration, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or '\u03bcs'), 'ms', 's', 'm', 'h'."`
		MyStringMap   map[string]string `name:"stringMap"   usage:"string=string (can be used multiple times)"`
		MyStringArray []string          `name:"stringArray" usage:"a string (can be used multiple times)"`
		MyIntMap      map[string]int    `name:"intMap"      usage:"string=int (can be used multiple times)"`
		MyIntArray    []int             `name:"intArray"    usage:"an int (can be used multiple times)"`
		MyWeekday     weekday           `name:"weekday"     usage:"a weekday : Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday"`
	}

	func main() {
		// Create an instance to store the flags, and set some default values
		flags := &flag_t{MyFloat64: 3.14, MyString: "hello, world", MyDuration: 1000000000 * (60 * 60 * 24), MyWeekday: "Saturday"}
		args, err := flag.Parse(flags)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		} else {
			fmt.Printf("flags = %v\nargs = %v\n", flags, args)
		}
	}

	type weekday string

	var weekdays = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	func (w *weekday) Set(s string) error {
		for _, v := range weekdays {
			if strings.EqualFold(v, s) {
				*w = weekday(v)
				return nil
			}
		}
		// invalid value "{VALUE}" for flag -{FLAG}: {ERROR_MSG}
		return errors.New("Not a weekday")
	}

	func (w *weekday) String() string {
		return string(*w)
	}


EXPORTED FUNCTIONS
------------------

The following functions are exported from the flag package.
NOTE: 'goflag' refers to the original Go flag package

	// Parse processes command-line args (os.Args) according to the fields and tags
	// of the specified flags struct.
	func Parse(flags interface{}) ([]string, error)

	// ParseArgs processes the specified args according to the fields and tags
	// of the specified flags struct.
	func ParseArgs(flags interface{}, args []string) ([]string, error)

	// NewFlagSet creates a new flag.FlagSet from the fields and tags of the
	// specified flags struct.
	func NewFlagSet(flags interface{}, errorHandling goflag.ErrorHandling) (*goflag.FlagSet, error)

	// NewStringArrayValue returns a new go.Value from the specifed array.
	// If the specified array is nil, a new []string is created
	func NewStringArrayValue(a *[]string) goflag.Value

	// NewIntArrayValue returns a new go.Value from the specifed array.
	// If the specified array is nil, a new []int is created
	func NewIntArrayValue(a *[]int) goflag.Value

	// NewStringMapValue returns a new go.Value from the specifed map.
	// If the specified map is nil, a new map[string]string is created
	func NewStringMapValue(m *map[string]string) goflag.Value

	// NewIntMapValue returns a new go.Value from the specifed map.
	// If the specified map is nil, a new map[string]int is created
	func NewIntMapValue(m *map[string]int) goflag.Value {


INSTALL
-------

The package is built using the Go tool.  Assuming you have correctly set the
$GOPATH variable, you can run the following command:

	go get github.com/iNamik/go_flag


DEPENDENCIES
------------

go_flag currently has no dependencies outside of the standard Go packages


AUTHORS
-------

 * David Farrell
