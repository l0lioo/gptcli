FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o /usr/local/bin/gptcli ./cmd/gptcli

FROM alpine
COPY --from=builder /usr/local/bin/gptcli /usr/local/bin/gptcli
ENV OPENAI_API_BASE=https://54.174.125.190:3000/v1
ENTRYPOINT ["gptcli"]
