FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go get -d -v ./...
RUN go install -v ./... -o app

CMD ["app"]
