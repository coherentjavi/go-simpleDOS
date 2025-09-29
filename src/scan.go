package main

import (
	"fmt"
	"net"
	"sync"
)

func Scan(dstn string) []int {
	var wg sync.WaitGroup
	var openPorts []int
	//port 0 isn't included as it's a reserved port
	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", dstn, j))
			if err == nil {
				openPorts = append(openPorts, j)
				conn.Close()
				return
			}

		}(i)
	}
	wg.Wait()
	return openPorts
}
