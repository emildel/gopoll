FROM --platform=$BUILDPLATFORM golang:alpine AS build
ARG TARGETOS
ARG TARGETARCH

WORKDIR /gopoll

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags='-s -w' -o gopoll-build ./cmd && \
    chmod +x gopoll-build

CMD ["./gopoll-build", "-port=80", "-environment=production"]