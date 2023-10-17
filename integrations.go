package integrations

import (
	"github.com/gorilla/mux"
	"github.com/netbirdio/netbird/management/server"
)

func RegisterHandlers(router *mux.Router, accountManager server.AccountManager) *mux.Router {
	return router
}
