package main

import (
	"encoding/json"
	"log"
	"net"
	"sync"
	"time"
)

func startUDPServer() {
	// 启动时加载数据
	data, err := loadData(UDP)
	if err != nil {
		log.Fatal("UDP服务器启动失败:", err)
	}

	addr, err := net.ResolveUDPAddr("udp", ":8889")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	log.Println("UDP Server started on port 8889")

	clients := make(map[string]*net.UDPAddr)
	var mu sync.RWMutex

	// 接收客户端消息的 goroutine
	go func() {
		buffer := make([]byte, 1024)
		for {
			_, remoteAddr, err := conn.ReadFromUDP(buffer)
			if err != nil {
				continue
			}
			mu.Lock()
			clients[remoteAddr.String()] = remoteAddr
			mu.Unlock()
			log.Printf("New UDP client connected: %s", remoteAddr)
		}
	}()

	// 发送数据的主循环
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Error marshaling data:", err)
	}

	for {
		mu.RLock()
		for addr, client := range clients {
			_, err := conn.WriteToUDP(jsonData, client)
			if err != nil {
				log.Printf("Error sending to %s: %v", addr, err)
				mu.RUnlock()
				mu.Lock()
				delete(clients, addr)
				mu.Unlock()
				mu.RLock()
			}
		}
		mu.RUnlock()

		time.Sleep(500 * time.Millisecond)
	}
}
