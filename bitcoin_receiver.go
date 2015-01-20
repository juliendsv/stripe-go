package stripe

import "encoding/json"

// BitcoinReceiverListParams is the set of parameters that can be used when listing BitcoinReceivers.
// For more details see https://stripe.com/docs/api/#list_bitcoin_receivers.
type BitcoinReceiverListParams struct {
	ListParams
	NotFilled  bool
	NotActive  bool
	Uncaptured bool
}

// BitcoinReceiverParams is the set of parameters that can be used when creating a BitcoinReceiver.
// For more details see https://stripe.com/docs/api/#create_bitcoin_receiver.
type BitcoinReceiverParams struct {
	Params
	Amount   uint64
	Currency Currency
	Desc     string
	Email    string
	Meta     map[string]string `json:"metadata"`
}

// BitcoinReceiver is the resource representing a Stripe bitcoin receiver.
// For more details see https://stripe.com/docs/api/#bitcoin_receivers
type BitcoinReceiver struct {
	ID                    string                  `json:"id"`
	Created               int64                   `json:"created"`
	Currency              Currency                `json:"currency"`
	Amount                uint64                  `json:"amount"`
	AmountReceived        uint64                  `json:"amount_received"`
	BitcoinAmount         uint64                  `json:"bitcoin_amount"`
	BitcoinAmountReceived uint64                  `json:"bitcoin_amount_received"`
	Filled                bool                    `json:"filled"`
	Active                bool                    `json:"active"`
	RejectTransactions    bool                    `json:"reject_transactions"`
	Desc                  string                  `json:"description"`
	InboundAddress        string                  `json:"inbound_address"`
	RefundAddress         string                  `json:"refund_address"`
	BitcoinUri            string                  `json:"bitcoin_uri"`
	Meta                  map[string]string       `json:"metadata"`
	Email                 string                  `json:"email"`
	Payment               string                  `json:"payment"`
	Customer              string                  `json:"customer"`
	Transactions          *BitcoinTransactionList `json:"transactions"`
}

func (br BitcoinReceiver) PaymentType() string {
	return "bitcoin_receiver"
}

// UnmarshalJSON handles deserialization of a BitcoinReceiver.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (br *BitcoinReceiver) UnmarshalJSON(data []byte) error {
	type bitcoinReceiver BitcoinReceiver
	var r bitcoinReceiver
	err := json.Unmarshal(data, &r)
	if err == nil {
		*br = BitcoinReceiver(r)
	} else {
		// the id is surrounded by "\" characters, so strip them
		br.ID = string(data[1 : len(data)-1])
	}

	return nil
}
