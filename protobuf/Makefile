GOBIN=$(shell go env GOPATH)/bin

default: gen bench

gen: gen-gogoprotobuf gen-golangprotobuf gen-golangprotobuf-v2 gen-vtprotobuf gen-csproto gen-fastpb

bench: bench-gogoprotobuf bench-golangprotobuf bench-golangprotobuf-v2 bench-vtprotobuf bench-csproto bench-fastpb

gen-gogoprotobuf:
	go install github.com/gogo/protobuf/protoc-gen-gofast@latest
	protoc --gofast_out=gogoprotobuf proto/*.proto

gen-golangprotobuf:
	go install github.com/golang/protobuf/protoc-gen-go@v1.3.0
	protoc --go_out=golangprotobuf proto/*.proto

gen-golangprotobuf-v2:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	protoc --go_out=golangprotobuf-v2 proto/*.proto

gen-vtprotobuf:
	go install github.com/planetscale/vtprotobuf/cmd/protoc-gen-go-vtproto@latest
	protoc --go_out=vtprotobuf --go-vtproto_out=vtprotobuf --plugin protoc-gen-go-vtproto="$(GOBIN)/protoc-gen-go-vtproto" proto/*.proto

gen-csproto:
	go install github.com/CrowdStrike/csproto/cmd/protoc-gen-fastmarshal@latest
	protoc --go_out=csproto --fastmarshal_out=csproto --fastmarshal_opt=apiversion=v2 proto/*.proto

gen-fastpb:
	go install github.com/cloudwego/fastpb/protoc-gen-fastpb@latest
	protoc --go_out=fastpb --fastpb_out=fastpb proto/*.proto

bench-gogoprotobuf:
	@cd gogoprotobuf && go test -bench . -benchmem

bench-golangprotobuf:
	@cd golangprotobuf && go test -bench . -benchmem

bench-golangprotobuf-v2:
	@cd golangprotobuf-v2 && go test -bench . -benchmem

bench-vtprotobuf:
	@cd vtprotobuf && go test -bench . -benchmem

bench-csproto:
	@cd csproto && go test -bench . -benchmem

bench-fastpb:
	@cd fastpb && go test -bench . -benchmem

clean: 
	@rm -rf gogoprotobuf/*.pb.go
	@rm -rf golangprotobuf/*.pb.go
	@rm -rf golangprotobuf-v2/*.pb.go
	@rm -rf vtprotobuf/*.pb.go
	@rm -rf csproto/*.pb*.go
	@rm -rf fastpb/*.pb*.go
