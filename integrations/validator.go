package integrations

import (
	"github.com/netbirdio/netbird/management/server/account"
	"github.com/netbirdio/netbird/management/server/activity"
	nbpeer "github.com/netbirdio/netbird/management/server/peer"
)

type IntegratedValidatorImpl struct {
}

func NewIntegratedValidator(activity.Store) (*IntegratedValidatorImpl, error) {
	return &IntegratedValidatorImpl{}, nil
}

func (v *IntegratedValidatorImpl) ValidateExtraSettings(*account.ExtraSettings, *account.ExtraSettings, map[string]*nbpeer.Peer, string, string) error {
	return nil
}

func (v *IntegratedValidatorImpl) ValidatePeer(update *nbpeer.Peer, _ *nbpeer.Peer, _ string, _ string, _ string, _ []string, _ *account.ExtraSettings) (*nbpeer.Peer, error) {
	return update, nil
}

func (v *IntegratedValidatorImpl) PreparePeer(_ string, peer *nbpeer.Peer, _ []string, _ *account.ExtraSettings) *nbpeer.Peer {
	return peer.Copy()
}

func (v *IntegratedValidatorImpl) IsNotValidPeer(_ string, _ *nbpeer.Peer, _ []string, _ *account.ExtraSettings) bool {
	return false
}

func (v *IntegratedValidatorImpl) GetValidatedPeers(_ string, peers map[string]*nbpeer.Peer, _ *account.ExtraSettings) (map[string]struct{}, error) {
	validatedPeers := make(map[string]struct{})
	for p := range peers {
		validatedPeers[p] = struct{}{}
	}
	return validatedPeers, nil
}

func (v *IntegratedValidatorImpl) Stop() {
}
