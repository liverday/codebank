FROM golang:1.15

WORKDIR /go/src/github.com/liverday/codebank/desafio-1/

COPY . .

ENV GO111MODULE=on

RUN go mod tidy
RUN GOOS=linux go build -o entrypoint

RUN ls

ENTRYPOINT ["./entrypoint"]
