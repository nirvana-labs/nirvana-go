// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall_rules_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/firewall_rules"
	"github.com/nirvana-labs/nirvana-go/internal/testutil"
	"github.com/nirvana-labs/nirvana-go/option"
)

func TestNetworkingFirewallRuleNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := nirvana.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAuthToken("My Auth Token"),
	)
	_, err := client.Networking.FirewallRules.New(
		context.TODO(),
		"vpc_id",
		firewall_rules.NetworkingFirewallRuleNewParams{
			Destination: nirvana.F(firewall_rules.FirewallRuleEndpointParam{
				Address: nirvana.F("0.0.0.0/0"),
				Ports:   nirvana.F([]string{"22", "80", "443"}),
			}),
			Name:     nirvana.F("my-firewall-rule"),
			Protocol: nirvana.F("tcp"),
			Source: nirvana.F(firewall_rules.FirewallRuleEndpointParam{
				Address: nirvana.F("0.0.0.0/0"),
				Ports:   nirvana.F([]string{"22", "80", "443"}),
			}),
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

func TestNetworkingFirewallRuleUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := nirvana.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAuthToken("My Auth Token"),
	)
	_, err := client.Networking.FirewallRules.Update(
		context.TODO(),
		"vpc_id",
		"firewall_rule_id",
		firewall_rules.NetworkingFirewallRuleUpdateParams{
			Destination: nirvana.F(firewall_rules.FirewallRuleEndpointParam{
				Address: nirvana.F("0.0.0.0/0"),
				Ports:   nirvana.F([]string{"22", "80", "443"}),
			}),
			Name:     nirvana.F("my-firewall-rule"),
			Protocol: nirvana.F(firewall_rules.NetworkingFirewallRuleUpdateParamsProtocolTcp),
			Source: nirvana.F(firewall_rules.FirewallRuleEndpointParam{
				Address: nirvana.F("0.0.0.0/0"),
				Ports:   nirvana.F([]string{"22", "80", "443"}),
			}),
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

func TestNetworkingFirewallRuleList(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := nirvana.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAuthToken("My Auth Token"),
	)
	_, err := client.Networking.FirewallRules.List(context.TODO(), "vpc_id")
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkingFirewallRuleDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := nirvana.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAuthToken("My Auth Token"),
	)
	_, err := client.Networking.FirewallRules.Delete(
		context.TODO(),
		"vpc_id",
		"firewall_rule_id",
	)
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkingFirewallRuleGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := nirvana.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAuthToken("My Auth Token"),
	)
	_, err := client.Networking.FirewallRules.Get(
		context.TODO(),
		"vpc_id",
		"firewall_rule_id",
	)
	if err != nil {
		var apierr *nirvana.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
