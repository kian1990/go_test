package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "github.com/google/generative-ai-go/genai"
    "google.golang.org/api/option"
)

type ChatRequest struct {
    Message string `json:"message"`
}

type ChatResponse struct {
    Response string `json:"response"`
}

func main() {
    http.HandleFunc("/chat", chatHandler)
    log.Println("服务器运行在 http://127.0.0.1:9088/chat 上")
    log.Fatal(http.ListenAndServe("127.0.0.1:9088", nil))
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
    apiKey := os.Getenv("API_KEY")
    if apiKey == "" {
        http.Error(w, "API_KEY 环境变量是必须的", http.StatusUnauthorized)
        return
    }

    var req ChatRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    response, err := getChatResponse(req.Message, apiKey)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    res := ChatResponse{Response: response}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(res)
}

func getChatResponse(message string, apiKey string) (string, error) {
    ctx := context.Background()
    client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
    if err != nil {
        return "", err
    }
    defer client.Close()

    model := client.GenerativeModel("gemini-1.5-pro")
    cs := model.StartChat()

    res, err := cs.SendMessage(ctx, genai.Text(message))
    if err != nil {
        return "", err
    }

    return extractResponse(res), nil
}

func extractResponse(resp *genai.GenerateContentResponse) string {
    var response string
    for _, cand := range resp.Candidates {
        if cand.Content != nil {
            for _, part := range cand.Content.Parts {
                contentAsString := fmt.Sprintf("%v", part)
                response += contentAsString + "\n"
            }
        }
    }
    return response
}
