package integrations

import (
	"context"

	"github.com/gorilla/mux"

	"github.com/netbirdio/netbird/management/internals/modules/peers"

	"go.opentelemetry.io/otel/metric"

	"github.com/netbirdio/netbird/management/server/store"
	"github.com/netbirdio/netbird/management/server/telemetry"

	"github.com/netbirdio/netbird/management/server/account"
	"github.com/netbirdio/netbird/management/server/integrations/integrated_validator"
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
	settingsManager settings.Manager,
) (*mux.Router, error) {
	return router, nil
}

func InitPermissionsManager(store store.Store, metric metric.Meter) permissions.Manager {
	return permissions.NewManager(store)
}

type Metrics struct {
	telemetry.AppMetrics
}

func InitIntegrationMetrics(ctx context.Context, metrics telemetry.AppMetrics) (*Metrics, error) {
	return &Metrics{
		AppMetrics: metrics,
	}, nil
}

func IsValidChildAccount(_ context.Context, _, _, _ string) bool {
	return false
}
