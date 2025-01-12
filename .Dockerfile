FROM golang:latest
WORKDIR /java_code_test_task
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /java-code-test-task

EXPOSE 80
CMD ["/java-code-test-task"]