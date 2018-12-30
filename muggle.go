package muggle

import (
	"v2ray.com/core"
	// Required features. Can't remove unless there is replacements.
	_ "v2ray.com/core/app/dispatcher"
	_ "v2ray.com/core/app/proxyman/inbound"
	_ "v2ray.com/core/app/proxyman/outbound"
)

func StartInstance(pbConfigBytes []byte) error {
	_, err := core.StartInstance("protobuf", pbConfigBytes)
	return err
}

