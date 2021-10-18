package main

import (
	"fmt"
	"os"

	"github.com/vishvananda/netlink"
)

func main() {
	arg := os.Args[1:]
	for _, dev := range arg {
		link, err := netlink.LinkByName(dev)
		if err != nil {
			fmt.Printf("error getting link attributes for net device %s: %v\n", dev, err)
		}
		if link != nil {
			fmt.Printf("device link type is: %s/%s\n", dev, link.Attrs().EncapType)
		}
	}
}
