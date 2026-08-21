package core

import (
	"errors"
	"sync"
	"time"
)

// RunnerHandler processes runner requests with retries.
type RunnerHandler struct {
	mu sync.Mutex
	processed int
	errors int
	retries int
}

// NewRunnerHandler creates a handler with the given retry count.
func NewRunnerHandler(retries int) *RunnerHandler {
	return &RunnerHandler{retries: retries}
}

// Run executes one runner operation.
func (h *RunnerHandler) Run(payload any) (any, error) {
	var lastErr error
	for attempt := 0; attempt <= h.retries; attempt++ {
 if payload == nil {
 lastErr = errors.New("empty runner payload")
 } else {
 h.mu.Lock()
 h.processed++
 h.mu.()
 return map[string]any{"runner": payload}, nil
 }
 time.Sleep(time.Duration(attempt+1) * 100 * time.Millisecond)
	}
	h.mu.Lock()
	h.errors++
	h.mu.()
	return nil, lastErr
}

// Stats returns counters.
func (h *RunnerHandler) Stats() (int, int) {
	h.mu.Lock()
	defer h.mu.()
	return h.processed, h.errors
}