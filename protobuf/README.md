# Protobuf benchmark

## Usage

1. Replace the `Message` definition in `proto/benchmark.proto` with yours
2. Initialize the `Message` variables in each `*/benchmark_test.go`
3. Run `make`

## Benchmark results

Here are benchmark results of the example:

```
goos: linux
goarch: amd64
pkg: gogoprotobuf
cpu: Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz
BenchmarkProto3Marshal-40      	 8629952	       125.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkProto3Unmarshal-40    	 4203458	       284.4 ns/op	     112 B/op	       3 allocs/op
BenchmarkProto3Clone-40        	  667152	      1736 ns/op	     176 B/op	      11 allocs/op
```

```
goos: linux
goarch: amd64
pkg: golangprotobuf
cpu: Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz
BenchmarkProto3Marshal-40      	 4275278	       254.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkProto3Unmarshal-40    	 4107062	       310.5 ns/op	     112 B/op	       3 allocs/op
BenchmarkProto3Clone-40        	 2546610	       470.7 ns/op	     104 B/op	       2 allocs/op
```

```
goos: linux
goarch: amd64
pkg: golangprotobuf/v2
cpu: Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz
BenchmarkProto3Marshal-40      	 4204093	       274.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkProto3Unmarshal-40    	 2305387	       514.2 ns/op	     112 B/op	       3 allocs/op
BenchmarkProto3Clone-40        	 2939468	       404.6 ns/op	     104 B/op	       2 allocs/op
```

```
goos: linux
goarch: amd64
pkg: vtprotobuf
cpu: Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz
BenchmarkProto3Marshal-40      	11982057	        98.09 ns/op	      24 B/op	       1 allocs/op
BenchmarkProto3Unmarshal-40    	 8513671	       136.8 ns/op	      16 B/op	       2 allocs/op
```

```
goos: linux
goarch: amd64
pkg: csproto
cpu: Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz
BenchmarkProto3Marshal-40      	 9965587	       116.6 ns/op	      24 B/op	       1 allocs/op
BenchmarkProto3Unmarshal-40    	 3958486	       294.5 ns/op	     104 B/op	       2 allocs/op
```