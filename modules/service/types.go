package service

import (
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	sdk "github.com/irisnet/irishub-sdk-go/types"
	"github.com/irisnet/irishub-sdk-go/utils/json"
)

var (
	_ sdk.Msg = MsgDefineService{}
	_ sdk.Msg = MsgBindService{}
	_ sdk.Msg = MsgUpdateServiceBinding{}
	_ sdk.Msg = MsgSetWithdrawAddress{}
	_ sdk.Msg = MsgDisableService{}
	_ sdk.Msg = MsgEnableService{}
	_ sdk.Msg = MsgRefundServiceDeposit{}
	_ sdk.Msg = MsgRequestService{}
	_ sdk.Msg = MsgRespondService{}
	_ sdk.Msg = MsgPauseRequestContext{}
	_ sdk.Msg = MsgStartRequestContext{}
	_ sdk.Msg = MsgKillRequestContext{}
	_ sdk.Msg = MsgUpdateRequestContext{}
	_ sdk.Msg = MsgWithdrawEarnedFees{}
	_ sdk.Msg = MsgWithdrawTax{}

	TagServiceName      = "service-name"
	TagProvider         = "provider"
	TagConsumer         = "consumer"
	TagRequestID        = "request-id"
	TagRespondService   = "respond_service"
	TagRequestContextID = "request-context-id"

	RequestContextIDLen = 32 // length of the request context ID in bytes

	cdc = sdk.NewAminoCodec()
)

func init() {
	RegisterCodec(cdc)
}

// MsgDefineService defines a message to define a service
type MsgDefineService struct {
	Name              string         `json:"name"`
	Description       string         `json:"description"`
	Tags              []string       `json:"tags"`
	Author            sdk.AccAddress `json:"author"`
	AuthorDescription string         `json:"author_description"`
	Schemas           string         `json:"schemas"`
}

func (msg MsgDefineService) Type() string {
	return "define_service"
}

func (msg MsgDefineService) ValidateBasic() error {
	if len(msg.Author) == 0 {
		return errors.New("author missing")
	}

	if len(msg.Name) == 0 {
		return errors.New("author missing")
	}

	if len(msg.Schemas) == 0 {
		return errors.New("schemas missing")
	}

	return nil
}

func (msg MsgDefineService) GetSignBytes() []byte {
	if len(msg.Tags) == 0 {
		msg.Tags = nil
	}

	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}

	return json.MustSort(b)
}

func (msg MsgDefineService) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Author}
}

// MsgBindService defines a message to bind a service
type MsgBindService struct {
	ServiceName     string         `json:"service_name"`
	Provider        sdk.AccAddress `json:"provider"`
	Deposit         sdk.Coins      `json:"deposit"`
	Pricing         string         `json:"pricing"`
	WithdrawAddress sdk.AccAddress `json:"withdraw_address"`
}

func (msg MsgBindService) Type() string {
	return "bind_service"
}

func (msg MsgBindService) ValidateBasic() error {
	if len(msg.Provider) == 0 {
		return errors.New("provider missing")
	}

	if len(msg.ServiceName) == 0 {
		return errors.New("serviceName missing")
	}

	if len(msg.Pricing) == 0 {
		return errors.New("pricing missing")
	}
	return nil
}

func (msg MsgBindService) GetSignBytes() []byte {
	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}

	return json.MustSort(b)
}

func (msg MsgBindService) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Provider}
}

// MsgRequestService defines a message to request a service
type MsgRequestService struct {
	ServiceName       string           `json:"service_name"`
	Providers         []sdk.AccAddress `json:"providers"`
	Consumer          sdk.AccAddress   `json:"consumer"`
	Input             string           `json:"input"`
	ServiceFeeCap     sdk.Coins        `json:"service_fee_cap"`
	Timeout           int64            `json:"timeout"`
	SuperMode         bool             `json:"super_mode"`
	Repeated          bool             `json:"repeated"`
	RepeatedFrequency uint64           `json:"repeated_frequency"`
	RepeatedTotal     int64            `json:"repeated_total"`
}

func (msg MsgRequestService) Type() string {
	return "request_service"
}

