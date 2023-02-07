package entities

import (
	"net"
)

type Network struct {
	name        string
	ipv4Address net.IPNet
	ipv6Address net.IPNet
	subnets     []Network
}

// Accessors and Mutators
func (n *Network) Name() string {
	return n.name
}
func (n *Network) SetName() string {
	return n.name
}

func (n *Network) IPv4Address() net.IPNet {
	return n.ipv4Address
}
func (n *Network) SetIPv4Address(ipv4Address net.IPNet) {
	n.ipv4Address = ipv4Address
}

func (n *Network) IPv6Address() net.IPNet {
	return n.ipv6Address
}
func (n *Network) SetIPv6Address(ipv6Address net.IPNet) {
	n.ipv6Address = ipv6Address
}

func (n *Network) Subnets() []Network {
	return n.subnets
}
func (n *Network) SetSubnets(subnets []Network) {
	n.subnets = subnets
}
