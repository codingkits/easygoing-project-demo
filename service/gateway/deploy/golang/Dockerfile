FROM golang:1.18-alpine3.16

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn\ 
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone
ENV LANG C.UTF-8

ENV PATH $GOPATH/bin:$PATH
WORKDIR $GOPATH/kd

# COPY ../../go.mod /kd/
ADD . ./

RUN go build  -buildmode=plugin -o krakend-auth-plugin.so .

EXPOSE 8080

ENTRYPOINT  ["./kd-dmp-plugin"]