// +build go1.11,!go1.12

package stdlib

// Code generated by 'goexports hash/adler32'. DO NOT EDIT.

import (
	"hash/adler32"
	"reflect"
)

func init() {
	Value["hash/adler32"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Checksum": reflect.ValueOf(adler32.Checksum),
		"New":      reflect.ValueOf(adler32.New),
		"Size":     reflect.ValueOf(adler32.Size),

		// type definitions

		// interface wrapper definitions

	}
}
