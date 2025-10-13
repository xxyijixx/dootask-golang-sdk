package types

import (
	"encoding/json"
	"time"
)

type DateTime struct {
	time.Time
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
