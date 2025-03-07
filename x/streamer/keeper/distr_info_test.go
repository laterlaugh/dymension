package keeper_test

import (
	"testing"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/dymensionxyz/dymension/v3/x/streamer/types"
)

func (suite *KeeperTestSuite) TestAllocateToGauges() {
	tests := []struct {
		name                   string
		testingDistrRecord     []types.DistrRecord
		mintedCoins            sdk.Coin
		expectedGaugesBalances []sdk.Coins
		expectedCommunityPool  sdk.DecCoin
	}{
		// With minting 15000 stake to module, after AllocateAsset we get:
		// expectedCommunityPool = 0 (All reward will be transferred to the gauges)
		// 	expectedGaugesBalances in order:
		//    gauge1_balance = 15000 * 100/(100+200+300) = 2500
		//    gauge2_balance = 15000 * 200/(100+200+300) = 5000 (using the formula in the function gives the exact result 4999,9999999999995000. But TruncateInt return 4999. Is this the issue?)
		//    gauge3_balance = 15000 * 300/(100+200+300) = 7500
		{
			name: "Allocated to the gauges proportionally",
			testingDistrRecord: []types.DistrRecord{
				{
					GaugeId: 1,
					Weight:  math.NewInt(100),
				},
				{
					GaugeId: 2,
					Weight:  math.NewInt(200),
				},
				{
					GaugeId: 3,
					Weight:  math.NewInt(300),
				},
			},
			mintedCoins: sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(15000)),
			expectedGaugesBalances: []sdk.Coins{
				sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(2500))),
				sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(4999))),
				sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(7500))),
			},
			expectedCommunityPool: sdk.NewDecCoin(sdk.DefaultBondDenom, math.NewInt(0)),
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			suite.CreateGauges(3)

			// create a stream
			suite.CreateStream(test.testingDistrRecord, sdk.NewCoins(test.mintedCoins), time.Now().Add(-1*time.Hour), "day", 1)

			// move all created streams from upcoming to active
			suite.Ctx = suite.Ctx.WithBlockTime(time.Now())

			suite.DistributeAllRewards()

			for i := 0; i < len(test.testingDistrRecord); i++ {
				if test.testingDistrRecord[i].GaugeId == 0 {
					continue
				}
				gauge, err := suite.App.IncentivesKeeper.GetGaugeByID(suite.Ctx, test.testingDistrRecord[i].GaugeId)
				suite.Require().NoError(err)
				suite.Require().ElementsMatch(test.expectedGaugesBalances[i], gauge.Coins)
			}
		})
	}
}

func TestNewDistrInfo(t *testing.T) {
	// Test case: valid records
	records := []types.DistrRecord{
		{Weight: math.NewInt(1)},
		{Weight: math.NewInt(2)},
	}
	distrInfo, err := types.NewDistrInfo(records)
	require.NoError(t, err)
	require.Equal(t, distrInfo.TotalWeight, math.NewInt(3))

	// Test case: invalid record
	records = []types.DistrRecord{
		{Weight: math.NewInt(-1)},
	}
	_, err = types.NewDistrInfo(records)
	require.Error(t, err)

	// Test case: total weight not positive
	records = []types.DistrRecord{}
	_, err = types.NewDistrInfo(records)
	require.Error(t, err)
	require.Equal(t, err, types.ErrDistrInfoNotPositiveWeight)
}
