package funcs

import (
	"fmt"
	gmath "math"
	"strconv"
	"sync"

	"github.com/hairyhenderson/gomplate/conv"

	"github.com/hairyhenderson/gomplate/math"
)

var (
	mathNS     *MathFuncs
	mathNSInit sync.Once
)

// MathNS - the math namespace
func MathNS() *MathFuncs {
	mathNSInit.Do(func() { mathNS = &MathFuncs{} })
	return mathNS
}

// AddMathFuncs -
func AddMathFuncs(f map[string]interface{}) {
	f["math"] = MathNS

	f["add"] = MathNS().Add
	f["sub"] = MathNS().Sub
	f["mul"] = MathNS().Mul
	f["div"] = MathNS().Div
	f["rem"] = MathNS().Rem
	f["pow"] = MathNS().Pow
	f["seq"] = MathNS().Seq
}

// MathFuncs -
type MathFuncs struct{}

// IsInt -
func (f *MathFuncs) IsInt(n interface{}) bool {
	switch i := n.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	case string:
		_, err := strconv.ParseInt(i, 0, 64)
		return err == nil
	}
	return false
}

// IsFloat -
func (f *MathFuncs) IsFloat(n interface{}) bool {
	switch i := n.(type) {
	case float32, float64:
		return true
	case string:
		_, err := strconv.ParseFloat(i, 64)
		if err != nil {
			return false
		}
		if f.IsInt(i) {
			return false
		}
		return err == nil
	}
	return false
}

// IsNum -
func (f *MathFuncs) IsNum(n interface{}) bool {
	return f.IsInt(n) || f.IsFloat(n)
}

// Add -
func (f *MathFuncs) Add(n ...interface{}) int64 {
	return math.AddInt(conv.ToInt64s(n...)...)
}

// Mul -
func (f *MathFuncs) Mul(n ...interface{}) int64 {
	return math.MulInt(conv.ToInt64s(n...)...)
}

// Sub -
func (f *MathFuncs) Sub(a, b interface{}) int64 {
	return conv.ToInt64(a) - conv.ToInt64(b)
}

// Div -
func (f *MathFuncs) Div(a, b interface{}) (int64, error) {
	divisor := conv.ToInt64(a)
	dividend := conv.ToInt64(b)
	if dividend == 0 {
		return 0, fmt.Errorf("Error: division by 0")
	}
	return divisor / dividend, nil
}

// Rem -
func (f *MathFuncs) Rem(a, b interface{}) int64 {
	return conv.ToInt64(a) % conv.ToInt64(b)
}

// Pow -
func (f *MathFuncs) Pow(a, b interface{}) int64 {
	return conv.ToInt64(gmath.Pow(conv.ToFloat64(a), conv.ToFloat64(b)))
}

// Seq - return a sequence from `start` to `end`, in steps of `step`
// start and step are optional, and default to 1.
func (f *MathFuncs) Seq(n ...interface{}) ([]int64, error) {
	start := int64(1)
	end := int64(0)
	step := int64(1)
	if len(n) == 0 {
		return nil, fmt.Errorf("math.Seq must be given at least an 'end' value")
	}
	if len(n) == 1 {
		end = conv.ToInt64(n[0])
	}
	if len(n) == 2 {
		start = conv.ToInt64(n[0])
		end = conv.ToInt64(n[1])
	}
	if len(n) == 3 {
		start = conv.ToInt64(n[0])
		end = conv.ToInt64(n[1])
		step = conv.ToInt64(n[2])
	}
	return math.Seq(conv.ToInt64(start), conv.ToInt64(end), conv.ToInt64(step)), nil
}

// Max -
func (f *MathFuncs) Max(n ...interface{}) (int64, error) {
	return 0, nil
}

// Min -
func (f *MathFuncs) Min(n ...interface{}) (int64, error) {
	return 0, nil
}

// Ceil -
func (f *MathFuncs) Ceil(n ...interface{}) (int64, error) {
	return 0, nil
}

// Floor -
func (f *MathFuncs) Floor(n ...interface{}) (int64, error) {
	return 0, nil
}

// Round -
func (f *MathFuncs) Round(n interface{}) (int64, error) {
	return 0, nil
}
