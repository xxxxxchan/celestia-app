package cli

import (
    "context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
    "github.com/lazyledger/lazyledger-app/x/lazyledgerapp/types"
)

func CmdListPayForMessage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-PayForMessage",
		Short: "list all PayForMessage",
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)

            pageReq, err := client.ReadPageRequest(cmd.Flags())
            if err != nil {
                return err
            }

            queryClient := types.NewQueryClient(clientCtx)

            params := &types.QueryAllPayForMessageRequest{
                Pagination: pageReq,
            }

            res, err := queryClient.PayForMessageAll(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}

func CmdShowPayForMessage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-PayForMessage [id]",
		Short: "shows a PayForMessage",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx := client.GetClientContextFromCmd(cmd)

            queryClient := types.NewQueryClient(clientCtx)

            params := &types.QueryGetPayForMessageRequest{
                Id: args[0],
            }

            res, err := queryClient.PayForMessage(context.Background(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}
