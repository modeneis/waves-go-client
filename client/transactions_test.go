package client

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/modeneis/waves-go-client/model"
)

func TestGetTransactionsInfoID_200(t *testing.T) {

	expected := &model.Transactions{
		Type:            4,
		ID:              "Ee9Ji5PropqoDW2FzyyQsVEobmTHRRbtPTfEzufjkrsi",
		Sender:          "3P2u3LjeBXbVGfUVS6JnxknJBgXtAJRLttF",
		SenderPublicKey: "7Dq3SwFoojdvxkbvyUFH5WNDViyv8FD9FtCBwFuagaxE",
		Fee:             100000,
		Timestamp:       1521485621237,
		Signature:       "2hgKLeRwaxxCjwPBuGsY483pkyCSfDwJ3ZQS2PUssyP8DwogMTe4WWmcouixPkuYfXNshv1StztrWWiB2cxmSwmu",
		Recipient:       "3P9ZK7jG2shcoWvyjhbGrL8QoS5yBZAMzMn",
		AssetID:         "",
		Amount:          199900000,
		FeeAsset:        "",
		Attachment:      "",
		Height:          924610,
	}

	wavesClient := NewTransactionsService(MainNET)

	transaction, res, err := wavesClient.GetTransactionsInfoID(expected.ID)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, expected, transaction)
}

func TestGetTransactionsInfoID_404(t *testing.T) {
	wavesClient := NewTransactionsService(MainNET)
	_, res, _ := wavesClient.GetTransactionsInfoID("1234")
	assert.NotNil(t, res)
	assert.Equal(t, 404, res.StatusCode)
}

func TestGetTransactionsInfoID_404_TestNET(t *testing.T) {
	wavesClient := NewTransactionsService(TestNET)
	_, res, _ := wavesClient.GetTransactionsInfoID("1234")
	assert.NotNil(t, res)
	assert.Equal(t, 404, res.StatusCode)
}
