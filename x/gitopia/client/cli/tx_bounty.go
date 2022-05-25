package cli

import (
	"errors"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cosmosTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/gitopia/gitopia/x/gitopia/types"
	"github.com/spf13/cobra"
)

func CmdCreateBounty() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-bounty [amount] [expiry] [parent-id] [parent]",
		Short: "Create a new Bounty",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmount, err := cosmosTypes.ParseCoinsNormalized(args[0])
			if err != nil {
				return err
			}
			argExpiry, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}
			argParentId, err := strconv.ParseUint(args[2], 10, 64)
			if err != nil {
				return err
			}
			argParent := args[3]
			parent, ok := types.BountyParent_value[argParent]
			if !ok {
				return errors.New("invalid bounty parent")
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateBounty(clientCtx.GetFromAddress().String(), argAmount, argExpiry, argParentId, types.BountyParent(parent))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateBounty() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-bounty [id] [amount] [expiry]",
		Short: "Update a Bounty",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argAmount, err := cosmosTypes.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}
			argExpiry, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateBounty(clientCtx.GetFromAddress().String(), id, argAmount, argExpiry)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteBounty() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-bounty [id]",
		Short: "Delete a Bounty by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteBounty(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
