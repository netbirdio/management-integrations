package integrations

import (
	"context"

	"github.com/netbirdio/netbird/management/server/account"
	"github.com/netbirdio/netbird/management/server/activity"
	"github.com/netbirdio/netbird/management/server/group"
	nbpeer "github.com/netbirdio/netbird/management/server/peer"
)

type IntegratedValidatorImpl struct {
}

func NewIntegratedValidator(_ context.Context, store activity.Store) (*IntegratedValidatorImpl, error) {
	return &IntegratedValidatorImpl{}, nil
}

func (v *IntegratedValidatorImpl) ValidateExtraSettings(context.Context, *account.ExtraSettings, *account.ExtraSettings, map[string]*nbpeer.Peer, string, string) error {
	return nil
}

func (v *IntegratedValidatorImpl) ValidatePeer(_ context.Context, update *nbpeer.Peer, _ *nbpeer.Peer, _ string, _ string, _ string, _ []string, _ *account.ExtraSettings) (*nbpeer.Peer, bool, error) {
	return update, false, nil
}

func (v *IntegratedValidatorImpl) PreparePeer(_ context.Context, _ string, peer *nbpeer.Peer, _ []string, _ *account.ExtraSettings) *nbpeer.Peer {
	return peer.Copy()
}

func (v *IntegratedValidatorImpl) IsNotValidPeer(_ context.Context, _ string, _ *nbpeer.Peer, _ []string, _ *account.ExtraSettings) (bool, bool, error) {
	return false, false, nil
}

func (v *IntegratedValidatorImpl) GetValidatedPeers(_ string, _ map[string]*group.Group, peers map[string]*nbpeer.Peer, _ *account.ExtraSettings) (map[string]struct{}, error) {
	validatedPeers := make(map[string]struct{})
	for p := range peers {
		validatedPeers[p] = struct{}{}
	}
	return validatedPeers, nil
}

func (v *IntegratedValidatorImpl) PeerDeleted(ctx context.Context, _, _ string) error {
	return nil
}

func (v *IntegratedValidatorImpl) SetPeerInvalidationListener(_ func(accountID string)) {

}

func (v *IntegratedValidatorImpl) Stop(ctx context.Context) {
}
