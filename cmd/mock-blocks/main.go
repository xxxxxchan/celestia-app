package mockblock

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/celestiaorg/celestia-app/x/payment/types"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	sdktypes "github.com/cosmos/cosmos-sdk/types"
)

func MockBlockCmd() *cobra.Command {
	command := &cobra.Command{
		Use:  "mock-blocks [rpc-address] [keyring-user-name] [path-to-test-keyring] [flags]",
		Args: cobra.MinimumNArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			rpcAddr := args[0]
			userName := args[1]
			chainID := args[2]
			keyringPath := args[3]

			// connect to a node
			rpcClient, err := grpc.Dial(rpcAddr, grpc.WithInsecure())
			if err != nil {
				return err
			}

			// load the keyring
			ring, err := keyring.New("celestia-app", "test", keyringPath, nil)
			if err != nil {
				return err
			}

			k := types.NewKeyringSigner(ring, userName, chainID)
			for {
				// fetch the latest account
				err = k.QueryAccountNumber(context.TODO(), rpcClient)
				if err != nil {
					return err
				}

				builder := k.NewTxBuilder()

				builder.SetGasLimit(100000000)

				coin := sdktypes.Coin{
					Denom:  "celes",
					Amount: sdktypes.NewInt(1),
				}

				builder.SetFeeAmount(sdktypes.NewCoins(coin))

				namespace := []byte{1, 1, 1, 1, 1, 1, 1, 1}
				message := bytes.Repeat([]byte{1, 2, 3}, 1000)

				msg, err := types.NewWirePayForMessage(namespace, message, 8, 16, 32, 64, 128)
				if err != nil {
					return err
				}

				signedTx, err := k.BuildSignedTx(builder, msg)
				if err != nil {
					return err
				}

				encodedTx, err := k.EncodeTx(signedTx)
				if err != nil {
					return err
				}

				resp, err := types.BroadcastTx(context.TODO(), rpcClient, tx.BroadcastMode_BROADCAST_MODE_SYNC, encodedTx)
				if err != nil {
					return err
				}

				if resp.TxResponse.Code != 0 {
					log.Println(fmt.Errorf("error when broadcasting tx: %w", err))
				}
			}
		},
	}
	// connect to a node
	// randomly generate wirePayForMessages
	// randomly generate other transactions
	// broadcast from the node
	return command
}
