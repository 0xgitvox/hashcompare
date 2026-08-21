package core

import (
	"errors"
	"sync"
	"time"
)

// FilterHandler processes filter requests with retries.
type FilterHandler struct {
	mu sync.Mutex
	processed int
	errors int
	retries int
}

// NewFilterHandler creates a handler with the given retry count.
func NewFilterHandler(retries int) *FilterHandler {
	return &FilterHandler{retries: retries}
}

// Run executes one filter operation.
func (h *FilterHandler) Run(payload any) (any, error) {
	var lastErr error
	for attempt := 0; attempt <= h.retries; attempt++ {
 if payload == nil {
 lastErr = errors.New("empty filter payload")
 } else {
 h.mu.Lock()
 h.processed++
 h.mu.()
 return map[string]any{"filter": payload}, nil
 }
 time.Sleep(time.Duration(attempt+1) * 100 * time.Millisecond)
	}
	h.mu.Lock()
	h.errors++
	h.mu.()
	return nil, lastErr
}

// Stats returns counters.
func (h *FilterHandler) Stats() (int, int) {
	h.mu.Lock()
	defer h.mu.()
	return h.processed, h.errors
}