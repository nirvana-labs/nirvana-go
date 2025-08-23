// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vektor

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/nirvana-labs/nirvana-go/internal/apijson"
	"github.com/nirvana-labs/nirvana-go/internal/requestconfig"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/packages/param"
	"github.com/nirvana-labs/nirvana-go/packages/respjson"
)

// ExecutionStepService contains methods and other services that help with
// interacting with the Nirvana Labs API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExecutionStepService] method instead.
type ExecutionStepService struct {
	Options []option.RequestOption
}

// NewExecutionStepService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExecutionStepService(opts ...option.RequestOption) (r ExecutionStepService) {
	r = ExecutionStepService{}
	r.Options = opts
	return
}

// Get a step of an execution
func (r *ExecutionStepService) Get(ctx context.Context, executionID string, stepID string, opts ...option.RequestOption) (res *ExecutionStepGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if executionID == "" {
		err = errors.New("missing required execution_id parameter")
		return
	}
	if stepID == "" {
		err = errors.New("missing required step_id parameter")
		return
	}
	path := fmt.Sprintf("v1/vektor/executions/%s/steps/%s", executionID, stepID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Sign an EVM transaction step
func (r *ExecutionStepService) Sign(ctx context.Context, executionID string, stepID string, body ExecutionStepSignParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if executionID == "" {
		err = errors.New("missing required execution_id parameter")
		return
	}
	if stepID == "" {
		err = errors.New("missing required step_id parameter")
		return
	}
	path := fmt.Sprintf("v1/vektor/executions/%s/steps/%s/sign", executionID, stepID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// An execution step
type ExecutionStepGetResponse struct {
	// An execution step ID, represented as a TypeID with `execution_step` prefix
	ID ExecutionStepID `json:"id,required"`
	// ISO8601 Timestamp
	CreatedAt Timestamp `json:"created_at,required"`
	// The definition of an execution step
	Definition ExecutionStepGetResponseDefinitionUnion `json:"definition,required"`
	Index      int64                                   `json:"index,required"`
	// The type of an execution step
	//
	// Any of "evm_transaction_approve", "evm_transaction_borrow",
	// "evm_transaction_borrow_repay", "evm_transaction_buy", "evm_transaction_lend",
	// "evm_transaction_lend_set_collateral", "evm_transaction_lend_withdraw",
	// "evm_transaction_move", "evm_transaction_permission", "evm_transaction_wrap",
	// "evm_transaction_unwrap", "evm_transaction_sell".
	Type ExecutionStepGetResponseType `json:"type,required"`
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
func (r ExecutionStepGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ExecutionStepGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExecutionStepGetResponseDefinitionUnion contains all possible properties and
// values from [ExecutionStepGetResponseDefinitionExecutionEVMTransactionApprove],
// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrow],
// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowRepay],
// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionBuy],
// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionLend],
// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendSetCollateral],
// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendWithdraw],
// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionMove],
// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionPermission],
// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionUnwrap],
// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionWrap],
// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionSell].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ExecutionStepGetResponseDefinitionUnion struct {
	Amount string `json:"amount"`
	// This field is from variant
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionApprove].
	Asset              Asset  `json:"asset"`
	BlockNumber        int64  `json:"block_number"`
	BroadcastedAt      string `json:"broadcasted_at"`
	ConfirmationTarget int64  `json:"confirmation_target"`
	ConfirmedAt        string `json:"confirmed_at"`
	// This field is from variant
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionApprove].
	CreatedAt         Timestamp `json:"created_at"`
	Data              string    `json:"data"`
	EffectiveGasPrice string    `json:"effective_gas_price"`
	// This field is a union of
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionApproveError],
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowError],
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowRepayError],
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionBuyError],
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendError],
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendSetCollateralError],
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendWithdrawError],
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionMoveError],
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionPermissionError],
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionUnwrapError],
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionWrapError],
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionSellError]
	Error     ExecutionStepGetResponseDefinitionUnionError `json:"error"`
	ErroredAt string                                       `json:"errored_at"`
	GasUsed   string                                       `json:"gas_used"`
	Hash      string                                       `json:"hash"`
	// This field is from variant
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionApprove].
	Payload  ExecutionEVMTransactionEIP1559Payload `json:"payload"`
	SignedAt string                                `json:"signed_at"`
	// This field is from variant
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionApprove].
	Spender Account `json:"spender"`
	// This field is from variant
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionApprove].
	State ExecutionEVMTransactionState `json:"state"`
	// This field is from variant
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionApprove].
	TargetState ExecutionEVMTransactionState `json:"target_state"`
	To          string                       `json:"to"`
	Type        string                       `json:"type"`
	// This field is from variant
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionApprove].
	UpdatedAt Timestamp `json:"updated_at"`
	Value     string    `json:"value"`
	// This field is from variant
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrow].
	LendBorrowMarketID LendBorrowMarketID `json:"lend_borrow_market_id"`
	// This field is from variant
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrow].
	VenueSymbol VenueSymbol `json:"venue_symbol"`
	// This field is from variant
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionBuy].
	ApprovalTarget Account `json:"approval_target"`
	// This field is from variant
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionBuy].
	MaxSpendAmount Decimal `json:"max_spend_amount"`
	// This field is a union of [BuyQuote], [SellQuote]
	Quote ExecutionStepGetResponseDefinitionUnionQuote `json:"quote"`
	// This field is from variant
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionBuy].
	Slippage Decimal `json:"slippage"`
	// This field is from variant
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendSetCollateral].
	Status bool `json:"status"`
	// This field is from variant
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionPermission].
	ContractAddress AddressEVM `json:"contract_address"`
	// This field is from variant
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionPermission].
	Permission bool `json:"permission"`
	// This field is from variant
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionSell].
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

