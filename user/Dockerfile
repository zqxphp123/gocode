#FROM alpine
#
#RUN mkdir /go && cd /go \
#    && wget --no-check-certificate https://golang.google.cn/dl/go1.23.4.linux-arm64.tar.gz \
#    && tar -C /usr/local  -zxf  go1.23.4.linux-arm64.tar.gz \
#    && rm -rf /go/go1.23.4.linux-arm64.tar.gz \
#    && mkdir /lib64 \
#    && ln -s /lib/libc.musl-x86_64.so.1  /ib64/ld-linux-x86-64.s0.2 \
#ENV GOPATH /go
#ENV PATH /usr/local/go/bin:$G0PATH/bin:$PATH
#CMD ["ping" ,"baidu.com"]
FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY restful/etc /app/etc #检查此处的配置文件目录
RUN go build -ldflags="-s -w" -o /app/user restful/user.go #这个地方需要检查生成的go文件路径是否正确.


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/user /app/user
COPY --from=builder /app/etc /app/etc

CMD ["./user", "-f", "etc/user.yaml"] #此处配置文件的路径和名称是否符合要求