func (msg MsgRequestService) ValidateBasic() error {
	if len(msg.Consumer) == 0 {
		return errors.New("consumer missing")
	}
	if len(msg.Providers) == 0 {
		return errors.New("providers missing")
	}

	if len(msg.ServiceName) == 0 {
		return errors.New("serviceName missing")
	}

	if len(msg.Input) == 0 {
		return errors.New("input missing")
	}
	return nil
}

func (msg MsgRequestService) GetSignBytes() []byte {
	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}

	return json.MustSort(b)
}

func (msg MsgRequestService) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Consumer}
}

// MsgRespondService defines a message to respond a service request
type MsgRespondService struct {
	RequestID string         `json:"request_id"`
	Provider  sdk.AccAddress `json:"provider"`
	Output    string         `json:"output"`
	Error     string         `json:"error"`
}

func (msg MsgRespondService) Type() string {
	return "respond_service"
}

func (msg MsgRespondService) ValidateBasic() error {
	if len(msg.Provider) == 0 {
		return errors.New("provider missing")
	}

	if len(msg.Output) == 0 && len(msg.Error) == 0 {
		return errors.New("either output or error should be specified, but neither was provided")
	}

	if len(msg.Output) > 0 && len(msg.Error) > 0 {
		return errors.New("either output or error should be specified, but both were provided")
	}
	return nil
}

func (msg MsgRespondService) GetSignBytes() []byte {
	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}

	return json.MustSort(b)
}

func (msg MsgRespondService) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Provider}
}

//______________________________________________________________________

// MsgUpdateServiceBinding defines a message to update a service binding
type MsgUpdateServiceBinding struct {
	ServiceName string         `json:"service_name"`
	Provider    sdk.AccAddress `json:"provider"`
	Deposit     sdk.Coins      `json:"deposit"`
	Pricing     string         `json:"pricing"`
}

// Type implements Msg.
func (msg MsgUpdateServiceBinding) Type() string { return "update_service_binding" }

// GetSignBytes implements Msg.
func (msg MsgUpdateServiceBinding) GetSignBytes() []byte {
	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}

	return json.MustSort(b)
}

// ValidateBasic implements Msg.
func (msg MsgUpdateServiceBinding) ValidateBasic() error {
	if len(msg.Provider) == 0 {
		return errors.New("provider missing")
	}

	if len(msg.ServiceName) == 0 {
		return errors.New("service name missing")
	}

	if !msg.Deposit.Empty() {
		return errors.New(fmt.Sprintf("invalid deposit: %s", msg.Deposit))
	}

	return nil
}

// GetSigners implements Msg.
func (msg MsgUpdateServiceBinding) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Provider}
}

//______________________________________________________________________

// MsgSetWithdrawAddress defines a message to set the withdrawal address for a service binding
type MsgSetWithdrawAddress struct {
	ServiceName     string         `json:"service_name"`
	Provider        sdk.AccAddress `json:"provider"`
	WithdrawAddress sdk.AccAddress `json:"withdraw_address"`
}

// Type implements Msg.
func (msg MsgSetWithdrawAddress) Type() string { return "set_withdraw_address" }

// GetSignBytes implements Msg.
func (msg MsgSetWithdrawAddress) GetSignBytes() []byte {
	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}

	return json.MustSort(b)
}

// ValidateBasic implements Msg.
func (msg MsgSetWithdrawAddress) ValidateBasic() error {
	if len(msg.Provider) == 0 {
		return errors.New("provider missing")
	}

	if len(msg.ServiceName) == 0 {
		return errors.New("service name missing")
	}

	if len(msg.WithdrawAddress) == 0 {
		return errors.New("withdrawal address missing")
	}

	return nil
}

// GetSigners implements Msg.
func (msg MsgSetWithdrawAddress) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Provider}
}

//______________________________________________________________________

// MsgDisableService defines a message to disable a service binding
type MsgDisableService struct {
	ServiceName string         `json:"service_name"`
	Provider    sdk.AccAddress `json:"provider"`
}

// Type implements Msg.
func (msg MsgDisableService) Type() string { return "disable_service" }

// GetSignBytes implements Msg.
func (msg MsgDisableService) GetSignBytes() []byte {
	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}

	return json.MustSort(b)
}

// ValidateBasic implements Msg.
func (msg MsgDisableService) ValidateBasic() error {
	if len(msg.Provider) == 0 {
		return errors.New("provider missing")
	}

	if len(msg.ServiceName) == 0 {
		return errors.New("service name missing")
	}

	return nil
}

