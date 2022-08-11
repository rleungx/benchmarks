# Protobuf benchmark

## Usage

1. Replace the `Message` definition in `proto/benchmark.proto` with yours
2. Initialize the `Message` variables in each `*/benchmark_test.go`
3. Run `make`

More details can be found in [Makefile](./Makefile).

## Benchmark results

Here are benchmark results of the example:

```
goos: linux
goarch: amd64
pkg: gogoprotobuf
cpu: Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz
BenchmarkProto3Marshal-40      	10405879	       123.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkProto3Unmarshal-40    	 4785404	       250.4 ns/op	     112 B/op	       3 allocs/op
BenchmarkProto3Clone-40        	  733822	      1728 ns/op	     176 B/op	      11 allocs/op
```

```
goos: linux
goarch: amd64
pkg: golangprotobuf
cpu: Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz
BenchmarkProto3Marshal-40      	 4911534	       245.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkProto3Unmarshal-40    	 3983472	       297.4 ns/op	     112 B/op	       3 allocs/op
BenchmarkProto3Clone-40        	 2561892	       437.3 ns/op	     104 B/op	       2 allocs/op
```

```
goos: linux
goarch: amd64
pkg: golangprotobuf/v2
cpu: Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz
BenchmarkProto3Marshal-40      	 4404859	       273.8 ns/op	      24 B/op	       1 allocs/op
BenchmarkProto3Unmarshal-40    	 2322964	       462.9 ns/op	     112 B/op	       3 allocs/op
BenchmarkProto3Clone-40        	 3486973	       353.5 ns/op	     104 B/op	       2 allocs/op
```

```
goos: linux
goarch: amd64
pkg: vtprotobuf
cpu: Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz
BenchmarkProto3Marshal-40      	12136482	        94.50 ns/op	      24 B/op	       1 allocs/op
BenchmarkProto3Unmarshal-40    	10307458	       121.2 ns/op	      16 B/op	       2 allocs/op
```

```
goos: linux
goarch: amd64
pkg: csproto
cpu: Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz
BenchmarkProto3Marshal-40      	 9272617	       112.2 ns/op	      24 B/op	       1 allocs/op
BenchmarkProto3Unmarshal-40    	 3867860	       267.7 ns/op	     104 B/op	       2 allocs/op
```

```
goos: linux
goarch: amd64
pkg: fastpb
cpu: Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz
BenchmarkProto3Marshal-40      	 7118349	       177.7 ns/op	      24 B/op	       1 allocs/op
BenchmarkProto3Unmarshal-40    	 3937045	       313.4 ns/op	     112 B/op	       3 allocs/op
```