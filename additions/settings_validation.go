package additions

import (
	"github.com/netbirdio/netbird/management/server/account"
	"github.com/netbirdio/netbird/management/server/activity"
	nbpeer "github.com/netbirdio/netbird/management/server/peer"
	log "github.com/sirupsen/logrus"
)

func ValidateExtraSettings(newExtraSettings *account.ExtraSettings, oldExtraSettings *account.ExtraSettings, peers map[string]*nbpeer.Peer, userID string, accountID string, eventStore activity.Store) error {
	if newExtraSettings != nil {
		log.Info("extra settings are only supported on the cloud version of NetBird")
		if newExtraSettings.PeerApprovalEnabled {
			log.Info("setting peer approval to false")
			newExtraSettings.PeerApprovalEnabled = false
		}
	}
	return nil
}
