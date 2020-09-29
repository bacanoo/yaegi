// Code generated by 'yaegi extract syscall'. DO NOT EDIT.

// +build go1.14,!go1.15

package unrestricted

import (
	"reflect"
	"syscall"
)

func init() {
	Symbols["syscall"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Exit":         reflect.ValueOf(syscall.Exit),
		"Kill":         reflect.ValueOf(syscall.Kill),
		"RawSyscall":   reflect.ValueOf(syscall.RawSyscall),
		"RawSyscall6":  reflect.ValueOf(syscall.RawSyscall6),
		"Shutdown":     reflect.ValueOf(syscall.Shutdown),
		"StartProcess": reflect.ValueOf(syscall.StartProcess),
		"Syscall":      reflect.ValueOf(syscall.Syscall),
		"Syscall6":     reflect.ValueOf(syscall.Syscall6),
	}
}
