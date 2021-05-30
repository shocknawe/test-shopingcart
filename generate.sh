#!/bin/bash
protoc -I . --go_out=. transaction/product.proto
protoc -I . --go_out=. transaction/transaction.proto
protoc -I . --go-grpc_out=. transaction/transaction.proto
