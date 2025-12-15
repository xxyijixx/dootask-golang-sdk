package types

import (
	"encoding/json"
	"time"
)

type DateTime struct {
	time.Time
}

// BoolFromInt 支持从数字（0/1）和布尔值解析的布尔类型
type BoolFromInt bool

// UnmarshalJSON 自定义JSON解析，支持数字和布尔值
func (b *BoolFromInt) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch val := v.(type) {
	case bool:
		*b = BoolFromInt(val)
	case float64:
		*b = BoolFromInt(val != 0)
	case int:
		*b = BoolFromInt(val != 0)
	case int64:
		*b = BoolFromInt(val != 0)
	default:
		*b = false
	}
	return nil
}

// MarshalJSON 自定义JSON序列化
func (b BoolFromInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(bool(b))
}

// Bool 返回标准布尔值
func (b BoolFromInt) Bool() bool {
	return bool(b)
}

// IntArray 支持从数字和数组解析的整数数组类型
type IntArray []int

// UnmarshalJSON 自定义JSON解析，支持数字和数组
func (ia *IntArray) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch val := v.(type) {
	case []interface{}:
		result := make([]int, 0, len(val))
		for _, item := range val {
			switch num := item.(type) {
			case float64:
				result = append(result, int(num))
			case int:
				result = append(result, num)
			case int64:
				result = append(result, int(num))
			}
		}
		*ia = result
	case []int:
		*ia = val
	case float64:
		*ia = []int{int(val)}
	case int:
		*ia = []int{val}
	case int64:
		*ia = []int{int(val)}
	default:
		*ia = []int{}
	}
	return nil
}

// MarshalJSON 自定义JSON序列化
func (ia IntArray) MarshalJSON() ([]byte, error) {
	return json.Marshal([]int(ia))
}

// UnmarshalJSON 自定义JSON解析方法
func (ct *DateTime) UnmarshalJSON(data []byte) error {
	// 移除JSON字符串的引号
	if len(data) > 2 && data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}

	// 如果解析失败，尝试其他常见格式
	formats := []string{
		time.DateTime,
		time.RFC3339,
		time.RFC3339Nano,
	}

	for _, format := range formats {
		if t, err := time.Parse(format, string(data)); err == nil {
			ct.Time = t
			return nil
		}
	}

	return &time.ParseError{
		Layout:     time.DateTime,
		Value:      string(data),
		LayoutElem: time.DateTime,
		ValueElem:  string(data),
		Message:    "unable to parse time",
	}
}

// MarshalJSON 自定义JSON序列化方法
func (ct DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.Time.Format(time.DateTime))
}
