package main

import (
	"fmt"
)

func main() {
	var src, destn string
	var packets, port int
	fmt.Print("Please enter the IP being spoofed, destination IP and number of packets that'll be sent in the flood (all separated by a space):")
	fmt.Scan(&src, &destn, &packets)
	fmt.Println("We'll now scan the ports on the device with the destination IP")
	fmt.Scan(&port)
	
	start := time.Now()
	ScanPorts(destn)
	elapsed := time.Since(start)
	fmt.Printf("***The scan took %v***\n", elapsed)
	
	fmt.Println("Which port do you want SYN flood?")
	fmt.Scan(&port)

	Attack(src, destn, packets, port)

}
