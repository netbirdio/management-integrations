package integrations

import "github.com/netbirdio/netbird/management/server"

func ValidatePeersUpdateRequest(update *server.Peer, peer *server.Peer, manager server.AccountManager) (*server.Peer, error) {
	return update, nil
}

func ValidatePeers(peers []*server.Peer, account *server.Account) []*server.Peer {
	return peers
}
