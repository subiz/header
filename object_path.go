package header

import (
	"fmt"
	"reflect"
	"strings"
)

func ObjectPath(object interface{}, keyWithDots string) interface{} {
	if object == nil {
		return nil
	}

	keySlice := strings.Split(keyWithDots, ".")
	v := reflect.ValueOf(object)
	// iterate through field names ,ignore the first name as it might be the current instance name
	// you can make it recursive also if want to support types like slice,map etc along with struct
	for _, key := range keySlice {
		for v.Kind() == reflect.Interface {
			v = v.Elem()
		}

		fmt.Println("KKKKK", key, "KIND", v.Kind())

		for v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		if strings.Contains(key, "=") {
			if v.Kind() == reflect.Slice {
				found := false
				/*
				for _, item := range v {
					ks := strings.Split(key, "=")
					subkey := ks[0]
					var js any
					json.Unmarshal(byte(strings.Join(ks[1], "=")), &js)
					subval := ObjectPath(v, subkey)
					if subval == subjson {
						found = true
						break
					}
				}
				*/
				if found {
					continue
				}
				return nil
			}

			continue
		}

		if v.Kind() == reflect.Slice {
		}

		if v.Kind() == reflect.Map {
			found := false
			for _, e := range v.MapKeys() {
				if key == e.String() {
					v = v.MapIndex(e)
					found = true
					break
				}
			}

			if found {
				continue
			}
			return nil
		}

		// we only accept structs
		if v.Kind() == reflect.Struct {
			v = v.FieldByName(key)
			continue
		}
	}
	return v
}
