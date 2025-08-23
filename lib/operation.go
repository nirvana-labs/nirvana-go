package lib

import (
	"context"
	"errors"
	"time"

	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/operations"
)

const (
	// DefaultTimeout is the default maximum time to wait for an operation to complete.
	// Used by NewOperationWaiter() when no custom timeout is specified.
	DefaultTimeout = 10 * time.Minute

	// DefaultPollInterval is the default interval between operation status checks.
	// Used by NewOperationWaiter() when no custom polling strategy is specified.
	DefaultPollInterval = 1 * time.Second
)

// WaitStrategy defines how polling intervals are calculated for operation waiting.
// This interface allows for different backoff strategies like fixed intervals,
// linear backoff, exponential backoff, or custom implementations.
//
// Example usage:
//
//	strategy := NewLinearBackoffStrategy(1*time.Second, 500*time.Millisecond, 5*time.Second)
//	waiter := NewOperationWaiter().WithStrategy(strategy)
type WaitStrategy interface {
	// NextInterval returns the next polling interval given the current attempt number (0-based).
	// The attempt parameter starts at 0 for the first retry after the initial poll.
	NextInterval(attempt int) time.Duration
	// Reset is called when starting a new wait operation to allow strategies
	// to reset any internal state if needed.
	Reset()
}

// FixedIntervalStrategy implements WaitStrategy with constant polling intervals.
// This is the simplest strategy that polls at regular, unchanging intervals.
// It's ideal when you want consistent polling frequency regardless of how long
// the operation has been running.
type FixedIntervalStrategy struct {
	interval time.Duration
}

// NewFixedIntervalStrategy creates a strategy that polls at fixed intervals.
// The interval parameter determines how long to wait between each poll attempt.
//
// Example:
//
//	strategy := NewFixedIntervalStrategy(2 * time.Second)
//	waiter := NewOperationWaiter().WithStrategy(strategy)
func NewFixedIntervalStrategy(interval time.Duration) *FixedIntervalStrategy {
	return &FixedIntervalStrategy{interval: interval}
}

// NextInterval returns the fixed interval regardless of attempt number.
// This implements the WaitStrategy interface.
func (s *FixedIntervalStrategy) NextInterval(attempt int) time.Duration {
	return s.interval
}

// Reset is a no-op for fixed interval strategy as there's no state to reset.
// This implements the WaitStrategy interface.
func (s *FixedIntervalStrategy) Reset() {}

// LinearBackoffStrategy implements WaitStrategy with linearly increasing intervals.
// Each successive poll waits longer by a fixed increment, up to a maximum interval.
// This is useful when you want to gradually reduce polling frequency to avoid
// overwhelming the service while still getting reasonably quick updates initially.
//
// Formula: interval = baseInterval + (attempt * increment), capped at maxInterval
type LinearBackoffStrategy struct {
	baseInterval time.Duration // Initial polling interval
	increment    time.Duration // Amount to increase each attempt
	maxInterval  time.Duration // Maximum interval to prevent infinite growth
}

// NewLinearBackoffStrategy creates a strategy with linear backoff.
// The baseInterval is the starting poll interval, increment is added for each
// subsequent attempt, and maxInterval caps the maximum wait time.
//
// Example - start at 1s, increase by 500ms each attempt, max 5s:
//
//	strategy := NewLinearBackoffStrategy(1*time.Second, 500*time.Millisecond, 5*time.Second)
//	// Attempt 0: 1s, Attempt 1: 1.5s, Attempt 2: 2s, ..., Attempt 8+: 5s
func NewLinearBackoffStrategy(baseInterval, increment, maxInterval time.Duration) *LinearBackoffStrategy {
	return &LinearBackoffStrategy{
		baseInterval: baseInterval,
		increment:    increment,
		maxInterval:  maxInterval,
	}
}

// NextInterval calculates the next polling interval using linear backoff.
// Returns baseInterval + (attempt * increment), capped at maxInterval.
// This implements the WaitStrategy interface.
func (s *LinearBackoffStrategy) NextInterval(attempt int) time.Duration {
	interval := s.baseInterval + time.Duration(attempt)*s.increment
	if interval > s.maxInterval {
		return s.maxInterval
	}
	return interval
}

// Reset is a no-op for linear backoff strategy as there's no state to reset.
// This implements the WaitStrategy interface.
func (s *LinearBackoffStrategy) Reset() {}