// GetSigners implements Msg.
func (msg MsgDisableService) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Provider}
}

//______________________________________________________________________

// MsgEnableService defines a message to enable a service binding
type MsgEnableService struct {
	ServiceName string         `json:"service_name"`
	Provider    sdk.AccAddress `json:"provider"`
	Deposit     sdk.Coins      `json:"deposit"`
}

// Type implements Msg.
func (msg MsgEnableService) Type() string { return "enable_service" }

// GetSignBytes implements Msg.
func (msg MsgEnableService) GetSignBytes() []byte {
	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}

	return json.MustSort(b)
}

// ValidateBasic implements Msg.
func (msg MsgEnableService) ValidateBasic() error {
	if len(msg.Provider) == 0 {
		return errors.New("provider missing")
	}

	if len(msg.ServiceName) == 0 {
		return errors.New("service name missing")
	}

	if !msg.Deposit.Empty() {
		return errors.New(fmt.Sprintf("invalid deposit: %s", msg.Deposit))
	}

	return nil
}

// GetSigners implements Msg.
func (msg MsgEnableService) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Provider}
}

//______________________________________________________________________

// MsgRefundServiceDeposit defines a message to refund deposit from a service binding
type MsgRefundServiceDeposit struct {
	ServiceName string         `json:"service_name"`
	Provider    sdk.AccAddress `json:"provider"`
}

// Type implements Msg.
func (msg MsgRefundServiceDeposit) Type() string { return "refund_service_deposit" }

// GetSignBytes implements Msg.
func (msg MsgRefundServiceDeposit) GetSignBytes() []byte {
	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}

	return json.MustSort(b)
}

// ValidateBasic implements Msg.
func (msg MsgRefundServiceDeposit) ValidateBasic() error {
	if len(msg.Provider) == 0 {
		return errors.New("provider missing")
	}

	if len(msg.ServiceName) == 0 {
		return errors.New("service name missing")
	}

	return nil
}

// GetSigners implements Msg.
func (msg MsgRefundServiceDeposit) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Provider}
}

//______________________________________________________________________

// MsgPauseRequestContext defines a message to suspend a request context
type MsgPauseRequestContext struct {
	RequestContextID []byte         `json:"request_context_id"`
	Consumer         sdk.AccAddress `json:"consumer"`
}

// Type implements Msg.
func (msg MsgPauseRequestContext) Type() string { return "pause_request_context" }

// GetSignBytes implements Msg.
func (msg MsgPauseRequestContext) GetSignBytes() []byte {
	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}

	return json.MustSort(b)
}

// ValidateBasic implements Msg.
func (msg MsgPauseRequestContext) ValidateBasic() error {
	if len(msg.Consumer) == 0 {
		return errors.New("consumer missing")
	}

	if len(msg.RequestContextID) != RequestContextIDLen {
		return errors.New(fmt.Sprintf("length of the request context ID must be %d in bytes", RequestContextIDLen))
	}

	return nil
}

// GetSigners implements Msg.
func (msg MsgPauseRequestContext) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Consumer}
}

//______________________________________________________________________

// MsgStartRequestContext defines a message to resume a request context
type MsgStartRequestContext struct {
	RequestContextID []byte         `json:"request_context_id"`
	Consumer         sdk.AccAddress `json:"consumer"`
}

// Type implements Msg.
func (msg MsgStartRequestContext) Type() string { return "start_request_context" }

// GetSignBytes implements Msg.
func (msg MsgStartRequestContext) GetSignBytes() []byte {
	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}

	return json.MustSort(b)
}

// ValidateBasic implements Msg.
func (msg MsgStartRequestContext) ValidateBasic() error {
	if len(msg.Consumer) == 0 {
		return errors.New("consumer missing")
	}

	if len(msg.RequestContextID) != RequestContextIDLen {
		return errors.New(fmt.Sprintf("length of the request context ID must be %d in bytes", RequestContextIDLen))
	}

	return nil
}

// GetSigners implements Msg.
func (msg MsgStartRequestContext) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Consumer}
}

//______________________________________________________________________

// MsgKillRequestContext defines a message to terminate a request context
type MsgKillRequestContext struct {
	RequestContextID []byte         `json:"request_context_id"`
	Consumer         sdk.AccAddress `json:"consumer"`
}

