package tunInterface

import (
	"fmt"

	"golang.zx2c4.com/wireguard/tun"
)

func CreateTUN() (tun.Device, error) {
	device, err := tun.CreateTUN("homeTunnel", 1420)
	if err != nil {
		return nil, fmt.Errorf("Failed creating TUN: %w", err)
	}
	return device, nil
}

func ReadPacket(device tun.Device) ([]byte, error) {

}

func WritePacket(device tun.Device, packet, []byte) error {

}
