/*
Package flag implements command-line flag parsing by way of an alternate
interface to the standard Go flag package.

We define a struct to store the flag values, and use struct tags to define the
flag parameters.

Maps (map[string]string) and Arrays ([]string) are supported via flags that
can be used multiple times (i.e. -map key1=value1 -map key2=value2)

Custom types can be used as long as they implement the Go flag.Value interface.

see examples/flags for an example that shows all of the currently supported
flag types, including support for Maps (map[string][string] and map[string]int),
Arrays ([]string and []int), as well as custom types that implement the
Go flag.Value interface.
*/
package flag

import (
	"errors"
	goflag "flag"
	"fmt"
	"os"
	"reflect"
	"time"
	"unsafe"
)

// Parse processes command-line args (os.Args) according to the fields and tags
// of the specified flags struct.
func Parse(flags interface{}) ([]string, error) {
	return ParseArgs(flags, os.Args[1:])
}

// ParseArgs processes the specified args according to the fields and tags
// of the specified flags struct.
func ParseArgs(flags interface{}, args []string) ([]string, error) {

	flagSet, err := NewFlagSet(flags, goflag.ExitOnError)

	if err != nil {
		return nil, err
	}

	err = flagSet.Parse(args)

	if err != nil {
		return nil, err
	}

	return flagSet.Args(), nil
}

// NewFlagSet creates a new flag.FlagSet from the fields and tags of the
// specified flags struct.
func NewFlagSet(flags interface{}, errorHandling goflag.ErrorHandling) (*goflag.FlagSet, error) {
	ptrValue := reflect.ValueOf(flags)
	if ptrValue.Kind() != reflect.Ptr {
		return nil, errors.New("Parameter 'flags' must be pointer to struct")
	}
	structType := ptrValue.Type().Elem()
	if structType.Kind() != reflect.Struct {
		return nil, errors.New("Parameter 'flags' must be pointer to struct")
	}
	structValue := reflect.Indirect(ptrValue)

	flagSet := goflag.NewFlagSet(os.Args[0], errorHandling)

	for i, n := 0, structType.NumField(); i < n; i++ {
		fieldMeta := structType.Field(i)
		tags := fieldMeta.Tag
		flagName := tags.Get("name")
		flagUsage := tags.Get("usage")

		// Only process fields with flag names
		if flagName != "" {
			field := structValue.Field(i)
			fieldType := field.Type()

			//fmt.Printf("field[%d]: fieldName:'%s' pkg:'%s' kind:'%s' typeName:'%s' typeString:'%s' flagName:'%s' flagUsage:'%s'\n", i, fieldMeta.Name, field.Type().PkgPath(), field.Kind().String(), fieldType.Name(), fieldType.String(), flagName, flagUsage)

			switch fieldType.String() {
			case "bool":
				boolVal := field.Bool()
				boolPtr := (*bool)(unsafe.Pointer(field.Addr().Pointer()))
				flagSet.BoolVar(boolPtr, flagName, boolVal, flagUsage)
			case "int":
				intVal := (int)(field.Int())
				intPtr := (*int)(unsafe.Pointer(field.Addr().Pointer()))
				flagSet.IntVar(intPtr, flagName, intVal, flagUsage)
			case "int64":
				int64Val := field.Int()
				int64Ptr := (*int64)(unsafe.Pointer(field.Addr().Pointer()))
				flagSet.Int64Var(int64Ptr, flagName, int64Val, flagUsage)
			case "uint":
				uintVal := (uint)(field.Uint())
				uintPtr := (*uint)(unsafe.Pointer(field.Addr().Pointer()))
				flagSet.UintVar(uintPtr, flagName, uintVal, flagUsage)
			case "uint64":
				uint64Val := field.Uint()
				uint64Ptr := (*uint64)(unsafe.Pointer(field.Addr().Pointer()))
				flagSet.Uint64Var(uint64Ptr, flagName, uint64Val, flagUsage)
			case "float64":
				float64Val := field.Float()
				float64Ptr := (*float64)(unsafe.Pointer(field.Addr().Pointer()))
				flagSet.Float64Var(float64Ptr, flagName, float64Val, flagUsage)
			case "string":
				strVal := (field.String())
				strPtr := (*string)(unsafe.Pointer(field.Addr().Pointer()))
				flagSet.StringVar(strPtr, flagName, strVal, flagUsage)
			case "time.Duration":
				durationVal := time.Duration(field.Int())
				durationPtr := (*time.Duration)(unsafe.Pointer(field.Addr().Pointer()))
				flagSet.DurationVar(durationPtr, flagName, durationVal, flagUsage)
			case "map[string]string":
				mapPtr := (*map[string]string)(unsafe.Pointer(field.Addr().Pointer()))
				flagSet.Var(NewStringMapValue(mapPtr), flagName, flagUsage)
			case "map[string]int":
				mapPtr := (*map[string]int)(unsafe.Pointer(field.Addr().Pointer()))
				flagSet.Var(NewIntMapValue(mapPtr), flagName, flagUsage)
			case "[]string":
				arrayPtr := (*[]string)(unsafe.Pointer(field.Addr().Pointer()))
				flagSet.Var(NewStringArrayValue(arrayPtr), flagName, flagUsage)
			case "[]int":
				arrayPtr := (*[]int)(unsafe.Pointer(field.Addr().Pointer()))
				flagSet.Var(NewIntArrayValue(arrayPtr), flagName, flagUsage)
			default:
				// If type implements *flag.Value then use setVar()
				if field.Type().Implements(reflect.TypeOf((*goflag.Value)(nil)).Elem()) {
					flagSet.Var(field.Interface().(goflag.Value), flagName, flagUsage)
					// If pointer to type implements *flag.Value then use setVar()
				} else if field.Addr().Type().Implements(reflect.TypeOf((*goflag.Value)(nil)).Elem()) {
					flagSet.Var(field.Addr().Interface().(goflag.Value), flagName, flagUsage)
				} else {
					return nil, errors.New(fmt.Sprintf("Field '%s' has unsupported type '%s' which does not implement flag.Value", fieldMeta.Name, fieldType.String()))
				}
			}
		}
	}

	return flagSet, nil
}
