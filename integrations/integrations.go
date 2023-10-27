package integrations

import (
	"github.com/gorilla/mux"
	"github.com/netbirdio/netbird/management/server"
	"github.com/netbirdio/netbird/management/server/jwtclaims"
)

func RegisterHandlers(
	router *mux.Router,
	accountManager server.AccountManager,
	extractor *jwtclaims.ClaimsExtractor,
) *mux.Router {
	return router
}
