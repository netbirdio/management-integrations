package integrations

import (
	"github.com/netbirdio/netbird/management/server/account"
	nbpeer "github.com/netbirdio/netbird/management/server/peer"
)

type IntegratedValidatorImpl struct {
}

func NewIntegratedValidator() *IntegratedValidatorImpl {
	return &IntegratedValidatorImpl{}
}

func (v *IntegratedValidatorImpl) PreparePeer(peer *nbpeer.Peer, _ *account.ExtraSettings, _ []string) *nbpeer.Peer {
	return peer.Copy()
}

func (v *IntegratedValidatorImpl) ValidatePeer(*nbpeer.Peer, []string) (bool, error) {
	return true, nil
}
