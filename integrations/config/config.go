package config

import (
	"github.com/netbirdio/netbird/shared/management/proto"
	"github.com/netbirdio/netbird/management/server/types"
)

func ExtendNetBirdConfig(peerID string, peerGroups []string, config *proto.NetbirdConfig, extraSettings *types.ExtraSettings) *proto.NetbirdConfig {
	return config
}