// Type implements Msg.
func (msg MsgKillRequestContext) Type() string { return "kill_request_context" }

// GetSignBytes implements Msg.
func (msg MsgKillRequestContext) GetSignBytes() []byte {
	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}

	return json.MustSort(b)
}

// ValidateBasic implements Msg.
func (msg MsgKillRequestContext) ValidateBasic() error {
	if len(msg.Consumer) == 0 {
		return errors.New("consumer missing")
	}

	if len(msg.RequestContextID) != RequestContextIDLen {
		return errors.New(fmt.Sprintf("length of the request context ID must be %d in bytes", RequestContextIDLen))
	}

	return nil
}

// GetSigners implements Msg.
func (msg MsgKillRequestContext) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Consumer}
}

//______________________________________________________________________

// MsgUpdateRequestContext defines a message to update a request context
type MsgUpdateRequestContext struct {
	RequestContextID  []byte           `json:"request_context_id"`
	Providers         []sdk.AccAddress `json:"providers"`
	ServiceFeeCap     sdk.Coins        `json:"service_fee_cap"`
	Timeout           int64            `json:"timeout"`
	RepeatedFrequency uint64           `json:"repeated_frequency"`
	RepeatedTotal     int64            `json:"repeated_total"`
	Consumer          sdk.AccAddress   `json:"consumer"`
}

// Type implements Msg.
func (msg MsgUpdateRequestContext) Type() string { return "update_request_context" }

// GetSignBytes implements Msg.
func (msg MsgUpdateRequestContext) GetSignBytes() []byte {
	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}

	return json.MustSort(b)
}

// ValidateBasic implements Msg.
func (msg MsgUpdateRequestContext) ValidateBasic() error {
	if len(msg.Consumer) == 0 {
		return errors.New("consumer missing")
	}

	if len(msg.RequestContextID) != RequestContextIDLen {
		return errors.New(fmt.Sprintf("length of the request context ID must be %d in bytes", RequestContextIDLen))
	}

	return nil
}

// GetSigners implements Msg.
func (msg MsgUpdateRequestContext) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Consumer}
}

//______________________________________________________________________

// MsgWithdrawEarnedFees defines a message to withdraw the fees earned by the provider
type MsgWithdrawEarnedFees struct {
	Provider sdk.AccAddress `json:"provider"`
}

// Type implements Msg.
func (msg MsgWithdrawEarnedFees) Type() string { return "withdraw_earned_fees" }

// GetSignBytes implements Msg.
func (msg MsgWithdrawEarnedFees) GetSignBytes() []byte {
	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}

	return json.MustSort(b)
}

// ValidateBasic implements Msg.
func (msg MsgWithdrawEarnedFees) ValidateBasic() error {
	if len(msg.Provider) == 0 {
		return errors.New("provider missing")
	}

	return nil
}

// GetSigners implements Msg.
func (msg MsgWithdrawEarnedFees) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Provider}
}

//______________________________________________________________________

// MsgWithdrawTax defines a message to withdraw the service tax
type MsgWithdrawTax struct {
	Trustee     sdk.AccAddress `json:"trustee"`
	DestAddress sdk.AccAddress `json:"dest_address"`
	Amount      sdk.Coins      `json:"amount"`
}

// Type implements Msg.
func (msg MsgWithdrawTax) Type() string { return "withdraw_tax" }

// GetSignBytes implements Msg.
func (msg MsgWithdrawTax) GetSignBytes() []byte {
	b, err := cdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}

	return json.MustSort(b)
}

// ValidateBasic implements Msg.
func (msg MsgWithdrawTax) ValidateBasic() error {
	if len(msg.Trustee) == 0 {
		return errors.New("trustee missing")
	}

	if len(msg.DestAddress) == 0 {
		return errors.New("destination address missing")
	}

	return nil
}

// GetSigners implements Msg.
func (msg MsgWithdrawTax) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Trustee}
}

//==========================================for Query==========================================

// Definition represents a service definition
type Definition struct {
	Name              string         `json:"name"`
	Description       string         `json:"description"`
	Tags              []string       `json:"tags"`
	Author            sdk.AccAddress `json:"author"`
	AuthorDescription string         `json:"author_description"`
	Schemas           string         `json:"schemas"`
}

