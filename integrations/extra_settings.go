package integrations

import (
	"context"

	"github.com/netbirdio/netbird/management/server/account"

	"github.com/netbirdio/netbird/management/server/integrations/flow"
)

type ManagerImpl struct {
}

func NewManager() flow.Manager {
	return &ManagerImpl{}
}

func (m *ManagerImpl) GetExtraSettings(ctx context.Context, accountID string) (*account.ExtraSettings, error) {
	return &account.ExtraSettings{}, nil
}

func (m *ManagerImpl) UpdateExtraSettings(ctx context.Context, accountID string, accountExtraSettings *account.ExtraSettings) error {
	return nil
}
