// Package implementation for privacy transformation and sensitive-value protection.
package config

import (
	"os"
	"strings"
)

func ApplyEnvironment(c Config) Config {
	if c.Metadata == nil {
		c.Metadata = map[string]string{}
	}
	for _, item := range []struct {
		key   string
		apply func(string)
	}{{"PRIVACY_TRANSFORM_HTTP_ADDR", func(v string) { c.HTTPAddr = v }}, {"PRIVACY_TRANSFORM_ENVIRONMENT", func(v string) { c.Environment = v }}} {
		if v := strings.TrimSpace(os.Getenv(item.key)); v != "" {
			item.apply(v)
		}
	}
	return c
}
func (c Config) Public() map[string]any {
	metadata := make(map[string]string, len(c.Metadata))
	for k, v := range c.Metadata {
		metadata[k] = v
	}
	return map[string]any{"http_addr": c.HTTPAddr, "environment": c.Environment, "shutdown_seconds": c.ShutdownSeconds, "metadata": metadata}
}
