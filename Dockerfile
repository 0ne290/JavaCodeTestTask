FROM golang:latest as builder
ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod ./
RUN go mod tidy
COPY . .

RUN go build ./cmd/main.go

FROM scratch
COPY --from=builder /app /server
EXPOSE 80
ENTRYPOINT ["/server"]