func (r Definition) toSDKDefinition() sdk.ServiceDefinition {
	return sdk.ServiceDefinition{
		Name:              r.Name,
		Description:       r.Description,
		Tags:              r.Tags,
		Author:            r.Author,
		AuthorDescription: r.AuthorDescription,
		Schemas:           r.Schemas,
	}
}

// Binding defines a struct for service binding
type Binding struct {
	ServiceName     string         `json:"service_name"`
	Provider        sdk.AccAddress `json:"provider"`
	Deposit         sdk.Coins      `json:"deposit"`
	Pricing         string         `json:"pricing"`
	WithdrawAddress sdk.AccAddress `json:"withdraw_address"`
	Available       bool           `json:"available"`
	DisabledTime    time.Time      `json:"disabled_time"`
}

func (r Binding) toSDKBinding() sdk.ServiceBinding {
	return sdk.ServiceBinding{
		ServiceName:     r.ServiceName,
		Provider:        r.Provider,
		Deposit:         r.Deposit,
		Pricing:         r.Pricing,
		WithdrawAddress: r.WithdrawAddress,
		Available:       r.Available,
		DisabledTime:    r.DisabledTime,
	}

}

type Bindings []Binding

func (rs Bindings) toSDKBindings() (bindings []sdk.ServiceBinding) {
	for _, binding := range rs {
		bindings = append(bindings, binding.toSDKBinding())
	}
	return
}

// Request defines a request which contains the detailed request data
type Request struct {
	ServiceName                string         `json:"service_name"`
	Provider                   sdk.AccAddress `json:"provider"`
	Consumer                   sdk.AccAddress `json:"consumer"`
	Input                      string         `json:"input"`
	ServiceFee                 sdk.Coins      `json:"service_fee"`
	SuperMode                  bool           `json:"super_mode"`
	RequestHeight              int64          `json:"request_height"`
	ExpirationHeight           int64          `json:"expiration_height"`
	RequestContextID           []byte         `json:"request_context_id"`
	RequestContextBatchCounter uint64         `json:"request_context_batch_counter"`
}

func (r Request) toSDKRequest() sdk.Request {
	return sdk.Request{
		ServiceName:                r.ServiceName,
		Provider:                   r.Provider,
		Consumer:                   r.Consumer,
		Input:                      r.Input,
		ServiceFee:                 r.ServiceFee,
		SuperMode:                  r.SuperMode,
		RequestHeight:              r.RequestHeight,
		ExpirationHeight:           r.ExpirationHeight,
		RequestContextID:           RequestContextIDToString(r.RequestContextID),
		RequestContextBatchCounter: r.RequestContextBatchCounter,
	}
}

type Requests []Request

func (rs Requests) toSDKRequests() (requests []sdk.Request) {
	for _, request := range rs {
		requests = append(requests, request.toSDKRequest())
	}
	return
}

// Response defines a response
type Response struct {
	Provider                   sdk.AccAddress `json:"provider"`
	Consumer                   sdk.AccAddress `json:"consumer"`
	Output                     string         `json:"output"`
	Error                      string         `json:"error"`
	RequestContextID           []byte         `json:"request_context_id"`
	RequestContextBatchCounter uint64         `json:"request_context_batch_counter"`
}

func (r Response) toSDKResponse() sdk.Response {
	return sdk.Response{
		Provider:                   r.Provider,
		Consumer:                   r.Consumer,
		Output:                     r.Output,
		Error:                      r.Error,
		RequestContextID:           RequestContextIDToString(r.RequestContextID),
		RequestContextBatchCounter: r.RequestContextBatchCounter,
	}
}

type Responses []Response

func (rs Responses) toSDKResponses() (responses []sdk.Response) {
	for _, response := range rs {
		responses = append(responses, response.toSDKResponse())
	}
	return
}

// RequestContext defines a context which holds request-related data
type RequestContext struct {
	ServiceName        string           `json:"service_name"`
	Providers          []sdk.AccAddress `json:"providers"`
	Consumer           sdk.AccAddress   `json:"consumer"`
	Input              string           `json:"input"`
	ServiceFeeCap      sdk.Coins        `json:"service_fee_cap"`
	Timeout            int64            `json:"timeout"`
	SuperMode          bool             `json:"super_mode"`
	Repeated           bool             `json:"repeated"`
	RepeatedFrequency  uint64           `json:"repeated_frequency"`
	RepeatedTotal      int64            `json:"repeated_total"`
	BatchCounter       uint64           `json:"batch_counter"`
	BatchRequestCount  uint16           `json:"batch_request_count"`
	BatchResponseCount uint16           `json:"batch_response_count"`
	BatchState         string           `json:"batch_state"`
	State              string           `json:"state"`
	ResponseThreshold  uint16           `json:"response_threshold"`
	ModuleName         string           `json:"module_name"`
}

