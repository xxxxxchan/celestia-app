package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
    // this line is used by starport scaffolding # 1
)

const (
    MethodGet = "GET"
)

// RegisterRoutes registers lazyledgerapp-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 3
    r.HandleFunc("/lazyledgerapp/PayForMessages/{id}", getPayForMessageHandler(clientCtx)).Methods("GET")
    r.HandleFunc("/lazyledgerapp/PayForMessages", listPayForMessageHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
    // this line is used by starport scaffolding # 4
    r.HandleFunc("/lazyledgerapp/PayForMessages", createPayForMessageHandler(clientCtx)).Methods("POST")
    r.HandleFunc("/lazyledgerapp/PayForMessages/{id}", updatePayForMessageHandler(clientCtx)).Methods("POST")
    r.HandleFunc("/lazyledgerapp/PayForMessages/{id}", deletePayForMessageHandler(clientCtx)).Methods("POST")

}

