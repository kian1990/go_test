package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"sync"
)

type PageData struct {
	Messages []Message
	sync.Mutex
}

type Message struct {
	UserMessage string
	Response    string
}

var (
	pageData PageData
	tmpl     *template.Template
)

func main() {
	var err error
	tmpl, err = template.ParseFiles("index.html")
	if err != nil {
		fmt.Printf("解析模板文件时出错：%v\n", err)
		return
	}

	pageData = PageData{
		Messages: []Message{},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			renderTemplate(w, pageData)
		} else if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "解析表单时出错", http.StatusBadRequest)
				return
			}
			message := r.FormValue("message")
			if message == "" {
				http.Error(w, "消息不能为空", http.StatusBadRequest)
				return
			}

			responseBody, err := sendPostRequest(message)
			if err != nil {
				http.Error(w, fmt.Sprintf("发送POST请求时出错：%v", err), http.StatusInternalServerError)
				return
			}

			pageData.Lock()
			defer pageData.Unlock()
			pageData.Messages = append(pageData.Messages, Message{
				UserMessage: message,
				Response:    responseBody,
			})

			jsonResponse, err := json.Marshal(pageData.Messages)
			if err != nil {
				http.Error(w, "序列化JSON时出错", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResponse)
		}
	})

	fmt.Println("服务器运行在 http://0.0.0.0:9080 上")
	http.ListenAndServe(":9080", nil)
}

func renderTemplate(w http.ResponseWriter, data PageData) {
	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "执行模板时出错", http.StatusInternalServerError)
	}
}

func sendPostRequest(message string) (string, error) {
	url := "http://127.0.0.1:9088/chat"
	data := []byte(fmt.Sprintf(`{"message":"%s"}`, message))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return "", fmt.Errorf("创建请求时出错：%v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("发送请求时出错：%v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应时出错：%v", err)
	}

	var result map[string]string
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("解析响应JSON时出错：%v", err)
	}

	response, ok := result["response"]
	if !ok {
		return "", fmt.Errorf("响应JSON中没有找到'response'字段")
	}

	return response, nil
}
