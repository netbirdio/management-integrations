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

func (v *IntegratedApprovalImpl) UpdatePeerApprovalSetting(*account.ExtraSettings, *account.ExtraSettings, map[string]*nbpeer.Peer, string, string) error {
	return nil
}

func (v *IntegratedApprovalImpl) ApprovePeer(update *nbpeer.Peer, _ *nbpeer.Peer, _ string, _ string, _ string, _ []string, _ *account.ExtraSettings) (*nbpeer.Peer, error) {
	return update, nil
}

func (v *IntegratedApprovalImpl) PreparePeer(_ string, peer *nbpeer.Peer, _ []string, _ *account.ExtraSettings) *nbpeer.Peer {
	return peer.Copy()
}

func (v *IntegratedApprovalImpl) IsRequiresApproval(_ string, _ *nbpeer.Peer, _ []string, _ *account.ExtraSettings) bool {
	return false
}

func (v *IntegratedApprovalImpl) GetApprovedPeers(_ string, peers map[string]*nbpeer.Peer, _ *account.ExtraSettings) (map[string]struct{}, error) {
	approvedPeers := make(map[string]struct{})
	for p := range peers {
		approvedPeers[p] = struct{}{}
	}
	return approvedPeers, nil
}

func (v *IntegratedApprovalImpl) Stop() {
}
