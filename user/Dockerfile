FROM alpine:3.18

RUN mkdir /go && cd /go \
    && wget --no-check-certificate https://golang.google.cn/dl/go1.23.4.linux-arm64.tar.gz \
    && tar -C /usr/local  -zxf  go1.23.4.linux-arm64.tar.gz \
    && rm -rf /go/go1.23.4.linux-arm64.tar.gz \
    && mkdir /lib64 \
    && ln -s /lib/libc.musl-x86_64.so.1  /ib64/ld-linux-x86-64.s0.2 \
ENV GOPATH /go
ENV PATH /usr/local/go/bin:$G0PATH/bin:$PATH
CMD ["ping" ,"baidu.com"]