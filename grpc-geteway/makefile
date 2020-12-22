.PHONY: test

run:
	protoc -I . \
			--go_out ./ --go_opt=paths=import  \
			--go-grpc_out ./ --go-grpc_opt=paths=import \
			--grpc-gateway_out ./  \
			--grpc-gateway_opt paths=import  \
			proto/api.proto

test:
	# 这是测试
