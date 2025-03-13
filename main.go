package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// 启动 TCP 服务器
	go func() {
		defer wg.Done()
		startTCPServer()
	}()

	// 启动 UDP 服务器
	go func() {
		defer wg.Done()
		startUDPServer()
	}()

	wg.Wait()
}