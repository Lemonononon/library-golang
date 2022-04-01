package numberu

import (
	"fmt"
	"reflect"
	"strconv"
)

func BoolToInt(i bool) int64 {
	if i {
		return 1
	}
	return 0
}

func AddrInt(i int) *int {
	return &i
}

func AddrInt64(i int64) *int64 {
	return &i
}

func AddrFloat64(f float64) *float64 {
	return &f
}

func ElemInt(p *int) int {
	if p != nil {
		return *p
	}
	return 0
}

func ElemInt64(p *int64) int64 {
	if p != nil {
		return *p
	}
	return 0
}

func ElemFloat64(p *float64) float64 {
	if p != nil {
		return *p
	}
	return 0
}

func ToInt64(i interface{}) int64 {
	if i == nil {
		return 0
	}

	var r int64
	var v = reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.String:
		r, _ = strconv.ParseInt(i.(string), 10, 64)

	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		r = v.Int()

	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		r = int64(v.Uint())

	case reflect.Float32, reflect.Float64:
		r = int64(v.Float())
	}
	return r
}

func ToFloat64(i interface{}) float64 {
	if i == nil {
		return 0
	}

	var r float64
	var v = reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.String:
		r, _ = strconv.ParseFloat(i.(string), 64)
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		r = float64(v.Int())

	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		r = float64(v.Uint())

	case reflect.Float32, reflect.Float64:
		r = v.Float()
	}
	return r
}

func FormatInt(i int) string {
	return strconv.FormatInt(int64(i), 10)
}

func FormatInt32(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

func FormatInt64(i int64) string {
	return strconv.FormatInt(i, 10)
}

func FormatFloat32(f float32) string {
	return fmt.Sprintf("%f", f)
}

func FormatFloat64(f float64) string {
	return fmt.Sprintf("%f", f)
}

func FromInt64Ptr(i *int64) int64 {
	if i == nil {
		return 0
	}
	return *i
}

func FromFloat64Ptr(i *float64) float64 {
	if i == nil {
		return 0
	}
	return *i
}

func ParseInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}

func ParseFloat64(s string) float64 {
	i, _ := strconv.ParseFloat(s, 64)
	return i
}
