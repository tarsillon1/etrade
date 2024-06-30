package etrade

import (
	"context"
	"fmt"
)

const (
	orderPlacePath   = "/v1/accounts/%s/orders/place"
	orderPreviewPath = "/v1/accounts/%s/orders/preview"
	orderListPath    = "/v1/accounts/%s/orders?marker=%s"
)

type MfTransaction string

const (
	MfTransactionBuy  MfTransaction = "BUY"
	MfTransactionSell MfTransaction = "SELL"
)

type Currency string

const (
	USD Currency = "USD"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
	HKD Currency = "HKD"
	JPY Currency = "JPY"
	CAD Currency = "CAD"
)

type QuantityType string

const (
	Quantity QuantityType = "QUANTITY"
	Dollar   QuantityType = "DOLLAR"
	AllIOwn  QuantityType = "ALL_I_OWN"
)

type OrderAction string

const (
	Buy        OrderAction = "BUY"
	Sell       OrderAction = "SELL"
	BuyToCover OrderAction = "BUY_TO_COVER"
	SellShort  OrderAction = "SELL_SHORT"
	BuyOpen    OrderAction = "BUY_OPEN"
	BuyClose   OrderAction = "BUY_CLOSE"
	SellOpen   OrderAction = "SELL_OPEN"
	SellClose  OrderAction = "SELL_CLOSE"
	Exchange   OrderAction = "EXCHANGE"
)

type ProductTypeCode string

const (
	Equity                    ProductTypeCode = "EQUITY"
	Option                    ProductTypeCode = "OPTION"
	MutualFund                ProductTypeCode = "MUTUAL_FUND"
	Index                     ProductTypeCode = "INDEX"
	MoneyMarketFund           ProductTypeCode = "MONEY_MARKET_FUND"
	Bond                      ProductTypeCode = "BOND"
	Unknown                   ProductTypeCode = "UNKNOWN"
	Wildcard                  ProductTypeCode = "WILDCARD"
	Move                      ProductTypeCode = "MOVE"
	Etf                       ProductTypeCode = "ETF"
	EquityOptionEtf           ProductTypeCode = "EQUITY_OPTION_ETF"
	EquityEtf                 ProductTypeCode = "EQUITY_ETF"
	ClosedEndFund             ProductTypeCode = "CLOSED_END_FUND"
	Preferred                 ProductTypeCode = "PREFERRED"
	EquityOptn                ProductTypeCode = "EQUITY_OPTN"
	ExchangeTradedFund        ProductTypeCode = "EXCHANGE_TRADED_FUND"
	MutualFundMoneyMarketFund ProductTypeCode = "MUTUAL_FUND_MONEY_MARKET_FUND"
)

type CallPut string

const (
	Call CallPut = "CALL"
	Put  CallPut = "PUT"
)

type SecurityType string

const (
	SecurityTypeEq   SecurityType = "EQ"
	SecurityTypeOptn SecurityType = "OPTN"
	SecurityTypeIndx SecurityType = "INDX"
	SecurityTypeMf   SecurityType = "MF"
	SecurityTypeMmf  SecurityType = "MMF"
)

type ReInvestOption string

const (
	Reinvest       ReInvestOption = "REINVEST"
	Deposit        ReInvestOption = "DEPOSIT"
	CurrentHolding ReInvestOption = "CURRENT_HOLDING"
)

type PositionQuantity string

const (
	PositionQuantityEntirePosition PositionQuantity = "ENTIRE_POSITION"
	PositionQuantityCash           PositionQuantity = "CASH"
	PositionQuantityMargin         PositionQuantity = "MARGIN"
)

type ConditionFollowPrice string

const (
	Ask  ConditionFollowPrice = "ASK"
	Bid  ConditionFollowPrice = "BID"
	Last ConditionFollowPrice = "LAST"
)

type ConditionType string

const (
	ContingentGte ConditionType = "CONTINGENT_GTE"
	ContingentLte ConditionType = "CONTINGENT_LTE"
)

type RoutingDestination string

const (
	Auto RoutingDestination = "AUTO"
	Amex RoutingDestination = "AMEX"
	Box  RoutingDestination = "BOX"
	Cboe RoutingDestination = "CBOE"
	Ise  RoutingDestination = "ISE"
	Nom  RoutingDestination = "NOM"
	Nyse RoutingDestination = "NYSE"
	Phx  RoutingDestination = "PHX"
)

type MarketSession string

