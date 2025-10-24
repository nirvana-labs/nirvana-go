// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package networking

import (
	"time"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
	"github.com/nirvana-labs/nirvana-go/shared"
)

// ConnectService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectService] method instead.
type ConnectService struct {
	Options     []option.RequestOption
	Connections ConnectConnectionService
	Routes      ConnectRouteService
}

// NewConnectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewConnectService(opts ...option.RequestOption) (r ConnectService) {
	r = ConnectService{}
	r.Options = opts
	r.Connections = NewConnectConnectionService(opts...)
	r.Routes = NewConnectRouteService(opts...)
	return
}

// Connect Connection speed in Mbps
type ConnectBandwidthMbps int64

const (
	ConnectBandwidthMbps50   ConnectBandwidthMbps = 50
	ConnectBandwidthMbps200  ConnectBandwidthMbps = 200
	ConnectBandwidthMbps500  ConnectBandwidthMbps = 500
	ConnectBandwidthMbps1000 ConnectBandwidthMbps = 1000
	ConnectBandwidthMbps2000 ConnectBandwidthMbps = 2000
)

// Connect Connection details.
type ConnectConnection struct {
	// Unique identifier for the Connect Connection
	ID string `json:"id,required"`
	// ASN
	ASN int64 `json:"asn,required"`
	// AWS provider configuration
	AWS ConnectConnectionAWSConfig `json:"aws,required"`
	// Connect Connection speed in Mbps
	//
	// Any of 50, 200, 500, 1000, 2000.
	BandwidthMbps int64 `json:"bandwidth_mbps,required"`
	// CIDRs for the Connect Connection
	CIDRs []string `json:"cidrs,required"`
	// When the Connect Connection was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Name of the Connect Connection
	Name string `json:"name,required"`
	// Provider ASN
	ProviderASN int64 `json:"provider_asn,required"`
	// Provider CIDRs for the Connect Connection
	ProviderCIDRs []string `json:"provider_cidrs,required"`
	// Provider Router IP for the Connect Connection
	ProviderRouterIP string `json:"provider_router_ip,required"`
	// Region the resource is in.
	//
	// Any of "us-sea-1", "us-sva-1", "us-chi-1", "us-wdc-1", "eu-frk-1", "ap-sin-1",
	// "ap-seo-1", "ap-tyo-1".
	Region shared.RegionName `json:"region,required"`
	// Router IP
	RouterIP string `json:"router_ip,required"`
	// Status of the resource.
	//
	// Any of "pending", "creating", "updating", "ready", "deleting", "deleted",
	// "error".
	Status shared.ResourceStatus `json:"status,required"`
	// Tags to attach to the Connect Connection
	Tags []string `json:"tags,required"`
	// When the Connect Connection was updated
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ASN              respjson.Field
		AWS              respjson.Field
		BandwidthMbps    respjson.Field
		CIDRs            respjson.Field
		CreatedAt        respjson.Field
		Name             respjson.Field
		ProviderASN      respjson.Field
		ProviderCIDRs    respjson.Field
		ProviderRouterIP respjson.Field
		Region           respjson.Field
		RouterIP         respjson.Field
		Status           respjson.Field
		Tags             respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectConnection) RawJSON() string { return r.JSON.raw }
func (r *ConnectConnection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AWS provider configuration
type ConnectConnectionAWSConfig struct {
	// AWS region where the connection is established
	Region string `json:"region,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectConnectionAWSConfig) RawJSON() string { return r.JSON.raw }
func (r *ConnectConnectionAWSConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AWS provider configuration
//
// The properties AccountID, Region are required.
type ConnectConnectionAWSConfigRequestParam struct {
	// AWS account id
	AccountID string `json:"account_id,required"`
	// AWS region where the connection will be established
	Region string `json:"region,required"`
	paramObj
}

func (r ConnectConnectionAWSConfigRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ConnectConnectionAWSConfigRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConnectConnectionAWSConfigRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectConnectionList struct {
	Items []ConnectConnection `json:"items,required"`
	// Pagination response details.
	Pagination shared.Pagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectConnectionList) RawJSON() string { return r.JSON.raw }
func (r *ConnectConnectionList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Routes supported for Connect.
type ConnectRoute struct {
	// Region the resource is in.
	//
	// Any of "us-sea-1", "us-sva-1", "us-chi-1", "us-wdc-1", "eu-frk-1", "ap-sin-1",
	// "ap-seo-1", "ap-tyo-1".
	NirvanaRegion shared.RegionName `json:"nirvana_region,required"`
	// Provider name.
	Provider string `json:"provider,required"`
	// Provider region name.
	ProviderRegion string `json:"provider_region,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NirvanaRegion  respjson.Field
		Provider       respjson.Field
		ProviderRegion respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectRoute) RawJSON() string { return r.JSON.raw }
func (r *ConnectRoute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectRouteList struct {
	Items []ConnectRoute `json:"items,required"`
	// Pagination response details.
	Pagination shared.Pagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectRouteList) RawJSON() string { return r.JSON.raw }
func (r *ConnectRouteList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
