package gvk

import (
	"encoding/json"
	"net/url"
	"reflect"
	"strconv"
)

func toString(v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return v.String()

	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)

	case reflect.Int, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)

	case reflect.Bool:
		return strconv.FormatBool(v.Bool())

	case reflect.Struct, reflect.Interface, reflect.Slice, reflect.Array:
		b, _ := json.Marshal(v.Interface())
		return string(b)

	default:
		return ""
	}
}

func scan(i any, v url.Values) url.Values {
	e := reflect.ValueOf(i)

	if e.Kind() == reflect.Pointer {
		e = e.Elem()
	}

	if e.Kind() == reflect.Invalid {
		return v
	}

	for i := 0; i < e.NumField(); i++ {
		fTag := e.Type().Field(i).Tag
		// TODO: peer_id = 0 !!! is valid!
		if name := fTag.Get("query"); name != "" && !e.Field(i).IsZero() {
			v.Set(name, toString(e.Field(i)))
		}
	}

	return v
}

func querify(i any) string {
	return scan(i, url.Values{}).Encode()
}

func urlValues(i any) url.Values {
	if i == nil {
		return nil
	}
	return scan(i, url.Values{})
}

func addValues(vals url.Values, i any) url.Values {
	if i == nil {
		return vals
	}
	if vals == nil {
		vals = make(url.Values)
	}
	return scan(i, vals)
}
