package etrade

import (
	"context"
	"fmt"
)

const (
	listAccountsPath   = "/v1/accounts/list"
	balanceAccountPath = "/v1/accounts/%s/balance?instType=BROKERAGE&realTimeNAV=true"
)

type accountClient struct {
	apiUrl string
	httpClientSource
}

// https://apisb.etrade.com/docs/api/account/api-account-v1.html#/definitions/Account
type Account struct {
	AccountIdKey string `json:"accountIdKey"`
	AccountId    string `json:"accountId"`
}

// https://apisb.etrade.com/docs/api/account/api-account-v1.html#/definitions/AccountListResponse
type AccountListResponse struct {
	Accounts []Account `json:"accounts"`
}

type AccountBalanceRequest struct {
	AccountIdKey string `json:"-"`
}

// https://apisb.etrade.com/docs/api/account/api-balance-v1.html#/definitions/ComputedBalance
type ComputedBalance struct {
	CashAvailableForInvestment     float64 `json:"cashAvailableForInvestment"`
	CashAvailableForWithdrawal     float64 `json:"cashAvailableForWithdrawal"`
	TotalAvailableForWithdrawal    float64 `json:"totalAvailableForWithdrawal"`
	NetCash                        float64 `json:"netCash"`
	CashBalance                    float64 `json:"cashBalance"`
	SettledCashForInvestment       float64 `json:"settledCashForInvestment"`
	UnSettledCashForInvestment     float64 `json:"unSettledCashForInvestment"`
	FundsWithheldFromPurchasePower float64 `json:"fundsWithheldFromPurchasePower"`
	FundsWithheldFromWithdrawal    float64 `json:"fundsWithheldFromWithdrawal"`
	MarginBuyingPower              float64 `json:"marginBuyingPower"`
	CashBuyingPower                float64 `json:"cashBuyingPower"`
	DtMarginBuyingPower            float64 `json:"dtMarginBuyingPower"`
	DtCashBuyingPower              float64 `json:"dtCashBuyingPower"`
	MarginBalance                  float64 `json:"marginBalance"`
	ShortAdjustBalance             float64 `json:"shortAdjustBalance"`
	RegtEquity                     float64 `json:"regtEquity"`
	RegtEquityPercent              float64 `json:"regtEquityPercent"`
	AccountBalance                 float64 `json:"accountBalance"`
}

// https://apisb.etrade.com/docs/api/account/api-balance-v1.html#/definitions/BalanceResponse
type BalanceResponse struct {
	AccountId       string          `json:"accountId"`
	ComputedBalance ComputedBalance `json:"computedBalance"`
}

// https://apisb.etrade.com/docs/api/account/api-account-v1.html
func (a *accountClient) List(ctx context.Context) (AccountListResponse, error) {
	return do[AccountListResponse](ctx, a.httpClientSource, "GET", a.apiUrl+listAccountsPath, nil)
}

// https://apisb.etrade.com/docs/api/account/api-balance-v1.html
func (a *accountClient) Balance(ctx context.Context, input AccountBalanceRequest) (BalanceResponse, error) {
	return do[BalanceResponse](ctx, a.httpClientSource, "GET", a.apiUrl+fmt.Sprintf(balanceAccountPath, input.AccountIdKey), nil)
}
