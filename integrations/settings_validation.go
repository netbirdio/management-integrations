package integrations

import (
	"github.com/netbirdio/netbird/management/server"
	"github.com/netbirdio/netbird/management/server/status"
)

func ValidateExtraSettings(extraSettings *server.ExtraSettings, account *server.Account, userID string, manager server.AccountManager) error {
	if extraSettings != nil {
		return status.Errorf(status.InvalidArgument, "extra settings are only supported on the cloud version of NetBird")
	}
	return nil
}
