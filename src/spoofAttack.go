package main

import(
	"fmt"
	"syscall"
)


type ipHeader struct := {

} 

func SpoofAttack(src string, dstn string, packets int, port int)  {

	//we're creating a raw socket so that we can directly add our own IP headers to the packets sent
	//AF_PACKET and ETH_P_ALL are linux specific network constants
	fd, err := syscall.Socket(syscall.AF_PACKET, syscall.SOCK_RAW, syscall.ETH_P_ALL)

	if err!= nil {
		fmt.Println("There was an error creating the raw socket needed for the spoof attack")
		fmt.Println("Error:" +  err.Error())
		return;
	}

	for i := 0; i <= packets; i++ {

	}
	
}