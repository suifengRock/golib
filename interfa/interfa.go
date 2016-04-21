package interfa

import (
	"encoding/json"
	"errors"
	"reflect"
)

func String(data interface{}) (s string, err error) {
	if s, ok := (data).(string); ok {
		return s, nil
	}
	return "", errors.New("type assertion to string failed")
}

func Bool(data interface{}) (b bool, err error) {
	if b, ok := (data).(bool); ok {
		return b, nil
	}
	return false, errors.New("type assertion to bool failed")
}

func Bytes(data interface{}) ([]byte, error) {
	if b, ok := (data).(string); ok {
		return []byte(b), nil
	}
	return nil, errors.New("type assertion to []byte failed")
}

func Map(data interface{}) (map[string]interface{}, error) {
	if m, ok := (data).(map[string]interface{}); ok {
		return m, nil
	}
	return nil, errors.New("type assertion to map[] failed")
}

func Array(data interface{}) ([]interface{}, error) {
	if a, ok := (data).([]interface{}); ok {
		return a, nil
	}
	return nil, errors.New("type assertion to array[] failed")
}

func Int(data interface{}) (int, error) {
	switch data.(type) {
	case json.Number:
		i, err := data.(json.Number).Int64()
		return int(i), err
	case float32, float64:
		return int(reflect.ValueOf(data).Float()), nil
	case int, int8, int16, int32, int64:
		return int(reflect.ValueOf(data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(data).Uint()), nil
	}
	return 0, errors.New("type assertion to Int failed")
}

func Float64(data interface{}) (i float64, err error) {
	switch data.(type) {
	case json.Number:
		return data.(json.Number).Float64()
	case float32, float64:
		return float64(reflect.ValueOf(data).Float()), nil
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(data).Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(data).Uint()), nil
	}
	return 0, errors.New("type assertion to float64 failed")
}
