package types

import (
	"bytes"
	"fmt"
	"testing"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMountainRange(t *testing.T) {
	type test struct {
		l, k     uint64
		expected []uint64
	}
	tests := []test{
		{
			l:        11,
			k:        4,
			expected: []uint64{4, 4, 2, 1},
		},
		{
			l:        2,
			k:        64,
			expected: []uint64{2},
		},
		{ //should this test throw an error? we
			l:        64,
			k:        8,
			expected: []uint64{8, 8, 8, 8, 8, 8, 8, 8},
		},
	}
	for _, tt := range tests {
		res := powerOf2MountainRange(tt.l, tt.k)
		assert.Equal(t, tt.expected, res)
	}
}

func TestNextPowerOf2(t *testing.T) {
	type test struct {
		input    uint64
		expected uint64
	}
	tests := []test{
		{
			input:    2,
			expected: 2,
		},
		{
			input:    11,
			expected: 8,
		},
		{
			input:    511,
			expected: 256,
		},
		{
			input:    1,
			expected: 1,
		},
		{
			input:    0,
			expected: 0,
		},
	}
	for _, tt := range tests {
		res := NextPowerOf2(tt.input)
		assert.Equal(t, tt.expected, res)
	}
}

func TestPowerOf2(t *testing.T) {
	type test struct {
		input    uint64
		expected bool
	}
	tests := []test{
		{
			input:    1,
			expected: true,
		},
		{
			input:    2,
			expected: true,
		},
		{
			input:    256,
			expected: true,
		},
		{
			input:    3,
			expected: false,
		},
		{
			input:    79,
			expected: false,
		},
		{
			input:    0,
			expected: false,
		},
	}
	for _, tt := range tests {
		res := powerOf2(tt.input)
		assert.Equal(t, tt.expected, res)
	}
}

// TestCreateCommit only shows if something changed, it doesn't actually show
// the commit is being created correctly todo(evan): fix me.
func TestCreateCommitment(t *testing.T) {
	type test struct {
		k         uint64
		namespace []byte
		message   []byte
		expected  []byte
		expectErr bool
	}
	tests := []test{
		{
			k:         4,
			namespace: bytes.Repeat([]byte{0xFF}, 8),
			message:   bytes.Repeat([]byte{0xFF}, 11*ShareSize),
			expected:  []byte{0x1c, 0x57, 0x89, 0x2f, 0xbe, 0xbf, 0xa2, 0xa4, 0x4c, 0x41, 0x9e, 0x2d, 0x88, 0xd5, 0x87, 0xc0, 0xbd, 0x37, 0xc0, 0x85, 0xbd, 0x10, 0x3c, 0x36, 0xd9, 0xa2, 0x4d, 0x4e, 0x31, 0xa2, 0xf8, 0x4e},
		},
		{
			k:         2,
			namespace: bytes.Repeat([]byte{0xFF}, 8),
			message:   bytes.Repeat([]byte{0xFF}, 100*ShareSize),
			expectErr: true,
		},
	}
	for _, tt := range tests {
		res, err := CreateCommitment(tt.k, tt.namespace, tt.message)
		if tt.expectErr {
			assert.Error(t, err)
			continue
		}
		assert.NoError(t, err)
		assert.Equal(t, tt.expected, res)
	}
}

func TestPadMessage(t *testing.T) {
	type test struct {
		input    []byte
		expected []byte
	}
	tests := []test{
		{
			input:    []byte{1},
			expected: append([]byte{1}, bytes.Repeat([]byte{0}, ShareSize-1)...),
		},
		{
			input:    []byte{},
			expected: []byte{},
		},
		{
			input:    bytes.Repeat([]byte{1}, ShareSize),
			expected: bytes.Repeat([]byte{1}, ShareSize),
		},
		{
			input:    bytes.Repeat([]byte{1}, (3*ShareSize)-10),
			expected: append(bytes.Repeat([]byte{1}, (3*ShareSize)-10), bytes.Repeat([]byte{0}, 10)...),
		},
	}
	for _, tt := range tests {
		res := padMessage(tt.input)
		assert.Equal(t, tt.expected, res)
	}
}

// TestSignMalleatedTxs checks to see that the signatures that are generated for
// the PayForDatas malleated from the original WirePayForData are actually
// valid.
func TestSignMalleatedTxs(t *testing.T) {
	type test struct {
		name    string
		ns, msg []byte
		ss      []uint64
		options []TxBuilderOption
	}

	kb := generateKeyring(t, "test")

	signer := NewKeyringSigner(kb, "test", "chain-id")

	tests := []test{
		{
			name:    "single share",
			ns:      []byte{1, 1, 1, 1, 1, 1, 1, 1},
			msg:     bytes.Repeat([]byte{1}, ShareSize-8),
			ss:      []uint64{2, 4, 8, 16},
			options: []TxBuilderOption{SetGasLimit(2000000)},
		},
		{
			name: "15 shares",
			ns:   []byte{1, 1, 1, 1, 1, 1, 1, 2},
			msg:  bytes.Repeat([]byte{2}, ShareSize*12),
			ss:   []uint64{4, 8, 16, 64},
			options: []TxBuilderOption{
				SetGasLimit(123456789),
				SetFeeAmount(sdk.NewCoins(sdk.NewCoin("tio", sdk.NewInt(987654321))))},
		},
	}

	for _, tt := range tests {
		wpfd, err := NewWirePayForData(tt.ns, tt.msg, tt.ss...)
		require.NoError(t, err, tt.name)
		err = wpfd.SignShareCommitments(signer, tt.options...)
		// there should be no error
		assert.NoError(t, err)
		// the signature should exist
		assert.Equal(t, len(wpfd.MessageShareCommitment[0].Signature), 64)

		// verify the signature for the PayForDatas derived from the
		// WirePayForData
		for i, size := range tt.ss {
			unsignedPFD, err := wpfd.unsignedPayForData(size)
			require.NoError(t, err)

			// create a new tx builder to create an unsigned PayForData
			builder := applyOptions(signer.NewTxBuilder(), tt.options...)
			tx, err := signer.BuildSignedTx(builder, unsignedPFD)
			require.NoError(t, err)

			// Generate the bytes to be signed.
			bytesToSign, err := signer.encCfg.TxConfig.SignModeHandler().GetSignBytes(
				signing.SignMode_SIGN_MODE_DIRECT,
				authsigning.SignerData{
					ChainID:       signer.chainID,
					AccountNumber: signer.accountNumber,
					Sequence:      signer.sequence,
				},
				tx,
			)
			require.NoError(t, err)

			// compare the commitments generated by the WirePayForData with
			// that of independently generated PayForData
			assert.Equal(t, unsignedPFD.MessageShareCommitment, wpfd.MessageShareCommitment[i].ShareCommitment)

			// verify the signature
			assert.True(t, signer.GetSignerInfo().GetPubKey().VerifySignature(
				bytesToSign,
				wpfd.MessageShareCommitment[i].Signature,
			),
				fmt.Sprintf("test: %s size: %d", tt.name, size),
			)
		}
	}
}

func TestProcessMessage(t *testing.T) {
	type test struct {
		name      string
		ns, msg   []byte
		ss        uint64
		expectErr bool
		modify    func(*MsgWirePayForData) *MsgWirePayForData
	}

	dontModify := func(in *MsgWirePayForData) *MsgWirePayForData {
		return in
	}

	kb := generateKeyring(t, "test")

	signer := NewKeyringSigner(kb, "test", "chain-id")

	tests := []test{
		{
			name:   "single share square size 2",
			ns:     []byte{1, 1, 1, 1, 1, 1, 1, 1},
			msg:    bytes.Repeat([]byte{1}, ShareSize),
			ss:     2,
			modify: dontModify,
		},
		{
			name:   "15 shares square size 4",
			ns:     []byte{1, 1, 1, 1, 1, 1, 1, 2},
			msg:    bytes.Repeat([]byte{2}, ShareSize*15),
			ss:     4,
			modify: dontModify,
		},
		{
			name: "",
			ns:   []byte{1, 1, 1, 1, 1, 1, 1, 2},
			msg:  bytes.Repeat([]byte{2}, ShareSize*15),
			ss:   4,
			modify: func(wpfd *MsgWirePayForData) *MsgWirePayForData {
				wpfd.MessageShareCommitment[0].K = 99999
				return wpfd
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		wpfd, err := NewWirePayForData(tt.ns, tt.msg, tt.ss)
		require.NoError(t, err, tt.name)
		err = wpfd.SignShareCommitments(signer)
		assert.NoError(t, err)

		wpfd = tt.modify(wpfd)

		message, spfd, sig, err := ProcessWirePayForData(wpfd, tt.ss)
		if tt.expectErr {
			assert.Error(t, err, tt.name)
			continue
		}

		// ensure that the shared fields are identical
		assert.Equal(t, tt.msg, message.Data, tt.name)
		assert.Equal(t, tt.ns, message.NamespaceId, tt.name)
		assert.Equal(t, wpfd.Signer, spfd.Signer, tt.name)
		assert.Equal(t, wpfd.MessageNameSpaceId, spfd.MessageNamespaceId, tt.name)
		assert.Equal(t, wpfd.MessageShareCommitment[0].ShareCommitment, spfd.MessageShareCommitment, tt.name)
		assert.Equal(t, wpfd.MessageShareCommitment[0].Signature, sig, tt.name)
	}
}

func validWirePayForData(t *testing.T) *MsgWirePayForData {
	msg, err := NewWirePayForData(
		[]byte{1, 2, 3, 4, 5, 6, 7, 8},
		bytes.Repeat([]byte{1}, 2000),
		16, 32, 64,
	)
	if err != nil {
		panic(err)
	}

	signer := generateKeyringSigner(t)

	err = msg.SignShareCommitments(signer)
	if err != nil {
		panic(err)
	}
	return msg
}

func applyOptions(builder sdkclient.TxBuilder, options ...TxBuilderOption) sdkclient.TxBuilder {
	for _, option := range options {
		builder = option(builder)
	}
	return builder
}
