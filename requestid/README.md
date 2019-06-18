# 记录协程上下文的 RequestID

## Usage
```go
import (
	"github.com/ncfwx/x/requestid"
)
```

在 goroutine 中使用，一般会在 http 的 middleware 中Set，在 logger 中 Get
```go
go func() {
    requestid.Set("my-request-id")
    defer requestid.Delete()
    
    requestid.Get()
}()
```

type CostDuration time.Duration

func (c CostDuration) String() string {
    return c/1000"ns"
}

log.Printf("%s", time.Now().Sub(s).Seconds())