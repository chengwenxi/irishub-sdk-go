module github.com/irisnet/record-sdk-go

go 1.16

require (
	github.com/gogo/protobuf v1.3.3
	github.com/irisnet/core-sdk-go v0.0.0-20210712061654-e9201a163ea3
	google.golang.org/genproto v0.0.0-20201119123407-9b1e624d6bc4
	google.golang.org/grpc v1.37.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
