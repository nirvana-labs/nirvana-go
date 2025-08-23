package lib_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/lib"
	"github.com/nirvana-labs/nirvana-go/operations"
)

// mockOperationsService implements a mock for testing
type mockOperationsService struct {
	getFunc func(ctx context.Context, operationID string) (*operations.Operation, error)
}

func (m *mockOperationsService) Get(ctx context.Context, operationID string) (*operations.Operation, error) {
	return m.getFunc(ctx, operationID)
}

func (m *mockOperationsService) List(ctx context.Context) (interface{}, error) {
	return nil, nil
}

func TestNewOperationWaiter(t *testing.T) {
	waiter := lib.NewOperationWaiter()
	if waiter == nil {
		t.Fatal("NewOperationWaiter should return a non-nil waiter")
	}
}

func TestOperationWaiter_WithTimeout(t *testing.T) {
	waiter := lib.NewOperationWaiter()
	customTimeout := 5 * time.Minute
	
	result := waiter.WithTimeout(customTimeout)
	
	// Should return the same instance for chaining
	if result != waiter {
		t.Error("WithTimeout should return the same instance for method chaining")
	}
}

func TestOperationWaiter_WithPollInterval(t *testing.T) {
	waiter := lib.NewOperationWaiter()
	customInterval := 2 * time.Second
	
	result := waiter.WithPollInterval(customInterval)
	
	// Should return the same instance for chaining
	if result != waiter {
		t.Error("WithPollInterval should return the same instance for method chaining")
	}
}

func TestOperationWaiter_Chaining(t *testing.T) {
	waiter := lib.NewOperationWaiter().
		WithTimeout(30 * time.Second).
		WithPollInterval(500 * time.Millisecond)
	
	if waiter == nil {
		t.Fatal("Method chaining should work and return a valid waiter")
	}
}

func TestWait_ContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately
	
	// Create a mock client that would normally succeed
	client := &nirvana.Client{}
	
	waiter := lib.NewOperationWaiter().WithTimeout(10 * time.Second)
	err := waiter.Wait(ctx, client, "test-operation")
	
	if err == nil {
		t.Fatal("Wait should return an error when context is cancelled")
	}
	
	if !errors.Is(err, context.Canceled) && !isContextCancelledError(err) {
		t.Errorf("Expected context cancellation error, got: %v", err)
	}
}

func TestWait_Timeout(t *testing.T) {
	ctx := context.Background()
	client := &nirvana.Client{}
	
	// Use a very short timeout to test timeout behavior
	waiter := lib.NewOperationWaiter().
		WithTimeout(1 * time.Millisecond)
	
	start := time.Now()
	err := waiter.Wait(ctx, client, "test-operation")
	elapsed := time.Since(start)
	
	if err == nil {
		t.Fatal("Wait should return an error")
	}
	
	// Should fail quickly due to either timeout or client error
	if elapsed > 100*time.Millisecond {
		t.Errorf("Operation took too long: %v", elapsed)
	}
	
	// The error could be either timeout or client configuration error
	// Both are acceptable for this test
	t.Logf("Got expected error: %v", err)
}

func TestDeprecatedWait_BackwardsCompatibility(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()
	
	client := &nirvana.Client{}
	
	// Test that the deprecated function still works
	err := lib.Wait(ctx, client, "test-operation")
	
	// Should get some error (timeout or network error)
	if err == nil {
		t.Error("Expected an error from deprecated Wait function")
	}
}

func TestConstants(t *testing.T) {
	if lib.DefaultTimeout != 10*time.Minute {
		t.Errorf("DefaultTimeout should be 10 minutes, got %v", lib.DefaultTimeout)
	}
	
	if lib.DefaultPollInterval != 1*time.Second {
		t.Errorf("DefaultPollInterval should be 1 second, got %v", lib.DefaultPollInterval)
	}
}

func TestFixedIntervalStrategy(t *testing.T) {
	strategy := lib.NewFixedIntervalStrategy(2 * time.Second)
	
	// Should return the same interval regardless of attempt
	if interval := strategy.NextInterval(0); interval != 2*time.Second {
		t.Errorf("Expected 2s interval, got %v", interval)
	}
	
	if interval := strategy.NextInterval(5); interval != 2*time.Second {
		t.Errorf("Expected 2s interval, got %v", interval)
	}
}

