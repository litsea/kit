# kit

## Graceful

```golang
import (
	"github.com/litsea/kit/graceful"
)

g := graceful.New(
	graceful.WithService(srv),
	graceful.WithLogger(l),
	graceful.WithStopTimeout(time.Second*5),
)

err := g.Run(context.Background())
if err != nil {
	l.Error("service.gracefulRun", "error", err)
}
```

See: [graceful.go](graceful/graceful.go)

## pprof

```golang
import (
	"net/http"

	"github.com/litsea/kit/pprof"
)

mux := http.NewServeMux()

pprof.Register(mux, func() string {
	return "token"
})
```

## Health Check

```golang
import (
	"net/http"

	"github.com/litsea/kit/health"
)

mux := http.NewServeMux()

health.Register(mux, "/v1/health")
```
