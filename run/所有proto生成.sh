#!/bin/bash
cd .. && cd service/sys/rpc && goctl rpc protoc sys.proto --go_out=. --go-grpc_out=. --zrpc_out=.
cd ../../pms/rpc && goctl rpc protoc pms.proto --go_out=. --go-grpc_out=. --zrpc_out=.
#cd ../../oms/rpc && goctl rpc protoc oms.proto --go_out=. --go-grpc_out=. --zrpc_out=.
cd ../../sms/rpc && goctl rpc protoc sms.proto --go_out=. --go-grpc_out=. --zrpc_out=.
cd ../../ums/rpc && goctl rpc protoc ums.proto --go_out=. --go-grpc_out=. --zrpc_out=.