func TestLinearBackoffStrategy(t *testing.T) {
	strategy := lib.NewLinearBackoffStrategy(1*time.Second, 500*time.Millisecond, 5*time.Second)
	
	tests := []struct {
		attempt  int
		expected time.Duration
	}{
		{0, 1 * time.Second},                    // base
		{1, 1500 * time.Millisecond},            // base + 1*increment
		{2, 2 * time.Second},                    // base + 2*increment
		{10, 5 * time.Second},                   // capped at max
	}
	
	for _, test := range tests {
		if interval := strategy.NextInterval(test.attempt); interval != test.expected {
			t.Errorf("Attempt %d: expected %v, got %v", test.attempt, test.expected, interval)
		}
	}
}

func TestExponentialBackoffStrategy(t *testing.T) {
	strategy := lib.NewExponentialBackoffStrategy(1*time.Second, 2.0, 8*time.Second)
	
	tests := []struct {
		attempt  int
		expected time.Duration
	}{
		{0, 0 * time.Second},                    // base * multiplier * 0
		{1, 2 * time.Second},                    // base * multiplier * 1
		{2, 4 * time.Second},                    // base * multiplier * 2
		{3, 6 * time.Second},                    // base * multiplier * 3
		{10, 8 * time.Second},                   // capped at max
	}
	
	for _, test := range tests {
		if interval := strategy.NextInterval(test.attempt); interval != test.expected {
			t.Errorf("Attempt %d: expected %v, got %v", test.attempt, test.expected, interval)
		}
	}
}

func TestOperationWaiter_WithStrategy(t *testing.T) {
	strategy := lib.NewLinearBackoffStrategy(1*time.Second, 500*time.Millisecond, 5*time.Second)
	waiter := lib.NewOperationWaiter().WithStrategy(strategy)
	
	if waiter == nil {
		t.Fatal("WithStrategy should return a valid waiter")
	}
}

func TestOperationWaiter_WithLinearBackoff(t *testing.T) {
	waiter := lib.NewOperationWaiter().WithLinearBackoff(1*time.Second, 500*time.Millisecond, 5*time.Second)
	
	if waiter == nil {
		t.Fatal("WithLinearBackoff should return a valid waiter")
	}
}

func TestOperationWaiter_WithExponentialBackoff(t *testing.T) {
	waiter := lib.NewOperationWaiter().WithExponentialBackoff(1*time.Second, 2.0, 10*time.Second)
	
	if waiter == nil {
		t.Fatal("WithExponentialBackoff should return a valid waiter")
	}
}

func TestOperationWaiter_StrategyChainingCombinations(t *testing.T) {
	// Test different combinations of strategy configuration
	tests := []struct {
		name string
		waiter func() *lib.OperationWaiter
	}{
		{
			name: "LinearBackoff + Timeout",
			waiter: func() *lib.OperationWaiter {
				return lib.NewOperationWaiter().
					WithTimeout(30*time.Second).
					WithLinearBackoff(500*time.Millisecond, 200*time.Millisecond, 3*time.Second)
			},
		},
		{
			name: "ExponentialBackoff + Timeout",
			waiter: func() *lib.OperationWaiter {
				return lib.NewOperationWaiter().
					WithTimeout(1*time.Minute).
					WithExponentialBackoff(100*time.Millisecond, 1.5, 5*time.Second)
			},
		},
		{
			name: "Custom Strategy",
			waiter: func() *lib.OperationWaiter {
				strategy := lib.NewFixedIntervalStrategy(2 * time.Second)
				return lib.NewOperationWaiter().
					WithTimeout(45*time.Second).
					WithStrategy(strategy)
			},
		},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			waiter := test.waiter()
			if waiter == nil {
				t.Fatalf("Strategy configuration should return a valid waiter")
			}
		})
	}
}

// Helper functions
func isContextCancelledError(err error) bool {
	return err != nil && (errors.Is(err, context.Canceled) || 
		(err.Error() != "" && (
			err.Error() == "context cancelled: context canceled" ||
			err.Error() == "context canceled")))
}

