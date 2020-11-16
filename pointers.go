package pointers

import (
	"time"
)

// ---- Public Functions ----

// Bool receives an input of bool and returns a pointer to that type.
func Bool(b bool) *bool {
	return &b
}

// Byte receives an input of byte and returns a pointer to that type.
func Byte(b byte) *byte {
	return &b
}

// Complex64 receives an input of complex64 and returns a pointer to that type.
func Complex64(c complex64) *complex64 {
	return &c
}

// Complex128 receives an input of complex128 and returns a pointer to that type.
func Complex128(c complex128) *complex128 {
	return &c
}

// Error receives an input of error and returns a pointer to that type.
func Error(e error) *error {
	return &e
}

// Float32 receives an input of float32 and returns a pointer to that type.
func Float32(f float32) *float32 {
	return &f
}

// Float64 receives an input of float64 and returns a pointer to that type.
func Float64(f float64) *float64 {
	return &f
}

// Int receives an input of int and returns a pointer to that type.
func Int(i int) *int {
	return &i
}

// Int8 receives an input of int8 and returns a pointer to that type.
func Int8(i int8) *int8 {
	return &i
}

// Int16 receives an input of int16 and returns a pointer to that type.
func Int16(i int16) *int16 {
	return &i
}

// Int32 receives an input of int32 and returns a pointer to that type.
func Int32(i int32) *int32 {
	return &i
}

// Int64 receives an input of int64 and returns a pointer to that type.
func Int64(i int64) *int64 {
	return &i
}

// Rune receives an input of rune and returns a pointer to that type.
func Rune(r rune) *rune {
	return &r
}

// String receives an input of string and returns a pointer to that type.
func String(s string) *string {
	return &s
}

// Uint receives an input of uint and returns a pointer to that type.
func Uint(u uint) *uint {
	return &u
}

// Uint8 receives an input of uint8 and returns a pointer to that type.
func Uint8(u uint8) *uint8 {
	return &u
}

// Uint16 receives an input of uint16 and returns a pointer to that type.
func Uint16(u uint16) *uint16 {
	return &u
}

// Uint32 receives an input of uint32 and returns a pointer to that type.
func Uint32(u uint32) *uint32 {
	return &u
}

// Uint64 receives an input of uint64 and returns a pointer to that type.
func Uint64(u uint64) *uint64 {
	return &u
}

// Uintptr receives an input of uintptr and returns a pointer to that type.
func Uintptr(u uintptr) *uintptr {
	return &u
}

// Duration receives an input of time.Duration and returns a pointer to that type.
func Duration(d time.Duration) *time.Duration {
	return &d
}

// Time receives an input of time.Time and returns a pointer to that type.
func Time(t time.Time) *time.Time {
	return &t
}
