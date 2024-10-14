FROM golang:1.21.4-alpine

WORKDIR /gopoll

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o gopoll-build ./cmd

EXPOSE 8080

RUN chmod +x gopoll-build

CMD ["./gopoll-build"]