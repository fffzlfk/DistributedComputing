FROM golang:1.16.3-alpine3.13

RUN mkdir /bookmgtapp
ADD . /bookmgtapp/

WORKDIR /bookmgtapp
RUN go env -w GOPROXY=https://goproxy.cn
RUN go mod download
RUN go build /server.server.go

CMD ["/bookmgtapp/bookmanagement"]
