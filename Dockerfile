FROM ccr.ccs.tencentyun.com/woniu/alpine:latest
WORKDIR /
ADD server /server

ENTRYPOINT ["/server"]