func (u ExecutionStepGetResponseDefinitionUnion) AsExecutionEVMTransactionApprove() (v ExecutionStepGetResponseDefinitionExecutionEVMTransactionApprove) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionStepGetResponseDefinitionUnion) AsExecutionEVMTransactionBorrow() (v ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionStepGetResponseDefinitionUnion) AsExecutionEVMTransactionBorrowRepay() (v ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowRepay) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionStepGetResponseDefinitionUnion) AsExecutionEVMTransactionBuy() (v ExecutionStepGetResponseDefinitionExecutionEVMTransactionBuy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionStepGetResponseDefinitionUnion) AsExecutionEVMTransactionLend() (v ExecutionStepGetResponseDefinitionExecutionEVMTransactionLend) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionStepGetResponseDefinitionUnion) AsExecutionEVMTransactionLendSetCollateral() (v ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendSetCollateral) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionStepGetResponseDefinitionUnion) AsExecutionEVMTransactionLendWithdraw() (v ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendWithdraw) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionStepGetResponseDefinitionUnion) AsExecutionEVMTransactionMove() (v ExecutionStepGetResponseDefinitionExecutionEVMTransactionMove) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionStepGetResponseDefinitionUnion) AsExecutionEVMTransactionPermission() (v ExecutionStepGetResponseDefinitionExecutionEVMTransactionPermission) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionStepGetResponseDefinitionUnion) AsExecutionEVMTransactionUnwrap() (v ExecutionStepGetResponseDefinitionExecutionEVMTransactionUnwrap) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionStepGetResponseDefinitionUnion) AsExecutionEVMTransactionWrap() (v ExecutionStepGetResponseDefinitionExecutionEVMTransactionWrap) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExecutionStepGetResponseDefinitionUnion) AsExecutionEVMTransactionSell() (v ExecutionStepGetResponseDefinitionExecutionEVMTransactionSell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExecutionStepGetResponseDefinitionUnion) RawJSON() string { return u.JSON.raw }

func (r *ExecutionStepGetResponseDefinitionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExecutionStepGetResponseDefinitionUnionError is an implicit subunion of
// [ExecutionStepGetResponseDefinitionUnion].
// ExecutionStepGetResponseDefinitionUnionError provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ExecutionStepGetResponseDefinitionUnion].
type ExecutionStepGetResponseDefinitionUnionError struct {
	// This field is a union of
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionApproveErrorContext],
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowErrorContext],
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowRepayErrorContext],
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionBuyErrorContext],
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendErrorContext],
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendSetCollateralErrorContext],
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendWithdrawErrorContext],
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionMoveErrorContext],
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionPermissionErrorContext],
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionUnwrapErrorContext],
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionWrapErrorContext],
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionSellErrorContext]
	Context   ExecutionStepGetResponseDefinitionUnionErrorContext `json:"context"`
	Message   string                                              `json:"message"`
	RequestID string                                              `json:"request_id"`
	Resource  string                                              `json:"resource"`
	// This field is from variant
	// [ExecutionStepGetResponseDefinitionExecutionEVMTransactionApproveError].
	Timestamp Timestamp `json:"timestamp"`
	Type      string    `json:"type"`
	JSON      struct {
		Context   respjson.Field
		Message   respjson.Field
		RequestID respjson.Field
		Resource  respjson.Field
		Timestamp respjson.Field
		Type      respjson.Field
		raw       string
	} `json:"-"`
}

