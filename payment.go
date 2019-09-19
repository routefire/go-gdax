package gdax

import (
	"fmt"
)

// source: https://mholt.github.io/json-to-go/
type PaymentMethod struct {
	ID            string `json:"id"`
	Type          string `json:"type"`
	Name          string `json:"name"`
	Currency      string `json:"currency"`
	PrimaryBuy    bool   `json:"primary_buy"`
	PrimarySell   bool   `json:"primary_sell"`
	AllowBuy      bool   `json:"allow_buy"`
	AllowSell     bool   `json:"allow_sell"`
	AllowDeposit  bool   `json:"allow_deposit"`
	AllowWithdraw bool   `json:"allow_withdraw"`
	Limits        struct {
		Buy []struct {
			PeriodInDays int `json:"period_in_days"`
			Total        struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"total"`
			Remaining struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"remaining"`
		} `json:"buy"`
		InstantBuy []struct {
			PeriodInDays int `json:"period_in_days"`
			Total        struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"total"`
			Remaining struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"remaining"`
		} `json:"instant_buy"`
		Sell []struct {
			PeriodInDays int `json:"period_in_days"`
			Total        struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"total"`
			Remaining struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"remaining"`
		} `json:"sell"`
		Deposit []struct {
			PeriodInDays int `json:"period_in_days"`
			Total        struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"total"`
			Remaining struct {
				Amount   string `json:"amount"`
				Currency string `json:"currency"`
			} `json:"remaining"`
		} `json:"deposit"`
	} `json:"limits"`
}

type CoinbaseAccount struct {
	ID                     string `json:"id"`
	Name                   string `json:"name"`
	Balance                string `json:"balance"`
	Currency               string `json:"currency"`
	Type                   string `json:"type"`
	Primary                bool   `json:"primary"`
	Active                 bool   `json:"active"`
	WireDepositInformation struct {
		AccountNumber string `json:"account_number"`
		RoutingNumber string `json:"routing_number"`
		BankName      string `json:"bank_name"`
		BankAddress   string `json:"bank_address"`
		BankCountry   struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"bank_country"`
		AccountName    string `json:"account_name"`
		AccountAddress string `json:"account_address"`
		Reference      string `json:"reference"`
	} `json:"wire_deposit_information,omitempty"`
	SepaDepositInformation struct {
		Iban            string `json:"iban"`
		Swift           string `json:"swift"`
		BankName        string `json:"bank_name"`
		BankAddress     string `json:"bank_address"`
		BankCountryName string `json:"bank_country_name"`
		AccountName     string `json:"account_name"`
		AccountAddress  string `json:"account_address"`
		Reference       string `json:"reference"`
	} `json:"sepa_deposit_information,omitempty"`
}

func (c *Client) GetPaymentMethods() ([]PaymentMethod, error){
	url := fmt.Sprintf("/payment-methods")
	var paymentMethods []PaymentMethod

	_, e := c.Request("GET", url, nil, &paymentMethods)

	return paymentMethods, e
}

func (c *Client) GetCoinbaseAccounts() ([]CoinbaseAccount, error) {
	url := fmt.Sprintf("/coinbase-accounts")

	var coinbaseAccounts []CoinbaseAccount

	_, e := c.Request("GET", url, nil, &coinbaseAccounts)

	return coinbaseAccounts, e
}

// map[address:0x70618c76C1217010f476839E4F19775dD03f5396 address_info:map[address:0x70618c76C1217010f476839E4F19775dD03f5396] callback_url:<nil> created_at:2019-09-18T08:23:56Z deposit_uri:ethereum:0x1985365e9f78359a9B6AD760e32412f4a445E862/transfer?address=0x70618c76C1217010f476839E4F19775dD03f5396 exchange_deposit_address:true id:938a70b1-0bb2-51dd-aa76-f759b1e8310a name:New exchange deposit address network:ethereum resource:address resource_path:/v2/accounts/9d0d7e9b-5f98-5773-9067-208a11f7c663/addresses/938a70b1-0bb2-51dd-aa76-f759b1e8310a updated_at:2019-09-18T08:23:56Z uri_scheme:ethereum warning_details:Sending any other digital asset, including Ethereum (ETH), will result in permanent loss. warning_title:Only send Augur (REP) to this address warnings:[map[details:Sending any other digital asset, including Ethereum (ETH), will result in permanent loss. image_url:https://dynamic-assets.coinbase.com/4ff758b0e7221284dd912b1bf8b35b56766a32c1b91307eb562e15e19ca68d253bc7ac83b818bcf237c1d72365fb5dd2994ecf2b437d678d612cf86596a0736c/asset_icons/bc09823603c6f5bafc30291b79302ebdda6eba1c49bf21ad076e0f887a56643e.png title:Only send Augur (REP) to this address]]]
func (c *Client) GetCoinbaseAccountAddress(cbAccountId string) (string, error) {
	url := fmt.Sprintf("/coinbase-accounts/%v/addresses", cbAccountId)

	//var addresses interface{}

	res := make(map[string]interface{})

	_, e := c.Request("POST", url, nil, &res)

	return fmt.Sprintf("%v", res["address"]), e

	//maps.GetValuesKV(addresses)
	//return addresses, e
}
