// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_keys_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/api_keys"
	"github.com/nirvana-labs/nirvana-go/internal/testutil"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/shared"
)

func TestAPIKeyNewWithOptionalParams(t *testing.T) {
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
	_, err := client.APIKeys.New(context.TODO(), api_keys.APIKeyNewParams{
		ExpiresAt: time.Now(),
		Name:      "My API Key",
		SourceIPRule: shared.SourceIPRuleParam{
			Allowed: []string{"192.168.1.0/24", "10.0.0.0/8"},
			Blocked: []string{"192.168.1.100/32"},
		},
		StartsAt: nirvana.Time(time.Now()),
		Tags:     []string{"production", "ethereum"},
	})
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAPIKeyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.APIKeys.Update(
		context.TODO(),
		"api_key_id",
		api_keys.APIKeyUpdateParams{
			Name: nirvana.String("My Updated API Key"),
			SourceIPRule: shared.SourceIPRuleParam{
				Allowed: []string{"192.168.1.0/24", "10.0.0.0/8"},
				Blocked: []string{"192.168.1.100/32"},
			},
			Tags: []string{"production", "ethereum"},
		},
	)
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAPIKeyListWithOptionalParams(t *testing.T) {
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
	_, err := client.APIKeys.List(context.TODO(), api_keys.APIKeyListParams{
		Cursor: nirvana.String("cursor"),
		Limit:  nirvana.Int(10),
	})
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAPIKeyDelete(t *testing.T) {
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
	err := client.APIKeys.Delete(context.TODO(), "api_key_id")
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAPIKeyGet(t *testing.T) {
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
	_, err := client.APIKeys.Get(context.TODO(), "api_key_id")
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