func (r *ExecutionStepGetResponseDefinitionUnionError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExecutionStepGetResponseDefinitionUnionErrorContext is an implicit subunion of
// [ExecutionStepGetResponseDefinitionUnion].
// ExecutionStepGetResponseDefinitionUnionErrorContext provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ExecutionStepGetResponseDefinitionUnion].
type ExecutionStepGetResponseDefinitionUnionErrorContext struct {
	Parameters any `json:"parameters"`
	JSON       struct {
		Parameters respjson.Field
		raw        string
	} `json:"-"`
}

func (r *ExecutionStepGetResponseDefinitionUnionErrorContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExecutionStepGetResponseDefinitionUnionQuote is an implicit subunion of
// [ExecutionStepGetResponseDefinitionUnion].
// ExecutionStepGetResponseDefinitionUnionQuote provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ExecutionStepGetResponseDefinitionUnion].
type ExecutionStepGetResponseDefinitionUnionQuote struct {
	// This field is from variant [BuyQuote].
	Blockchain Blockchain `json:"blockchain"`
	// This field is a union of [BuyQuoteFeeEstimate], [SellQuoteFeeEstimate]
	FeeEstimate      ExecutionStepGetResponseDefinitionUnionQuoteFeeEstimate `json:"fee_estimate"`
	QuoteAssetSymbol string                                                  `json:"quote_asset_symbol"`
	// This field is a union of [BuyQuoteQuoteInfoUnion], [SellQuoteQuoteInfoUnion]
	QuoteInfo  ExecutionStepGetResponseDefinitionUnionQuoteQuoteInfo `json:"quote_info"`
	QuoteValue string                                                `json:"quote_value"`
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

func (r *ExecutionStepGetResponseDefinitionUnionQuote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExecutionStepGetResponseDefinitionUnionQuoteFeeEstimate is an implicit subunion
// of [ExecutionStepGetResponseDefinitionUnion].
// ExecutionStepGetResponseDefinitionUnionQuoteFeeEstimate provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ExecutionStepGetResponseDefinitionUnion].
type ExecutionStepGetResponseDefinitionUnionQuoteFeeEstimate struct {
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

func (r *ExecutionStepGetResponseDefinitionUnionQuoteFeeEstimate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExecutionStepGetResponseDefinitionUnionQuoteQuoteInfo is an implicit subunion of
// [ExecutionStepGetResponseDefinitionUnion].
// ExecutionStepGetResponseDefinitionUnionQuoteQuoteInfo provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ExecutionStepGetResponseDefinitionUnion].
type ExecutionStepGetResponseDefinitionUnionQuoteQuoteInfo struct {
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

func (r *ExecutionStepGetResponseDefinitionUnionQuoteQuoteInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An approval of an asset
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionApprove struct {
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
	// Vektor error
	Error ExecutionStepGetResponseDefinitionExecutionEVMTransactionApproveError `json:"error,required"`
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
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionApprove) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionApprove) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Vektor error
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionApproveError struct {
	// Error context
	Context ExecutionStepGetResponseDefinitionExecutionEVMTransactionApproveErrorContext `json:"context,required"`
	// Error message
	Message string `json:"message,required"`
	// Request ID
	RequestID string `json:"request_id,required"`
	// Error resource
	Resource string `json:"resource,required"`
	// ISO8601 Timestamp
	Timestamp Timestamp `json:"timestamp,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context     respjson.Field
		Message     respjson.Field
		RequestID   respjson.Field
		Resource    respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionApproveError) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionApproveError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error context
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionApproveErrorContext struct {
	// Error parameters
	Parameters map[string]any `json:"parameters,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parameters  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionApproveErrorContext) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionApproveErrorContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Borrowing an asset
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrow struct {
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
	// Vektor error
	Error ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowError `json:"error,required"`
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
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrow) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Vektor error
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowError struct {
	// Error context
	Context ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowErrorContext `json:"context,required"`
	// Error message
	Message string `json:"message,required"`
	// Request ID
	RequestID string `json:"request_id,required"`
	// Error resource
	Resource string `json:"resource,required"`
	// ISO8601 Timestamp
	Timestamp Timestamp `json:"timestamp,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context     respjson.Field
		Message     respjson.Field
		RequestID   respjson.Field
		Resource    respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowError) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error context
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowErrorContext struct {
	// Error parameters
	Parameters map[string]any `json:"parameters,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parameters  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowErrorContext) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowErrorContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Repaying a borrowed asset
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowRepay struct {
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
	// Vektor error
	Error ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowRepayError `json:"error,required"`
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
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowRepay) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowRepay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Vektor error
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowRepayError struct {
	// Error context
	Context ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowRepayErrorContext `json:"context,required"`
	// Error message
	Message string `json:"message,required"`
	// Request ID
	RequestID string `json:"request_id,required"`
	// Error resource
	Resource string `json:"resource,required"`
	// ISO8601 Timestamp
	Timestamp Timestamp `json:"timestamp,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context     respjson.Field
		Message     respjson.Field
		RequestID   respjson.Field
		Resource    respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowRepayError) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowRepayError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error context
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowRepayErrorContext struct {
	// Error parameters
	Parameters map[string]any `json:"parameters,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parameters  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowRepayErrorContext) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionBorrowRepayErrorContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Buying an asset with another asset
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionBuy struct {
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
	// Vektor error
	Error ExecutionStepGetResponseDefinitionExecutionEVMTransactionBuyError `json:"error,required"`
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
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionBuy) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionBuy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Vektor error
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionBuyError struct {
	// Error context
	Context ExecutionStepGetResponseDefinitionExecutionEVMTransactionBuyErrorContext `json:"context,required"`
	// Error message
	Message string `json:"message,required"`
	// Request ID
	RequestID string `json:"request_id,required"`
	// Error resource
	Resource string `json:"resource,required"`
	// ISO8601 Timestamp
	Timestamp Timestamp `json:"timestamp,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context     respjson.Field
		Message     respjson.Field
		RequestID   respjson.Field
		Resource    respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionBuyError) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionBuyError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error context
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionBuyErrorContext struct {
	// Error parameters
	Parameters map[string]any `json:"parameters,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parameters  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionBuyErrorContext) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionBuyErrorContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lending an asset
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionLend struct {
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
	// Vektor error
	Error ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendError `json:"error,required"`
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
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionLend) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionLend) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Vektor error
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendError struct {
	// Error context
	Context ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendErrorContext `json:"context,required"`
	// Error message
	Message string `json:"message,required"`
	// Request ID
	RequestID string `json:"request_id,required"`
	// Error resource
	Resource string `json:"resource,required"`
	// ISO8601 Timestamp
	Timestamp Timestamp `json:"timestamp,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context     respjson.Field
		Message     respjson.Field
		RequestID   respjson.Field
		Resource    respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendError) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error context
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendErrorContext struct {
	// Error parameters
	Parameters map[string]any `json:"parameters,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parameters  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendErrorContext) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendErrorContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Setting/unsetting a position as collateral
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendSetCollateral struct {
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
	// Vektor error
	Error ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendSetCollateralError `json:"error,required"`
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
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendSetCollateral) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendSetCollateral) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Vektor error
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendSetCollateralError struct {
	// Error context
	Context ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendSetCollateralErrorContext `json:"context,required"`
	// Error message
	Message string `json:"message,required"`
	// Request ID
	RequestID string `json:"request_id,required"`
	// Error resource
	Resource string `json:"resource,required"`
	// ISO8601 Timestamp
	Timestamp Timestamp `json:"timestamp,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context     respjson.Field
		Message     respjson.Field
		RequestID   respjson.Field
		Resource    respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendSetCollateralError) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendSetCollateralError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error context
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendSetCollateralErrorContext struct {
	// Error parameters
	Parameters map[string]any `json:"parameters,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parameters  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendSetCollateralErrorContext) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendSetCollateralErrorContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Withdrawing an asset
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendWithdraw struct {
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
	// Vektor error
	Error ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendWithdrawError `json:"error,required"`
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
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendWithdraw) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendWithdraw) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Vektor error
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendWithdrawError struct {
	// Error context
	Context ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendWithdrawErrorContext `json:"context,required"`
	// Error message
	Message string `json:"message,required"`
	// Request ID
	RequestID string `json:"request_id,required"`
	// Error resource
	Resource string `json:"resource,required"`
	// ISO8601 Timestamp
	Timestamp Timestamp `json:"timestamp,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context     respjson.Field
		Message     respjson.Field
		RequestID   respjson.Field
		Resource    respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendWithdrawError) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendWithdrawError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error context
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendWithdrawErrorContext struct {
	// Error parameters
	Parameters map[string]any `json:"parameters,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parameters  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendWithdrawErrorContext) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionLendWithdrawErrorContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A move of assets from one account to another
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionMove struct {
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
	// Vektor error
	Error ExecutionStepGetResponseDefinitionExecutionEVMTransactionMoveError `json:"error,required"`
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
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionMove) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionMove) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Vektor error
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionMoveError struct {
	// Error context
	Context ExecutionStepGetResponseDefinitionExecutionEVMTransactionMoveErrorContext `json:"context,required"`
	// Error message
	Message string `json:"message,required"`
	// Request ID
	RequestID string `json:"request_id,required"`
	// Error resource
	Resource string `json:"resource,required"`
	// ISO8601 Timestamp
	Timestamp Timestamp `json:"timestamp,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context     respjson.Field
		Message     respjson.Field
		RequestID   respjson.Field
		Resource    respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionMoveError) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionMoveError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error context
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionMoveErrorContext struct {
	// Error parameters
	Parameters map[string]any `json:"parameters,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parameters  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionMoveErrorContext) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionMoveErrorContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A permission to a contract
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionPermission struct {
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
	// Vektor error
	Error ExecutionStepGetResponseDefinitionExecutionEVMTransactionPermissionError `json:"error,required"`
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
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionPermission) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Vektor error
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionPermissionError struct {
	// Error context
	Context ExecutionStepGetResponseDefinitionExecutionEVMTransactionPermissionErrorContext `json:"context,required"`
	// Error message
	Message string `json:"message,required"`
	// Request ID
	RequestID string `json:"request_id,required"`
	// Error resource
	Resource string `json:"resource,required"`
	// ISO8601 Timestamp
	Timestamp Timestamp `json:"timestamp,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context     respjson.Field
		Message     respjson.Field
		RequestID   respjson.Field
		Resource    respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionPermissionError) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionPermissionError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error context
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionPermissionErrorContext struct {
	// Error parameters
	Parameters map[string]any `json:"parameters,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parameters  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionPermissionErrorContext) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionPermissionErrorContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An unwrap of the wrapped native asset
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionUnwrap struct {
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
	// Vektor error
	Error ExecutionStepGetResponseDefinitionExecutionEVMTransactionUnwrapError `json:"error,required"`
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
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionUnwrap) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionUnwrap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Vektor error
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionUnwrapError struct {
	// Error context
	Context ExecutionStepGetResponseDefinitionExecutionEVMTransactionUnwrapErrorContext `json:"context,required"`
	// Error message
	Message string `json:"message,required"`
	// Request ID
	RequestID string `json:"request_id,required"`
	// Error resource
	Resource string `json:"resource,required"`
	// ISO8601 Timestamp
	Timestamp Timestamp `json:"timestamp,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context     respjson.Field
		Message     respjson.Field
		RequestID   respjson.Field
		Resource    respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionUnwrapError) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionUnwrapError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error context
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionUnwrapErrorContext struct {
	// Error parameters
	Parameters map[string]any `json:"parameters,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parameters  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionUnwrapErrorContext) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionUnwrapErrorContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A wrap of the native asset
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionWrap struct {
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
	// Vektor error
	Error ExecutionStepGetResponseDefinitionExecutionEVMTransactionWrapError `json:"error,required"`
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
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionWrap) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionWrap) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Vektor error
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionWrapError struct {
	// Error context
	Context ExecutionStepGetResponseDefinitionExecutionEVMTransactionWrapErrorContext `json:"context,required"`
	// Error message
	Message string `json:"message,required"`
	// Request ID
	RequestID string `json:"request_id,required"`
	// Error resource
	Resource string `json:"resource,required"`
	// ISO8601 Timestamp
	Timestamp Timestamp `json:"timestamp,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context     respjson.Field
		Message     respjson.Field
		RequestID   respjson.Field
		Resource    respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionWrapError) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionWrapError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error context
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionWrapErrorContext struct {
	// Error parameters
	Parameters map[string]any `json:"parameters,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parameters  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionWrapErrorContext) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionWrapErrorContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Selling an asset for another asset
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionSell struct {
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
	// Vektor error
	Error ExecutionStepGetResponseDefinitionExecutionEVMTransactionSellError `json:"error,required"`
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
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionSell) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionSell) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Vektor error
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionSellError struct {
	// Error context
	Context ExecutionStepGetResponseDefinitionExecutionEVMTransactionSellErrorContext `json:"context,required"`
	// Error message
	Message string `json:"message,required"`
	// Request ID
	RequestID string `json:"request_id,required"`
	// Error resource
	Resource string `json:"resource,required"`
	// ISO8601 Timestamp
	Timestamp Timestamp `json:"timestamp,required"`
	// Error type
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Context     respjson.Field
		Message     respjson.Field
		RequestID   respjson.Field
		Resource    respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionSellError) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionSellError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error context
type ExecutionStepGetResponseDefinitionExecutionEVMTransactionSellErrorContext struct {
	// Error parameters
	Parameters map[string]any `json:"parameters,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parameters  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExecutionStepGetResponseDefinitionExecutionEVMTransactionSellErrorContext) RawJSON() string {
	return r.JSON.raw
}
func (r *ExecutionStepGetResponseDefinitionExecutionEVMTransactionSellErrorContext) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of an execution step
type ExecutionStepGetResponseType string

const (
	ExecutionStepGetResponseTypeEVMTransactionApprove           ExecutionStepGetResponseType = "evm_transaction_approve"
	ExecutionStepGetResponseTypeEVMTransactionBorrow            ExecutionStepGetResponseType = "evm_transaction_borrow"
	ExecutionStepGetResponseTypeEVMTransactionBorrowRepay       ExecutionStepGetResponseType = "evm_transaction_borrow_repay"
	ExecutionStepGetResponseTypeEVMTransactionBuy               ExecutionStepGetResponseType = "evm_transaction_buy"
	ExecutionStepGetResponseTypeEVMTransactionLend              ExecutionStepGetResponseType = "evm_transaction_lend"
	ExecutionStepGetResponseTypeEVMTransactionLendSetCollateral ExecutionStepGetResponseType = "evm_transaction_lend_set_collateral"
	ExecutionStepGetResponseTypeEVMTransactionLendWithdraw      ExecutionStepGetResponseType = "evm_transaction_lend_withdraw"
	ExecutionStepGetResponseTypeEVMTransactionMove              ExecutionStepGetResponseType = "evm_transaction_move"
	ExecutionStepGetResponseTypeEVMTransactionPermission        ExecutionStepGetResponseType = "evm_transaction_permission"
	ExecutionStepGetResponseTypeEVMTransactionWrap              ExecutionStepGetResponseType = "evm_transaction_wrap"
	ExecutionStepGetResponseTypeEVMTransactionUnwrap            ExecutionStepGetResponseType = "evm_transaction_unwrap"
	ExecutionStepGetResponseTypeEVMTransactionSell              ExecutionStepGetResponseType = "evm_transaction_sell"
)

type ExecutionStepSignParams struct {
	// A hex string starting with 0x
	SignedPayload HexString `json:"signed_payload,required"`
	paramObj
}

func (r ExecutionStepSignParams) MarshalJSON() (data []byte, err error) {
	type shadow ExecutionStepSignParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExecutionStepSignParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
