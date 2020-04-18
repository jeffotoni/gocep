package config

var (
	Port = ":8084"
)

var JsonDefault = `{"cidade":"","uf":"","logradouro":"","bairro":""}`

var (
	NumCounters int64 = 1e7     // Num keys to track frequency of (30M).
	MaxCost     int64 = 1 << 30 // Maximum cost of cache (2GB).
	BufferItems int64 = 64      // Number of keys per Get buffer.
)
