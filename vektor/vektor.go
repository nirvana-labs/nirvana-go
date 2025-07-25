// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vektor

import (
	"encoding/json"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
)

// VektorService contains methods and other services that help with interacting
// with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVektorService] method instead.
type VektorService struct {
	Options     []option.RequestOption
	Registry    RegistryService
	Balances    BalanceService
	Prices      PriceService
	Lend        LendService
	Borrow      BorrowService
	LP          LPService
	Buy         BuyService
	Sell        SellService
	Move        MoveService
	Wrap        WrapService
	Bridge      BridgeService
	Lock        LockService
	Vote        VoteService
	Incentivize IncentivizeService
	Executions  ExecutionService
}

// NewVektorService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVektorService(opts ...option.RequestOption) (r VektorService) {
	r = VektorService{}
	r.Options = opts
	r.Registry = NewRegistryService(opts...)
	r.Balances = NewBalanceService(opts...)
	r.Prices = NewPriceService(opts...)
	r.Lend = NewLendService(opts...)
	r.Borrow = NewBorrowService(opts...)
	r.LP = NewLPService(opts...)
	r.Buy = NewBuyService(opts...)
	r.Sell = NewSellService(opts...)
	r.Move = NewMoveService(opts...)
	r.Wrap = NewWrapService(opts...)
	r.Bridge = NewBridgeService(opts...)
	r.Lock = NewLockService(opts...)
	r.Vote = NewVoteService(opts...)
	r.Incentivize = NewIncentivizeService(opts...)
	r.Executions = NewExecutionService(opts...)
	return
}

type Account = string

type AddressEVM = string

// APY Data
type APY struct {
	// An arbitrary precision decimal represented as a string
	APY Decimal `json:"apy,required"`
	// On-chain asset (aka token)
	Asset Asset `json:"asset,required"`
	// APY Type
	//
	// Any of "fixed", "variable".
	Type APYType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APY         respjson.Field
		Asset       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APY) RawJSON() string { return r.JSON.raw }
func (r *APY) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APY Type
type APYType string

const (
	APYTypeFixed    APYType = "fixed"
	APYTypeVariable APYType = "variable"
)

// On-chain asset (aka token)
type Asset struct {
	// An asset ID, represented as a TypeID with `asset` prefix
	ID AssetID `json:"id,required"`
	// An EVM address
	Address string `json:"address,required"`
	// Data about a blockchain
	Blockchain Blockchain `json:"blockchain,required"`
	// Asset's decimal places
	Decimals int64 `json:"decimals,required"`
	// Asset's name
	Name string `json:"name,required"`
	// Asset's symbol
	Symbol string `json:"symbol,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Address     respjson.Field
		Blockchain  respjson.Field
		Decimals    respjson.Field
		Name        respjson.Field
		Symbol      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Asset) RawJSON() string { return r.JSON.raw }
func (r *Asset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetID = string

type AssetIDOrAddressEVMOrAssetSymbol = string

type AssetSymbol = string

// Balance properties
type Balance struct {
	// An EVM address
	Account AddressEVM `json:"account,required"`
	// An arbitrary precision decimal represented as a string
	Amount Decimal `json:"amount,required"`
	// An asset ID, represented as a TypeID with `asset` prefix
	AssetID AssetID `json:"asset_id,required"`
	// Information about a specific block number and its timestamp
	Blockstamp Blockstamp `json:"blockstamp,required"`
	// An arbitrary precision decimal represented as a string
	Value Decimal `json:"value,required"`
	// An asset symbol
	QuoteAssetSymbol string `json:"quote_asset_symbol,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account          respjson.Field
		Amount           respjson.Field
		AssetID          respjson.Field
		Blockstamp       respjson.Field
		Value            respjson.Field
		QuoteAssetSymbol respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Balance) RawJSON() string { return r.JSON.raw }
func (r *Balance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlockNumber = int64

// Data about a blockchain
type Blockchain struct {
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	ID BlockchainID `json:"id,required"`
	// Data about an EVM blockchain
	ChainData EVMChainData `json:"chain_data,required"`
	// Blockchain ecosystem
	//
	// Any of "evm".
	ChainType ChainType `json:"chain_type,required"`
	// The blockchain's explorer URL
	ExplorerURL string `json:"explorer_url,required"`
	// BlockchainName
	Name BlockchainName `json:"name,required"`
	// Blockchain's network
	Network NetworkName `json:"network,required"`
	// A blockchain symbol
	Symbol BlockchainSymbol `json:"symbol,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ChainData   respjson.Field
		ChainType   respjson.Field
		ExplorerURL respjson.Field
		Name        respjson.Field
		Network     respjson.Field
		Symbol      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Blockchain) RawJSON() string { return r.JSON.raw }
func (r *Blockchain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BlockchainID = string

type BlockchainIDOrBlockchainSymbol = string

type BlockchainName = string

type BlockchainSymbol = string

// Information about a specific block number and its timestamp
type Blockstamp struct {
	// A block number
	BlockNumber BlockNumber `json:"block_number,required"`
	// Data about a blockchain
	Blockchain Blockchain `json:"blockchain,required"`
	// ISO8601 Timestamp
	Timestamp Timestamp `json:"timestamp,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlockNumber respjson.Field
		Blockchain  respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Blockstamp) RawJSON() string { return r.JSON.raw }
func (r *Blockstamp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A borrow account
type BorrowAccount struct {
	// An EVM address
	Account Account `json:"account,required"`
	// An arbitrary precision decimal represented as a string
	AvailableBorrow Decimal `json:"available_borrow,required"`
	// Information about a specific block number and its timestamp
	Blockstamp Blockstamp `json:"blockstamp,required"`
	// An arbitrary precision decimal represented as a string
	HealthFactor Decimal `json:"health_factor,required"`
	// An asset symbol
	QuoteAssetSymbol string `json:"quote_asset_symbol,required"`
	// An arbitrary precision decimal represented as a string
	TotalCollateral Decimal `json:"total_collateral,required"`
	// An arbitrary precision decimal represented as a string
	TotalDebt Decimal `json:"total_debt,required"`
	// On-chain venue
	Venue Venue `json:"venue,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account          respjson.Field
		AvailableBorrow  respjson.Field
		Blockstamp       respjson.Field
		HealthFactor     respjson.Field
		QuoteAssetSymbol respjson.Field
		TotalCollateral  respjson.Field
		TotalDebt        respjson.Field
		Venue            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BorrowAccount) RawJSON() string { return r.JSON.raw }
func (r *BorrowAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A borrow market
type BorrowMarket struct {
	// A lend/borrow market ID, represented as a TypeID with `lend_borrow_market`
	// prefix
	ID LendBorrowMarketID `json:"id,required"`
	// APY Data for lend/borrow markets
	APYs LendBorrowAPYs `json:"apys,required"`
	// On-chain asset (aka token)
	Asset Asset `json:"asset,required"`
	// Information about a specific block number and its timestamp
	Blockstamp Blockstamp `json:"blockstamp,required"`
	// An arbitrary precision decimal represented as a string
	Liquidity Decimal `json:"liquidity,required"`
	// An asset symbol
	QuoteAssetSymbol string `json:"quote_asset_symbol,required"`
	// An arbitrary precision decimal represented as a string
	TotalDebt Decimal `json:"total_debt,required"`
	// On-chain venue
	Venue Venue `json:"venue,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		APYs             respjson.Field
		Asset            respjson.Field
		Blockstamp       respjson.Field
		Liquidity        respjson.Field
		QuoteAssetSymbol respjson.Field
		TotalDebt        respjson.Field
		Venue            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BorrowMarket) RawJSON() string { return r.JSON.raw }
func (r *BorrowMarket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A borrow position
type BorrowPosition struct {
	// An EVM address
	Account Account `json:"account,required"`
	// Information about a specific block number and its timestamp
	Blockstamp Blockstamp `json:"blockstamp,required"`
	// An arbitrary precision decimal represented as a string
	DebtAmount Decimal `json:"debt_amount,required"`
	// A borrow market
	Market BorrowMarket `json:"market,required"`
	// An asset symbol
	QuoteAssetSymbol string `json:"quote_asset_symbol,required"`
	// An arbitrary precision decimal represented as a string
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account          respjson.Field
		Blockstamp       respjson.Field
		DebtAmount       respjson.Field
		Market           respjson.Field
		QuoteAssetSymbol respjson.Field
		Value            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BorrowPosition) RawJSON() string { return r.JSON.raw }
func (r *BorrowPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A bridge quote
type BridgeQuote struct {
	// Estimated cost of an operation
	FeeEstimate BridgeQuoteFeeEstimate `json:"fee_estimate,required"`
	// An arbitrary precision decimal represented as a string
	ReceiveAmount Decimal `json:"receive_amount,required"`
	// On-chain asset (aka token)
	ReceiveAsset Asset `json:"receive_asset,required"`
	// Estimated time to receive the asset (in seconds)
	ReceiveTimeEst int64 `json:"receive_time_est,required"`
	// The route name
	Route string `json:"route,required"`
	// An arbitrary precision decimal represented as a string
	SendAmount Decimal `json:"send_amount,required"`
	// On-chain asset (aka token)
	SendAsset Asset `json:"send_asset,required"`
	// On-chain venue
	Venue Venue `json:"venue,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FeeEstimate    respjson.Field
		ReceiveAmount  respjson.Field
		ReceiveAsset   respjson.Field
		ReceiveTimeEst respjson.Field
		Route          respjson.Field
		SendAmount     respjson.Field
		SendAsset      respjson.Field
		Venue          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeQuote) RawJSON() string { return r.JSON.raw }
func (r *BridgeQuote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Estimated cost of an operation
type BridgeQuoteFeeEstimate struct {
	// An arbitrary precision decimal represented as a string
	Amount Decimal `json:"amount,required"`
	// On-chain asset (aka token)
	Asset Asset `json:"asset,required"`
	// An arbitrary precision decimal represented as a string
	Cost string `json:"cost,required"`
	// An asset symbol
	QuoteAssetSymbol string `json:"quote_asset_symbol,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount           respjson.Field
		Asset            respjson.Field
		Cost             respjson.Field
		QuoteAssetSymbol respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BridgeQuoteFeeEstimate) RawJSON() string { return r.JSON.raw }
func (r *BridgeQuoteFeeEstimate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A buy quote
type BuyQuote struct {
	// Data about a blockchain
	Blockchain Blockchain `json:"blockchain,required"`
	// Estimated cost of an operation
	FeeEstimate BuyQuoteFeeEstimate `json:"fee_estimate,required"`
	// An asset symbol
	QuoteAssetSymbol string `json:"quote_asset_symbol,required"`
	// Quote info for buy quotes
	QuoteInfo BuyQuoteQuoteInfoUnion `json:"quote_info,required"`
	// An arbitrary precision decimal represented as a string
	QuoteValue string `json:"quote_value,required"`
	// An arbitrary precision decimal represented as a string
	ReceiveAmount Decimal `json:"receive_amount,required"`
	// On-chain asset (aka token)
	ReceiveAsset Asset `json:"receive_asset,required"`
	// An arbitrary precision decimal represented as a string
	SpendAmount Decimal `json:"spend_amount,required"`
	// On-chain asset (aka token)
	SpendAsset Asset `json:"spend_asset,required"`
	// On-chain venue
	Venue Venue `json:"venue,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blockchain       respjson.Field
		FeeEstimate      respjson.Field
		QuoteAssetSymbol respjson.Field
		QuoteInfo        respjson.Field
		QuoteValue       respjson.Field
		ReceiveAmount    respjson.Field
		ReceiveAsset     respjson.Field
		SpendAmount      respjson.Field
		SpendAsset       respjson.Field
		Venue            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuyQuote) RawJSON() string { return r.JSON.raw }
func (r *BuyQuote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Estimated cost of an operation
type BuyQuoteFeeEstimate struct {
	// An arbitrary precision decimal represented as a string
	Amount Decimal `json:"amount,required"`
	// On-chain asset (aka token)
	Asset Asset `json:"asset,required"`
	// An arbitrary precision decimal represented as a string
	Cost string `json:"cost,required"`
	// An asset symbol
	QuoteAssetSymbol string `json:"quote_asset_symbol,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount           respjson.Field
		Asset            respjson.Field
		Cost             respjson.Field
		QuoteAssetSymbol respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BuyQuoteFeeEstimate) RawJSON() string { return r.JSON.raw }
func (r *BuyQuoteFeeEstimate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// BuyQuoteQuoteInfoUnion contains all possible properties and values from
// [QuoteInfoUniswapV2], [QuoteInfoUniswapV3].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type BuyQuoteQuoteInfoUnion struct {
	PoolIDs []string `json:"pool_ids"`
	JSON    struct {
		PoolIDs respjson.Field
		raw     string
	} `json:"-"`
}

func (u BuyQuoteQuoteInfoUnion) AsQuoteInfoUniswapV2() (v QuoteInfoUniswapV2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u BuyQuoteQuoteInfoUnion) AsQuoteInfoUniswapV3() (v QuoteInfoUniswapV3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u BuyQuoteQuoteInfoUnion) RawJSON() string { return u.JSON.raw }

func (r *BuyQuoteQuoteInfoUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Blockchain ecosystem
type ChainType string

const (
	ChainTypeEVM ChainType = "evm"
)

type Decimal = string

// Data about an EVM blockchain
type EVMChainData struct {
	// Chain ID
	ChainID int64 `json:"chain_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChainID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EVMChainData) RawJSON() string { return r.JSON.raw }
func (r *EVMChainData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EVMChainData to a EVMChainDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EVMChainDataParam.Overrides()
func (r EVMChainData) ToParam() EVMChainDataParam {
	return param.Override[EVMChainDataParam](json.RawMessage(r.RawJSON()))
}

// Data about an EVM blockchain
//
// The property ChainID is required.
type EVMChainDataParam struct {
	// Chain ID
	ChainID int64 `json:"chain_id,required"`
	paramObj
}

func (r EVMChainDataParam) MarshalJSON() (data []byte, err error) {
	type shadow EVMChainDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EVMChainDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An execution
type Execution struct {
	// An execution ID, represented as a TypeID with `execution` prefix
	ID ExecutionID `json:"id,required"`
	// ISO8601 Timestamp
	CreatedAt Timestamp `json:"created_at,required"`
	// An execution request
	Request ExecutionRequestUnion `json:"request,required"`
	// An execution step ID, represented as a TypeID with `execution_step` prefix
	RequestActionStepID string `json:"request_action_step_id,required"`
	// The state of an execution
	//
	// Any of "not_started", "started", "success", "aborted", "error".
	State         ExecutionState `json:"state,required"`
	StepActionURL string         `json:"step_action_url,required"`
	// A list of execution steps
	Steps []ExecutionStep `json:"steps,required"`
	// ISO8601 Timestamp
	UpdatedAt Timestamp `json:"updated_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		Request             respjson.Field
		RequestActionStepID respjson.Field
		State               respjson.Field
		StepActionURL       respjson.Field
		Steps               respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Execution) RawJSON() string { return r.JSON.raw }
func (r *Execution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExecutionRequestUnion contains all possible properties and values from
// [ExecutionRequestBorrowRequestFull], [ExecutionRequestBuyRequestFull],
// [ExecutionRequestLendRequestFull],
// [ExecutionRequestLendSetCollateralRequestFull],
// [ExecutionRequestLendWithdrawRequestFull], [ExecutionRequestMoveRequestFull],
// [ExecutionRequestSellRequestFull], [ExecutionRequestWrapRequestFull],
// [ExecutionRequestUnwrapRequestFull].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ExecutionRequestUnion struct {
	Amount string `json:"amount"`
	// This field is a union of [Asset], [ExecutionRequestLendWithdrawRequestFullAsset]
	Asset ExecutionRequestUnionAsset `json:"asset"`
	// This field is from variant [ExecutionRequestBorrowRequestFull].
	Blockchain Blockchain `json:"blockchain"`
	// This field is from variant [ExecutionRequestBorrowRequestFull].
	From   Account `json:"from"`
	Venues []Venue `json:"venues"`
	// This field is from variant [ExecutionRequestBuyRequestFull].
	ReceiveAmount Decimal `json:"receive_amount"`
	// This field is from variant [ExecutionRequestBuyRequestFull].
	ReceiveAsset Asset `json:"receive_asset"`
	// This field is from variant [ExecutionRequestBuyRequestFull].
	Slippage Decimal `json:"slippage"`
	// This field is from variant [ExecutionRequestBuyRequestFull].
	SpendAsset Asset `json:"spend_asset"`
	// This field is from variant [ExecutionRequestLendSetCollateralRequestFull].
	MarketID LendBorrowMarketID `json:"market_id"`
	// This field is from variant [ExecutionRequestLendSetCollateralRequestFull].
	Status bool `json:"status"`
	// This field is from variant [ExecutionRequestMoveRequestFull].
	To Account `json:"to"`
	// This field is from variant [ExecutionRequestSellRequestFull].
	SpendAmount Decimal `json:"spend_amount"`
	JSON        struct {
		Amount        respjson.Field
		Asset         respjson.Field
		Blockchain    respjson.Field
		From          respjson.Field
		Venues        respjson.Field
		ReceiveAmount respjson.Field
		ReceiveAsset  respjson.Field
		Slippage      respjson.Field
		SpendAsset    respjson.Field
		MarketID      respjson.Field
		Status        respjson.Field
		To            respjson.Field
		SpendAmount   respjson.Field
		raw           string
	} `json:"-"`
}

func (u ExecutionRequestUnion) AsBorrowRequestFull() (v ExecutionRequestBorrowRequestFull) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionRequestUnion) AsBuyRequestFull() (v ExecutionRequestBuyRequestFull) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionRequestUnion) AsLendRequestFull() (v ExecutionRequestLendRequestFull) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionRequestUnion) AsLendSetCollateralRequestFull() (v ExecutionRequestLendSetCollateralRequestFull) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionRequestUnion) AsLendWithdrawRequestFull() (v ExecutionRequestLendWithdrawRequestFull) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionRequestUnion) AsMoveRequestFull() (v ExecutionRequestMoveRequestFull) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionRequestUnion) AsSellRequestFull() (v ExecutionRequestSellRequestFull) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionRequestUnion) AsWrapRequestFull() (v ExecutionRequestWrapRequestFull) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionRequestUnion) AsUnwrapRequestFull() (v ExecutionRequestUnwrapRequestFull) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExecutionRequestUnion) RawJSON() string { return u.JSON.raw }

