GOPATH:=$(shell go env GOPATH)

#生成docker镜像，请确保已安装docker
.PHONY: docker
docker:
	docker build -t mall/cart:latest -f services/cart/Dockerfile .
	docker build -t mall/center:latest -f services/center/Dockerfile .
	docker build -t mall/market:latest -f services/market/Dockerfile .
	docker build -t mall/member:latest -f services/member/Dockerfile .
	docker build -t mall/order:latest -f services/order/Dockerfile .
	docker build -t mall/product:latest -f services/product/Dockerfile .
	docker build -t mall/seckill:latest -f services/seckill/Dockerfile .
	docker build -t mall/seckill:latest -f services/seckill/Dockerfile .
	docker build -t mall/third-party:latest -f services/third-party/Dockerfile .
	docker build -t mall/warehouse:latest -f services/warehouse/Dockerfile .