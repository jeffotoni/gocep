package config

import (
	"runtime"
)

var (
	Port = "0.0.0.0:8080"
)

var JsonDefault = `{"cidade":"","uf":"","logradouro":"","bairro":""}`

var (
	NumCounters      int64 = 1e7     // Num keys to track frequency of (30M).
	MaxCost          int64 = 1 << 30 // Maximum cost of cache (2GB).
	BufferItems      int64 = 64      // Number of keys per Get buffer.
	NumCPU           int   = runtime.NumCPU()
	TimeOutSearchCep int   = 15     // secouds
	TTlCache         int   = 172800 // secouds => 2 days
)
