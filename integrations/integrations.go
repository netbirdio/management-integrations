package integrations

import (
	"github.com/gorilla/mux"
	"github.com/netbirdio/netbird/management/server"
	"github.com/netbirdio/netbird/management/server/activity"
	"github.com/netbirdio/netbird/management/server/activity/sqlite"
	"github.com/netbirdio/netbird/management/server/jwtclaims"
	log "github.com/sirupsen/logrus"
)

func RegisterHandlers(
	router *mux.Router,
	accountManager server.AccountManager,
	extractor *jwtclaims.ClaimsExtractor,
) *mux.Router {
	return router
}

func InitEventStore(dataDir string, key string) (activity.Store, string, error) {
	var err error
	if key == "" {
		log.Debugf("generate new activity store encryption key")
		key, err = sqlite.GenerateKey()
		if err != nil {
			return nil, "", err
		}
	}
	store, err := sqlite.NewSQLiteStore(dataDir, key)
	return store, key, err
}
