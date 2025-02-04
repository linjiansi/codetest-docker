FROM golang:1.23.5

WORKDIR /app

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.63.4 && \
   go install github.com/air-verse/air@latest && \
   go install github.com/go-delve/delve/cmd/dlv@latest

COPY . .
RUN go mod download

# For debugging
#CMD ["air", "-c", ".air.toml"]

RUN go build -o /go/bin/app ./src/cmd/api
CMD ["/go/bin/app"]