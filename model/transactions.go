package model

// Transactions is transaction response body
type Transactions struct {
	Type            int    `json:"type"`
	ID              string `json:"id"`
	Sender          string `json:"sender"`
	SenderPublicKey string `json:"senderPublicKey"`
	Recipient       string `json:"recipient"`
	AssetID         string `json:"assetId"`
	Amount          uint64 `json:"amount"`
	FeeAsset        string `json:"feeAsset"`
	Fee             uint64 `json:"fee"`
	Timestamp       uint64 `json:"timestamp"`
	Attachment      string `json:"attachment"`
	Signature       string `json:"signature"`
	Height          uint64 `json:"height"`
}
