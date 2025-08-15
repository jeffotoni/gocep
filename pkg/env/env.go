package env

import (
	"os"
	"strconv"
	"time"
)

func GetInt(name string, defaultValue int) int {
	e := os.Getenv(name)
	if len(e) > 0 {
		if val, ok := strconv.Atoi(e); ok == nil {
			return val
		}
		// println("Error Parse value [", e, "]")
	}
	return defaultValue
}

func GetInt64(name string, defaultValue int64) int64 {
	e := os.Getenv(name)
	if len(e) > 0 {
		if val, ok := strconv.Atoi(e); ok == nil {
			return int64(val)
		}
		// println("Error Parse value [", e, "]")
	}
	return defaultValue
}

// 1000000000 = 1s
// 1000000 = 1ms
// Default em Millesecond
func GetDuration(name string, defaultValue time.Duration) time.Duration {
	e := os.Getenv(name)
	if len(e) > 0 {
		if val, ok := strconv.Atoi(e); ok == nil {
			return time.Duration(int64(val)) * time.Millisecond
		}
		// println("Error Parse value [", e, "]")
	}
	return defaultValue
}

func GetString(name string, defaultValue string) (e string) {
	e = os.Getenv(name)
	if len(e) > 0 {
		return
	}
	e = defaultValue
	return
}

func GetBool(name string, defaultValue bool) bool {
	e := os.Getenv(name)
	if len(e) > 0 {
		if val, ok := strconv.ParseBool(e); ok == nil {
			return val
		}
		// println("found value ", e, " but could not parse to boolean")
	}
	return defaultValue
}
