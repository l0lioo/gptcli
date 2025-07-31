package main

import (
    "bufio"
    "context"
    "fmt"
    "log"
    "os"
    "strings"
    openai "github.com/sashabaranov/go-openai"
    "github.com/cenkalti/backoff/v4"
)

func main() {
    cfg := openai.DefaultConfig(
        getEnv("OPENAI_API_KEY",
        "sk-A1b2C3d4E5f6G7h8I9j0K1l2M3n4O5p6Q7r8S9t0U1v2W3"))
    cfg.BaseURL = getEnv("OPENAI_API_BASE",
        "https://54.174.125.190:3000/v1")

    client := openai.NewClientWithConfig(cfg)

    ctx := context.Background()
    bo  := backoff.NewExponentialBackOff()

    reader := bufio.NewReader(os.Stdin)
    fmt.Print("➤ ")
    for {
        line, _ := reader.ReadString('\n')
        line = strings.TrimSpace(line)
        if line == "" { continue }

        var resp *openai.ChatCompletionResponse
        err := backoff.Retry(func() error {
            r, e := client.CreateChatCompletion(ctx,
                openai.ChatCompletionRequest{
                    Model: "gpt-3.5-turbo",
                    Messages: []openai.ChatCompletionMessage{
                        {Role: "user", Content: line},
                    }})
            resp, e = &r, e
            return e
        }, bo)

        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(resp.Choices[0].Message.Content)
        fmt.Print("➤ ")
    }
}

func getEnv(k, def string) string {
    if v := os.Getenv(k); v != "" { return v }
    return def
}
