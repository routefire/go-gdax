package gdax

import "fmt"

type ConversionRequest struct {
	From   string `json:"from"` // USDC, USD
	To     string `json:"to"` // USDC, USD
	Amount string `json:"amount"`
}

type ConversionResponse struct {
	ID            string `json:"id"`
	Amount        string `json:"amount"`
	FromAccountID string `json:"from_account_id"`
	ToAccountID   string `json:"to_account_id"`
	From          string `json:"from"`
	To            string `json:"to"`
}

func (c *Client) Conversion(newConversion ConversionRequest) (ConversionResponse, error) {
	var response ConversionResponse

	url := fmt.Sprintf("/conversions")
	_, err := c.Request("POST", url, newConversion, &response)
	return response, err
}