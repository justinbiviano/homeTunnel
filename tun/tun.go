package tunInterface

import (
	"fmt"

	"github.com/vishvananda/netlink"
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
	buf := make([]byte, 1500)
	bufs := [][]byte{buf}
	sizes := []int{0}

	_, err := device.Read(bufs, sizes, 0)
	if err != nil {
		return nil, fmt.Errorf("Failed to read packet %w", err)
	}

	return bufs[0][:sizes[0]], nil
}

func WritePacket(device tun.Device, packet []byte) error {
	bufs := [][]byte{packet}

	_, err := device.Write(bufs, 0)
	if err != nil {
		return fmt.Errorf("Failed writing packet %w:", err)
	}

	return nil
}

func ConfigureTUN(name string, ip string) error {
	link, err := netlink.LinkByName(name)
	if err != nil {
		return fmt.Errorf("Failed to find interface %s: %w", name, err)
	}

	addr, err := netlink.ParseAddr(ip)
	if err != nil {
		return fmt.Errorf("Failed to pharse address: %w", err)
	}

	err = netlink.AddrAdd(link, addr)
	if err != nil {
		return fmt.Errorf("Failed to add address: %w", err)
	}

	err = netlink.LinkSetUp(link)
	if err != nil {
		return fmt.Errorf("Failed to being interface up: %w", err)
	}

	return nil
}
