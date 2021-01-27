package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/lazyledger/lazyledger-app/x/lazyledgerapp/types"
)

func CmdCreatePayForMessage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-PayForMessage [TxType] [TxFee] [Nonce] [MsgNamespaceID] [MsgSize] [MsgShareCommitment]",
		Short: "Creates a new PayForMessage",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsTxType, _ := strconv.ParseInt(args[0], 10, 64)
			argsTxFee, _ := strconv.ParseInt(args[1], 10, 64)
			argsNonce, _ := strconv.ParseInt(args[2], 10, 64)
			argsMsgNamespaceID := string(args[3])
			argsMsgSize, _ := strconv.ParseInt(args[4], 10, 64)
			argsMsgShareCommitment := string(args[5])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePayForMessage(clientCtx.GetFromAddress().String(), uint32(argsTxType), uint64(argsTxFee), uint64(argsNonce), []byte(argsMsgNamespaceID), uint64(argsMsgSize), []byte(argsMsgShareCommitment))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdatePayForMessage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-PayForMessage [id] [TxType] [TxFee] [Nonce] [MsgNamespaceID] [MsgSize] [MsgShareCommitment]",
		Short: "Update a PayForMessage",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsTxType, _ := strconv.ParseInt(args[1], 10, 64)
			argsTxFee, _ := strconv.ParseInt(args[2], 10, 64)
			argsNonce, _ := strconv.ParseInt(args[3], 10, 64)
			argsMsgNamespaceID := string(args[4])
			argsMsgSize, _ := strconv.ParseInt(args[5], 10, 64)
			argsMsgShareCommitment := string(args[6])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdatePayForMessage(clientCtx.GetFromAddress().String(), id, uint32(argsTxType), uint64(argsTxFee), uint64(argsNonce), []byte(argsMsgNamespaceID), uint64(argsMsgSize), []byte(argsMsgShareCommitment))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeletePayForMessage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-PayForMessage [id] [TxType] [TxFee] [Nonce] [MsgNamespaceID] [MsgSize] [MsgShareCommitment]",
		Short: "Delete a PayForMessage by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeletePayForMessage(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
