// Package implementation for privacy transformation and sensitive-value protection.
package metrics

import (
	"fmt"
	"io"
	"net/http"
	"sync/atomic"
)

type Metrics struct{ requests uint64 }

func New() *Metrics     { return &Metrics{} }
func (m *Metrics) Inc() { atomic.AddUint64(&m.requests, 1) }
func (m *Metrics) WriteSnapshot(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "privacy_transform_requests_total %d\n", atomic.LoadUint64(&m.requests)); err != nil {
		return fmt.Errorf("write metrics snapshot: %v", err)
	}
	return nil
}
func (m *Metrics) Handler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; version=0.0.4")
	_ = m.WriteSnapshot(w)
}
