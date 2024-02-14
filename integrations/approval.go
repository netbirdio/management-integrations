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

func (v *IntegratedApprovalImpl) PreparePeer(peer *nbpeer.Peer, _ *account.ExtraSettings) *nbpeer.Peer {
	return peer.Copy()
}

func (v *IntegratedApprovalImpl) ValidatePeer(*nbpeer.Peer) (bool, error) {
	return true, nil
}
