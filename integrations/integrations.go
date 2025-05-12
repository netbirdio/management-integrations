package integrations

import (
	"context"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/metric"

	"github.com/netbirdio/netbird/management/server/store"
	"github.com/netbirdio/netbird/management/server/telemetry"

	"github.com/netbirdio/netbird/management/server/account"
	"github.com/netbirdio/netbird/management/server/activity"
	"github.com/netbirdio/netbird/management/server/activity/sqlite"
	"github.com/netbirdio/netbird/management/server/integrations/integrated_validator"
	"github.com/netbirdio/netbird/management/server/integrations/port_forwarding"
	"github.com/netbirdio/netbird/management/server/peers"
	"github.com/netbirdio/netbird/management/server/permissions"
	"github.com/netbirdio/netbird/management/server/settings"
)

func RegisterHandlers(
	ctx context.Context,
	prefix string,
	router *mux.Router,
	accountManager account.Manager,
	integratedValidator integrated_validator.IntegratedValidator,
	meter metric.Meter,
	permissionsManager permissions.Manager,
	peersManager peers.Manager,
	proxyController port_forwarding.Controller,
	settingsManager settings.Manager,
) (*mux.Router, error) {
	return router, nil
}

func InitEventStore(ctx context.Context, dataDir string, key string, _ IntegrationMetrics) (activity.Store, string, error) {
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

func InitPermissionsManager(store store.Store) permissions.Manager {
	return permissions.NewManager(store)
}

type IntegrationMetrics interface {
	telemetry.AppMetrics
}

func InitIntegrationMetrics(ctx context.Context, metrics telemetry.AppMetrics) (IntegrationMetrics, error) {
	return metrics, nil
}
