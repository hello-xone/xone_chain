FROM golang:1.21 AS builder

ENV SRC_DIR /code
COPY . $SRC_DIR

RUN go env -w GOPROXY=https://proxy.golang.org,direct
RUN apt-get update && apt-get install -y ca-certificates make build-essential
RUN cd $SRC_DIR && make

FROM ubuntu:22.04

ENV SRC_DIR /code

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder $SRC_DIR/xoned /usr/local/bin/xoned

ENV HOME_PATH /data
VOLUME $HOME_PATH

# evm json-rpc port
EXPOSE 8545
# evm websocket port
EXPOSE 8546
# p2p port
EXPOSE 26656
# node port
EXPOSE 26657
# proxy port
EXPOSE 26658
# prometheus port
EXPOSE 26660
# app api port
EXPOSE 1317
# app grpc port
EXPOSE 9090
# app grpc web port
EXPOSE 9091