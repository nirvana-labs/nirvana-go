// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vektor

import (
	"github.com/nirvana-labs/nirvana-go/option"
)

// RegistryService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegistryService] method instead.
type RegistryService struct {
	Options       []option.RequestOption
	Assets        RegistryAssetService
	Blockchains   RegistryBlockchainService
	Venues        RegistryVenueService
	Errors        RegistryErrorService
	LendMarkets   RegistryLendMarketService
	BorrowMarkets RegistryBorrowMarketService
	LPPools       RegistryLPPoolService
}

// NewRegistryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRegistryService(opts ...option.RequestOption) (r RegistryService) {
	r = RegistryService{}
	r.Options = opts
	r.Assets = NewRegistryAssetService(opts...)
	r.Blockchains = NewRegistryBlockchainService(opts...)
	r.Venues = NewRegistryVenueService(opts...)
	r.Errors = NewRegistryErrorService(opts...)
	r.LendMarkets = NewRegistryLendMarketService(opts...)
	r.BorrowMarkets = NewRegistryBorrowMarketService(opts...)
	r.LPPools = NewRegistryLPPoolService(opts...)
	return
}