const (
	Regular  MarketSession = "REGULAR"
	Extended MarketSession = "EXTENDED"
)

type OffsetType string

const (
	OffsetTypeTrailingStopCnst OffsetType = "TRAILING_STOP_CNST"
	OffsetTypeTrailingStopPrct OffsetType = "TRAILING_STOP_PRCT"
)

type PriceType string

const (
	Market                         PriceType = "MARKET"
	Limit                          PriceType = "LIMIT"
	Stop                           PriceType = "STOP"
	StopLimit                      PriceType = "STOP_LIMIT"
	TrailingStopCnstByLowerTrigger PriceType = "TRAILING_STOP_CNST_BY_LOWER_TRIGGER"
	UpperTriggerByTrailingStopCnst PriceType = "UPPER_TRIGGER_BY_TRAILING_STOP_CNST"
	TrailingStopPrctByLowerTrigger PriceType = "TRAILING_STOP_PRCT_BY_LOWER_TRIGGER"
	UpperTriggerByTrailingStopPrct PriceType = "UPPER_TRIGGER_BY_TRAILING_STOP_PRCT"
	TrailingStopCnst               PriceType = "TRAILING_STOP_CNST"
	TrailingStopPrct               PriceType = "TRAILING_STOP_PRCT"
	HiddenStop                     PriceType = "HIDDEN_STOP"
	HiddenStopByLowerTrigger       PriceType = "HIDDEN_STOP_BY_LOWER_TRIGGER"
	UpperTriggerByHiddenStop       PriceType = "UPPER_TRIGGER_BY_HIDDEN_STOP"
	NetDebit                       PriceType = "NET_DEBIT"
	NetCredit                      PriceType = "NET_CREDIT"
	NetEven                        PriceType = "NET_EVEN"
	MarketOnOpen                   PriceType = "MARKET_ON_OPEN"
	MarketOnClose                  PriceType = "MARKET_ON_CLOSE"
	LimitOnOpen                    PriceType = "LIMIT_ON_OPEN"
	LimitOnClose                   PriceType = "LIMIT_ON_CLOSE"
)

type OrderTerm string

const (
	GoodUntilCancel   OrderTerm = "GOOD_UNTIL_CANCEL"
	GoodForDay        OrderTerm = "GOOD_FOR_DAY"
	GoodTillDate      OrderTerm = "GOOD_TILL_DATE"
	ImmediateOrCancel OrderTerm = "IMMEDIATE_OR_CANCEL"
	FillOrKill        OrderTerm = "FILL_OR_KILL"
)

type OrderType string

const (
	Eq            OrderType = "EQ"
	Optn          OrderType = "OPTN"
	Spreads       OrderType = "SPREADS"
	BuyWrites     OrderType = "BUY_WRITES"
	Butterfly     OrderType = "BUTTERFLY"
	IronButterfly OrderType = "IRON_BUTTERFLY"
	Condor        OrderType = "CONDOR"
	IronCondor    OrderType = "IRON_CONDOR"
	Mf            OrderType = "MF"
	Mmf           OrderType = "MMF"
)

type CashMargin string

const (
	Cash   CashMargin = "CASH"
	Margin CashMargin = "MARGIN"
)

type OrderStatus string

const (
	Open            OrderStatus = "OPEN"
	Executed        OrderStatus = "EXECUTED"
	Cancelled       OrderStatus = "CANCELLED"
	IndividualFills OrderStatus = "INDIVIDUAL_FILLS"
	CancelRequested OrderStatus = "CANCEL_REQUESTED"
	Expired         OrderStatus = "EXPIRED"
	Rejected        OrderStatus = "REJECTED"
)

type orderClient struct {
	apiUrl string
	httpClientSource
}

// https://apisb.etrade.com/docs/api/order/api-order-v1.html#/definitions/PreviewId
type PreviewId struct {
	PreviewId  int64      `json:"previewId"`
	CashMargin CashMargin `json:"cashMargin"`
}

type ProductId struct {
	Symbol   string          `json:"symbol"`
	TypeCode ProductTypeCode `json:"typeCode"`
}

// https://apisb.etrade.com/docs/api/order/api-order-v1.html#/definitions/Product
type Product struct {
	Symbol       string       `json:"symbol"`
	SecurityType SecurityType `json:"securityType"`
	CallPut      CallPut      `json:"callPut"`
	ExpiryYear   int32        `json:"expiryYear"`
	ExpiryMonth  int32        `json:"expiryMonth"`
	ExpiryDay    int32        `json:"expiryDay"`
	StrikePrice  float64      `json:"strikePrice"`
	ExpiryType   string       `json:"expiryType"`
	ProductId    ProductId    `json:"productId"`
}

