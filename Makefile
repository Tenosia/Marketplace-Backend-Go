proto-auth:
	@protoc \
	--proto_path=protobuf "protobuf/auth.proto" \
	--go_out=services/common/genproto/auth \
	--go_opt=paths=source_relative \
	--go-grpc_out=services/common/genproto/auth \
	--go-grpc_opt=paths=source_relative

proto-user:
	@protoc \
	--proto_path=protobuf "protobuf/user.proto" \
	--go_out=services/common/genproto/user \
	--go_opt=paths=source_relative \
	--go-grpc_out=services/common/genproto/user \
	--go-grpc_opt=paths=source_relative

proto-notification:
	@protoc \
	--proto_path=protobuf "protobuf/notification.proto" \
	--go_out=services/common/genproto/notification \
	--go_opt=paths=source_relative \
	--go-grpc_out=services/common/genproto/notification \
	--go-grpc_opt=paths=source_relative

proto-product:
	@protoc \
	--proto_path=protobuf "protobuf/product.proto" \
	--go_out=services/common/genproto/product \
	--go_opt=paths=source_relative \
	--go-grpc_out=services/common/genproto/product \
	--go-grpc_opt=paths=source_relative

proto-order:
	@protoc \
	--proto_path=protobuf "protobuf/order.proto" \
	--go_out=services/common/genproto/order \
	--go_opt=paths=source_relative \
	--go-grpc_out=services/common/genproto/order \
	--go-grpc_opt=paths=source_relative

proto-payment:
	@protoc \
	--proto_path=protobuf "protobuf/payment.proto" \
	--go_out=services/common/genproto/payment \
	--go_opt=paths=source_relative \
	--go-grpc_out=services/common/genproto/payment \
	--go-grpc_opt=paths=source_relative

proto-review:
	@protoc \
	--proto_path=protobuf "protobuf/review.proto" \
	--go_out=services/common/genproto/review \
	--go_opt=paths=source_relative \
	--go-grpc_out=services/common/genproto/review \
	--go-grpc_opt=paths=source_relative

run-auth:
	@go run ./services/3-auth/*.go

run-gateway:
	@go run ./services/1-gateway/*.go

run-user:
	@go run ./services/4-user/*.go

run-product:
	@go run ./services/5-product/*.go

run-order:
	@go run ./services/6-order/*.go

run-payment:
	@go run ./services/7-payment/*.go

run-notification:
	@go run ./services/2-notification/*.go

run-review:
	@go run ./services/8-review/*.go
