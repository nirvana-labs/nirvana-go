package lib

import (
	"context"
	"errors"
	"time"

	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/operations"
)

// WaitForOperation polls the operation status until it's 'done' or times out.
func WaitForOperation(ctx context.Context, client *nirvana.Client, operationID string) error {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	timeoutDuration := 10 * time.Minute
	timeout := time.After(timeoutDuration)

	for {
		select {
		case <-ticker.C:
			op, err := client.Operations.Get(ctx, operationID)
			if err != nil {
				return errors.New("failed to get operation status: " + err.Error())
			}

			if op.Status == operations.OperationStatusDone {
				return nil
			}

			if op.Status == operations.OperationStatusFailed {
				return errors.New("operation failed")
			}
		case <-timeout:
			return errors.New("operation timed out after " + timeoutDuration.String() + " secs")
		case <-ctx.Done():
			return errors.New("context cancelled: " + ctx.Err().Error())
		}
	}
}
