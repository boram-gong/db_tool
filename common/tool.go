package common

import (
	"fmt"
	json "github.com/json-iterator/go"
	"strconv"
)

func Interface2Map(data interface{}) map[string]interface{} {
	switch data.(type) {
	case map[string]interface{}:
		return data.(map[string]interface{})
	case string:
		var m map[string]interface{}
		err := json.UnmarshalFromString(data.(string), &m)
		if err != nil {
			return nil
		} else {
			return m
		}
	default:
		return nil
	}
}

func Interface2Slice(data interface{}) []interface{} {
	switch data.(type) {
	case []interface{}:
		return data.([]interface{})
	default:
		return nil
	}
}

func Interface2String(data interface{}) string {
	switch data.(type) {
	case nil:
		return ""
	case string:
		return data.(string)
	case []byte:
		return string(data.([]byte))
	default:
		return fmt.Sprintf("%v", data)
	}
}

func Interface2Int(data interface{}) int {
	switch data.(type) {
	case int64:
		return int(data.(int64))
	case int:
		return data.(int)
	case string:
		i, _ := strconv.Atoi(data.(string))
		return i
	case float64:
		return int(data.(float64))
	case []byte:
		i, _ := strconv.Atoi(string(data.([]byte)))
		return i
	default:
		return 0
	}
}

func Interface2Float(data interface{}) float64 {
	switch data.(type) {
	case float64:
		return data.(float64)
	case int64:
		return float64(data.(int64))
	case int:
		return float64(data.(int))
	case string:
		i, _ := strconv.ParseFloat(data.(string), 64)
		return i
	case []byte:
		i, _ := strconv.ParseFloat(string(data.([]byte)), 64)
		return i
	default:
		return 0.0
	}
}
