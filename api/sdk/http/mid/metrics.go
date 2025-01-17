package mid

import (
	"context"
	"net/http"

	"github.com/wmuizelaar/myproject/app/sdk/mid"
	"github.com/wmuizelaar/myproject/foundation/web"
)

// Metrics updates program counters using the middleware functionality.
func Metrics() web.MidFunc {
	midFunc := func(ctx context.Context, r *http.Request, next mid.HandlerFunc) mid.Encoder {
		return mid.Metrics(ctx, next)
	}

	return addMidFunc(midFunc)
}
