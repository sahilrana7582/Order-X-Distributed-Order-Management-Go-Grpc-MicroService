package utils

import (
	"fmt"
	"os"
	"strconv"
)

func GetEnv[T any](key string, defaultVal T) T {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}

	var parsed any

	switch any(defaultVal).(type) {
	case int:
		parsedInt, e := strconv.Atoi(val)
		if e != nil {
			return defaultVal
		}
		parsed = parsedInt
	case bool:
		parsedBool, e := strconv.ParseBool(val)
		if e != nil {
			return defaultVal
		}
		parsed = parsedBool
	case float64:
		parsedFloat, e := strconv.ParseFloat(val, 64)
		if e != nil {
			return defaultVal
		}
		parsed = parsedFloat
	case string:
		parsed = val
	default:
		fmt.Printf("Unsupported type for GetEnv: %T\n", defaultVal)
		return defaultVal
	}

	return parsed.(T)
}
