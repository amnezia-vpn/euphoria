//go:build !linux

package device

import (
	"github.com/amnezia-vpn/euphoria/conn"
	"github.com/amnezia-vpn/euphoria/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
