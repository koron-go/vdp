## Benchmark result

### 2024/08/18 12:54

```console
$ go test -benchmem -bench Render ./tms9918/
goos: windows
goarch: amd64
pkg: github.com/koron-go/vdp/tms9918
cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
BenchmarkRenderScreen1-16           8746            134710 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/koron-go/vdp/tms9918 1.312s
```

### 2024/08/18 13:12

```console
$ go test -benchmem -bench Render ./tms9918/
goos: windows
goarch: amd64
pkg: github.com/koron-go/vdp/tms9918
cpu: Intel(R) Core(TM) i9-9900K CPU @ 3.60GHz
BenchmarkRenderScreen1-16           9060            130692 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/koron-go/vdp/tms9918 1.315s
```
