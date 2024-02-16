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

func (v *IntegratedApprovalImpl) ValidatePeer(string, *nbpeer.Peer, []string, []string) (bool, error) {
	return true, nil
}
