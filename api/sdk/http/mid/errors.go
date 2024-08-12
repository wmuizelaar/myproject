package mid

import (
	"context"
	"net/http"

	"github.com/wmuizelaar/myproject/app/sdk/mid"
	"github.com/wmuizelaar/myproject/foundation/logger"
	"github.com/wmuizelaar/myproject/foundation/web"
)

// Errors executes the errors middleware functionality.
func Errors(log *logger.Logger) web.MidFunc {
	midFunc := func(ctx context.Context, r *http.Request, next mid.HandlerFunc) mid.Encoder {
		return mid.Errors(ctx, log, next)
	}

	return addMidFunc(midFunc)
}