// ExponentialBackoffStrategy implements WaitStrategy with exponentially increasing intervals.
// Each successive poll waits exponentially longer, which is ideal for operations that
// might take a long time and where you want to quickly back off to avoid overwhelming
// the service while still eventually checking for completion.
//
// Formula: interval = baseInterval * multiplier * attempt, capped at maxInterval
type ExponentialBackoffStrategy struct {
	baseInterval time.Duration // Base interval for calculations
	multiplier   float64       // Exponential growth factor
	maxInterval  time.Duration // Maximum interval to prevent excessive delays
}

// NewExponentialBackoffStrategy creates a strategy with exponential backoff.
// The baseInterval is used in calculations, multiplier determines growth rate,
// and maxInterval caps the maximum wait time.
//
// Example - start calculations with 100ms base, 2x multiplier, max 30s:
//
//	strategy := NewExponentialBackoffStrategy(100*time.Millisecond, 2.0, 30*time.Second)
//	// Attempt 0: 0ms, Attempt 1: 200ms, Attempt 2: 400ms, Attempt 3: 600ms, etc.
func NewExponentialBackoffStrategy(baseInterval time.Duration, multiplier float64, maxInterval time.Duration) *ExponentialBackoffStrategy {
	return &ExponentialBackoffStrategy{
		baseInterval: baseInterval,
		multiplier:   multiplier,
		maxInterval:  maxInterval,
	}
}

// NextInterval calculates the next polling interval using exponential backoff.
// Returns baseInterval * multiplier * attempt, capped at maxInterval.
// This implements the WaitStrategy interface.
func (s *ExponentialBackoffStrategy) NextInterval(attempt int) time.Duration {
	interval := time.Duration(float64(s.baseInterval) * s.multiplier * float64(attempt))
	if interval > s.maxInterval {
		return s.maxInterval
	}
	return interval
}

// Reset is a no-op for exponential backoff strategy as there's no state to reset.
// This implements the WaitStrategy interface.
func (s *ExponentialBackoffStrategy) Reset() {}

// OperationWaiter provides configurable waiting for operations with support for
// different polling strategies. It allows you to customize timeout duration and
// polling behavior (fixed intervals, linear backoff, exponential backoff, etc.).
//
// Example usage:
//
//	// Default: 10-minute timeout with 1-second fixed intervals
//	waiter := NewOperationWaiter()
//
//	// Custom timeout with linear backoff
//	waiter := NewOperationWaiter().
//		WithTimeout(5*time.Minute).
//		WithLinearBackoff(1*time.Second, 500*time.Millisecond, 10*time.Second)
//
//	err := waiter.Wait(ctx, client, "operation-id")
type OperationWaiter struct {
	timeout  time.Duration // Maximum time to wait for operation completion
	strategy WaitStrategy  // Polling strategy (fixed, linear, exponential, custom)
}

// NewOperationWaiter creates a new OperationWaiter with default settings.
// Default configuration:
//   - Timeout: 10 minutes (DefaultTimeout)
//   - Strategy: Fixed 1-second intervals (DefaultPollInterval)
//
// The returned waiter can be customized using the With* methods before calling Wait.
//
// Example with exponential backoff:
//
//	waiter := NewOperationWaiter().
//		WithTimeout(5*time.Minute).
//		WithExponentialBackoff(100*time.Millisecond, 2.0, 30*time.Second)
//	err := waiter.Wait(ctx, client, operationID)
func NewOperationWaiter() *OperationWaiter {
	return &OperationWaiter{
		timeout:  DefaultTimeout,
		strategy: NewFixedIntervalStrategy(DefaultPollInterval),
	}
}

// WithTimeout sets the maximum duration to wait for operation completion.
// If the operation doesn't complete within this timeout, Wait() will return a timeout error.
//
// Example:
//
//	waiter := NewOperationWaiter().WithTimeout(5 * time.Minute)
func (w *OperationWaiter) WithTimeout(timeout time.Duration) *OperationWaiter {
	w.timeout = timeout
	return w
}

// WithPollInterval sets the polling interval using a fixed interval strategy.
// This is a convenience method that replaces the current strategy with a fixed interval.
// For more advanced polling strategies, use WithStrategy, WithLinearBackoff, or WithExponentialBackoff.
//
// Example:
//
//	waiter := NewOperationWaiter().WithPollInterval(2 * time.Second)
func (w *OperationWaiter) WithPollInterval(interval time.Duration) *OperationWaiter {
	w.strategy = NewFixedIntervalStrategy(interval)
	return w
}

