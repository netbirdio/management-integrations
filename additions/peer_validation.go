package additions

import (
	"github.com/netbirdio/netbird/management/server/account"
	"github.com/netbirdio/netbird/management/server/activity"
	nbpeer "github.com/netbirdio/netbird/management/server/peer"
)

func ValidatePeersUpdateRequest(update *nbpeer.Peer, peer *nbpeer.Peer, userID string, accountID string, eventStore activity.Store, dnsDomain string) (*nbpeer.Peer, error) {
	return update, nil
}

func ValidatePeers(peers []*nbpeer.Peer) []*nbpeer.Peer {
	return peers
}

func PreparePeer(peer *nbpeer.Peer, extraSettings *account.ExtraSettings) *nbpeer.Peer {
	return peer.Copy()
}
