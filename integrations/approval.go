package integrations

import (
	"github.com/netbirdio/netbird/management/server/account"
	nbpeer "github.com/netbirdio/netbird/management/server/peer"
)

type IntegratedApprovalImpl struct {
}

func NewIntegratedApproval() (*IntegratedApprovalImpl, error) {
	return &IntegratedApprovalImpl{}, nil
}

func (v *IntegratedApprovalImpl) PreparePeer(_ string, peer *nbpeer.Peer, _ []string, _ *account.ExtraSettings) *nbpeer.Peer {
	return peer.Copy()
}

func (v *IntegratedApprovalImpl) SyncPeer(accountID string, peer *nbpeer.Peer, peersGroup []string, extraSettings *account.ExtraSettings) (*nbpeer.Peer, bool) {
	return peer.Copy(), false
}
