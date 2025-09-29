package main

import (
	"fmt"
)

func main() {
	var src, destn string
	var packets int
	fmt.Print("Please enter source IP being spoofed, destination IP and number of packets that'll be sent (all separated by a space):")
	fmt.Scan(&src, &destn, &packets)
	fmt.Println(Scan(destn))
}
