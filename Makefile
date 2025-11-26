#SHELL=/usr/bin/env bash

CLEAN:=
BINS:=
DATE_TIME=`date +'%Y%m%d %H:%M:%S'`
COMMIT_ID=`git rev-parse --short HEAD`
TAG_ID=`git describe --tags --abbrev=0`
DOCKER := $(shell which docker)

build-rocksdb:
	CGO_ENABLED=1 \
	CGO_CFLAGS="-I/usr/include" \
	CGO_LDFLAGS="/usr/lib/x86_64-linux-gnu/librocksdb.a \
	/usr/lib/x86_64-linux-gnu/libsnappy.a \
	/usr/lib/x86_64-linux-gnu/liblz4.a \
	/usr/lib/x86_64-linux-gnu/libzstd.a \
	/usr/lib/x86_64-linux-gnu/libbz2.a \
	/usr/lib/x86_64-linux-gnu/libz.a \
	-lstdc++ -lm -ldl -lpthread" \
	go build -tags "rocksdb" \
	-ldflags "-s -w" \
	-o xoned cmd/xoned/main.go
build:
	go mod tidy \
	&& go build -ldflags "-s -w -X 'main.BuildTime=${DATE_TIME}' -X 'main.GitCommit=${COMMIT_ID}' -X 'main.Version=${TAG_ID}'" -o xoned cmd/xoned/main.go
.PHONY: build
BINS+=xoned

install:build
	cp -f xoned ${GOPATH}/bin && xoned version

init:
	ignite chain init --skip-proto -y

ignite-build:
	ignite chain build -y --debug --clear-cache --check-dependencies -v

# legacy version 0.11.6
protoVer=0.13.0
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace --user 0 $(protoImageName)

proto-gen:
	@echo "Generating Protobuf files"
	@$(protoImage) sh ./scripts/protocgen.sh
.PHONY: proto-gen

debug:
	xoned start --pruning=nothing --evm.tracer=json --log_level trace \
                 --json-rpc.api eth,txpool,personal,net,debug,web3,miner \
                 --api.enable --json-rpc.enable --json-rpc.address 0.0.0.0:8545 \
                 --json-rpc.ws-address 0.0.0.0:8546

start:
	xoned start --pruning=nothing --json-rpc.api eth,txpool,personal,net,debug,web3,miner \
                 --api.enable --json-rpc.enable --json-rpc.address 0.0.0.0:8545  --json-rpc.ws-address 0.0.0.0:8546

docker: clean
	docker build --tag xone-node -f ./Dockerfile .

# 构建支持 RocksDB 的静态链接二进制（通过 Docker）
docker-rocksdb:
	docker build --tag xone-node-rocksdb \
		--build-arg VERSION=${TAG_ID} \
		--build-arg BUILD_TIME="${DATE_TIME}" \
		--build-arg GIT_COMMIT=${COMMIT_ID} \
		-f ./Dockerfile.rocksdb .

# 从 Docker 镜像中提取静态链接的二进制文件
extract-rocksdb-bin: docker-rocksdb
	docker create --name temp-xone xone-node-rocksdb
	docker cp temp-xone:/usr/local/bin/xoned ./xoned-rocksdb
	docker rm temp-xone
	@echo "Binary extracted to ./xoned-rocksdb"

reset: install init start

docker-test: install
	docker build --tag xone-node -f ./Dockerfile.test .

clean:
	rm -rf $(BINS) $(CLEAN)

