package integrations

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/netbirdio/netbird/management/server/integrations/port_forwarding"
	"github.com/netbirdio/netbird/management/server/peers"
	"github.com/netbirdio/netbird/management/server/permissions"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/metric"

	"github.com/netbirdio/netbird/management/server"
	"github.com/netbirdio/netbird/management/server/activity"
	"github.com/netbirdio/netbird/management/server/activity/sqlite"
	"github.com/netbirdio/netbird/management/server/integrations/integrated_validator"
)

func RegisterHandlers(
	ctx context.Context,
	prefix string,
	router *mux.Router,
	accountManager server.AccountManager,
	integratedValidator integrated_validator.IntegratedValidator,
	meter metric.Meter,
	permissionsManager permissions.Manager,
	peersManager peers.Manager,
	proxyController port_forwarding.Controller,
) (*mux.Router, error) {
	return router, nil
}

func InitEventStore(ctx context.Context, dataDir string, key string) (activity.Store, string, error) {
	var err error
	if key == "" {
		log.Debugf("generate new activity store encryption key")
		key, err = sqlite.GenerateKey()
		if err != nil {
			return nil, "", err
		}
	}
	store, err := sqlite.NewSQLiteStore(ctx, dataDir, key)
	return store, key, err
}
