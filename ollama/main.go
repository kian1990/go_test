package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Ollama 响应结构体
type OllamaResponse struct {
	Response string `json:"response"`
}

func main() {
	url := "http://localhost:11434/api/generate"
	requestData := map[string]interface{}{
		"model":  "deepseek-r1:1.5b",
		"prompt": "你好，介绍一下 DeepSeek 模型。",
		"stream": false, // 设为 false 以获取完整 JSON
	}

	// 将请求数据编码为 JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("JSON 序列化失败:", err)
		return
	}

	// 发送 HTTP POST 请求
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 创建一个 Scanner 逐行解析 JSON
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		var result OllamaResponse
		if err := json.Unmarshal([]byte(line), &result); err != nil {
			fmt.Println("解析 JSON 失败:", err)
			fmt.Println("返回内容:", line)
			continue
		}
		fmt.Println("Ollama 响应:", result.Response)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取响应失败:", err)
	}
}
