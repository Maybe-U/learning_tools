.PHONY: test

run:
	protoc -I . \
			--go_out ./gen/go/ --go_opt=paths=./proto/  \
			--go-grpc_out ./gen/go/ --go-grpc_opt=paths=./proto  proto/api.proto

test:
	# 这是测试
