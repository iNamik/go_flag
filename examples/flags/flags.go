package main

import (
	"fmt"
	"github.com/inamik/go_flag"
	"time"
)

type flag_t struct {
	myBool     bool          `name:"bool"     usage:"a bool"`
	myInt      int           `name:"int"      usage:"an int"`
	myInt64    int64         `name:"int64"    usage:"an int64"`
	myUint     uint          `name:"uint"     usage:"an uint"`
	myUint64   uint64        `name:"uint64"   usage:"an uint64"`
	myFloat64  float64       `name:"float64"  usage:"a float64"`
	myString   string        `name:"string"   usage:"a string"`
	myDuration time.Duration `name:"duration" usage:"a duration, such as '300ms', '-1.5h' or '2h45m'. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'."`
}

func main() {
	// Create an instance to store the flags, and set some default values
	flags := &flag_t{true, -1, -2, 1, 2, 3.14, "hello, world", 1000000000 * (60 * 60 * 24)}
	args, err := flag.Parse(flags)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	} else {
		fmt.Printf("opts = %v\nargs = %v\n", flags, args)
	}
}
