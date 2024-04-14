FROM golang:1.22 as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main main.go

FROM alpine:3.19.1
WORKDIR /

COPY --from=builder /build/main /main

CMD [ "/main" ]
