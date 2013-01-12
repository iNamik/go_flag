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
