package api

import (
	"fmt"
	"reflect"

	"github.com/google/go-querystring/query"
)

// addOptions adds the parameters in opts as URL query params
// must be a struct whose fields may contain url tags
// Copied from https://github.com/google/go-github
func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	return fmt.Sprintf("%s?%s", s, qs.Encode()), nil
}
