package resource

import (
	"context"
	"testing"

	"github.com/stellar/go/services/horizon/internal/db2/assets"
	"github.com/stretchr/testify/assert"
)

func TestLargeAmount(t *testing.T) {
	row := assets.AssetStatsR{
		SortKey:     "",
		Type:        "credit_alphanum4",
		Code:        "XIM",
		Issuer:      "GBZ35ZJRIKJGYH5PBKLKOZ5L6EXCNTO7BKIL7DAVVDFQ2ODJEEHHJXIM",
		Amount:      "100000000000000000000", // 10T
		NumAccounts: 429,
		Flags:       0,
		Toml:        "https://xim.com/.well-known/stellar.toml",
	}
	var res AssetStat
	err := res.Populate(context.Background(), row)
	assert.NoError(t, err)

	assert.Equal(t, "credit_alphanum4", res.Type)
	assert.Equal(t, "XIM", res.Code)
	assert.Equal(t, "GBZ35ZJRIKJGYH5PBKLKOZ5L6EXCNTO7BKIL7DAVVDFQ2ODJEEHHJXIM", res.Issuer)
	assert.Equal(t, "10000000000000.0000000", res.Amount)
	assert.Equal(t, int32(429), res.NumAccounts)
	assert.Equal(t, "https://xim.com/.well-known/stellar.toml", res.Links.Toml.Href)
}
