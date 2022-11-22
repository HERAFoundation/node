package cli

import (
    "strconv"
	
	"github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"sporechain/x/sporechain/types"
)

var _ = strconv.Itoa(0)

func CmdCreateSporechainEntry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-sporechain-entry [user-address] [content-hash] [time]",
		Short: "Broadcast message createSporechainEntry",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
      		 argUserAddress := args[0]
             argContentHash := args[1]
             argTime := args[2]
            
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSporechainEntry(
				clientCtx.GetFromAddress().String(),
				argUserAddress,
				argContentHash,
				argTime,
				
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

    return cmd
}