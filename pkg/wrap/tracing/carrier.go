package tracing

import (
	"google.golang.org/grpc/metadata"
)

// MetadataCarrier adapts metadata.MD to satisfy the TextMapCarrier interface.
type MetadataCarrier metadata.MD

// Get returns the value associated with the passed key.
func (m MetadataCarrier) Get(key string) string {
	values := metadata.MD(m).Get(key)
	if len(values) == 0 {
		return ""
	}
	return values[0]
}

// Set stores the key-value pair.
func (m MetadataCarrier) Set(key string, value string) {
	metadata.MD(m).Set(key, value)
}

// Keys lists the keys stored in this carrier.
func (m MetadataCarrier) Keys() []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
