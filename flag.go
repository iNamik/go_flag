/*
Package flag implements command-line flag parsing by way of an alternate
interface to the standard Go flag package.

We define a struct to store the flag values, and use struct tags to define the
flag parameters.

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
*/
package flag

import (
	"errors"
	goflag "flag"
	"fmt"
	"reflect"
	"os"
	"time"
	"unsafe"
)

// Parse processes command-line flags according to the fields and tags of
// the specified struct.
func Parse(flags interface{}) ([]string, error) {

	flagSet, err := NewFlagSet(flags, goflag.ExitOnError)

	if err != nil {
		return nil, err
	}

	err = flagSet.Parse(os.Args[1:])

	if err != nil {
		return nil, err
	}

	return flagSet.Args(), nil
}

// NewFlagSet creates a new flag.FlagSet from the fields and tags of the
// specified struct
func NewFlagSet(flags interface{}, errorHandling goflag.ErrorHandling) (*goflag.FlagSet, error) {
	ptrValue := reflect.ValueOf(flags);
	if ptrValue.Kind() != reflect.Ptr {
		return nil, errors.New("Parameter 'flags' must be pointer to struct")
	}
	structType := ptrValue.Type().Elem()
	if structType.Kind() != reflect.Struct {
		return nil, errors.New("Parameter 'flags' must be pointer to struct")
	}
	structValue := reflect.Indirect(ptrValue)

	flagSet := goflag.NewFlagSet(os.Args[0], errorHandling)

	for i, n := 0, structType.NumField(); i < n ; i++ {
		fieldMeta := structType.Field(i)
		tags      := fieldMeta.Tag
		flagName  := tags.Get("name")
		flagUsage := tags.Get("usage")

		// Only process fields with flag names
		if flagName != "" {
			field     := structValue.Field(i)
			fieldPkg  := field.Type().PkgPath()
			fieldType := field.Type().Name()
			fieldKind := field.Kind()

			var fullType string

			if fieldPkg == "" {
				fullType = fieldType
			} else {
				fullType = fmt.Sprintf("%s.%s", fieldPkg, fieldType)
			}

			//fmt.Printf("field[%d]: fieldName:'%s' pkg:'%s' kind:'%s' type:'%s' fullType='%s' flagName:'%s' flagUsage:'%s'\n", i, fieldMeta.Name, fieldPkg, fieldKind.String(), fieldType, fullType, flagName, flagUsage)

			switch fullType {
				case "bool":
				{
					boolVal := field.Bool()
					boolPtr := (*bool)(unsafe.Pointer(field.Addr().Pointer()))
					flagSet.BoolVar(boolPtr, flagName, boolVal, flagUsage)
				}
				case "int":
				{
					intVal := (int)(field.Int())
					intPtr := (*int)(unsafe.Pointer(field.Addr().Pointer()))
					flagSet.IntVar(intPtr, flagName, intVal, flagUsage)
				}
				case "int64":
				{
					int64Val := field.Int()
					int64Ptr := (*int64)(unsafe.Pointer(field.Addr().Pointer()))
					flagSet.Int64Var(int64Ptr, flagName, int64Val, flagUsage)
				}
				case "uint":
				{
					uintVal := (uint)(field.Uint())
					uintPtr := (*uint)(unsafe.Pointer(field.Addr().Pointer()))
					flagSet.UintVar(uintPtr, flagName, uintVal, flagUsage)
				}
				case "uint64":
				{
					uint64Val := field.Uint()
					uint64Ptr := (*uint64)(unsafe.Pointer(field.Addr().Pointer()))
					flagSet.Uint64Var(uint64Ptr, flagName, uint64Val, flagUsage)
				}
				case "float64":
				{
					float64Val := field.Float()
					float64Ptr := (*float64)(unsafe.Pointer(field.Addr().Pointer()))
					flagSet.Float64Var(float64Ptr, flagName, float64Val, flagUsage)
				}
				case "string":
					strVal := (field.String())
					strPtr := (*string)(unsafe.Pointer(field.Addr().Pointer()))
					flagSet.StringVar(strPtr, flagName, strVal, flagUsage)
				case "time.Duration":
					durationVal := time.Duration(field.Int())
					durationPtr := (*time.Duration)(unsafe.Pointer(field.Addr().Pointer()))
					flagSet.DurationVar(durationPtr, flagName, durationVal, flagUsage)
				default:
					return nil, errors.New(fmt.Sprintf("Field '%s' has unsupported type '%s'", fieldMeta.Name, fieldKind))
			}
		}
	}

	return flagSet, nil
}
