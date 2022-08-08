default: gen bench

gen: gen-gogoprotobuf gen-golangprotobuf gen-golangprotobuf-v2

bench: bench-gogoprotobuf bench-golangprotobuf bench-golangprotobuf-v2

gen-gogoprotobuf:
	go install github.com/gogo/protobuf/protoc-gen-gofast@latest
	protoc --gofast_out=gogoprotobuf proto/benchmark.proto

gen-golangprotobuf:
	go install github.com/golang/protobuf/protoc-gen-go@v1.3.0
	protoc --go_out=golangprotobuf proto/benchmark.proto

gen-golangprotobuf-v2:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	protoc --go_out=golangprotobuf-v2 proto//benchmark.proto

bench-gogoprotobuf: gen-gogoprotobuf
	cd gogoprotobuf && go test -bench . -benchmem

bench-golangprotobuf: gen-golangprotobuf
	cd golangprotobuf && go test -bench . -benchmem

bench-golangprotobuf-v2: gen-golangprotobuf-v2
	cd golangprotobuf-v2 && go test -bench . -benchmem

clean: 
	rm gogoprotobuf/benchmark.pb.go
	rm golangprotobuf/benchmark.pb.go
	rm golangprotobuf-v2/benchmark.pb.go