package integrations

import (
	"context"

	"github.com/netbirdio/netbird/management/server/integrations/port_forwarding"
	"github.com/netbirdio/netbird/management/server/store"
	"github.com/netbirdio/netbird/management/server/types"
)

type controllerImpl struct {
	store store.Store
}

func NewController(store store.Store) port_forwarding.Controller {
	return &controllerImpl{
		store: store,
	}
}

func (c *controllerImpl) SendUpdate(ctx context.Context, accountID string, affectedProxyID string, affectedPeerIDs []string) {
	// noop
}

func (c *controllerImpl) GetProxyNetworkMaps(ctx context.Context, accountID string) (map[string]*types.NetworkMap, error) {
	return make(map[string]*types.NetworkMap), nil
}

func (c *controllerImpl) IsPeerInIngressPorts(ctx context.Context, accountID, peerID string) (bool, error) {
	return false, nil
}
