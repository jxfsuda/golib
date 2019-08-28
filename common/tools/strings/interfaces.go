package strings

import "time"

//判断对象是否为空(或者原始值)
func IfZero(arg interface{}) bool {
	if arg == nil {
		return true
	}
	switch v := arg.(type) {
	case int, float64, int32, int16, int64, float32:
		if v == 0 {
			return true
		}
	case string:
		if v == "" {
			return true
		}
	case *string, *int, *int64, *int32, *int16, *int8, *float32, *float64:
		if v == nil {
			return true
		}
	case time.Time:
		return v.IsZero()
	default:
		return false
	}
	return false
}