// https://apisb.etrade.com/docs/api/order/api-order-v1.html#/definitions/Instrument
type Instrument struct {
	Product               Product       `json:"product"`
	SymbolDescription     string        `json:"symbolDescription"`
	OrderAction           OrderAction   `json:"orderAction"`
	QuantityType          QuantityType  `json:"quantityType"`
	Quantity              float64       `json:"quantity"`
	CancelQuantity        float64       `json:"cancelQuantity"`
	OrderedQuantity       float64       `json:"orderedQuantity"`
	FilledQuantity        float64       `json:"filledQuantity"`
	AverageExecutionPrice float64       `json:"averageExecutionPrice"`
	EstimatedCommission   float64       `json:"estimatedCommission"`
	EstimatedFees         float64       `json:"estimatedFees"`
	Bid                   float64       `json:"bid"`
	Ask                   float64       `json:"ask"`
	LastPrice             float64       `json:"lastprice"`
	Currency              Currency      `json:"currency"`
	OsiKey                string        `json:"osiKey"`
	MfTransaction         MfTransaction `json:"mfTransaction"`
	ReserveOrder          bool          `json:"reserveOrder"`
	ReserveQuantity       float64       `json:"reserveQuantity"`
}

// https://apisb.etrade.com/docs/api/order/api-order-v1.html#/definitions/OrderDetail
type OrderDetail struct {
	OrderNumber           int                  `json:"orderNumber"`
	AccountId             string               `json:"accountId"`
	PreviewTime           int64                `json:"previewTime"`
	PlacedTime            int64                `json:"placedTime"`
	ExecutedTime          int64                `json:"executedTime"`
	OrderValue            float64              `json:"orderValue"`
	Status                OrderStatus          `json:"status"`
	OrderType             OrderType            `json:"orderType"`
	OrderTerm             OrderTerm            `json:"orderTerm"`
	PriceType             PriceType            `json:"priceType"`
	PriceValue            string               `json:"priceValue"`
	LimitPrice            float64              `json:"limitPrice"`
	StopPrice             float64              `json:"stopPrice"`
	StopLimitPrice        float64              `json:"stopLimitPrice"`
	OffsetType            OffsetType           `json:"offsetType"`
	OffsetValue           float64              `json:"offsetValue"`
	MarketSession         MarketSession        `json:"marketSession"`
	RoutingDestination    RoutingDestination   `json:"routingDestination"`
	BracketedLimitPrice   float64              `json:"bracketedLimitPrice"`
	InitialStopPrice      float64              `json:"initialStopPrice"`
	TrailPrice            float64              `json:"trailPrice"`
	TriggerPrice          float64              `json:"triggerPrice"`
	ConditionPrice        float64              `json:"conditionPrice"`
	ConditionSymbol       string               `json:"conditionSymbol"`
	ConditionType         ConditionType        `json:"ConditionType"`
	ConditionFollowPrice  ConditionFollowPrice `json:"ConditionFollowPrice"`
	ConditionSecurityType string               `json:"conditionSecurityType"`
	ReplacedByOrderId     int                  `json:"replacedByOrderId"`
	ReplacesOrderId       int                  `json:"replacesOrderId"`
	AllOrNone             bool                 `json:"allOrNone"`
	PreviewId             int64                `json:"previewId"`
	InvestmentAmount      float64              `json:"investmentAmount"`
	PositionQuantity      PositionQuantity     `json:"positionQuantity"`
	AipFlag               bool                 `json:"aipFlag"`
	ReInvestOption        ReInvestOption       `json:"reInvestOption"`
	EstimatedCommission   float64              `json:"estimatedCommission"`
	EstimatedFees         float64              `json:"estimatedFees"`
	EstimatedTotalAmount  float64              `json:"estimatedTotalAmount"`
	NetPrice              float64              `json:"netPrice"`
	NetBid                float64              `json:"netBid"`
	NetAsk                float64              `json:"netAsk"`
	Gcd                   int32                `json:"gcd"`
	Ratio                 string               `json:"ratio"`
	MfpriceType           string               `json:"mfpriceType"`
	Instrument            []Instrument         `json:"instrument"`
}

