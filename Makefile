.PHONY: proto
proto:
	protoc -I.:${PROTO_INCLUDE} -I./proto --go_out=. --go-grpc_out=. --validate_out="lang=go:." proto/*.proto && \
	protoc -I.:${PROTO_INCLUDE} -I./proto --grpc-gateway_out=logtostderr=true:. proto/*.proto
