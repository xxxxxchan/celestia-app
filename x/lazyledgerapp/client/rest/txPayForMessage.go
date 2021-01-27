package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/lazyledger/lazyledger-app/x/lazyledgerapp/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createPayForMessageRequest struct {
	BaseReq            rest.BaseReq `json:"base_req"`
	Creator            string       `json:"creator"`
	TxType             string       `json:"TxType"`
	TxFee              string       `json:"TxFee"`
	Nonce              string       `json:"Nonce"`
	MsgNamespaceID     string       `json:"MsgNamespaceID"`
	MsgSize            string       `json:"MsgSize"`
	MsgShareCommitment string       `json:"MsgShareCommitment"`
}

func createPayForMessageHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createPayForMessageRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		parsedTxType64, err := strconv.ParseUint(req.TxType, 10, 32)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		parsedTxType := uint32(parsedTxType64)

		parsedTxFee, err := strconv.ParseUint(req.TxFee, 10, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		parsedNonce64, err := strconv.ParseUint(req.Nonce, 10, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		parsedNonce := uint64(parsedNonce64)

		parsedMsgNamespaceID := req.MsgNamespaceID

		parsedMsgSize, err := strconv.ParseUint(req.MsgSize, 10, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		parsedMsgShareCommitment := req.MsgShareCommitment

		msg := types.NewMsgCreatePayForMessage(
			req.Creator,
			parsedTxType,
			parsedTxFee,
			parsedNonce,
			[]byte(parsedMsgNamespaceID),
			parsedMsgSize,
			[]byte(parsedMsgShareCommitment),
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type updatePayForMessageRequest struct {
	BaseReq            rest.BaseReq `json:"base_req"`
	Creator            string       `json:"creator"`
	TxType             string       `json:"TxType"`
	TxFee              string       `json:"TxFee"`
	Nonce              string       `json:"Nonce"`
	MsgNamespaceID     string       `json:"MsgNamespaceID"`
	MsgSize            string       `json:"MsgSize"`
	MsgShareCommitment string       `json:"MsgShareCommitment"`
}

func updatePayForMessageHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		var req updatePayForMessageRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		parsedTxType64, err := strconv.ParseUint(req.TxType, 10, 32)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		parsedTxType := uint32(parsedTxType64)

		parsedTxFee, err := strconv.ParseUint(req.TxFee, 10, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		parsedNonce, err := strconv.ParseUint(req.TxFee, 10, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		parsedMsgNamespaceID := req.MsgNamespaceID

		parsedMsgSize64, err := strconv.ParseInt(req.MsgSize, 10, 32)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		parsedMsgSize := uint64(parsedMsgSize64)

		parsedMsgShareCommitment := req.MsgShareCommitment

		msg := types.NewMsgUpdatePayForMessage(
			req.Creator,
			id,
			parsedTxType,
			parsedTxFee,
			parsedNonce,
			[]byte(parsedMsgNamespaceID),
			parsedMsgSize,
			[]byte(parsedMsgShareCommitment),
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type deletePayForMessageRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string       `json:"creator"`
}

func deletePayForMessageHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		var req deletePayForMessageRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgDeletePayForMessage(
			req.Creator,
			id,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
