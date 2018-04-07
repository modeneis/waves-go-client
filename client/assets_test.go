package client_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/modeneis/waves-go-client/model"
	"github.com/modeneis/waves-go-client/client"
)

func TestGetAssetsBalanceAddress_200(t *testing.T) {

	expected := &model.Assets{
		Address: "3PPAMDKQz2yncVEJN6NSQAzzfYhNkqwbpZ4",
		Balances: []model.Balances{
			{AssetID: "3TqM722VJ4ieVDAm5dTt6yT8EBXvhXmAqGkx1G2tmE6b", Balance: 500000000, Issued: false},
			{AssetID: "2uAJHj2fEDZsmir45cESPkAzGiR1cDJTE5mt3iiBDmf4", Balance: 2200000000, Issued: false},
			{AssetID: "9kFnHTRG73xpnh1xeFd2qeCkVWhvxt9hTtj2uU36ytKm", Balance: 1000, Issued: false},
			{AssetID: "3DXyxUGhQZoAm5R2dJDFsC4Fh9F8kEkDnFSZCsTurzwd", Balance: 100000000, Issued: false},
			{AssetID: "6XNbqxyp5yPPUCaoCk4JBHGLWvabt6JwnMvW5No76XfJ", Balance: 200000000000, Issued: false},
			{AssetID: "GDbV3k7CbGjUWqWrdFLvVkqR9xdAzbjdhszsquf6Hdjd", Balance: 67500000000, Issued: false},
			{AssetID: "GiPB4PyRZBzbDHFtRFs6v3vbZ9TnGsevVaYtee6qaAWM", Balance: 25, Issued: false},
		},
	}

	wavesClient := client.NewAssetsService(client.MainNET)

	assets, res, err := wavesClient.GetAssetsBalanceAddress(expected.Address)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, expected, assets)
}

func TestGetAssetsBalanceAddressAssetID_200(t *testing.T) {

	expected := &model.Balances{Address: "3PPAMDKQz2yncVEJN6NSQAzzfYhNkqwbpZ4", AssetID: "3TqM722VJ4ieVDAm5dTt6yT8EBXvhXmAqGkx1G2tmE6b", Issued: false, Balance: 500000000}

	wavesClient := client.NewAssetsService(client.MainNET)

	balance, res, err := wavesClient.GetAssetsBalanceAddressAssetID(expected.Address, expected.AssetID)
	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)
	assert.NotEmpty(t, balance, "Blocksize cant be null")
	assert.Equal(t, expected, balance)

}
