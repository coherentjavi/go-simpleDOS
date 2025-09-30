package main

import (
	"fmt"
)

func main() {
	var src, destn string
	var packets, port int
	fmt.Print("Please enter the IP being spoofed, destination IP and number of packets that'll be sent in the flood (all separated by a space):")
	fmt.Scan(&src, &destn, &packets)
	/*fmt.Println("The following ports are open in the destination IP. Which port do you want SYN flood?:", Scan(destn))*/
	fmt.Scan(&port)

	Attack(src, destn, packets, port)

}
