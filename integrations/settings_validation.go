package integrations

import (
	"github.com/netbirdio/netbird/management/server/account"
	"github.com/netbirdio/netbird/management/server/activity"
	nbpeer "github.com/netbirdio/netbird/management/server/peer"
	"github.com/netbirdio/netbird/management/server/status"
)

func ValidateExtraSettings(newExtraSettings *account.ExtraSettings, oldExtraSettings *account.ExtraSettings, peers map[string]*nbpeer.Peer, userID string, accountID string, eventStore activity.Store) error {
	if newExtraSettings != nil {
		return status.Errorf(status.InvalidArgument, "extra settings are only supported on the cloud version of NetBird")
	}
	return nil
}
