package integrations

import "github.com/netbirdio/netbird/management/server"

func IsPeerAllowedToConnect(peer *server.Peer) bool {
	return true
}
