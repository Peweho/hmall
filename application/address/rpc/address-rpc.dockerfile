FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
#ENV GOPROXY https://goproxy.cn,direct
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

#ADD go.mod .
#ADD go.sum .
#RUN go mod download
#COPY . .
COPY application/address/rpc/etc /app/etc
COPY application/address/rpc/address-rpc /app/address
#RUN go build -ldflags="-s -w" -o /app/address application/address/rpc/address.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /usr/share/zoneinfo/Asia/Shanghai
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/address /app/address
COPY --from=builder /app/etc /app/etc

CMD ["./address", "-f", "etc/address.yaml"]