// https://apisb.etrade.com/docs/api/order/api-order-v1.html#/definitions/PlaceOrderRequest
type PlaceOrderRequest struct {
	AccountIdKey  string        `json:"-"`
	OrderType     string        `json:"orderType"`
	ClientOrderId string        `json:"clientOrderId"`
	Order         []OrderDetail `json:"order"`
	PreviewIds    []PreviewId   `json:"previewIds"`
}

// https://apisb.etrade.com/docs/api/order/api-order-v1.html#/definitions/OrderId
type OrderId struct {
	OrderId    int64      `json:"orderId"`
	CashMargin CashMargin `json:"cashMargin"`
}

// https://apisb.etrade.com/docs/api/order/api-order-v1.html#/definitions/PlaceOrderResponse
type PlaceOrderResponse struct {
	OrderType       OrderType     `json:"orderType"`
	TotalOrderValue float64       `json:"totalOrderValue"`
	TotalCommission float64       `json:"totalCommission"`
	OrderId         int64         `json:"orderId"`
	Order           []OrderDetail `json:"order"`
	DstFlag         bool          `json:"dstFlag"`
	OptionLevelCd   int32         `json:"optionLevelCd"`
	OrderIds        []OrderId     `json:"orderIds"`
	PlacedTime      int64         `json:"placedTime"`
	AccountId       string        `json:"accountId"`
	ClientOrderId   string        `json:"clientOrderId"`
}

// https://apisb.etrade.com/docs/api/order/api-order-v1.html#/definitions/Order
type Order struct {
	OrderId         int64         `json:"orderId"`
	Details         string        `json:"details"`
	OrderType       OrderType     `json:"orderType"`
	TotalOrderValue float64       `json:"totalOrderValue"`
	TotalCommission float64       `json:"totalCommission"`
	OrderDetail     []OrderDetail `json:"orderDetail"`
}

// https://apisb.etrade.com/docs/api/order/api-order-v1.html#/definition/getOrders
type OrdersRequest struct {
	AccountIdKey string `json:"-"`
	Marker       string `json:"-"`
}

type OrdersResponse struct {
	Marker string  `json:"marker"`
	Next   string  `json:"next"`
	Order  []Order `json:"order"`
}

// https://apisb.etrade.com/docs/api/order/api-order-v1.html#/definitions/PreviewOrderRequest
type PreviewOrderRequest struct {
	AccountIdKey  string    `json:"-"`
	OrderType     OrderType `json:"orderType"`
	Order         []Order   `json:"order"`
	ClientOrderId string    `json:"clientOrderId"`
}

// https://apisb.etrade.com/docs/api/order/api-order-v1.html#/definitions/PreviewOrderResponse
type PreviewOrderResponse struct {
	OrderType       OrderType     `json:"orderType"`
	TotalOrderValue float64       `json:"totalOrderValue"`
	TotalCommission float64       `json:"totalCommission"`
	OrderId         int64         `json:"orderId"`
	Order           []OrderDetail `json:"order"`
	DstFlag         bool          `json:"dstFlag"`
	OptionLevelCd   int32         `json:"optionLevelCd"`
	PreviewIds      []PreviewId   `json:"previewIds"`
	PreviewTime     int64         `json:"previewTime"`
	AccountId       string        `json:"accountId"`
	ClientOrderId   string        `json:"clientOrderId"`
}

// Place is used to submit an order after it has been successfully previewed.
func (o *orderClient) Place(ctx context.Context, input PlaceOrderRequest) (PlaceOrderResponse, error) {
	return do[PlaceOrderResponse](ctx, o.httpClientSource, "POST", o.apiUrl+fmt.Sprintf(orderPlacePath, input.AccountIdKey), input)
}

// Preview is used to submit an order request for preview before placing it.
func (o *orderClient) Preview(ctx context.Context, input PreviewOrderRequest) (PreviewOrderResponse, error) {
	return do[PreviewOrderResponse](ctx, o.httpClientSource, "POST", o.apiUrl+fmt.Sprintf(orderPreviewPath, input.AccountIdKey), input)
}

// https://apisb.etrade.com/docs/api/order/api-order-v1.html#/definition/getOrders
func (o *orderClient) List(ctx context.Context, input OrdersRequest) (OrdersResponse, error) {
	return do[OrdersResponse](ctx, o.httpClientSource, "GET", o.apiUrl+fmt.Sprintf(orderListPath, input.AccountIdKey, input.Marker), nil)
}
