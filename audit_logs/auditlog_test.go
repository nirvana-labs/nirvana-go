// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package audit_logs_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/audit_logs"
	"github.com/nirvana-labs/nirvana-go/internal/testutil"
	"github.com/nirvana-labs/nirvana-go/option"
)

func TestAuditLogListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := nirvana.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.AuditLogs.List(context.TODO(), audit_logs.AuditLogListParams{
		Action:        nirvana.String("action"),
		ActorID:       nirvana.String("actor_id"),
		ActorType:     audit_logs.AuditLogListParamsActorTypeUser,
		ClientIP:      nirvana.String("client_ip"),
		CreatedAtMax:  nirvana.Time(time.Now()),
		CreatedAtMin:  nirvana.Time(time.Now()),
		Cursor:        nirvana.String("cursor"),
		Limit:         nirvana.Int(10),
		Method:        nirvana.String("method"),
		Path:          nirvana.String("path"),
		Sort:          nirvana.String("sort"),
		StatusCodeMax: nirvana.Int(0),
		StatusCodeMin: nirvana.Int(0),
		TargetID:      nirvana.String("target_id"),
		TargetType:    nirvana.String("target_type"),
	})
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAuditLogGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := nirvana.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.AuditLogs.Get(context.TODO(), "audit_log_id")
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
