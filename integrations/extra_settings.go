package integrations

import (
	"context"

	"github.com/netbirdio/netbird/management/server/activity"
	"github.com/netbirdio/netbird/management/server/types"

	"github.com/netbirdio/netbird/management/server/integrations/extra_settings"
)

type ManagerImpl struct {
}

func NewManager(eventStore activity.Store) extra_settings.Manager {
	return &ManagerImpl{}
}

func (m *ManagerImpl) GetExtraSettings(ctx context.Context, accountID string) (*types.ExtraSettings, error) {
	return &types.ExtraSettings{}, nil
}

func (m *ManagerImpl) UpdateExtraSettings(ctx context.Context, accountID, userID string, accountExtraSettings *types.ExtraSettings) error {
	return nil
}
