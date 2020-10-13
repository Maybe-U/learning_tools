module learning_tools/go-kit/v3

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-kit/kit v0.10.0
	github.com/natefinch/lumberjack v2.0.0+incompatible // indirect
	github.com/satori/go.uuid v1.2.0
	go.uber.org/zap v1.16.0
	learning_tools/all_packaged_library v0.0.0-00010101000000-000000000000
)

replace learning_tools/all_packaged_library => ../../all_packaged_library
