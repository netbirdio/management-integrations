package integrations

import (
	"github.com/netbirdio/netbird/management/server/account"
	"github.com/netbirdio/netbird/management/server/activity"
	nbpeer "github.com/netbirdio/netbird/management/server/peer"
)

type IntegratedApprovalImpl struct {
}

func NewIntegratedApproval(activity.Store) (*IntegratedApprovalImpl, error) {
	return &IntegratedApprovalImpl{}, nil
}

func (v *IntegratedApprovalImpl) PreparePeer(_ string, peer *nbpeer.Peer, _ []string, _ *account.ExtraSettings) *nbpeer.Peer {
	return peer.Copy()
}

func (v *IntegratedApprovalImpl) IsRequiresApproval(accountID string, peer *nbpeer.Peer, peersGroup []string, extraSettings *account.ExtraSettings) bool {
	return false
}

func (v *IntegratedApprovalImpl) Stop() {
}
