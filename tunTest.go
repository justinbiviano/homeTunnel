package main

import (
	"fmt"

	tunInterface "github.com/justinbiviano/homeTunnel/tun"
)

func main() {
	device, err := tunInterface.CreateTUN()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = tunInterface.ConfigureTUN("homeTunnel", "10.0.0.1/24")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer device.Close()

	fmt.Println("TUN Created and Up")

	for {
		packet, err := tunInterface.ReadPacket(device)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Packet: %d bytes\n", len(packet))
	}
}
