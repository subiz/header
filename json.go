package header

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// Unmarshal like json.Umarshal but more fault tolerrant
// useful when convert back from csv
//
// var s string
// header.Unmarshal([]byte("string"), &s) -> s = "string"
// header.Unmarshal([]byte(""string""), &s) -> s = "string"
//
// var i int64
// header.Unmarshal([]byte("4"), &s) -> s = 4
// header.Unmarshal([]byte(""4""), &s) -> s = 4
//
// var b boolean
// header.Unmarshal([]byte("true"), &s) -> s = true
// header.Unmarshal([]byte(""true""), &s) -> s = true
//
func Unmarshal(data []byte, v any) error {
	err := json.Unmarshal(data, v)
	if err == nil {
		return nil
	}

	// json.Unmarshal failed. Let's try to be more tolerant.

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &json.InvalidUnmarshalError{Type: reflect.TypeOf(v)}
	}

	s := string(bytes.TrimSpace(data))

	// If data is a JSON string (e.g., ""hello""), unquote it.
	// After this, s will be `hello`.
	if unquoted, e := strconv.Unquote(s); e == nil {
		s = unquoted
	}

	elem := rv.Elem()
	switch elem.Kind() {
	case reflect.String:
		elem.SetString(s)
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, parseErr := strconv.ParseInt(s, 10, 64)
		if parseErr != nil {
			return err // return original error
		}
		if elem.OverflowInt(i) {
			return err // return original error
		}
		elem.SetInt(i)
		return nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		u, parseErr := strconv.ParseUint(s, 10, 64)
		if parseErr != nil {
			return err // return original error
		}
		if elem.OverflowUint(u) {
			return err // return original error
		}
		elem.SetUint(u)
		return nil
	case reflect.Float32, reflect.Float64:
		f, parseErr := strconv.ParseFloat(s, 64)
		if parseErr != nil {
			return err // return original error
		}
		if elem.OverflowFloat(f) {
			return err // return original error
		}
		elem.SetFloat(f)
		return nil
	case reflect.Bool:
		b, parseErr := strconv.ParseBool(s)
		if parseErr != nil {
			return err // return original error
		}
		elem.SetBool(b)
		return nil
	}

	// For other types we don't handle, just return the original error.
	return err
}


// Log, print values to stdout
// for primitive values like string, int, boolean, float, ... just print normally
// for arrya, map or struct value, json marshal then print the output
//
// examples: Log("a", map[string]string{"b": 5}) -> a {"b": 5}
func Log(args ...any) {
	for i, arg := range args {
		if i > 0 {
			fmt.Print(" ")
		}

		if arg == nil {
			fmt.Print("null")
			continue
		}

		val := reflect.ValueOf(arg)
		kind := val.Kind()
		if kind == reflect.Ptr {
			if val.IsNil() {
				fmt.Print("null")
				continue
			}
			kind = val.Elem().Kind()
		}

		switch kind {
		case reflect.Array, reflect.Map, reflect.Slice, reflect.Struct:
			jsonBytes, err := json.Marshal(arg)
			if err != nil {
				fmt.Print(arg) // fallback to default printing
			} else {
				fmt.Print(string(jsonBytes))
			}
		default:
			fmt.Print(arg)
		}
	}
	fmt.Println()
}
