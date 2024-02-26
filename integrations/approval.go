package integrations

import (
	"github.com/netbirdio/netbird/management/server/account"
	nbpeer "github.com/netbirdio/netbird/management/server/peer"
)

type IntegratedApprovalImpl struct {
}

func NewIntegratedApproval() *IntegratedApprovalImpl {
	return &IntegratedApprovalImpl{}
}

func (v *IntegratedApprovalImpl) PreparePeer(_ string, peer *nbpeer.Peer, _ []string, _ *account.ExtraSettings) *nbpeer.Peer {
	return peer.Copy()
}

func (v *IntegratedApprovalImpl) SyncPeer(string, *nbpeer.Peer, []string, *account.ExtraSettings) (bool, error) {
	return true, nil
}