func (r *ExecutionRequestUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExecutionRequestUnionAsset is an implicit subunion of [ExecutionRequestUnion].
// ExecutionRequestUnionAsset provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [ExecutionRequestUnion].
type ExecutionRequestUnionAsset struct {
	// This field is from variant [Asset].
	ID      AssetID `json:"id"`
	Address string  `json:"address"`
	// This field is from variant [Asset].
	Blockchain Blockchain `json:"blockchain"`
	Decimals   int64      `json:"decimals"`
	Name       string     `json:"name"`
	Symbol     string     `json:"symbol"`
	JSON       struct {
		ID         respjson.Field
		Address    respjson.Field
		Blockchain respjson.Field
		Decimals   respjson.Field
		Name       respjson.Field
		Symbol     respjson.Field
		raw        string
	} `json:"-"`
}

func (r *ExecutionRequestUnionAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request to borrow an asset
type ExecutionRequestBorrowRequestFull struct {
	// An arbitrary precision decimal represented as a string
	Amount Decimal `json:"amount,required"`
	// On-chain asset (aka token)
	Asset Asset `json:"asset,required"`
	// Data about a blockchain
	Blockchain Blockchain `json:"blockchain,required"`
	// An EVM address
	From Account `json:"from,required"`
	// A list of venues
	Venues []Venue `json:"venues,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Asset       respjson.Field
		Blockchain  respjson.Field
		From        respjson.Field
		Venues      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionRequestBorrowRequestFull) RawJSON() string { return r.JSON.raw }
func (r *ExecutionRequestBorrowRequestFull) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request to buy an asset
type ExecutionRequestBuyRequestFull struct {
	// Data about a blockchain
	Blockchain Blockchain `json:"blockchain,required"`
	// An EVM address
	From Account `json:"from,required"`
	// An arbitrary precision decimal represented as a string
	ReceiveAmount Decimal `json:"receive_amount,required"`
	// On-chain asset (aka token)
	ReceiveAsset Asset `json:"receive_asset,required"`
	// An arbitrary precision decimal represented as a string
	Slippage Decimal `json:"slippage,required"`
	// On-chain asset (aka token)
	SpendAsset Asset `json:"spend_asset,required"`
	// A list of venues
	Venues []Venue `json:"venues,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blockchain    respjson.Field
		From          respjson.Field
		ReceiveAmount respjson.Field
		ReceiveAsset  respjson.Field
		Slippage      respjson.Field
		SpendAsset    respjson.Field
		Venues        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionRequestBuyRequestFull) RawJSON() string { return r.JSON.raw }
func (r *ExecutionRequestBuyRequestFull) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request to lend an asset
type ExecutionRequestLendRequestFull struct {
	// An arbitrary precision decimal represented as a string
	Amount Decimal `json:"amount,required"`
	// On-chain asset (aka token)
	Asset Asset `json:"asset,required"`
	// Data about a blockchain
	Blockchain Blockchain `json:"blockchain,required"`
	// An EVM address
	From Account `json:"from,required"`
	// A list of venues
	Venues []Venue `json:"venues,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Asset       respjson.Field
		Blockchain  respjson.Field
		From        respjson.Field
		Venues      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionRequestLendRequestFull) RawJSON() string { return r.JSON.raw }
func (r *ExecutionRequestLendRequestFull) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request to set/unset a position as collateral
type ExecutionRequestLendSetCollateralRequestFull struct {
	// Data about a blockchain
	Blockchain Blockchain `json:"blockchain,required"`
	// An EVM address
	From Account `json:"from,required"`
	// A lend/borrow market ID, represented as a TypeID with `lend_borrow_market`
	// prefix
	MarketID LendBorrowMarketID `json:"market_id,required"`
	Status   bool               `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blockchain  respjson.Field
		From        respjson.Field
		MarketID    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionRequestLendSetCollateralRequestFull) RawJSON() string { return r.JSON.raw }
func (r *ExecutionRequestLendSetCollateralRequestFull) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request to withdraw an asset
type ExecutionRequestLendWithdrawRequestFull struct {
	// An arbitrary precision decimal represented as a string
	Amount string `json:"amount,required"`
	// Data about a blockchain
	Blockchain Blockchain `json:"blockchain,required"`
	// An EVM address
	From Account `json:"from,required"`
	// A lend/borrow market ID, represented as a TypeID with `lend_borrow_market`
	// prefix
	MarketID LendBorrowMarketID `json:"market_id,required"`
	// On-chain asset (aka token)
	Asset ExecutionRequestLendWithdrawRequestFullAsset `json:"asset,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Blockchain  respjson.Field
		From        respjson.Field
		MarketID    respjson.Field
		Asset       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionRequestLendWithdrawRequestFull) RawJSON() string { return r.JSON.raw }
func (r *ExecutionRequestLendWithdrawRequestFull) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// On-chain asset (aka token)
type ExecutionRequestLendWithdrawRequestFullAsset struct {
	// An asset ID, represented as a TypeID with `asset` prefix
	ID AssetID `json:"id,required"`
	// An EVM address
	Address string `json:"address,required"`
	// Data about a blockchain
	Blockchain Blockchain `json:"blockchain,required"`
	// Asset's decimal places
	Decimals int64 `json:"decimals,required"`
	// Asset's name
	Name string `json:"name,required"`
	// Asset's symbol
	Symbol string `json:"symbol,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Address     respjson.Field
		Blockchain  respjson.Field
		Decimals    respjson.Field
		Name        respjson.Field
		Symbol      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionRequestLendWithdrawRequestFullAsset) RawJSON() string { return r.JSON.raw }
func (r *ExecutionRequestLendWithdrawRequestFullAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request to move assets from one account to another
type ExecutionRequestMoveRequestFull struct {
	// An arbitrary precision decimal represented as a string
	Amount Decimal `json:"amount,required"`
	// On-chain asset (aka token)
	Asset Asset `json:"asset,required"`
	// Data about a blockchain
	Blockchain Blockchain `json:"blockchain,required"`
	// An EVM address
	From Account `json:"from,required"`
	// An EVM address
	To Account `json:"to,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Asset       respjson.Field
		Blockchain  respjson.Field
		From        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionRequestMoveRequestFull) RawJSON() string { return r.JSON.raw }
func (r *ExecutionRequestMoveRequestFull) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request to sell an asset
type ExecutionRequestSellRequestFull struct {
	// Data about a blockchain
	Blockchain Blockchain `json:"blockchain,required"`
	// An EVM address
	From Account `json:"from,required"`
	// On-chain asset (aka token)
	ReceiveAsset Asset `json:"receive_asset,required"`
	// An arbitrary precision decimal represented as a string
	Slippage Decimal `json:"slippage,required"`
	// An arbitrary precision decimal represented as a string
	SpendAmount Decimal `json:"spend_amount,required"`
	// On-chain asset (aka token)
	SpendAsset Asset `json:"spend_asset,required"`
	// A list of venues
	Venues []Venue `json:"venues,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blockchain   respjson.Field
		From         respjson.Field
		ReceiveAsset respjson.Field
		Slippage     respjson.Field
		SpendAmount  respjson.Field
		SpendAsset   respjson.Field
		Venues       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionRequestSellRequestFull) RawJSON() string { return r.JSON.raw }
func (r *ExecutionRequestSellRequestFull) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request to wrap the native asset
type ExecutionRequestWrapRequestFull struct {
	// An arbitrary precision decimal represented as a string
	Amount Decimal `json:"amount,required"`
	// Data about a blockchain
	Blockchain Blockchain `json:"blockchain,required"`
	// An EVM address
	From Account `json:"from,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Blockchain  respjson.Field
		From        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionRequestWrapRequestFull) RawJSON() string { return r.JSON.raw }
func (r *ExecutionRequestWrapRequestFull) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A request to unwrap the wrapped native asset
type ExecutionRequestUnwrapRequestFull struct {
	// An arbitrary precision decimal represented as a string
	Amount Decimal `json:"amount,required"`
	// Data about a blockchain
	Blockchain Blockchain `json:"blockchain,required"`
	// An EVM address
	From Account `json:"from,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount      respjson.Field
		Blockchain  respjson.Field
		From        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionRequestUnwrapRequestFull) RawJSON() string { return r.JSON.raw }
func (r *ExecutionRequestUnwrapRequestFull) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An execution step
type ExecutionStep struct {
	// An execution step ID, represented as a TypeID with `execution_step` prefix
	ID ExecutionStepID `json:"id,required"`
	// ISO8601 Timestamp
	CreatedAt Timestamp `json:"created_at,required"`
	// The definition of an execution step
	Definition ExecutionStepDefinitionUnion `json:"definition,required"`
	Index      int64                        `json:"index,required"`
	// The type of an execution step
	//
	// Any of "evm_transaction_approve", "evm_transaction_borrow",
	// "evm_transaction_buy", "evm_transaction_lend",
	// "evm_transaction_lend_set_collateral", "evm_transaction_lend_withdraw",
	// "evm_transaction_move", "evm_transaction_permission", "evm_transaction_wrap",
	// "evm_transaction_unwrap", "evm_transaction_sell".
	Type string `json:"type,required"`
	// ISO8601 Timestamp
	UpdatedAt Timestamp `json:"updated_at,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Definition  respjson.Field
		Index       respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStep) RawJSON() string { return r.JSON.raw }
func (r *ExecutionStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExecutionStepDefinitionUnion contains all possible properties and values from
// [ExecutionStepDefinitionExecutionEVMTransactionApprove],
// [ExecutionStepDefinitionExecutionEVMTransactionBorrow],
// [ExecutionStepDefinitionExecutionEVMTransactionBuy],
// [ExecutionStepDefinitionExecutionEVMTransactionLend],
// [ExecutionStepDefinitionExecutionEVMTransactionLendSetCollateral],
// [ExecutionStepDefinitionExecutionEVMTransactionLendWithdraw],
// [ExecutionStepDefinitionExecutionEVMTransactionMove],
// [ExecutionStepDefinitionExecutionEVMTransactionPermission],
// [ExecutionStepDefinitionExecutionEVMTransactionUnwrap],
// [ExecutionStepDefinitionExecutionEVMTransactionWrap],
// [ExecutionStepDefinitionExecutionEVMTransactionSell].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ExecutionStepDefinitionUnion struct {
	Amount string `json:"amount"`
	// This field is from variant
	// [ExecutionStepDefinitionExecutionEVMTransactionApprove].
	Asset              Asset  `json:"asset"`
	BlockNumber        int64  `json:"block_number"`
	BroadcastedAt      string `json:"broadcasted_at"`
	ConfirmationTarget int64  `json:"confirmation_target"`
	ConfirmedAt        string `json:"confirmed_at"`
	// This field is from variant
	// [ExecutionStepDefinitionExecutionEVMTransactionApprove].
	CreatedAt         Timestamp `json:"created_at"`
	Data              string    `json:"data"`
	EffectiveGasPrice string    `json:"effective_gas_price"`
	// This field is a union of
	// [ExecutionStepDefinitionExecutionEVMTransactionApproveError],
	// [ExecutionStepDefinitionExecutionEVMTransactionBorrowError],
	// [ExecutionStepDefinitionExecutionEVMTransactionBuyError],
	// [ExecutionStepDefinitionExecutionEVMTransactionLendError],
	// [ExecutionStepDefinitionExecutionEVMTransactionLendSetCollateralError],
	// [ExecutionStepDefinitionExecutionEVMTransactionLendWithdrawError],
	// [ExecutionStepDefinitionExecutionEVMTransactionMoveError],
	// [ExecutionStepDefinitionExecutionEVMTransactionPermissionError],
	// [ExecutionStepDefinitionExecutionEVMTransactionUnwrapError],
	// [ExecutionStepDefinitionExecutionEVMTransactionWrapError],
	// [ExecutionStepDefinitionExecutionEVMTransactionSellError]
	Error     ExecutionStepDefinitionUnionError `json:"error"`
	ErroredAt string                            `json:"errored_at"`
	GasUsed   string                            `json:"gas_used"`
	Hash      string                            `json:"hash"`
	// This field is from variant
	// [ExecutionStepDefinitionExecutionEVMTransactionApprove].
	Payload  ExecutionEVMTransactionEIP1559Payload `json:"payload"`
	SignedAt string                                `json:"signed_at"`
	// This field is from variant
	// [ExecutionStepDefinitionExecutionEVMTransactionApprove].
	Spender Account `json:"spender"`
	// This field is from variant
	// [ExecutionStepDefinitionExecutionEVMTransactionApprove].
	State ExecutionEVMTransactionState `json:"state"`
	// This field is from variant
	// [ExecutionStepDefinitionExecutionEVMTransactionApprove].
	TargetState ExecutionEVMTransactionState `json:"target_state"`
	To          string                       `json:"to"`
	Type        string                       `json:"type"`
	// This field is from variant
	// [ExecutionStepDefinitionExecutionEVMTransactionApprove].
	UpdatedAt Timestamp `json:"updated_at"`
	Value     string    `json:"value"`
	// This field is from variant
	// [ExecutionStepDefinitionExecutionEVMTransactionBorrow].
	LendBorrowMarketID LendBorrowMarketID `json:"lend_borrow_market_id"`
	// This field is from variant
	// [ExecutionStepDefinitionExecutionEVMTransactionBorrow].
	VenueSymbol VenueSymbol `json:"venue_symbol"`
	// This field is from variant [ExecutionStepDefinitionExecutionEVMTransactionBuy].
	ApprovalTarget Account `json:"approval_target"`
	// This field is from variant [ExecutionStepDefinitionExecutionEVMTransactionBuy].
	MaxSpendAmount Decimal `json:"max_spend_amount"`
	// This field is a union of [BuyQuote], [SellQuote]
	Quote ExecutionStepDefinitionUnionQuote `json:"quote"`
	// This field is from variant [ExecutionStepDefinitionExecutionEVMTransactionBuy].
	Slippage Decimal `json:"slippage"`
	// This field is from variant
	// [ExecutionStepDefinitionExecutionEVMTransactionLendSetCollateral].
	Status bool `json:"status"`
	// This field is from variant
	// [ExecutionStepDefinitionExecutionEVMTransactionPermission].
	ContractAddress AddressEVM `json:"contract_address"`
	// This field is from variant
	// [ExecutionStepDefinitionExecutionEVMTransactionPermission].
	Permission bool `json:"permission"`
	// This field is from variant [ExecutionStepDefinitionExecutionEVMTransactionSell].
	MinReceiveAmount Decimal `json:"min_receive_amount"`
	JSON             struct {
		Amount             respjson.Field
		Asset              respjson.Field
		BlockNumber        respjson.Field
		BroadcastedAt      respjson.Field
		ConfirmationTarget respjson.Field
		ConfirmedAt        respjson.Field
		CreatedAt          respjson.Field
		Data               respjson.Field
		EffectiveGasPrice  respjson.Field
		Error              respjson.Field
		ErroredAt          respjson.Field
		GasUsed            respjson.Field
		Hash               respjson.Field
		Payload            respjson.Field
		SignedAt           respjson.Field
		Spender            respjson.Field
		State              respjson.Field
		TargetState        respjson.Field
		To                 respjson.Field
		Type               respjson.Field
		UpdatedAt          respjson.Field
		Value              respjson.Field
		LendBorrowMarketID respjson.Field
		VenueSymbol        respjson.Field
		ApprovalTarget     respjson.Field
		MaxSpendAmount     respjson.Field
		Quote              respjson.Field
		Slippage           respjson.Field
		Status             respjson.Field
		ContractAddress    respjson.Field
		Permission         respjson.Field
		MinReceiveAmount   respjson.Field
		raw                string
	} `json:"-"`
}

func (u ExecutionStepDefinitionUnion) AsExecutionEVMTransactionApprove() (v ExecutionStepDefinitionExecutionEVMTransactionApprove) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionStepDefinitionUnion) AsExecutionEVMTransactionBorrow() (v ExecutionStepDefinitionExecutionEVMTransactionBorrow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionStepDefinitionUnion) AsExecutionEVMTransactionBuy() (v ExecutionStepDefinitionExecutionEVMTransactionBuy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionStepDefinitionUnion) AsExecutionEVMTransactionLend() (v ExecutionStepDefinitionExecutionEVMTransactionLend) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionStepDefinitionUnion) AsExecutionEVMTransactionLendSetCollateral() (v ExecutionStepDefinitionExecutionEVMTransactionLendSetCollateral) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionStepDefinitionUnion) AsExecutionEVMTransactionLendWithdraw() (v ExecutionStepDefinitionExecutionEVMTransactionLendWithdraw) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionStepDefinitionUnion) AsExecutionEVMTransactionMove() (v ExecutionStepDefinitionExecutionEVMTransactionMove) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionStepDefinitionUnion) AsExecutionEVMTransactionPermission() (v ExecutionStepDefinitionExecutionEVMTransactionPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionStepDefinitionUnion) AsExecutionEVMTransactionUnwrap() (v ExecutionStepDefinitionExecutionEVMTransactionUnwrap) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionStepDefinitionUnion) AsExecutionEVMTransactionWrap() (v ExecutionStepDefinitionExecutionEVMTransactionWrap) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionStepDefinitionUnion) AsExecutionEVMTransactionSell() (v ExecutionStepDefinitionExecutionEVMTransactionSell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExecutionStepDefinitionUnion) RawJSON() string { return u.JSON.raw }

func (r *ExecutionStepDefinitionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExecutionStepDefinitionUnionError is an implicit subunion of
// [ExecutionStepDefinitionUnion]. ExecutionStepDefinitionUnionError provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ExecutionStepDefinitionUnion].
type ExecutionStepDefinitionUnionError struct {
	Message string `json:"message"`
	Params  any    `json:"params"`
	Type    string `json:"type"`
	JSON    struct {
		Message respjson.Field
		Params  respjson.Field
		Type    respjson.Field
		raw     string
	} `json:"-"`
}

func (r *ExecutionStepDefinitionUnionError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExecutionStepDefinitionUnionQuote is an implicit subunion of
// [ExecutionStepDefinitionUnion]. ExecutionStepDefinitionUnionQuote provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ExecutionStepDefinitionUnion].
type ExecutionStepDefinitionUnionQuote struct {
	// This field is from variant [BuyQuote].
	Blockchain Blockchain `json:"blockchain"`
	// This field is a union of [BuyQuoteFeeEstimate], [SellQuoteFeeEstimate]
	FeeEstimate      ExecutionStepDefinitionUnionQuoteFeeEstimate `json:"fee_estimate"`
	QuoteAssetSymbol string                                       `json:"quote_asset_symbol"`
	// This field is a union of [BuyQuoteQuoteInfoUnion], [SellQuoteQuoteInfoUnion]
	QuoteInfo  ExecutionStepDefinitionUnionQuoteQuoteInfo `json:"quote_info"`
	QuoteValue string                                     `json:"quote_value"`
	// This field is from variant [BuyQuote].
	ReceiveAmount Decimal `json:"receive_amount"`
	// This field is from variant [BuyQuote].
	ReceiveAsset Asset `json:"receive_asset"`
	// This field is from variant [BuyQuote].
	SpendAmount Decimal `json:"spend_amount"`
	// This field is from variant [BuyQuote].
	SpendAsset Asset `json:"spend_asset"`
	// This field is from variant [BuyQuote].
	Venue Venue `json:"venue"`
	JSON  struct {
		Blockchain       respjson.Field
		FeeEstimate      respjson.Field
		QuoteAssetSymbol respjson.Field
		QuoteInfo        respjson.Field
		QuoteValue       respjson.Field
		ReceiveAmount    respjson.Field
		ReceiveAsset     respjson.Field
		SpendAmount      respjson.Field
		SpendAsset       respjson.Field
		Venue            respjson.Field
		raw              string
	} `json:"-"`
}

func (r *ExecutionStepDefinitionUnionQuote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExecutionStepDefinitionUnionQuoteFeeEstimate is an implicit subunion of
// [ExecutionStepDefinitionUnion]. ExecutionStepDefinitionUnionQuoteFeeEstimate
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ExecutionStepDefinitionUnion].
type ExecutionStepDefinitionUnionQuoteFeeEstimate struct {
	// This field is from variant [BuyQuoteFeeEstimate].
	Amount Decimal `json:"amount"`
	// This field is from variant [BuyQuoteFeeEstimate].
	Asset            Asset  `json:"asset"`
	Cost             string `json:"cost"`
	QuoteAssetSymbol string `json:"quote_asset_symbol"`
	JSON             struct {
		Amount           respjson.Field
		Asset            respjson.Field
		Cost             respjson.Field
		QuoteAssetSymbol respjson.Field
		raw              string
	} `json:"-"`
}

func (r *ExecutionStepDefinitionUnionQuoteFeeEstimate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExecutionStepDefinitionUnionQuoteQuoteInfo is an implicit subunion of
// [ExecutionStepDefinitionUnion]. ExecutionStepDefinitionUnionQuoteQuoteInfo
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ExecutionStepDefinitionUnion].
type ExecutionStepDefinitionUnionQuoteQuoteInfo struct {
	PoolIDs []string `json:"pool_ids"`
	// This field is from variant [SellQuoteQuoteInfoUnion].
	IIndex int64 `json:"i_index"`
	// This field is from variant [SellQuoteQuoteInfoUnion].
	JIndex int64 `json:"j_index"`
	// This field is from variant [SellQuoteQuoteInfoUnion].
	PoolID string `json:"pool_id"`
	// This field is from variant [SellQuoteQuoteInfoUnion].
	SwapType string `json:"swap_type"`
	// This field is from variant [SellQuoteQuoteInfoUnion].
	EstimatedGasUsed string `json:"estimated_gas_used"`
	// This field is from variant [SellQuoteQuoteInfoUnion].
	Route QuoteInfo0xRoute `json:"route"`
	JSON  struct {
		PoolIDs          respjson.Field
		IIndex           respjson.Field
		JIndex           respjson.Field
		PoolID           respjson.Field
		SwapType         respjson.Field
		EstimatedGasUsed respjson.Field
		Route            respjson.Field
		raw              string
	} `json:"-"`
}

func (r *ExecutionStepDefinitionUnionQuoteQuoteInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An approval of an asset
type ExecutionStepDefinitionExecutionEVMTransactionApprove struct {
	// An arbitrary precision decimal represented as a string
	Amount Decimal `json:"amount,required"`
	// On-chain asset (aka token)
	Asset       Asset `json:"asset,required"`
	BlockNumber int64 `json:"block_number,required"`
	// ISO8601 Timestamp
	BroadcastedAt      string `json:"broadcasted_at,required"`
	ConfirmationTarget int64  `json:"confirmation_target,required"`
	// ISO8601 Timestamp
	ConfirmedAt string `json:"confirmed_at,required"`
	// ISO8601 Timestamp
	CreatedAt Timestamp `json:"created_at,required"`
	// A hex string starting with 0x
	Data string `json:"data,required"`
	// An arbitrary precision decimal represented as a string
	EffectiveGasPrice string `json:"effective_gas_price,required"`
	// An error
	Error ExecutionStepDefinitionExecutionEVMTransactionApproveError `json:"error,required"`
	// ISO8601 Timestamp
	ErroredAt string `json:"errored_at,required"`
	// An arbitrary precision decimal represented as a string
	GasUsed string `json:"gas_used,required"`
	// A transaction hash
	Hash string `json:"hash,required"`
	// The payload of an EIP-1559 transaction
	Payload ExecutionEVMTransactionEIP1559Payload `json:"payload,required"`
	// ISO8601 Timestamp
	SignedAt string `json:"signed_at,required"`
	// An EVM address
	Spender Account `json:"spender,required"`
	// The state of an EVM transaction
	//
	// Any of "not_started", "signature_required", "signed", "broadcasted",
	// "confirmed", "error".
	State ExecutionEVMTransactionState `json:"state,required"`
	// The state of an EVM transaction
	//
	// Any of "not_started", "signature_required", "signed", "broadcasted",
	// "confirmed", "error".
	TargetState ExecutionEVMTransactionState `json:"target_state,required"`
	// An EVM address
	To string `json:"to,required"`
	// The type of approval
	//
	// Any of "spend_erc20", "borrow_erc20", "spend_erc721", "spend_erc721_collection".
	Type string `json:"type,required"`
	// ISO8601 Timestamp
	UpdatedAt Timestamp `json:"updated_at,required"`
	// An arbitrary precision decimal represented as a string
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount             respjson.Field
		Asset              respjson.Field
		BlockNumber        respjson.Field
		BroadcastedAt      respjson.Field
		ConfirmationTarget respjson.Field
		ConfirmedAt        respjson.Field
		CreatedAt          respjson.Field
		Data               respjson.Field
		EffectiveGasPrice  respjson.Field
		Error              respjson.Field
		ErroredAt          respjson.Field
		GasUsed            respjson.Field
		Hash               respjson.Field
		Payload            respjson.Field
		SignedAt           respjson.Field
		Spender            respjson.Field
		State              respjson.Field
		TargetState        respjson.Field
		To                 respjson.Field
		Type               respjson.Field
		UpdatedAt          respjson.Field
		Value              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepDefinitionExecutionEVMTransactionApprove) RawJSON() string { return r.JSON.raw }
func (r *ExecutionStepDefinitionExecutionEVMTransactionApprove) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An error
type ExecutionStepDefinitionExecutionEVMTransactionApproveError struct {
	// Error message
	Message string `json:"message,required"`
	// Error parameters
	Params map[string]any `json:"params,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Params      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepDefinitionExecutionEVMTransactionApproveError) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepDefinitionExecutionEVMTransactionApproveError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Borrowing an asset
type ExecutionStepDefinitionExecutionEVMTransactionBorrow struct {
	// An arbitrary precision decimal represented as a string
	Amount Decimal `json:"amount,required"`
	// On-chain asset (aka token)
	Asset       Asset `json:"asset,required"`
	BlockNumber int64 `json:"block_number,required"`
	// ISO8601 Timestamp
	BroadcastedAt      string `json:"broadcasted_at,required"`
	ConfirmationTarget int64  `json:"confirmation_target,required"`
	// ISO8601 Timestamp
	ConfirmedAt string `json:"confirmed_at,required"`
	// ISO8601 Timestamp
	CreatedAt Timestamp `json:"created_at,required"`
	// A hex string starting with 0x
	Data string `json:"data,required"`
	// An arbitrary precision decimal represented as a string
	EffectiveGasPrice string `json:"effective_gas_price,required"`
	// An error
	Error ExecutionStepDefinitionExecutionEVMTransactionBorrowError `json:"error,required"`
	// ISO8601 Timestamp
	ErroredAt string `json:"errored_at,required"`
	// An arbitrary precision decimal represented as a string
	GasUsed string `json:"gas_used,required"`
	// A transaction hash
	Hash string `json:"hash,required"`
	// A lend/borrow market ID, represented as a TypeID with `lend_borrow_market`
	// prefix
	LendBorrowMarketID LendBorrowMarketID `json:"lend_borrow_market_id,required"`
	// The payload of an EIP-1559 transaction
	Payload ExecutionEVMTransactionEIP1559Payload `json:"payload,required"`
	// ISO8601 Timestamp
	SignedAt string `json:"signed_at,required"`
	// The state of an EVM transaction
	//
	// Any of "not_started", "signature_required", "signed", "broadcasted",
	// "confirmed", "error".
	State ExecutionEVMTransactionState `json:"state,required"`
	// The state of an EVM transaction
	//
	// Any of "not_started", "signature_required", "signed", "broadcasted",
	// "confirmed", "error".
	TargetState ExecutionEVMTransactionState `json:"target_state,required"`
	// An EVM address
	To string `json:"to,required"`
	// ISO8601 Timestamp
	UpdatedAt Timestamp `json:"updated_at,required"`
	// An arbitrary precision decimal represented as a string
	Value string `json:"value,required"`
	// A venue symbol
	VenueSymbol VenueSymbol `json:"venue_symbol,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount             respjson.Field
		Asset              respjson.Field
		BlockNumber        respjson.Field
		BroadcastedAt      respjson.Field
		ConfirmationTarget respjson.Field
		ConfirmedAt        respjson.Field
		CreatedAt          respjson.Field
		Data               respjson.Field
		EffectiveGasPrice  respjson.Field
		Error              respjson.Field
		ErroredAt          respjson.Field
		GasUsed            respjson.Field
		Hash               respjson.Field
		LendBorrowMarketID respjson.Field
		Payload            respjson.Field
		SignedAt           respjson.Field
		State              respjson.Field
		TargetState        respjson.Field
		To                 respjson.Field
		UpdatedAt          respjson.Field
		Value              respjson.Field
		VenueSymbol        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepDefinitionExecutionEVMTransactionBorrow) RawJSON() string { return r.JSON.raw }
func (r *ExecutionStepDefinitionExecutionEVMTransactionBorrow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An error
type ExecutionStepDefinitionExecutionEVMTransactionBorrowError struct {
	// Error message
	Message string `json:"message,required"`
	// Error parameters
	Params map[string]any `json:"params,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Params      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepDefinitionExecutionEVMTransactionBorrowError) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepDefinitionExecutionEVMTransactionBorrowError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Buying an asset with another asset
type ExecutionStepDefinitionExecutionEVMTransactionBuy struct {
	// An EVM address
	ApprovalTarget Account `json:"approval_target,required"`
	BlockNumber    int64   `json:"block_number,required"`
	// ISO8601 Timestamp
	BroadcastedAt      string `json:"broadcasted_at,required"`
	ConfirmationTarget int64  `json:"confirmation_target,required"`
	// ISO8601 Timestamp
	ConfirmedAt string `json:"confirmed_at,required"`
	// ISO8601 Timestamp
	CreatedAt Timestamp `json:"created_at,required"`
	// A hex string starting with 0x
	Data string `json:"data,required"`
	// An arbitrary precision decimal represented as a string
	EffectiveGasPrice string `json:"effective_gas_price,required"`
	// An error
	Error ExecutionStepDefinitionExecutionEVMTransactionBuyError `json:"error,required"`
	// ISO8601 Timestamp
	ErroredAt string `json:"errored_at,required"`
	// An arbitrary precision decimal represented as a string
	GasUsed string `json:"gas_used,required"`
	// A transaction hash
	Hash string `json:"hash,required"`
	// An arbitrary precision decimal represented as a string
	MaxSpendAmount Decimal `json:"max_spend_amount,required"`
	// The payload of an EIP-1559 transaction
	Payload ExecutionEVMTransactionEIP1559Payload `json:"payload,required"`
	// A buy quote
	Quote BuyQuote `json:"quote,required"`
	// ISO8601 Timestamp
	SignedAt string `json:"signed_at,required"`
	// An arbitrary precision decimal represented as a string
	Slippage Decimal `json:"slippage,required"`
	// The state of an EVM transaction
	//
	// Any of "not_started", "signature_required", "signed", "broadcasted",
	// "confirmed", "error".
	State ExecutionEVMTransactionState `json:"state,required"`
	// The state of an EVM transaction
	//
	// Any of "not_started", "signature_required", "signed", "broadcasted",
	// "confirmed", "error".
	TargetState ExecutionEVMTransactionState `json:"target_state,required"`
	// An EVM address
	To string `json:"to,required"`
	// ISO8601 Timestamp
	UpdatedAt Timestamp `json:"updated_at,required"`
	// An arbitrary precision decimal represented as a string
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApprovalTarget     respjson.Field
		BlockNumber        respjson.Field
		BroadcastedAt      respjson.Field
		ConfirmationTarget respjson.Field
		ConfirmedAt        respjson.Field
		CreatedAt          respjson.Field
		Data               respjson.Field
		EffectiveGasPrice  respjson.Field
		Error              respjson.Field
		ErroredAt          respjson.Field
		GasUsed            respjson.Field
		Hash               respjson.Field
		MaxSpendAmount     respjson.Field
		Payload            respjson.Field
		Quote              respjson.Field
		SignedAt           respjson.Field
		Slippage           respjson.Field
		State              respjson.Field
		TargetState        respjson.Field
		To                 respjson.Field
		UpdatedAt          respjson.Field
		Value              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepDefinitionExecutionEVMTransactionBuy) RawJSON() string { return r.JSON.raw }
func (r *ExecutionStepDefinitionExecutionEVMTransactionBuy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An error
type ExecutionStepDefinitionExecutionEVMTransactionBuyError struct {
	// Error message
	Message string `json:"message,required"`
	// Error parameters
	Params map[string]any `json:"params,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Params      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepDefinitionExecutionEVMTransactionBuyError) RawJSON() string { return r.JSON.raw }
func (r *ExecutionStepDefinitionExecutionEVMTransactionBuyError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lending an asset
type ExecutionStepDefinitionExecutionEVMTransactionLend struct {
	// An arbitrary precision decimal represented as a string
	Amount Decimal `json:"amount,required"`
	// On-chain asset (aka token)
	Asset       Asset `json:"asset,required"`
	BlockNumber int64 `json:"block_number,required"`
	// ISO8601 Timestamp
	BroadcastedAt      string `json:"broadcasted_at,required"`
	ConfirmationTarget int64  `json:"confirmation_target,required"`
	// ISO8601 Timestamp
	ConfirmedAt string `json:"confirmed_at,required"`
	// ISO8601 Timestamp
	CreatedAt Timestamp `json:"created_at,required"`
	// A hex string starting with 0x
	Data string `json:"data,required"`
	// An arbitrary precision decimal represented as a string
	EffectiveGasPrice string `json:"effective_gas_price,required"`
	// An error
	Error ExecutionStepDefinitionExecutionEVMTransactionLendError `json:"error,required"`
	// ISO8601 Timestamp
	ErroredAt string `json:"errored_at,required"`
	// An arbitrary precision decimal represented as a string
	GasUsed string `json:"gas_used,required"`
	// A transaction hash
	Hash string `json:"hash,required"`
	// A lend/borrow market ID, represented as a TypeID with `lend_borrow_market`
	// prefix
	LendBorrowMarketID LendBorrowMarketID `json:"lend_borrow_market_id,required"`
	// The payload of an EIP-1559 transaction
	Payload ExecutionEVMTransactionEIP1559Payload `json:"payload,required"`
	// ISO8601 Timestamp
	SignedAt string `json:"signed_at,required"`
	// The state of an EVM transaction
	//
	// Any of "not_started", "signature_required", "signed", "broadcasted",
	// "confirmed", "error".
	State ExecutionEVMTransactionState `json:"state,required"`
	// The state of an EVM transaction
	//
	// Any of "not_started", "signature_required", "signed", "broadcasted",
	// "confirmed", "error".
	TargetState ExecutionEVMTransactionState `json:"target_state,required"`
	// An EVM address
	To string `json:"to,required"`
	// ISO8601 Timestamp
	UpdatedAt Timestamp `json:"updated_at,required"`
	// An arbitrary precision decimal represented as a string
	Value string `json:"value,required"`
	// A venue symbol
	VenueSymbol VenueSymbol `json:"venue_symbol,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount             respjson.Field
		Asset              respjson.Field
		BlockNumber        respjson.Field
		BroadcastedAt      respjson.Field
		ConfirmationTarget respjson.Field
		ConfirmedAt        respjson.Field
		CreatedAt          respjson.Field
		Data               respjson.Field
		EffectiveGasPrice  respjson.Field
		Error              respjson.Field
		ErroredAt          respjson.Field
		GasUsed            respjson.Field
		Hash               respjson.Field
		LendBorrowMarketID respjson.Field
		Payload            respjson.Field
		SignedAt           respjson.Field
		State              respjson.Field
		TargetState        respjson.Field
		To                 respjson.Field
		UpdatedAt          respjson.Field
		Value              respjson.Field
		VenueSymbol        respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepDefinitionExecutionEVMTransactionLend) RawJSON() string { return r.JSON.raw }
func (r *ExecutionStepDefinitionExecutionEVMTransactionLend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An error
type ExecutionStepDefinitionExecutionEVMTransactionLendError struct {
	// Error message
	Message string `json:"message,required"`
	// Error parameters
	Params map[string]any `json:"params,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Params      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepDefinitionExecutionEVMTransactionLendError) RawJSON() string { return r.JSON.raw }
func (r *ExecutionStepDefinitionExecutionEVMTransactionLendError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Setting/unsetting a position as collateral
type ExecutionStepDefinitionExecutionEVMTransactionLendSetCollateral struct {
	BlockNumber int64 `json:"block_number,required"`
	// ISO8601 Timestamp
	BroadcastedAt      string `json:"broadcasted_at,required"`
	ConfirmationTarget int64  `json:"confirmation_target,required"`
	// ISO8601 Timestamp
	ConfirmedAt string `json:"confirmed_at,required"`
	// ISO8601 Timestamp
	CreatedAt Timestamp `json:"created_at,required"`
	// A hex string starting with 0x
	Data string `json:"data,required"`
	// An arbitrary precision decimal represented as a string
	EffectiveGasPrice string `json:"effective_gas_price,required"`
	// An error
	Error ExecutionStepDefinitionExecutionEVMTransactionLendSetCollateralError `json:"error,required"`
	// ISO8601 Timestamp
	ErroredAt string `json:"errored_at,required"`
	// An arbitrary precision decimal represented as a string
	GasUsed string `json:"gas_used,required"`
	// A transaction hash
	Hash string `json:"hash,required"`
	// A lend/borrow market ID, represented as a TypeID with `lend_borrow_market`
	// prefix
	LendBorrowMarketID LendBorrowMarketID `json:"lend_borrow_market_id,required"`
	// The payload of an EIP-1559 transaction
	Payload ExecutionEVMTransactionEIP1559Payload `json:"payload,required"`
	// ISO8601 Timestamp
	SignedAt string `json:"signed_at,required"`
	// The state of an EVM transaction
	//
	// Any of "not_started", "signature_required", "signed", "broadcasted",
	// "confirmed", "error".
	State  ExecutionEVMTransactionState `json:"state,required"`
	Status bool                         `json:"status,required"`
	// The state of an EVM transaction
	//
	// Any of "not_started", "signature_required", "signed", "broadcasted",
	// "confirmed", "error".
	TargetState ExecutionEVMTransactionState `json:"target_state,required"`
	// An EVM address
	To string `json:"to,required"`
	// ISO8601 Timestamp
	UpdatedAt Timestamp `json:"updated_at,required"`
	// An arbitrary precision decimal represented as a string
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlockNumber        respjson.Field
		BroadcastedAt      respjson.Field
		ConfirmationTarget respjson.Field
		ConfirmedAt        respjson.Field
		CreatedAt          respjson.Field
		Data               respjson.Field
		EffectiveGasPrice  respjson.Field
		Error              respjson.Field
		ErroredAt          respjson.Field
		GasUsed            respjson.Field
		Hash               respjson.Field
		LendBorrowMarketID respjson.Field
		Payload            respjson.Field
		SignedAt           respjson.Field
		State              respjson.Field
		Status             respjson.Field
		TargetState        respjson.Field
		To                 respjson.Field
		UpdatedAt          respjson.Field
		Value              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepDefinitionExecutionEVMTransactionLendSetCollateral) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepDefinitionExecutionEVMTransactionLendSetCollateral) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An error
type ExecutionStepDefinitionExecutionEVMTransactionLendSetCollateralError struct {
	// Error message
	Message string `json:"message,required"`
	// Error parameters
	Params map[string]any `json:"params,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Params      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepDefinitionExecutionEVMTransactionLendSetCollateralError) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepDefinitionExecutionEVMTransactionLendSetCollateralError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Withdrawing an asset
type ExecutionStepDefinitionExecutionEVMTransactionLendWithdraw struct {
	// An arbitrary precision decimal represented as a string
	Amount string `json:"amount,required"`
	// On-chain asset (aka token)
	Asset       Asset `json:"asset,required"`
	BlockNumber int64 `json:"block_number,required"`
	// ISO8601 Timestamp
	BroadcastedAt      string `json:"broadcasted_at,required"`
	ConfirmationTarget int64  `json:"confirmation_target,required"`
	// ISO8601 Timestamp
	ConfirmedAt string `json:"confirmed_at,required"`
	// ISO8601 Timestamp
	CreatedAt Timestamp `json:"created_at,required"`
	// A hex string starting with 0x
	Data string `json:"data,required"`
	// An arbitrary precision decimal represented as a string
	EffectiveGasPrice string `json:"effective_gas_price,required"`
	// An error
	Error ExecutionStepDefinitionExecutionEVMTransactionLendWithdrawError `json:"error,required"`
	// ISO8601 Timestamp
	ErroredAt string `json:"errored_at,required"`
	// An arbitrary precision decimal represented as a string
	GasUsed string `json:"gas_used,required"`
	// A transaction hash
	Hash string `json:"hash,required"`
	// A lend/borrow market ID, represented as a TypeID with `lend_borrow_market`
	// prefix
	LendBorrowMarketID LendBorrowMarketID `json:"lend_borrow_market_id,required"`
	// The payload of an EIP-1559 transaction
	Payload ExecutionEVMTransactionEIP1559Payload `json:"payload,required"`
	// ISO8601 Timestamp
	SignedAt string `json:"signed_at,required"`
	// The state of an EVM transaction
	//
	// Any of "not_started", "signature_required", "signed", "broadcasted",
	// "confirmed", "error".
	State ExecutionEVMTransactionState `json:"state,required"`
	// The state of an EVM transaction
	//
	// Any of "not_started", "signature_required", "signed", "broadcasted",
	// "confirmed", "error".
	TargetState ExecutionEVMTransactionState `json:"target_state,required"`
	// An EVM address
	To string `json:"to,required"`
	// ISO8601 Timestamp
	UpdatedAt Timestamp `json:"updated_at,required"`
	// An arbitrary precision decimal represented as a string
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount             respjson.Field
		Asset              respjson.Field
		BlockNumber        respjson.Field
		BroadcastedAt      respjson.Field
		ConfirmationTarget respjson.Field
		ConfirmedAt        respjson.Field
		CreatedAt          respjson.Field
		Data               respjson.Field
		EffectiveGasPrice  respjson.Field
		Error              respjson.Field
		ErroredAt          respjson.Field
		GasUsed            respjson.Field
		Hash               respjson.Field
		LendBorrowMarketID respjson.Field
		Payload            respjson.Field
		SignedAt           respjson.Field
		State              respjson.Field
		TargetState        respjson.Field
		To                 respjson.Field
		UpdatedAt          respjson.Field
		Value              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepDefinitionExecutionEVMTransactionLendWithdraw) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepDefinitionExecutionEVMTransactionLendWithdraw) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An error
type ExecutionStepDefinitionExecutionEVMTransactionLendWithdrawError struct {
	// Error message
	Message string `json:"message,required"`
	// Error parameters
	Params map[string]any `json:"params,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Params      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepDefinitionExecutionEVMTransactionLendWithdrawError) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepDefinitionExecutionEVMTransactionLendWithdrawError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A move of assets from one account to another
type ExecutionStepDefinitionExecutionEVMTransactionMove struct {
	// An arbitrary precision decimal represented as a string
	Amount Decimal `json:"amount,required"`
	// On-chain asset (aka token)
	Asset       Asset `json:"asset,required"`
	BlockNumber int64 `json:"block_number,required"`
	// ISO8601 Timestamp
	BroadcastedAt      string `json:"broadcasted_at,required"`
	ConfirmationTarget int64  `json:"confirmation_target,required"`
	// ISO8601 Timestamp
	ConfirmedAt string `json:"confirmed_at,required"`
	// ISO8601 Timestamp
	CreatedAt Timestamp `json:"created_at,required"`
	// A hex string starting with 0x
	Data string `json:"data,required"`
	// An arbitrary precision decimal represented as a string
	EffectiveGasPrice string `json:"effective_gas_price,required"`
	// An error
	Error ExecutionStepDefinitionExecutionEVMTransactionMoveError `json:"error,required"`
	// ISO8601 Timestamp
	ErroredAt string `json:"errored_at,required"`
	// An arbitrary precision decimal represented as a string
	GasUsed string `json:"gas_used,required"`
	// A transaction hash
	Hash string `json:"hash,required"`
	// The payload of an EIP-1559 transaction
	Payload ExecutionEVMTransactionEIP1559Payload `json:"payload,required"`
	// ISO8601 Timestamp
	SignedAt string `json:"signed_at,required"`
	// The state of an EVM transaction
	//
	// Any of "not_started", "signature_required", "signed", "broadcasted",
	// "confirmed", "error".
	State ExecutionEVMTransactionState `json:"state,required"`
	// The state of an EVM transaction
	//
	// Any of "not_started", "signature_required", "signed", "broadcasted",
	// "confirmed", "error".
	TargetState ExecutionEVMTransactionState `json:"target_state,required"`
	// An EVM address
	To Account `json:"to,required"`
	// ISO8601 Timestamp
	UpdatedAt Timestamp `json:"updated_at,required"`
	// An arbitrary precision decimal represented as a string
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount             respjson.Field
		Asset              respjson.Field
		BlockNumber        respjson.Field
		BroadcastedAt      respjson.Field
		ConfirmationTarget respjson.Field
		ConfirmedAt        respjson.Field
		CreatedAt          respjson.Field
		Data               respjson.Field
		EffectiveGasPrice  respjson.Field
		Error              respjson.Field
		ErroredAt          respjson.Field
		GasUsed            respjson.Field
		Hash               respjson.Field
		Payload            respjson.Field
		SignedAt           respjson.Field
		State              respjson.Field
		TargetState        respjson.Field
		To                 respjson.Field
		UpdatedAt          respjson.Field
		Value              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepDefinitionExecutionEVMTransactionMove) RawJSON() string { return r.JSON.raw }
func (r *ExecutionStepDefinitionExecutionEVMTransactionMove) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An error
type ExecutionStepDefinitionExecutionEVMTransactionMoveError struct {
	// Error message
	Message string `json:"message,required"`
	// Error parameters
	Params map[string]any `json:"params,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Params      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepDefinitionExecutionEVMTransactionMoveError) RawJSON() string { return r.JSON.raw }
func (r *ExecutionStepDefinitionExecutionEVMTransactionMoveError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A permission to a contract
type ExecutionStepDefinitionExecutionEVMTransactionPermission struct {
	BlockNumber int64 `json:"block_number,required"`
	// ISO8601 Timestamp
	BroadcastedAt      string `json:"broadcasted_at,required"`
	ConfirmationTarget int64  `json:"confirmation_target,required"`
	// ISO8601 Timestamp
	ConfirmedAt string `json:"confirmed_at,required"`
	// An EVM address
	ContractAddress AddressEVM `json:"contract_address,required"`
	// ISO8601 Timestamp
	CreatedAt Timestamp `json:"created_at,required"`
	// A hex string starting with 0x
	Data string `json:"data,required"`
	// An arbitrary precision decimal represented as a string
	EffectiveGasPrice string `json:"effective_gas_price,required"`
	// An error
	Error ExecutionStepDefinitionExecutionEVMTransactionPermissionError `json:"error,required"`
	// ISO8601 Timestamp
	ErroredAt string `json:"errored_at,required"`
	// An arbitrary precision decimal represented as a string
	GasUsed string `json:"gas_used,required"`
	// A transaction hash
	Hash string `json:"hash,required"`
	// A lend/borrow market ID, represented as a TypeID with `lend_borrow_market`
	// prefix
	LendBorrowMarketID LendBorrowMarketID `json:"lend_borrow_market_id,required"`
	// The payload of an EIP-1559 transaction
	Payload    ExecutionEVMTransactionEIP1559Payload `json:"payload,required"`
	Permission bool                                  `json:"permission,required"`
	// ISO8601 Timestamp
	SignedAt string `json:"signed_at,required"`
	// An EVM address
	Spender Account `json:"spender,required"`
	// The state of an EVM transaction
	//
	// Any of "not_started", "signature_required", "signed", "broadcasted",
	// "confirmed", "error".
	State ExecutionEVMTransactionState `json:"state,required"`
	// The state of an EVM transaction
	//
	// Any of "not_started", "signature_required", "signed", "broadcasted",
	// "confirmed", "error".
	TargetState ExecutionEVMTransactionState `json:"target_state,required"`
	// An EVM address
	To string `json:"to,required"`
	// The type of a permission
	//
	// Any of "compound_v3_comet".
	Type string `json:"type,required"`
	// ISO8601 Timestamp
	UpdatedAt Timestamp `json:"updated_at,required"`
	// An arbitrary precision decimal represented as a string
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlockNumber        respjson.Field
		BroadcastedAt      respjson.Field
		ConfirmationTarget respjson.Field
		ConfirmedAt        respjson.Field
		ContractAddress    respjson.Field
		CreatedAt          respjson.Field
		Data               respjson.Field
		EffectiveGasPrice  respjson.Field
		Error              respjson.Field
		ErroredAt          respjson.Field
		GasUsed            respjson.Field
		Hash               respjson.Field
		LendBorrowMarketID respjson.Field
		Payload            respjson.Field
		Permission         respjson.Field
		SignedAt           respjson.Field
		Spender            respjson.Field
		State              respjson.Field
		TargetState        respjson.Field
		To                 respjson.Field
		Type               respjson.Field
		UpdatedAt          respjson.Field
		Value              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepDefinitionExecutionEVMTransactionPermission) RawJSON() string { return r.JSON.raw }
func (r *ExecutionStepDefinitionExecutionEVMTransactionPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An error
type ExecutionStepDefinitionExecutionEVMTransactionPermissionError struct {
	// Error message
	Message string `json:"message,required"`
	// Error parameters
	Params map[string]any `json:"params,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Params      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepDefinitionExecutionEVMTransactionPermissionError) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepDefinitionExecutionEVMTransactionPermissionError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An unwrap of the wrapped native asset
type ExecutionStepDefinitionExecutionEVMTransactionUnwrap struct {
	// An arbitrary precision decimal represented as a string
	Amount      Decimal `json:"amount,required"`
	BlockNumber int64   `json:"block_number,required"`
	// ISO8601 Timestamp
	BroadcastedAt      string `json:"broadcasted_at,required"`
	ConfirmationTarget int64  `json:"confirmation_target,required"`
	// ISO8601 Timestamp
	ConfirmedAt string `json:"confirmed_at,required"`
	// ISO8601 Timestamp
	CreatedAt Timestamp `json:"created_at,required"`
	// A hex string starting with 0x
	Data string `json:"data,required"`
	// An arbitrary precision decimal represented as a string
	EffectiveGasPrice string `json:"effective_gas_price,required"`
	// An error
	Error ExecutionStepDefinitionExecutionEVMTransactionUnwrapError `json:"error,required"`
	// ISO8601 Timestamp
	ErroredAt string `json:"errored_at,required"`
	// An arbitrary precision decimal represented as a string
	GasUsed string `json:"gas_used,required"`
	// A transaction hash
	Hash string `json:"hash,required"`
	// The payload of an EIP-1559 transaction
	Payload ExecutionEVMTransactionEIP1559Payload `json:"payload,required"`
	// ISO8601 Timestamp
	SignedAt string `json:"signed_at,required"`
	// The state of an EVM transaction
	//
	// Any of "not_started", "signature_required", "signed", "broadcasted",
	// "confirmed", "error".
	State ExecutionEVMTransactionState `json:"state,required"`
	// The state of an EVM transaction
	//
	// Any of "not_started", "signature_required", "signed", "broadcasted",
	// "confirmed", "error".
	TargetState ExecutionEVMTransactionState `json:"target_state,required"`
	// An EVM address
	To string `json:"to,required"`
	// ISO8601 Timestamp
	UpdatedAt Timestamp `json:"updated_at,required"`
	// An arbitrary precision decimal represented as a string
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount             respjson.Field
		BlockNumber        respjson.Field
		BroadcastedAt      respjson.Field
		ConfirmationTarget respjson.Field
		ConfirmedAt        respjson.Field
		CreatedAt          respjson.Field
		Data               respjson.Field
		EffectiveGasPrice  respjson.Field
		Error              respjson.Field
		ErroredAt          respjson.Field
		GasUsed            respjson.Field
		Hash               respjson.Field
		Payload            respjson.Field
		SignedAt           respjson.Field
		State              respjson.Field
		TargetState        respjson.Field
		To                 respjson.Field
		UpdatedAt          respjson.Field
		Value              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepDefinitionExecutionEVMTransactionUnwrap) RawJSON() string { return r.JSON.raw }
func (r *ExecutionStepDefinitionExecutionEVMTransactionUnwrap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An error
type ExecutionStepDefinitionExecutionEVMTransactionUnwrapError struct {
	// Error message
	Message string `json:"message,required"`
	// Error parameters
	Params map[string]any `json:"params,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Params      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepDefinitionExecutionEVMTransactionUnwrapError) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepDefinitionExecutionEVMTransactionUnwrapError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A wrap of the native asset
type ExecutionStepDefinitionExecutionEVMTransactionWrap struct {
	// An arbitrary precision decimal represented as a string
	Amount      Decimal `json:"amount,required"`
	BlockNumber int64   `json:"block_number,required"`
	// ISO8601 Timestamp
	BroadcastedAt      string `json:"broadcasted_at,required"`
	ConfirmationTarget int64  `json:"confirmation_target,required"`
	// ISO8601 Timestamp
	ConfirmedAt string `json:"confirmed_at,required"`
	// ISO8601 Timestamp
	CreatedAt Timestamp `json:"created_at,required"`
	// A hex string starting with 0x
	Data string `json:"data,required"`
	// An arbitrary precision decimal represented as a string
	EffectiveGasPrice string `json:"effective_gas_price,required"`
	// An error
	Error ExecutionStepDefinitionExecutionEVMTransactionWrapError `json:"error,required"`
	// ISO8601 Timestamp
	ErroredAt string `json:"errored_at,required"`
	// An arbitrary precision decimal represented as a string
	GasUsed string `json:"gas_used,required"`
	// A transaction hash
	Hash string `json:"hash,required"`
	// The payload of an EIP-1559 transaction
	Payload ExecutionEVMTransactionEIP1559Payload `json:"payload,required"`
	// ISO8601 Timestamp
	SignedAt string `json:"signed_at,required"`
	// The state of an EVM transaction
	//
	// Any of "not_started", "signature_required", "signed", "broadcasted",
	// "confirmed", "error".
	State ExecutionEVMTransactionState `json:"state,required"`
	// The state of an EVM transaction
	//
	// Any of "not_started", "signature_required", "signed", "broadcasted",
	// "confirmed", "error".
	TargetState ExecutionEVMTransactionState `json:"target_state,required"`
	// An EVM address
	To string `json:"to,required"`
	// ISO8601 Timestamp
	UpdatedAt Timestamp `json:"updated_at,required"`
	// An arbitrary precision decimal represented as a string
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount             respjson.Field
		BlockNumber        respjson.Field
		BroadcastedAt      respjson.Field
		ConfirmationTarget respjson.Field
		ConfirmedAt        respjson.Field
		CreatedAt          respjson.Field
		Data               respjson.Field
		EffectiveGasPrice  respjson.Field
		Error              respjson.Field
		ErroredAt          respjson.Field
		GasUsed            respjson.Field
		Hash               respjson.Field
		Payload            respjson.Field
		SignedAt           respjson.Field
		State              respjson.Field
		TargetState        respjson.Field
		To                 respjson.Field
		UpdatedAt          respjson.Field
		Value              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepDefinitionExecutionEVMTransactionWrap) RawJSON() string { return r.JSON.raw }
func (r *ExecutionStepDefinitionExecutionEVMTransactionWrap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An error
type ExecutionStepDefinitionExecutionEVMTransactionWrapError struct {
	// Error message
	Message string `json:"message,required"`
	// Error parameters
	Params map[string]any `json:"params,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Params      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepDefinitionExecutionEVMTransactionWrapError) RawJSON() string { return r.JSON.raw }
func (r *ExecutionStepDefinitionExecutionEVMTransactionWrapError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Selling an asset for another asset
type ExecutionStepDefinitionExecutionEVMTransactionSell struct {
	// An EVM address
	ApprovalTarget Account `json:"approval_target,required"`
	BlockNumber    int64   `json:"block_number,required"`
	// ISO8601 Timestamp
	BroadcastedAt      string `json:"broadcasted_at,required"`
	ConfirmationTarget int64  `json:"confirmation_target,required"`
	// ISO8601 Timestamp
	ConfirmedAt string `json:"confirmed_at,required"`
	// ISO8601 Timestamp
	CreatedAt Timestamp `json:"created_at,required"`
	// A hex string starting with 0x
	Data string `json:"data,required"`
	// An arbitrary precision decimal represented as a string
	EffectiveGasPrice string `json:"effective_gas_price,required"`
	// An error
	Error ExecutionStepDefinitionExecutionEVMTransactionSellError `json:"error,required"`
	// ISO8601 Timestamp
	ErroredAt string `json:"errored_at,required"`
	// An arbitrary precision decimal represented as a string
	GasUsed string `json:"gas_used,required"`
	// A transaction hash
	Hash string `json:"hash,required"`
	// An arbitrary precision decimal represented as a string
	MinReceiveAmount Decimal `json:"min_receive_amount,required"`
	// The payload of an EIP-1559 transaction
	Payload ExecutionEVMTransactionEIP1559Payload `json:"payload,required"`
	// A sell quote
	Quote SellQuote `json:"quote,required"`
	// ISO8601 Timestamp
	SignedAt string `json:"signed_at,required"`
	// An arbitrary precision decimal represented as a string
	Slippage Decimal `json:"slippage,required"`
	// The state of an EVM transaction
	//
	// Any of "not_started", "signature_required", "signed", "broadcasted",
	// "confirmed", "error".
	State ExecutionEVMTransactionState `json:"state,required"`
	// The state of an EVM transaction
	//
	// Any of "not_started", "signature_required", "signed", "broadcasted",
	// "confirmed", "error".
	TargetState ExecutionEVMTransactionState `json:"target_state,required"`
	// An EVM address
	To string `json:"to,required"`
	// ISO8601 Timestamp
	UpdatedAt Timestamp `json:"updated_at,required"`
	// An arbitrary precision decimal represented as a string
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApprovalTarget     respjson.Field
		BlockNumber        respjson.Field
		BroadcastedAt      respjson.Field
		ConfirmationTarget respjson.Field
		ConfirmedAt        respjson.Field
		CreatedAt          respjson.Field
		Data               respjson.Field
		EffectiveGasPrice  respjson.Field
		Error              respjson.Field
		ErroredAt          respjson.Field
		GasUsed            respjson.Field
		Hash               respjson.Field
		MinReceiveAmount   respjson.Field
		Payload            respjson.Field
		Quote              respjson.Field
		SignedAt           respjson.Field
		Slippage           respjson.Field
		State              respjson.Field
		TargetState        respjson.Field
		To                 respjson.Field
		UpdatedAt          respjson.Field
		Value              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepDefinitionExecutionEVMTransactionSell) RawJSON() string { return r.JSON.raw }
func (r *ExecutionStepDefinitionExecutionEVMTransactionSell) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An error
type ExecutionStepDefinitionExecutionEVMTransactionSellError struct {
	// Error message
	Message string `json:"message,required"`
	// Error parameters
	Params map[string]any `json:"params,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Params      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepDefinitionExecutionEVMTransactionSellError) RawJSON() string { return r.JSON.raw }
func (r *ExecutionStepDefinitionExecutionEVMTransactionSellError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The payload of an EIP-1559 transaction
type ExecutionEVMTransactionEIP1559Payload struct {
	// A hex string starting with 0x
	Data HexString `json:"data,required"`
	// An EVM address
	From AddressEVM `json:"from,required"`
	// A hex string starting with 0x
	GasLimit HexString `json:"gasLimit,required"`
	// A hex string starting with 0x
	MaxFeePerGas HexString `json:"maxFeePerGas,required"`
	// A hex string starting with 0x
	MaxPriorityFeePerGas HexString `json:"maxPriorityFeePerGas,required"`
	// A hex string starting with 0x
	Nonce HexString `json:"nonce,required"`
	// An EVM address
	To AddressEVM `json:"to,required"`
	// A hex string starting with 0x
	Value HexString `json:"value,required"`
	// A hex string starting with 0x
	ChainID HexString `json:"chainId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data                 respjson.Field
		From                 respjson.Field
		GasLimit             respjson.Field
		MaxFeePerGas         respjson.Field
		MaxPriorityFeePerGas respjson.Field
		Nonce                respjson.Field
		To                   respjson.Field
		Value                respjson.Field
		ChainID              respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionEVMTransactionEIP1559Payload) RawJSON() string { return r.JSON.raw }
func (r *ExecutionEVMTransactionEIP1559Payload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The state of an EVM transaction
type ExecutionEVMTransactionState string

const (
	ExecutionEVMTransactionStateNotStarted        ExecutionEVMTransactionState = "not_started"
	ExecutionEVMTransactionStateSignatureRequired ExecutionEVMTransactionState = "signature_required"
	ExecutionEVMTransactionStateSigned            ExecutionEVMTransactionState = "signed"
	ExecutionEVMTransactionStateBroadcasted       ExecutionEVMTransactionState = "broadcasted"
	ExecutionEVMTransactionStateConfirmed         ExecutionEVMTransactionState = "confirmed"
	ExecutionEVMTransactionStateError             ExecutionEVMTransactionState = "error"
)

type ExecutionID = string

// The state of an execution
type ExecutionState string

const (
	ExecutionStateNotStarted ExecutionState = "not_started"
	ExecutionStateStarted    ExecutionState = "started"
	ExecutionStateSuccess    ExecutionState = "success"
	ExecutionStateAborted    ExecutionState = "aborted"
	ExecutionStateError      ExecutionState = "error"
)

type ExecutionStepID = string

type HexString = string

// A lp pool incentive market
type IncentivizeMarket struct {
	// A list of arbitrary precision decimals
	Amounts []Decimal `json:"amounts,required"`
	// Response for multiple assets
	Assets []Asset `json:"assets,required"`
	// Information about a specific block number and its timestamp
	Blockstamp Blockstamp `json:"blockstamp,required"`
	// ISO8601 Timestamp
	ClosedAt Timestamp `json:"closed_at,required"`
	// A liquidity pool
	Pool LPPool `json:"pool,required"`
	// An asset symbol
	QuoteAssetSymbol string `json:"quote_asset_symbol,required"`
	// The round number
	Round int64 `json:"round,required"`
	// An arbitrary precision decimal represented as a string
	Value string `json:"value,required"`
	// On-chain venue
	Venue Venue `json:"venue,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amounts          respjson.Field
		Assets           respjson.Field
		Blockstamp       respjson.Field
		ClosedAt         respjson.Field
		Pool             respjson.Field
		QuoteAssetSymbol respjson.Field
		Round            respjson.Field
		Value            respjson.Field
		Venue            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IncentivizeMarket) RawJSON() string { return r.JSON.raw }
func (r *IncentivizeMarket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// APY Data for lend/borrow markets
type LendBorrowAPYs struct {
	// APY Data
	Base APY `json:"base,required"`
	// Array of APY
	Rewards []APY `json:"rewards,required"`
	// An arbitrary precision decimal represented as a string
	Total Decimal `json:"total,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Base        respjson.Field
		Rewards     respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LendBorrowAPYs) RawJSON() string { return r.JSON.raw }
func (r *LendBorrowAPYs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LendBorrowMarketID = string

// A lend market
type LendMarket struct {
	// A lend/borrow market ID, represented as a TypeID with `lend_borrow_market`
	// prefix
	ID LendBorrowMarketID `json:"id,required"`
	// APY Data for lend/borrow markets
	APYs LendBorrowAPYs `json:"apys,required"`
	// On-chain asset (aka token)
	Asset Asset `json:"asset,required"`
	// Information about a specific block number and its timestamp
	Blockstamp Blockstamp `json:"blockstamp,required"`
	// An arbitrary precision decimal represented as a string
	Liquidity Decimal `json:"liquidity,required"`
	// An asset symbol
	QuoteAssetSymbol string `json:"quote_asset_symbol,required"`
	// An arbitrary precision decimal represented as a string
	TotalSupply Decimal `json:"total_supply,required"`
	// On-chain venue
	Venue Venue `json:"venue,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		APYs             respjson.Field
		Asset            respjson.Field
		Blockstamp       respjson.Field
		Liquidity        respjson.Field
		QuoteAssetSymbol respjson.Field
		TotalSupply      respjson.Field
		Venue            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LendMarket) RawJSON() string { return r.JSON.raw }
func (r *LendMarket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A lend position
type LendPosition struct {
	// An EVM address
	Account Account `json:"account,required"`
	// Information about a specific block number and its timestamp
	Blockstamp        Blockstamp `json:"blockstamp,required"`
	CollateralEnabled bool       `json:"collateral_enabled,required"`
	// A lend market
	Market LendMarket `json:"market,required"`
	// An asset symbol
	QuoteAssetSymbol string `json:"quote_asset_symbol,required"`
	// An arbitrary precision decimal represented as a string
	SuppliedAmount Decimal `json:"supplied_amount,required"`
	// An arbitrary precision decimal represented as a string
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account           respjson.Field
		Blockstamp        respjson.Field
		CollateralEnabled respjson.Field
		Market            respjson.Field
		QuoteAssetSymbol  respjson.Field
		SuppliedAmount    respjson.Field
		Value             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LendPosition) RawJSON() string { return r.JSON.raw }
func (r *LendPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A lock market
type LockMarket struct {
	// On-chain asset (aka token)
	Asset Asset `json:"asset,required"`
	// Data about a blockchain
	Blockchain Blockchain `json:"blockchain,required"`
	// On-chain venue
	Venue Venue `json:"venue,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asset       respjson.Field
		Blockchain  respjson.Field
		Venue       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LockMarket) RawJSON() string { return r.JSON.raw }
func (r *LockMarket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A lock position
type LockPosition struct {
	// An EVM address
	Account Account `json:"account,required"`
	// Information about a specific block number and its timestamp
	Blockstamp Blockstamp `json:"blockstamp,required"`
	// An arbitrary precision decimal represented as a string
	LockedAmount Decimal `json:"locked_amount,required"`
	// A NFT
	LockedAsset NFT `json:"locked_asset,required"`
	// A lock market
	Market LockMarket `json:"market,required"`
	// ISO8601 Timestamp
	UnlockedAt Timestamp `json:"unlocked_at,required"`
	Used       bool      `json:"used,required"`
	// An arbitrary precision decimal represented as a string
	VotingPower Decimal `json:"voting_power,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account      respjson.Field
		Blockstamp   respjson.Field
		LockedAmount respjson.Field
		LockedAsset  respjson.Field
		Market       respjson.Field
		UnlockedAt   respjson.Field
		Used         respjson.Field
		VotingPower  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LockPosition) RawJSON() string { return r.JSON.raw }
func (r *LockPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A liquidity pool
type LPPool struct {
	// A LP pool ID, represented as a TypeID with `lp_pool` prefix
	ID string `json:"id,required"`
	// Response for multiple assets
	Assets []Asset `json:"assets,required"`
	// Pool attributes that are specific to particular venue types
	Attributes LPPoolAttributesUnion `json:"attributes,required"`
	// Information about a specific block number and its timestamp
	Blockstamp Blockstamp `json:"blockstamp,required"`
	// An EVM address
	PoolAddress Account `json:"pool_address,required"`
	// An arbitrary precision decimal represented as a string
	PoolFee Decimal `json:"pool_fee,required"`
	// Type of a liquidity pool
	//
	// Any of "pair", "multi", "weighted", "range".
	PoolType LPPoolPoolType `json:"pool_type,required"`
	// An asset symbol
	QuoteAssetSymbol string `json:"quote_asset_symbol,required"`
	// A list of arbitrary precision decimals
	ReserveAmounts []Decimal `json:"reserve_amounts,required"`
	// A list of arbitrary precision decimals
	ReserveRatios []Decimal `json:"reserve_ratios,required"`
	// An arbitrary precision decimal represented as a string
	TotalLiquidity string `json:"total_liquidity,required"`
	// On-chain venue
	Venue Venue `json:"venue,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Assets           respjson.Field
		Attributes       respjson.Field
		Blockstamp       respjson.Field
		PoolAddress      respjson.Field
		PoolFee          respjson.Field
		PoolType         respjson.Field
		QuoteAssetSymbol respjson.Field
		ReserveAmounts   respjson.Field
		ReserveRatios    respjson.Field
		TotalLiquidity   respjson.Field
		Venue            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LPPool) RawJSON() string { return r.JSON.raw }
func (r *LPPool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LPPoolAttributesUnion contains all possible properties and values from
// [LPPoolUniswapV3Attributes], [LPPoolSolidlyAttributes].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type LPPoolAttributesUnion struct {
	// This field is from variant [LPPoolUniswapV3Attributes].
	Fee string `json:"fee"`
	// This field is from variant [LPPoolSolidlyAttributes].
	Stable bool `json:"stable"`
	JSON   struct {
		Fee    respjson.Field
		Stable respjson.Field
		raw    string
	} `json:"-"`
}

func (u LPPoolAttributesUnion) AsLPPoolUniswapV3Attributes() (v LPPoolUniswapV3Attributes) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LPPoolAttributesUnion) AsLPPoolSolidlyAttributes() (v LPPoolSolidlyAttributes) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LPPoolAttributesUnion) RawJSON() string { return u.JSON.raw }

func (r *LPPoolAttributesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of a liquidity pool
type LPPoolPoolType string

const (
	LPPoolPoolTypePair     LPPoolPoolType = "pair"
	LPPoolPoolTypeMulti    LPPoolPoolType = "multi"
	LPPoolPoolTypeWeighted LPPoolPoolType = "weighted"
	LPPoolPoolTypeRange    LPPoolPoolType = "range"
)

// Attributes for Solidly venue type pools
type LPPoolSolidlyAttributes struct {
	// True if pool is stable, false if volatile
	Stable bool `json:"stable,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Stable      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LPPoolSolidlyAttributes) RawJSON() string { return r.JSON.raw }
func (r *LPPoolSolidlyAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attributes for Uniswap V3 venue type pools
type LPPoolUniswapV3Attributes struct {
	// Pool fee
	Fee string `json:"fee,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fee         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LPPoolUniswapV3Attributes) RawJSON() string { return r.JSON.raw }
func (r *LPPoolUniswapV3Attributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A position in a liquidity pool
type LPPosition struct {
	// An EVM address
	Account Account `json:"account,required"`
	// A list of arbitrary precision decimals
	Amounts []Decimal `json:"amounts,required"`
	// Attributes for Uniswap V3 venue type positions
	Attributes LPPositionAttributes `json:"attributes,required"`
	// Information about a specific block number and its timestamp
	Blockstamp Blockstamp `json:"blockstamp,required"`
	// An arbitrary precision decimal represented as a string
	LPAmount Decimal `json:"lp_amount,required"`
	// An LP asset, either an Asset or a NFT
	LPAsset LPPositionLPAssetUnion `json:"lp_asset,required"`
	// An arbitrary precision decimal represented as a string
	Ownership string `json:"ownership,required"`
	// A liquidity pool
	Pool LPPool `json:"pool,required"`
	// An asset symbol
	QuoteAssetSymbol string `json:"quote_asset_symbol,required"`
	// An arbitrary precision decimal represented as a string
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account          respjson.Field
		Amounts          respjson.Field
		Attributes       respjson.Field
		Blockstamp       respjson.Field
		LPAmount         respjson.Field
		LPAsset          respjson.Field
		Ownership        respjson.Field
		Pool             respjson.Field
		QuoteAssetSymbol respjson.Field
		Value            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LPPosition) RawJSON() string { return r.JSON.raw }
func (r *LPPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LPPositionLPAssetUnion contains all possible properties and values from [Asset],
// [NFT].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type LPPositionLPAssetUnion struct {
	// This field is a union of [AssetID], [int64]
	ID LPPositionLPAssetUnionID `json:"id"`
	// This field is from variant [Asset].
	Address string `json:"address"`
	// This field is from variant [Asset].
	Blockchain Blockchain `json:"blockchain"`
	// This field is from variant [Asset].
	Decimals int64 `json:"decimals"`
	// This field is from variant [Asset].
	Name string `json:"name"`
	// This field is from variant [Asset].
	Symbol string `json:"symbol"`
	// This field is from variant [NFT].
	Collection NFTCollection `json:"collection"`
	JSON       struct {
		ID         respjson.Field
		Address    respjson.Field
		Blockchain respjson.Field
		Decimals   respjson.Field
		Name       respjson.Field
		Symbol     respjson.Field
		Collection respjson.Field
		raw        string
	} `json:"-"`
}

func (u LPPositionLPAssetUnion) AsAsset() (v Asset) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LPPositionLPAssetUnion) AsNFT() (v NFT) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LPPositionLPAssetUnion) RawJSON() string { return u.JSON.raw }

func (r *LPPositionLPAssetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LPPositionLPAssetUnionID is an implicit subunion of [LPPositionLPAssetUnion].
// LPPositionLPAssetUnionID provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [LPPositionLPAssetUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInt]
type LPPositionLPAssetUnionID struct {
	// This field will be present if the value is a [AssetID] instead of an object.
	OfString AssetID `json:",inline"`
	// This field will be present if the value is a [int64] instead of an object.
	OfInt int64 `json:",inline"`
	JSON  struct {
		OfString respjson.Field
		OfInt    respjson.Field
		raw      string
	} `json:"-"`
}

func (r *LPPositionLPAssetUnionID) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Attributes for Uniswap V3 venue type positions
type LPPositionAttributes struct {
	// Whether the position is in range
	InRange bool `json:"in_range,required"`
	// An arbitrary precision decimal represented as a string
	Lower Decimal `json:"lower,required"`
	// An arbitrary precision decimal represented as a string
	Price Decimal `json:"price,required"`
	// Lower tick
	TickLower int64 `json:"tick_lower,required"`
	// Upper tick
	TickUpper int64 `json:"tick_upper,required"`
	// An arbitrary precision decimal represented as a string
	Upper Decimal `json:"upper,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InRange     respjson.Field
		Lower       respjson.Field
		Price       respjson.Field
		TickLower   respjson.Field
		TickUpper   respjson.Field
		Upper       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LPPositionAttributes) RawJSON() string { return r.JSON.raw }
func (r *LPPositionAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An LP quote
type LPQuote struct {
	// An EVM address
	Account Account `json:"account,required"`
	// A list of arbitrary precision decimals
	Amounts []Decimal `json:"amounts,required"`
	// A list of arbitrary precision decimals
	AmountsDelta []Decimal `json:"amounts_delta,required"`
	// Information about a specific block number and its timestamp
	Blockstamp Blockstamp `json:"blockstamp,required"`
	// An arbitrary precision decimal represented as a string
	Ownership string `json:"ownership,required"`
	// A liquidity pool
	Pool LPPool `json:"pool,required"`
	// An asset symbol
	QuoteAssetSymbol string `json:"quote_asset_symbol,required"`
	// An arbitrary precision decimal represented as a string
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account          respjson.Field
		Amounts          respjson.Field
		AmountsDelta     respjson.Field
		Blockstamp       respjson.Field
		Ownership        respjson.Field
		Pool             respjson.Field
		QuoteAssetSymbol respjson.Field
		Value            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LPQuote) RawJSON() string { return r.JSON.raw }
func (r *LPQuote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkName = string

// A NFT
type NFT struct {
	// The NFT id
	ID int64 `json:"id,required"`
	// A NFT Collection
	Collection NFTCollection `json:"collection,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Collection  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NFT) RawJSON() string { return r.JSON.raw }
func (r *NFT) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this NFT to a NFTParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// NFTParam.Overrides()
func (r NFT) ToParam() NFTParam {
	return param.Override[NFTParam](json.RawMessage(r.RawJSON()))
}

// A NFT
//
// The properties ID, Collection are required.
type NFTParam struct {
	// The NFT id
	ID int64 `json:"id,required"`
	// A NFT Collection
	Collection NFTCollectionParam `json:"collection,omitzero,required"`
	paramObj
}

func (r NFTParam) MarshalJSON() (data []byte, err error) {
	type shadow NFTParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NFTParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A NFT Collection
type NFTCollection struct {
	// The NFT Collection's address
	Address string `json:"address,required"`
	// Data about a blockchain
	Blockchain NFTCollectionBlockchain `json:"blockchain"`
	// The NFT Collection's name
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		Blockchain  respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NFTCollection) RawJSON() string { return r.JSON.raw }
func (r *NFTCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this NFTCollection to a NFTCollectionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// NFTCollectionParam.Overrides()
func (r NFTCollection) ToParam() NFTCollectionParam {
	return param.Override[NFTCollectionParam](json.RawMessage(r.RawJSON()))
}

// Data about a blockchain
type NFTCollectionBlockchain struct {
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	ID BlockchainID `json:"id,required"`
	// Data about an EVM blockchain
	ChainData EVMChainData `json:"chain_data,required"`
	// Blockchain ecosystem
	//
	// Any of "evm".
	ChainType ChainType `json:"chain_type,required"`
	// The blockchain's explorer URL
	ExplorerURL string `json:"explorer_url,required"`
	// BlockchainName
	Name BlockchainName `json:"name,required"`
	// Blockchain's network
	Network NetworkName `json:"network,required"`
	// A blockchain symbol
	Symbol BlockchainSymbol `json:"symbol,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ChainData   respjson.Field
		ChainType   respjson.Field
		ExplorerURL respjson.Field
		Name        respjson.Field
		Network     respjson.Field
		Symbol      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NFTCollectionBlockchain) RawJSON() string { return r.JSON.raw }
func (r *NFTCollectionBlockchain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A NFT Collection
//
// The property Address is required.
type NFTCollectionParam struct {
	// The NFT Collection's address
	Address string `json:"address,required"`
	// The NFT Collection's name
	Name param.Opt[string] `json:"name,omitzero"`
	// Data about a blockchain
	Blockchain NFTCollectionBlockchainParam `json:"blockchain,omitzero"`
	paramObj
}

func (r NFTCollectionParam) MarshalJSON() (data []byte, err error) {
	type shadow NFTCollectionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NFTCollectionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Data about a blockchain
//
// The properties ID, ChainData, ChainType, ExplorerURL, Name, Network, Symbol are
// required.
type NFTCollectionBlockchainParam struct {
	// A blockchain ID, represented as a TypeID with `blockchain` prefix
	ID BlockchainID `json:"id,required"`
	// Data about an EVM blockchain
	ChainData EVMChainDataParam `json:"chain_data,omitzero,required"`
	// Blockchain ecosystem
	//
	// Any of "evm".
	ChainType ChainType `json:"chain_type,omitzero,required"`
	// The blockchain's explorer URL
	ExplorerURL string `json:"explorer_url,required"`
	// BlockchainName
	Name BlockchainName `json:"name,required"`
	// Blockchain's network
	Network NetworkName `json:"network,required"`
	// A blockchain symbol
	Symbol BlockchainSymbol `json:"symbol,required"`
	paramObj
}

func (r NFTCollectionBlockchainParam) MarshalJSON() (data []byte, err error) {
	type shadow NFTCollectionBlockchainParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NFTCollectionBlockchainParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A range of timestamps
type OffChainHistoricalRange struct {
	// ISO8601 Timestamp
	From Timestamp `json:"from,required"`
	// ISO8601 Timestamp
	To Timestamp `json:"to,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OffChainHistoricalRange) RawJSON() string { return r.JSON.raw }
func (r *OffChainHistoricalRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A range of blockstamps
type OnChainHistoricalRange struct {
	// Information about a specific block number and its timestamp
	From Blockstamp `json:"from,required"`
	// Information about a specific block number and its timestamp
	To Blockstamp `json:"to,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OnChainHistoricalRange) RawJSON() string { return r.JSON.raw }
func (r *OnChainHistoricalRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A price
type Price struct {
	// An asset symbol
	AssetSymbol AssetSymbol `json:"asset_symbol,required"`
	// An arbitrary precision decimal represented as a string
	Change1h string `json:"change_1h,required"`
	// An arbitrary precision decimal represented as a string
	Change1y string `json:"change_1y,required"`
	// An arbitrary precision decimal represented as a string
	Change24h string `json:"change_24h,required"`
	// An arbitrary precision decimal represented as a string
	Change30d string `json:"change_30d,required"`
	// An arbitrary precision decimal represented as a string
	Change7d string `json:"change_7d,required"`
	// An arbitrary precision decimal represented as a string
	MarketCap string `json:"market_cap,required"`
	// An arbitrary precision decimal represented as a string
	Price Decimal `json:"price,required"`
	// An asset symbol
	QuoteAssetSymbol AssetSymbol `json:"quote_asset_symbol,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetSymbol      respjson.Field
		Change1h         respjson.Field
		Change1y         respjson.Field
		Change24h        respjson.Field
		Change30d        respjson.Field
		Change7d         respjson.Field
		MarketCap        respjson.Field
		Price            respjson.Field
		QuoteAssetSymbol respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Price) RawJSON() string { return r.JSON.raw }
func (r *Price) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Quote info for 0x
type QuoteInfo0x struct {
	// An arbitrary precision decimal represented as a string
	EstimatedGasUsed string `json:"estimated_gas_used,required"`
	// A route in a 0x quote
	Route QuoteInfo0xRoute `json:"route,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EstimatedGasUsed respjson.Field
		Route            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuoteInfo0x) RawJSON() string { return r.JSON.raw }
func (r *QuoteInfo0x) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A fill in a 0x quote
type QuoteInfo0xFill struct {
	// An EVM address
	From          AddressEVM `json:"from,required"`
	ProportionBps int64      `json:"proportion_bps,required"`
	Source        string     `json:"source,required"`
	// An EVM address
	To AddressEVM `json:"to,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From          respjson.Field
		ProportionBps respjson.Field
		Source        respjson.Field
		To            respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuoteInfo0xFill) RawJSON() string { return r.JSON.raw }
func (r *QuoteInfo0xFill) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A route in a 0x quote
type QuoteInfo0xRoute struct {
	// A list of fills in a 0x quote
	Fills []QuoteInfo0xFill `json:"fills,required"`
	// A list of tokens in a 0x quote
	Tokens []QuoteInfo0xToken `json:"tokens,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Fills       respjson.Field
		Tokens      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuoteInfo0xRoute) RawJSON() string { return r.JSON.raw }
func (r *QuoteInfo0xRoute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A token in a 0x quote
type QuoteInfo0xToken struct {
	// An EVM address
	Address AddressEVM `json:"address,required"`
	Symbol  string     `json:"symbol,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		Symbol      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuoteInfo0xToken) RawJSON() string { return r.JSON.raw }
func (r *QuoteInfo0xToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Quote info for Curve
type QuoteInfoCurve struct {
	IIndex int64 `json:"i_index,required"`
	JIndex int64 `json:"j_index,required"`
	// A LP pool ID, represented as a TypeID with `lp_pool` prefix
	PoolID   string `json:"pool_id,required"`
	SwapType string `json:"swap_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IIndex      respjson.Field
		JIndex      respjson.Field
		PoolID      respjson.Field
		SwapType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuoteInfoCurve) RawJSON() string { return r.JSON.raw }
func (r *QuoteInfoCurve) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Quote info for Uniswap V2
type QuoteInfoUniswapV2 struct {
	// A list of LP pool IDs
	PoolIDs []string `json:"pool_ids,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PoolIDs     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuoteInfoUniswapV2) RawJSON() string { return r.JSON.raw }
func (r *QuoteInfoUniswapV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Quote info for Uniswap V3
type QuoteInfoUniswapV3 struct {
	// A list of LP pool IDs
	PoolIDs []string `json:"pool_ids,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PoolIDs     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuoteInfoUniswapV3) RawJSON() string { return r.JSON.raw }
func (r *QuoteInfoUniswapV3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Registry data for a lend borrow market
type RegistryLendBorrowMarket struct {
	// A lend/borrow market ID, represented as a TypeID with `lend_borrow_market`
	// prefix
	ID LendBorrowMarketID `json:"id,required"`
	// On-chain asset (aka token)
	Asset Asset `json:"asset,required"`
	// Data about a blockchain
	Blockchain Blockchain `json:"blockchain,required"`
	// On-chain venue
	Venue Venue `json:"venue,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Asset       respjson.Field
		Blockchain  respjson.Field
		Venue       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryLendBorrowMarket) RawJSON() string { return r.JSON.raw }
func (r *RegistryLendBorrowMarket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Registry data for an LP pool
type RegistryLPPool struct {
	// A LP pool ID, represented as a TypeID with `lp_pool` prefix
	ID string `json:"id,required"`
	// Response for multiple assets
	Assets []Asset `json:"assets,required"`
	// Data about a blockchain
	Blockchain Blockchain `json:"blockchain,required"`
	// On-chain venue
	Venue Venue `json:"venue,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Assets      respjson.Field
		Blockchain  respjson.Field
		Venue       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegistryLPPool) RawJSON() string { return r.JSON.raw }
func (r *RegistryLPPool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A sell quote
type SellQuote struct {
	// Data about a blockchain
	Blockchain Blockchain `json:"blockchain,required"`
	// Estimated cost of an operation
	FeeEstimate SellQuoteFeeEstimate `json:"fee_estimate,required"`
	// An asset symbol
	QuoteAssetSymbol string `json:"quote_asset_symbol,required"`
	// Quote info for sell quotes
	QuoteInfo SellQuoteQuoteInfoUnion `json:"quote_info,required"`
	// An arbitrary precision decimal represented as a string
	QuoteValue string `json:"quote_value,required"`
	// An arbitrary precision decimal represented as a string
	ReceiveAmount Decimal `json:"receive_amount,required"`
	// On-chain asset (aka token)
	ReceiveAsset Asset `json:"receive_asset,required"`
	// An arbitrary precision decimal represented as a string
	SpendAmount Decimal `json:"spend_amount,required"`
	// On-chain asset (aka token)
	SpendAsset Asset `json:"spend_asset,required"`
	// On-chain venue
	Venue Venue `json:"venue,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blockchain       respjson.Field
		FeeEstimate      respjson.Field
		QuoteAssetSymbol respjson.Field
		QuoteInfo        respjson.Field
		QuoteValue       respjson.Field
		ReceiveAmount    respjson.Field
		ReceiveAsset     respjson.Field
		SpendAmount      respjson.Field
		SpendAsset       respjson.Field
		Venue            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SellQuote) RawJSON() string { return r.JSON.raw }
func (r *SellQuote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Estimated cost of an operation
type SellQuoteFeeEstimate struct {
	// An arbitrary precision decimal represented as a string
	Amount Decimal `json:"amount,required"`
	// On-chain asset (aka token)
	Asset Asset `json:"asset,required"`
	// An arbitrary precision decimal represented as a string
	Cost string `json:"cost,required"`
	// An asset symbol
	QuoteAssetSymbol string `json:"quote_asset_symbol,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amount           respjson.Field
		Asset            respjson.Field
		Cost             respjson.Field
		QuoteAssetSymbol respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SellQuoteFeeEstimate) RawJSON() string { return r.JSON.raw }
func (r *SellQuoteFeeEstimate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SellQuoteQuoteInfoUnion contains all possible properties and values from
// [QuoteInfoCurve], [QuoteInfoUniswapV2], [QuoteInfoUniswapV3], [QuoteInfo0x].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SellQuoteQuoteInfoUnion struct {
	// This field is from variant [QuoteInfoCurve].
	IIndex int64 `json:"i_index"`
	// This field is from variant [QuoteInfoCurve].
	JIndex int64 `json:"j_index"`
	// This field is from variant [QuoteInfoCurve].
	PoolID string `json:"pool_id"`
	// This field is from variant [QuoteInfoCurve].
	SwapType string   `json:"swap_type"`
	PoolIDs  []string `json:"pool_ids"`
	// This field is from variant [QuoteInfo0x].
	EstimatedGasUsed string `json:"estimated_gas_used"`
	// This field is from variant [QuoteInfo0x].
	Route QuoteInfo0xRoute `json:"route"`
	JSON  struct {
		IIndex           respjson.Field
		JIndex           respjson.Field
		PoolID           respjson.Field
		SwapType         respjson.Field
		PoolIDs          respjson.Field
		EstimatedGasUsed respjson.Field
		Route            respjson.Field
		raw              string
	} `json:"-"`
}

func (u SellQuoteQuoteInfoUnion) AsQuoteInfoCurve() (v QuoteInfoCurve) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SellQuoteQuoteInfoUnion) AsQuoteInfoUniswapV2() (v QuoteInfoUniswapV2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SellQuoteQuoteInfoUnion) AsQuoteInfoUniswapV3() (v QuoteInfoUniswapV3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SellQuoteQuoteInfoUnion) AsQuoteInfo0x() (v QuoteInfo0x) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SellQuoteQuoteInfoUnion) RawJSON() string { return u.JSON.raw }

func (r *SellQuoteQuoteInfoUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Timestamp = string

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TimestampOrBlockNumberUnionParam struct {
	OfString param.Opt[Timestamp]   `json:",omitzero,inline"`
	OfInt    param.Opt[BlockNumber] `json:",omitzero,inline"`
	paramUnion
}

func (u TimestampOrBlockNumberUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInt)
}
func (u *TimestampOrBlockNumberUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TimestampOrBlockNumberUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInt) {
		return &u.OfInt.Value
	}
	return nil
}

// An error
type VektorError struct {
	// Error message
	Message string `json:"message,required"`
	// Error parameters
	Params map[string]any `json:"params,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Params      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VektorError) RawJSON() string { return r.JSON.raw }
func (r *VektorError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// On-chain venue
type Venue struct {
	// A venue ID, represented as a TypeID with `venue` prefix
	ID VenueID `json:"id,required"`
	// A list of blockchain IDs
	BlockchainIDs []BlockchainID `json:"blockchain_ids,required"`
	// Venue name
	Name string `json:"name,required"`
	// A venue symbol
	Symbol VenueSymbol `json:"symbol,required"`
	// Venue type
	Type string `json:"type,required"`
	// Venue url
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		BlockchainIDs respjson.Field
		Name          respjson.Field
		Symbol        respjson.Field
		Type          respjson.Field
		URL           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Venue) RawJSON() string { return r.JSON.raw }
func (r *Venue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VenueID = string

type VenueIDOrVenueSymbol = string

type VenueSymbol = string

// A LP incentive vote market
type VoteMarket struct {
	// Information about a specific block number and its timestamp
	Blockstamp Blockstamp `json:"blockstamp,required"`
	// A lp pool incentive market
	Market IncentivizeMarket `json:"market,required"`
	// An asset symbol
	QuoteAssetSymbol string `json:"quote_asset_symbol,required"`
	// A list of arbitrary precision decimals
	TradingFeesAmounts []Decimal `json:"trading_fees_amounts,required"`
	// Response for multiple assets
	TradingFeesAssets []Asset `json:"trading_fees_assets,required"`
	// An arbitrary precision decimal represented as a string
	Value Decimal `json:"value,required"`
	// An arbitrary precision decimal represented as a string
	ValuePerVote Decimal `json:"value_per_vote,required"`
	// An arbitrary precision decimal represented as a string
	Votes Decimal `json:"votes,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blockstamp         respjson.Field
		Market             respjson.Field
		QuoteAssetSymbol   respjson.Field
		TradingFeesAmounts respjson.Field
		TradingFeesAssets  respjson.Field
		Value              respjson.Field
		ValuePerVote       respjson.Field
		Votes              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoteMarket) RawJSON() string { return r.JSON.raw }
func (r *VoteMarket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A LP vote reward
type VoteReward struct {
	// An EVM address
	Account Account `json:"account,required"`
	// Information about a specific block number and its timestamp
	Blockstamp Blockstamp `json:"blockstamp,required"`
	// A list of arbitrary precision decimals
	FeeAmounts []Decimal `json:"fee_amounts,required"`
	// Response for multiple assets
	FeeAssets []Asset `json:"fee_assets,required"`
	// A list of arbitrary precision decimals
	IncentiveAmounts []Decimal `json:"incentive_amounts,required"`
	// Response for multiple assets
	IncentiveAssets []Asset `json:"incentive_assets,required"`
	// A list of lock positions
	LockPositions []LockPosition `json:"lock_positions,required"`
	// A liquidity pool
	Pool LPPool `json:"pool,required"`
	// An asset symbol
	QuoteAssetSymbol string `json:"quote_asset_symbol,required"`
	// An arbitrary precision decimal represented as a string
	Value Decimal `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Account          respjson.Field
		Blockstamp       respjson.Field
		FeeAmounts       respjson.Field
		FeeAssets        respjson.Field
		IncentiveAmounts respjson.Field
		IncentiveAssets  respjson.Field
		LockPositions    respjson.Field
		Pool             respjson.Field
		QuoteAssetSymbol respjson.Field
		Value            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoteReward) RawJSON() string { return r.JSON.raw }
func (r *VoteReward) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
