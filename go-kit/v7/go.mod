module learning_tools/go-kit/v7

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-kit/kit v0.10.0
	github.com/golang/protobuf v1.4.2
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/prometheus/client_golang v1.3.0
	github.com/satori/go.uuid v1.2.0
	go.uber.org/zap v1.16.0
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e
	google.golang.org/grpc v1.33.0
	google.golang.org/grpc/examples v0.0.0-20201010204749-3c400e7fcc87 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	learning_tools/all_packaged_library v0.0.0-00010101000000-000000000000
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace learning_tools/all_packaged_library => ../../all_packaged_library
