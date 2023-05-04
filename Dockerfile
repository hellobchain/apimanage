FROM golang:1.15 as builder
ARG Version
ARG CommitVersion
ARG BuildTime
LABEL version=$Version comshbuimit=$CommitVersion create_time=$BuildTime

ADD . /api-manage
WORKDIR /api-manage
RUN  go version && go env && gcc -v && \
     CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build \
     --mod=vendor  -v -o api-manage cmd/main.go
#生成中间镜像后,将build之后的可执行文件考到新的镜像中
FROM alpine:3.14 as api-manage
ARG Version
ARG CommitVersion
ARG BuildTime
LABEL version=$Version commit=$CommitVersion create_time=$BuildTime
COPY --from=builder  /api-manage/api-manage /usr/local/bin
COPY --from=builder  /api-manage/deployments/config /api-manage/deployments/config
COPY --from=builder  /api-manage/html /api-manage/html
WORKDIR /api-manage
#容器内部开放端口
CMD ["api-manage"]