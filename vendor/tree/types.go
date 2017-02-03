package tree

import (
	"strconv"
	"tree/radix"
)

type Tree struct {
	tree *radix.Tree
}

type Type struct {
	Value interface{}
}

func (t Type) Int() int {
	return t.Value.(int)
}

func (t Type) Int8() int8 {
	return t.Value.(int8)
}

func (t Type) Int16() int16 {
	return t.Value.(int16)
}

func (t Type) Int32() int32 {
	return t.Value.(int32)
}

func (t Type) Int64() int64 {
	return t.Value.(int64)
}

func (t Type) String() string {
	return t.Value.(string)
}

func (t Type) Float32() float32 {
	return t.Value.(float32)
}

func (t Type) Float64() float64 {
	return t.Value.(float64)
}

func (t Type) Bool() bool {
	return t.Value.(bool)
}

func (t Type) ToString() string {
	switch t.Value.(type) {
	case bool:
		if t.Value.(bool) {
			return "true"
		}
		return "false"
	case int:
		return strconv.Itoa(t.Value.(int))
	case int64:
		return strconv.FormatInt(t.Value.(int64), 10)
	case int8:
		return strconv.Itoa(int(t.Value.(int8)))
	case uint8:
		return strconv.Itoa(int(t.Value.(uint8)))
	case uint16:
		return strconv.Itoa(int(t.Value.(uint16)))
	case int16:
		return strconv.Itoa(int(t.Value.(int16)))
	case float32:
		return strconv.FormatFloat(float64(t.Value.(float32)), 'f', 2, 32)
	case float64:
		return strconv.FormatFloat(t.Value.(float64), 'f', 2, 64)
	}
	return ""
}
