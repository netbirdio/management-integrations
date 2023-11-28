package integrations

import (
	"github.com/netbirdio/netbird/management/server"
	"github.com/netbirdio/netbird/management/server/activity"
	nbpeer "github.com/netbirdio/netbird/management/server/peer"
)

func ValidatePeersUpdateRequest(update *nbpeer.Peer, peer *nbpeer.Peer, userID string, eventStore activity.Store, dnsDomain string) (*nbpeer.Peer, error) {
	return update, nil
}

func ValidatePeers(peers []*nbpeer.Peer, account *server.Account) []*nbpeer.Peer {
	return peers
}
