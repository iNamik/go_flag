go_flag
=======

**struct-based interface to the standard Go flag package**


ABOUT
-----

Package flag implements command-line flag parsing by way of an alternate
interface to the standard Go flag package.

We define a struct to store the flag values, and use struct tags to define the
flag parameters.


EXAMPLE
-------

Below is an example that shows all of the currently supported flag types

	package main

	import ("fmt"; "time"; "github.com/inamik/go_flag")

	type flag_t struct {
		myBool     bool          `name:"bool"     usage: a bool"`
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
		flags := &flag_t{true, -1, -2, 1, 2, 3.14, "hello, world", 1000000000*(60*60*24)}
		args, err := flag.Parse(flags)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
		} else {
			fmt.Printf("opts = %v\nargs = %v\n", flags, args)
		}
	}


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
