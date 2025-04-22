FROM ubuntu:22.04

COPY ./xoned /usr/local/bin/xoned

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