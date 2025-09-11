package header

import (
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	t.Parallel()

	// Test cases
	var s string
	var i int
	var i64 int64
	var f64 float64
	var b bool
	type SimpleStruct struct {
		Name string
		Age  int
	}
	var st SimpleStruct

	testCases := []struct {
		name    string
		input   []byte
		output  any
		want    any
		wantErr bool
	}{
		// String tests
		{name: "string from unquoted", input: []byte("hello world"), output: &s, want: "hello world"},
		{name: "string from quoted", input: []byte(`"hello world"`), output: &s, want: "hello world"},
		{name: "string with spaces", input: []byte("  spaced out  "), output: &s, want: "spaced out"},
		{name: "string from quoted with spaces", input: []byte(`"  spaced out  "`), output: &s, want: "  spaced out  "},

		// Int tests
		{name: "int from number", input: []byte("123"), output: &i, want: 123},
		{name: "int from quoted number", input: []byte(`"456"`), output: &i, want: 456},
		{name: "int64 from number", input: []byte("789"), output: &i64, want: int64(789)},
		{name: "int64 from quoted number", input: []byte(`"101112"`), output: &i64, want: int64(101112)},
		{name: "int with spaces", input: []byte("  123  "), output: &i, want: 123},

		// Float tests
		{name: "float64 from number", input: []byte("123.45"), output: &f64, want: 123.45},
		{name: "float64 from quoted number", input: []byte(`"67.89"`), output: &f64, want: 67.89},
		{name: "float64 with spaces", input: []byte("  123.45  "), output: &f64, want: 123.45},

		// Bool tests
		{name: "bool from literal", input: []byte("true"), output: &b, want: true},
		{name: "bool from quoted literal", input: []byte(`"false"`), output: &b, want: false},
		{name: "bool with spaces", input: []byte("  true  "), output: &b, want: true},

		// Struct test
		{name: "struct from json", input: []byte(`{"Name": "Gemini", "Age": 1}`), output: &st, want: SimpleStruct{Name: "Gemini", Age: 1}},

		// Error cases
		{name: "error int from non-number", input: []byte("not-a-number"), output: &i, wantErr: true},
		{name: "error bool from non-bool", input: []byte("not-a-bool"), output: &b, wantErr: true},
		{name: "error struct from invalid json", input: []byte(`{invalid json}`), output: &st, wantErr: true},
		{name: "error nil pointer", input: []byte("any"), output: nil, wantErr: true},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := Unmarshal(tc.input, tc.output)

			if (err != nil) != tc.wantErr {
				t.Fatalf("Unmarshal() error = %v, wantErr %v", err, tc.wantErr)
			}

			if !tc.wantErr {
				// reflect.Indirect to get the value from pointer
				val := reflect.Indirect(reflect.ValueOf(tc.output)).Interface()
				if !reflect.DeepEqual(val, tc.want) {
					t.Errorf("Unmarshal() got = %v, want %v", val, tc.want)
				}
			}
		})
	}
}
