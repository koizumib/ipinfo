package netcalc

import (
	"errors"
	"fmt"
	"net"
)

type Row struct {
	IPAddress        string
	SubnetMask       string
	NetworkAddress   string
	BroadcastAddress string
	HostRange        string
	Hosts            uint64
}

func RowFromCIDR(cidr string) (*Row, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, fmt.Errorf("invalid CIDR: %w", err)
	}
	ip4 := ip.To4()
	if ip4 == nil {
		return nil, errors.New("only IPv4 is supported")
	}

	mask := ipnet.Mask
	ones, bits := mask.Size()
	if bits != 32 {
		return nil, errors.New("only IPv4 /32 supported")
	}

	ipU := ipToUint32(ip4)
	maskU := maskToUint32(mask)

	networkU := ipU & maskU
	broadcastU := networkU | (^maskU)

	var firstU, lastU uint32
	var hosts uint64
	switch {
	case ones == 32:
		firstU, lastU = ipU, ipU
		hosts = 1
	case ones == 31:
		firstU, lastU = networkU, broadcastU
		hosts = 2
	default:
		firstU, lastU = networkU+1, broadcastU-1
		hosts = (1 << (32 - ones)) - 2
	}

	return &Row{
		IPAddress:        ip4.String(),
		SubnetMask:       maskToDotted(mask),
		NetworkAddress:   uint32ToIP(networkU).String(),
		BroadcastAddress: uint32ToIP(broadcastU).String(),
		HostRange:        fmt.Sprintf("%s - %s", uint32ToIP(firstU), uint32ToIP(lastU)),
		Hosts:            hosts,
	}, nil
}

func ipToUint32(ip net.IP) uint32 {
	ip = ip.To4()
	return (uint32(ip[0]) << 24) | (uint32(ip[1]) << 16) | (uint32(ip[2]) << 8) | uint32(ip[3])
}

func uint32ToIP(u uint32) net.IP {
	return net.IPv4(
		byte(u>>24),
		byte((u>>16)&0xFF),
		byte((u>>8)&0xFF),
		byte(u&0xFF),
	)
}

func maskToUint32(m net.IPMask) uint32 {
	// IPv4マスク前提
	return (uint32(m[0]) << 24) | (uint32(m[1]) << 16) | (uint32(m[2]) << 8) | uint32(m[3])
}

func maskToDotted(m net.IPMask) string {
	return net.IPv4(m[0], m[1], m[2], m[3]).String()
}
