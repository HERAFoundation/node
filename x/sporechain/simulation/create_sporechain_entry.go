package simulation

import (
	"math/rand"

	"sporechain/x/sporechain/keeper"
	"sporechain/x/sporechain/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateSporechainEntry(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateSporechainEntry{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateSporechainEntry simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateSporechainEntry simulation not implemented"), nil, nil
	}
}
