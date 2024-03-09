FROM golang:1.21

ENV CGO_ENABLED 0
ENV GOOS "linux"

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY cmd/ ./cmd
COPY internal/ ./internal
COPY bootstrap/ ./bootstrap
COPY handlers ./handlers
COPY .env ./


RUN CGO_ENABLED=$CGO_ENABLED GOOS=$GOOS go build -o /val cmd/DiscordBot/main.go

EXPOSE 8080

# Run
CMD ["/val"]