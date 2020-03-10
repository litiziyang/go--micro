FROM golang:alpine as build
LABEL stage=builder

WORKDIR /micro-lzy/service-user

COPY . .

RUN GOPROXY="https://goproxy.cn" GO111MODULE=on go mod download
RUN CGO_ENABLED=0 GOOS=linux GOPROXY="https://goproxy.cn" GO111MODULE=on go build -a -installsuffix cgo -o service-user

FROM alpine:latest

RUN apk --no-cache add ca-certificates


RUN mkdir /app
WORKDIR /app
COPY --from=build /micro-lzy/service-user .


CMD ["./service-user"]

