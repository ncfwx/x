# 记录协程上下文的 RequestID

## Usage
```go
import (
	"github.com/ncfwx/x/requestid"
)
```

一般会在 http 的 middleware 中Set，在 logger 中 Get

在 goroutine 中使用示例
```go
go func() {
	requestid.Set("my-request-id")
	defer requestid.Delete()

    func() {
        requestid.Get()
    }()
}()
```

## 性能测试

因为 go 的 benchmark 实际是在一个 goroutine 中运行，并没有并发，所以实际性能可能会有点差别。
```
goos: darwin
goarch: amd64
BenchmarkSet-4       	20000000	        92.6 ns/op
BenchmarkGet-4       	20000000	        75.8 ns/op
BenchmarkDelete-4    	20000000	        75.6 ns/op
BenchmarkGetGoID-4   	2000000000	         1.67 ns/op
```
