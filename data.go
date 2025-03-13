package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// ProtocolType 限定协议类型只能是 TCP 或 UDP
type ProtocolType string

const (
	TCP ProtocolType = "tcp"
	UDP ProtocolType = "udp"
)

// Data 将根据实际 JSON 内容动态解析
type Data struct {
	data map[string]interface{}
}

func loadData(name ProtocolType) (*Data, error) {
	filePath := filepath.Join("json", string(name)+".json")
	log.Printf("正在读取文件: %s", filePath)

	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("读取文件失败: %v", err)
	}

	var rawData map[string]interface{}
	err = json.Unmarshal(file, &rawData)
	if err != nil {
		return nil, fmt.Errorf("解析JSON失败: %v", err)
	}

	return &Data{data: rawData}, nil
}

// GetString 获取字符串字段
func (d *Data) GetString(key string) string {
	if val, ok := d.data[key].(string); ok {
		return val
	}
	return ""
}

// GetInt 获取整数字段
func (d *Data) GetInt(key string) int {
	if val, ok := d.data[key].(float64); ok {
		return int(val)
	}
	return 0
}

// MarshalJSON 实现 json.Marshaler 接口
func (d *Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.data)
}
