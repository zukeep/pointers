package pointers

import (
	"errors"
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

func TestBool(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(*Bool(true)).To(Equal(true))
	g.Expect(*Bool(false)).To(Equal(false))
}

func TestByte(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(*Byte(127)).To(Equal(byte(127)))
	g.Expect(*Byte(0)).To(Equal(byte(0)))
}

func TestComplex64(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(*Complex64(127)).To(Equal(complex64(127)))
	g.Expect(*Complex64(0)).To(Equal(complex64(0)))
}

func TestComplex128(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(*Complex128(127)).To(Equal(complex128(127)))
	g.Expect(*Complex128(0)).To(Equal(complex128(0)))
}

func TestError(t *testing.T) {
	g := NewGomegaWithT(t)
	newErr := errors.New("new error")

	g.Expect(*Error(newErr)).To(Equal(newErr))
}

func TestFloat32(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(*Float32(127)).To(Equal(float32(127)))
	g.Expect(*Float32(0)).To(Equal(float32(0)))
}

func TestFloat64(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(*Float64(127)).To(Equal(float64(127)))
	g.Expect(*Float64(0)).To(Equal(float64(0)))
}

func TestInt(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(*Int(127)).To(Equal(127))
	g.Expect(*Int(0)).To(Equal(0))
}

func TestInt8(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(*Int8(127)).To(Equal(int8(127)))
	g.Expect(*Int8(0)).To(Equal(int8(0)))
}

func TestInt16(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(*Int16(127)).To(Equal(int16(127)))
	g.Expect(*Int16(0)).To(Equal(int16(0)))
}

func TestInt32(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(*Int32(127)).To(Equal(int32(127)))
	g.Expect(*Int32(0)).To(Equal(int32(0)))
}

func TestInt64(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(*Int64(127)).To(Equal(int64(127)))
	g.Expect(*Int64(0)).To(Equal(int64(0)))
}

func TestRune(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(*Rune(127)).To(Equal(rune(127)))
	g.Expect(*Rune(0)).To(Equal(rune(0)))
}

func TestString(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(*String("hello world")).To(Equal("hello world"))
	g.Expect(*String("")).To(Equal(""))
}

func TestUint(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(*Uint(127)).To(Equal(uint(127)))
	g.Expect(*Uint(0)).To(Equal(uint(0)))
}

func TestUint8(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(*Uint8(127)).To(Equal(uint8(127)))
	g.Expect(*Uint8(0)).To(Equal(uint8(0)))
}

func TestUint16(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(*Uint16(127)).To(Equal(uint16(127)))
	g.Expect(*Uint16(0)).To(Equal(uint16(0)))
}

func TestUint32(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(*Uint32(127)).To(Equal(uint32(127)))
	g.Expect(*Uint32(0)).To(Equal(uint32(0)))
}

func TestUint64(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(*Uint64(127)).To(Equal(uint64(127)))
	g.Expect(*Uint64(0)).To(Equal(uint64(0)))
}

func TestUintptr(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(*Uintptr(127)).To(Equal(uintptr(127)))
	g.Expect(*Uintptr(0)).To(Equal(uintptr(0)))
}

func TestDuration(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(*Duration(127)).To(Equal(time.Duration(127)))
	g.Expect(*Duration(0)).To(Equal(time.Duration(0)))
}

func TestTime(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(Time(time.Now())).ToNot(BeNil())
	tt := time.Now()
	g.Expect(*Time(tt)).To(Equal(tt))
}
