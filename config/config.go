package config

import (
	"runtime"
	"time"
)

var (
	Port        = "0.0.0.0:8080"
	JsonDefault = `{"cidade":"","uf":"","logradouro":"","bairro":""}`

	NumCounters      int64 = 1e7     // Num keys to track frequency of (30M).
	MaxCost          int64 = 1 << 30 // Maximum cost of cache (2GB).
	BufferItems      int64 = 64      // Number of keys per Get buffer.
	NumCPU           int   = runtime.NumCPU()
	TimeOutSearchCep int   = 15     // secouds
	TTlCache         int   = 172800 // secouds => 2 days

	// httpClient is used to make HTTP requests with TLS configuration
	InsecureSkipVerify  = true             // Skip TLS verification for testing purposes
	MaxIdleConns        = 100              // Maximum number of idle connections
	MaxIdleConnsPerHost = 10               // Maximum number of idle connections per host
	IdleConnTimeout     = 90 * time.Second // Idle connection timeout in seconds
	Timeout             = 30 * time.Second // Request timeout in seconds

	// ctx cancel search
	CancelCTXSearch = 30 * time.Second
)
