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

	requestid.Get()
}()
```
