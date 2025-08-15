package config

import (
	"runtime"
	"time"

	"github.com/jeffotoni/gocep/pkg/env"
)

var (
	Port        = env.GetString("PORT", "0.0.0.0:8080")
	JsonDefault = `{"cidade":"","uf":"","logradouro":"","bairro":""}`

	NumCounters      = env.GetInt64("NUM_COUNTERS", 1e7) // Num keys to track frequency of (30M).
	MaxCost          = env.GetInt64("MAX_COST", 1<<30)   // Maximum cost of cache (2GB).
	BufferItems      = env.GetInt64("BUFFER_ITEMS", 64)  // Number of keys per Get buffer.
	NumCPU           = env.GetInt("NUM_CPU", runtime.NumCPU())
	TimeOutSearchCep = env.GetInt("TIMEOUT_SEARCH_CEP", 15) // secouds
	TTlCache         = env.GetInt("TTL_CACHE", 172800)      // secouds => 2 days

	// httpClient is used to make HTTP requests with TLS configuration
	InsecureSkipVerify  = env.GetBool("INSECURE_SKIP_VERIFY", true)         // Skip TLS verification for testing purposes
	MaxIdleConns        = env.GetInt("HTTP_CLIENT_MAXIDLECONNS", 100)       // Maximum number of idle connections
	MaxIdleConnsPerHost = env.GetInt("HTTP_CLIENT_MAXIDLECONNSPERHOST", 10) // Maximum number of idle connections per host
	IdleConnTimeout     = time.Duration(env.GetInt("IDLE_CONN_TIMEOUT", 90)) * time.Second
	Timeout             = time.Duration(env.GetInt("TIMEOUT", 30)) * time.Second

	// ctx cancel search
	CancelCTXSearch = time.Duration(env.GetInt("CANCEL_CTX_SEARCH", 30)) * time.Second

	CACHE_ENABLE = env.GetBool("CACHE_ENABLE", true)
)