// WithStrategy sets a custom wait strategy that implements the WaitStrategy interface.
// This allows you to provide your own polling logic or use pre-built strategies.
//
// Example:
//
//	customStrategy := NewLinearBackoffStrategy(500*time.Millisecond, 100*time.Millisecond, 5*time.Second)
//	waiter := NewOperationWaiter().WithStrategy(customStrategy)
func (w *OperationWaiter) WithStrategy(strategy WaitStrategy) *OperationWaiter {
	w.strategy = strategy
	return w
}

// WithLinearBackoff sets a linear backoff strategy where polling intervals increase
// linearly with each attempt. This is useful when you want to gradually reduce polling
// frequency while still getting reasonably quick updates.
//
// Parameters:
//   - baseInterval: Starting polling interval
//   - increment: Amount to increase each attempt
//   - maxInterval: Maximum interval to prevent infinite growth
//
// Example:
//
//	waiter := NewOperationWaiter().WithLinearBackoff(1*time.Second, 500*time.Millisecond, 10*time.Second)
//	// Polls at: 1s, 1.5s, 2s, 2.5s, ..., up to 10s
func (w *OperationWaiter) WithLinearBackoff(baseInterval, increment, maxInterval time.Duration) *OperationWaiter {
	w.strategy = NewLinearBackoffStrategy(baseInterval, increment, maxInterval)
	return w
}

// WithExponentialBackoff sets an exponential backoff strategy where polling intervals
// grow exponentially with each attempt. This is ideal for long-running operations where
// you want to quickly back off to avoid overwhelming the service.
//
// Parameters:
//   - baseInterval: Base interval used in calculations
//   - multiplier: Growth factor for each attempt
//   - maxInterval: Maximum interval to prevent excessive delays
//
// Example:
//
//	waiter := NewOperationWaiter().WithExponentialBackoff(100*time.Millisecond, 2.0, 30*time.Second)
//	// Polls at: 0ms, 200ms, 400ms, 600ms, ..., up to 30s
func (w *OperationWaiter) WithExponentialBackoff(baseInterval time.Duration, multiplier float64, maxInterval time.Duration) *OperationWaiter {
	w.strategy = NewExponentialBackoffStrategy(baseInterval, multiplier, maxInterval)
	return w
}

// Wait polls the operation status until it completes, fails, times out, or the context is cancelled.
// It uses the configured polling strategy to determine wait intervals between status checks.
//
// The method will:
//   - Poll the operation status at intervals determined by the WaitStrategy
//   - Return nil when the operation status becomes 'done'
//   - Return an error if the operation fails, times out, or context is cancelled
//   - Return an error if any API call fails
//
// Parameters:
//   - ctx: Context for cancellation and deadlines
//   - client: Nirvana client for API calls
//   - operationID: ID of the operation to wait for
//
// Returns an error if the operation fails, times out, is cancelled, or if API calls fail.
func (w *OperationWaiter) Wait(ctx context.Context, client *nirvana.Client, operationID string) error {
	w.strategy.Reset()
	timeoutChan := time.After(w.timeout)

	attempt := 0
	var timer *time.Timer

	for {
		select {
		case <-timeoutChan:
			if timer != nil {
				timer.Stop()
			}
			return errors.New("operation timed out after " + w.timeout.String() + " secs")
		case <-ctx.Done():
			if timer != nil {
				timer.Stop()
			}
			return errors.New("context cancelled: " + ctx.Err().Error())
		default:
			// Check operation status
			op, err := client.Operations.Get(ctx, operationID)
			if err != nil {
				if timer != nil {
					timer.Stop()
				}
				return errors.New("failed to get operation status: " + err.Error())
			}

			if op.Status == operations.OperationStatusDone {
				if timer != nil {
					timer.Stop()
				}
				return nil
			}

			if op.Status == operations.OperationStatusFailed {
				if timer != nil {
					timer.Stop()
				}
				return errors.New("operation failed")
			}

			// Wait for next poll using strategy
			interval := w.strategy.NextInterval(attempt)
			if timer != nil {
				timer.Stop()
			}
			timer = time.NewTimer(interval)

			select {
			case <-timer.C:
				attempt++
				// Continue to next iteration
			case <-timeoutChan:
				timer.Stop()
				return errors.New("operation timed out after " + w.timeout.String() + " secs")
			case <-ctx.Done():
				timer.Stop()
				return errors.New("context cancelled: " + ctx.Err().Error())
			}
		}
	}
}
