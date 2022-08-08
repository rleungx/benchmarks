# Protobuf benchmark

### Benchmark results
```
goos: linux
goarch: amd64
pkg: gogoprotobuf
cpu: Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz
BenchmarkProto3Marshal-40      	 9471655	       130.1 ns/op	      24 B/op	       1 allocs/op
BenchmarkProto3Unmarshal-40    	 4183198	       257.1 ns/op	     112 B/op	       3 allocs/op
```

```
goos: linux
goarch: amd64
pkg: golangprotobuf
cpu: Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz
BenchmarkProto3Marshal-40      	 4978122	       257.9 ns/op	      24 B/op	       1 allocs/op
BenchmarkProto3Unmarshal-40    	 3602485	       296.1 ns/op	     112 B/op	       3 allocs/op
```

```
goos: linux
goarch: amd64
pkg: golang-protobuf/v2
cpu: Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz
BenchmarkProto3Marshal-40      	 3863858	       268.5 ns/op	      24 B/op	       1 allocs/op
BenchmarkProto3Unmarshal-40    	 2246197	       491.2 ns/op	     112 B/op	       3 allocs/op
```

```
goos: linux
goarch: amd64
pkg: vtprotobuf
cpu: Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz
BenchmarkProto3Marshal-40      	11712817	       102.3 ns/op	      24 B/op	       1 allocs/op
BenchmarkProto3Unmarshal-40    	 9382294	       137.6 ns/op	      16 B/op	       2 allocs/op
```