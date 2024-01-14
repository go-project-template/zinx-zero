FROM golang:1.21-alpine3.19

ENV TZ Asia/Shanghai

# 阿里云
# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
# 清华 https://mirrors.tuna.tsinghua.edu.cn/help/alpine/
# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories
# 中科大
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories

# 安装时区 并设置时区
RUN apk update --no-cache && apk add --no-cache tzdata
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime &&  echo $TZ >/etc/timezone

# goproxy
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct
RUN go install github.com/cortesi/modd/cmd/modd@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /go/src

CMD ["modd"]