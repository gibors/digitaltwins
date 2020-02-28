REM protoc -I simulator/ simulator.proto --go_out=plugins=grpc:simulator

REM protoc --go_out=plugins=grpc:. simulator\simulator.proto
protoc --go_out=plugins=grpc:. simulator/simulator.proto
