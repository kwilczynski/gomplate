package funcs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	m := MathNS()
	assert.Equal(t, int64(12), m.Add(1, 1, 2, 3, 5))
	assert.Equal(t, int64(2), m.Add(1, 1))
	assert.Equal(t, int64(1), m.Add(1))
	assert.Equal(t, int64(0), m.Add(-5, 5))
}

func TestMul(t *testing.T) {
	m := MathNS()
	assert.Equal(t, int64(30), m.Mul(1, 1, 2, 3, 5))
	assert.Equal(t, int64(1), m.Mul(1, 1))
	assert.Equal(t, int64(1), m.Mul(1))
	assert.Equal(t, int64(-25), m.Mul("-5", 5))
	assert.Equal(t, int64(28), m.Mul(14, "2"))
}

func TestSub(t *testing.T) {
	m := MathNS()
	assert.Equal(t, int64(0), m.Sub(1, 1))
	assert.Equal(t, int64(-10), m.Sub(-5, 5))
	assert.Equal(t, int64(-41), m.Sub(true, "42"))
}

func mustDiv(a, b interface{}) int64 {
	m := MathNS()
	r, err := m.Div(a, b)
	if err != nil {
		return -1
	}
	return r
}

func TestDiv(t *testing.T) {
	m := MathNS()
	_, err := m.Div(1, 0)
	assert.Error(t, err)
	assert.Equal(t, int64(1), mustDiv(1, 1))
	assert.Equal(t, int64(-1), mustDiv(-5, 5))
	assert.Equal(t, int64(0), mustDiv(true, "42"))
}

func TestRem(t *testing.T) {
	m := MathNS()
	assert.Equal(t, int64(0), m.Rem(1, 1))
	assert.Equal(t, int64(2), m.Rem(5, 3.0))
}

func TestPow(t *testing.T) {
	m := MathNS()
	assert.Equal(t, int64(4), m.Pow(2, "2"))
}

func mustSeq(n ...interface{}) []int64 {
	m := MathNS()
	s, err := m.Seq(n...)
	if err != nil {
		panic(err)
	}
	return s
}
func TestSeq(t *testing.T) {
	m := MathNS()
	assert.EqualValues(t, []int64{0, 1, 2, 3}, mustSeq(0, 3))
	assert.EqualValues(t, []int64{1, 0}, mustSeq(0))
	assert.EqualValues(t, []int64{0, 2, 4}, mustSeq(0, 4, 2))
	assert.EqualValues(t, []int64{0, 2, 4}, mustSeq(0, 5, 2))
	assert.EqualValues(t, []int64{0}, mustSeq(0, 5, 8))
	_, err := m.Seq()
	assert.Error(t, err)
}

func TestIsIntFloatNum(t *testing.T) {
	tests := []struct {
		in      interface{}
		isInt   bool
		isFloat bool
	}{
		{0, true, false},
		{1, true, false},
		{-1, true, false},
		{uint(42), true, false},
		{uint8(255), true, false},
		{uint16(42), true, false},
		{uint32(42), true, false},
		{uint64(42), true, false},
		{int(42), true, false},
		{int8(127), true, false},
		{int16(42), true, false},
		{int32(42), true, false},
		{int64(42), true, false},
		{float32(18.3), false, true},
		{float64(18.3), false, true},
		{1.5, false, true},
		{-18.6, false, true},
		{"42", true, false},
		{"052", true, false},
		{"0xff", true, false},
		{"-42", true, false},
		{"-0", true, false},
		{"3.14", false, true},
		{"-3.14", false, true},
		{"0.00", false, true},
		{"NaN", false, true},
		{"-Inf", false, true},
		{"+Inf", false, true},
		{"", false, false},
		{"foo", false, false},
		{nil, false, false},
		{true, false, false},
	}
	m := MathNS()
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%T(%#v)", tt.in, tt.in), func(t *testing.T) {
			assert.Equal(t, tt.isInt, m.IsInt(tt.in))
			assert.Equal(t, tt.isFloat, m.IsFloat(tt.in))
			assert.Equal(t, tt.isInt || tt.isFloat, m.IsNum(tt.in))
		})
	}
}

func BenchmarkIsFloat(b *testing.B) {
	data := []interface{}{
		0, 1, -1, uint(42), uint8(255), uint16(42), uint32(42), uint64(42), int(42), int8(127), int16(42), int32(42), int64(42), float32(18.3), float64(18.3), 1.5, -18.6, "42", "052", "0xff", "-42", "-0", "3.14", "-3.14", "0.00", "NaN", "-Inf", "+Inf", "", "foo", nil, true,
	}
	m := MathNS()
	for _, n := range data {
		b.Run(fmt.Sprintf("%T(%v)", n, n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				m.IsFloat(n)
			}
		})
	}
}
