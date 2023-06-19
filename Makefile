# Go parameters
GOCMD=GO111MODULE=on go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test

PROJECT_NAME := "mall"
PKG := "$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /example)

all: test build
build:
	rm -rf target/
	mkdir -p target/config
	cp warehouse/default.yaml target/config/warehouse.yaml
	$(GOBUILD) -o target/warehouse warehouse/main.go

test:
	$(GOTEST) -v ${PKG_LIST}

clean:
	rm -rf target/
	rm -rf nohup.out

# 运行服务
run:
	nohup target/warehouse --config target/config/warehouse.yaml &

# 停止服务
stop:
	pkill -f target/warehouse

#生成docker镜像，请确保已安装docker
docker:
	docker build -t mall/warehouse:latest -f services/warehouse/Dockerfile .
	docker build -t mall/center:latest -f services/center/Dockerfile .
	docker build -t mall/third-party:latest -f services/third-party/Dockerfile .
	docker build -t mall/cart:latest -f services/cart/Dockerfile .
	docker build -t mall/market:latest -f services/market/Dockerfile .
	docker build -t mall/member:latest -f services/member/Dockerfile .
	docker build -t mall/order:latest -f services/order/Dockerfile .
	docker build -t mall/product:latest -f services/product/Dockerfile .
	docker build -t mall/seckill:latest -f services/seckill/Dockerfile .
	docker build -t mall/gateway:latest -f gateway/Dockerfile .