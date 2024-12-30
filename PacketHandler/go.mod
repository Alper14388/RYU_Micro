module sdn/Ryu_go

go 1.22

toolchain go1.23.4

require (
	google.golang.org/grpc v1.69.2
	sdn v0.0.0-00010101000000-000000000000
)

require (
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241223144023-3abc09e42ca8 // indirect
	google.golang.org/protobuf v1.36.1 // indirect
)

replace sdn => ../
