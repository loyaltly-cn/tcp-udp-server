package main

import (
	"encoding/json"
	"log"
	"net"
	"time"
)

func startTCPServer() {
	// 启动时加载数据
	data, err := loadData(TCP)
	if err != nil {
		log.Fatal("TCP服务器启动失败:", err)
	}

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("TCP Server started on port 8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go handleTCPClient(conn, data)
	}
}

func handleTCPClient(conn net.Conn, data *Data) {
	defer conn.Close()
	log.Printf("New TCP client connected: %s", conn.RemoteAddr())

	for {
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Println("Error marshaling data:", err)
			return
		}

		_, err = conn.Write(jsonData)
		if err != nil {
			log.Println("Error sending data:", err)
			return
		}

		time.Sleep(500 * time.Millisecond)
	}
}
