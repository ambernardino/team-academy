FROM alpine:3.4
RUN apk --update upgrade && \
    apk add sqlite && \
    rm -rf /var/cache/apk/*
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
ADD res/ca-certificates.crt /etc/ssl/certs/
ADD main /
CMD ["/main"]
