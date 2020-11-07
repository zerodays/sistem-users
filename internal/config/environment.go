package config

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Tag name used for environment variables.
const tagName = "ini"

// getFromEnvironment sets fields in dest struct according to environment variables.
// Variables name are decided as `prefix_<field name>`.
func getFromEnvironment(prefix string, dest interface{}) {
	// Get struct as value and dereference it.
	s := reflect.ValueOf(dest)
	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}

	// Check that dest is struct.
	if s.Kind() != reflect.Struct {
		return
	}

	// Get type of dest.
	typ := reflect.TypeOf(dest)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	for i := 0; i < typ.NumField(); i++ {
		// Get field and check that it's not anonymous.
		f := typ.Field(i)
		if f.Anonymous {
			continue
		}

		// Get value from environment variable.
		tag := f.Tag.Get(tagName)
		var key string
		if len(tag) > 0 {
			key = fmt.Sprintf("%s_%s", prefix, tag)
		} else {
			key = fmt.Sprintf("%s_%s", prefix, f.Name)
		}
		key = strings.ToUpper(key)

		value := os.Getenv(key)
		if len(value) == 0 {
			continue
		}

		// Get field from struct to set the value to.
		field := s.FieldByName(f.Name)
		if !field.IsValid() || !field.CanSet() {
			continue
		}

		// Set the value to the field.
		kind := field.Kind()
		if kind == reflect.Int || kind == reflect.Int64 {
			setStringToInt(field, value, 64)
		} else if kind == reflect.Int32 {
			setStringToInt(field, value, 32)
		} else if kind == reflect.Int16 {
			setStringToInt(field, value, 16)
		} else if kind == reflect.Uint || kind == reflect.Uint64 {
			setStringToUInt(field, value, 64)
		} else if kind == reflect.Uint32 {
			setStringToUInt(field, value, 32)
		} else if kind == reflect.Uint16 {
			setStringToUInt(field, value, 16)
		} else if kind == reflect.Bool {
			setStringToBool(field, value)
		} else if kind == reflect.Float64 {
			setStringToFloat(field, value, 64)
		} else if kind == reflect.Float32 {
			setStringToFloat(field, value, 32)
		} else if kind == reflect.String {
			field.SetString(value)
		}
	}
}

func setStringToInt(f reflect.Value, value string, bitSize int) {
	convertedValue, err := strconv.ParseInt(value, 10, bitSize)

	if err == nil {
		if !f.OverflowInt(convertedValue) {
			f.SetInt(convertedValue)
		}
	}
}

func setStringToUInt(f reflect.Value, value string, bitSize int) {
	convertedValue, err := strconv.ParseUint(value, 10, bitSize)

	if err == nil {
		if !f.OverflowUint(convertedValue) {
			f.SetUint(convertedValue)
		}
	}
}

func setStringToBool(f reflect.Value, value string) {
	convertedValue, err := strconv.ParseBool(value)

	if err == nil {
		f.SetBool(convertedValue)
	}
}

func setStringToFloat(f reflect.Value, value string, bitSize int) {
	convertedValue, err := strconv.ParseFloat(value, bitSize)

	if err == nil {
		if !f.OverflowFloat(convertedValue) {
			f.SetFloat(convertedValue)
		}
	}
}