func (r RequestContext) toSDKRequestContext() sdk.RequestContext {
	return sdk.RequestContext{
		ServiceName:        r.ServiceName,
		Providers:          r.Providers,
		Consumer:           r.Consumer,
		Input:              r.Input,
		ServiceFeeCap:      r.ServiceFeeCap,
		Timeout:            r.Timeout,
		SuperMode:          r.SuperMode,
		Repeated:           r.Repeated,
		RepeatedFrequency:  r.RepeatedFrequency,
		RepeatedTotal:      r.RepeatedTotal,
		BatchCounter:       r.BatchCounter,
		BatchRequestCount:  r.BatchRequestCount,
		BatchResponseCount: r.BatchResponseCount,
		BatchState:         r.BatchState,
		State:              r.State,
		ResponseThreshold:  r.ResponseThreshold,
		ModuleName:         r.ModuleName,
	}
}

// EarnedFees defines a struct for the fees earned by the provider
type EarnedFees struct {
	Address sdk.AccAddress `json:"address"`
	Coins   sdk.Coins      `json:"coins"`
}

func (e EarnedFees) toSDKEarnedFees() sdk.EarnedFees {
	return sdk.EarnedFees{
		Address: e.Address,
		Coins:   e.Coins,
	}
}

func RequestContextIDToString(reqCtxID []byte) string {
	return hex.EncodeToString(reqCtxID)
}

func RequestContextIDToByte(reqCtxID string) []byte {
	dst, err := hex.DecodeString(reqCtxID)
	if err != nil {
		panic(err)
	}
	return dst
}

func RegisterCodec(cdc sdk.Codec) {
	cdc.RegisterConcrete(MsgDefineService{}, "irishub/service/MsgDefineService")
	cdc.RegisterConcrete(MsgBindService{}, "irishub/service/MsgBindService")
	cdc.RegisterConcrete(MsgUpdateServiceBinding{}, "irishub/service/MsgUpdateServiceBinding")
	cdc.RegisterConcrete(MsgSetWithdrawAddress{}, "irishub/service/MsgSetWithdrawAddress")
	cdc.RegisterConcrete(MsgDisableService{}, "irishub/service/MsgDisableService")
	cdc.RegisterConcrete(MsgEnableService{}, "irishub/service/MsgEnableService")
	cdc.RegisterConcrete(MsgRefundServiceDeposit{}, "irishub/service/MsgRefundServiceDeposit")
	cdc.RegisterConcrete(MsgRequestService{}, "irishub/service/MsgRequestService")
	cdc.RegisterConcrete(MsgRespondService{}, "irishub/service/MsgRespondService")
	cdc.RegisterConcrete(MsgPauseRequestContext{}, "irishub/service/MsgPauseRequestContext")
	cdc.RegisterConcrete(MsgStartRequestContext{}, "irishub/service/MsgStartRequestContext")
	cdc.RegisterConcrete(MsgKillRequestContext{}, "irishub/service/MsgKillRequestContext")
	cdc.RegisterConcrete(MsgUpdateRequestContext{}, "irishub/service/MsgUpdateRequestContext")
	cdc.RegisterConcrete(MsgWithdrawEarnedFees{}, "irishub/service/MsgWithdrawEarnedFees")
	cdc.RegisterConcrete(MsgWithdrawTax{}, "irishub/service/MsgWithdrawTax")

	cdc.RegisterConcrete(Definition{}, "irishub/service/ServiceDefinition")
	cdc.RegisterConcrete(Binding{}, "irishub/service/ServiceBinding")
	cdc.RegisterConcrete(RequestContext{}, "irishub/service/RequestContext")
	cdc.RegisterConcrete(Request{}, "irishub/service/Request")
	cdc.RegisterConcrete(Response{}, "irishub/service/Response")
	cdc.RegisterConcrete(EarnedFees{}, "irishub/service/EarnedFees")
}