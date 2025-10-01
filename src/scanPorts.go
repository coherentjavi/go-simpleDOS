package main

import (
	"fmt"
	"net"
	"sync"
)

func ScanPorts(dstn string) {
	fmt.Println("Please wait for scan to finish and find which ports are open...")
	var wg sync.WaitGroup
	//port 0 isn't included as it's a reserved port
	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", dstn, j))
			if err == nil {
				fmt.Printf("Port: %d\n", j)
				conn.Close()
				return
			}
		}(i)
	}
	wg.Wait()
}
