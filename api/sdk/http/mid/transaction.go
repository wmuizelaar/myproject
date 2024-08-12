package mid

import (
	"context"
	"net/http"

	"github.com/wmuizelaar/myproject/app/sdk/mid"
	"github.com/wmuizelaar/myproject/business/sdk/sqldb"
	"github.com/wmuizelaar/myproject/foundation/logger"
	"github.com/wmuizelaar/myproject/foundation/web"
)

// BeginCommitRollback executes the transaction middleware functionality.
func BeginCommitRollback(log *logger.Logger, bgn sqldb.Beginner) web.MidFunc {
	midFunc := func(ctx context.Context, r *http.Request, next mid.HandlerFunc) mid.Encoder {
		return mid.BeginCommitRollback(ctx, log, bgn, next)
	}

	return addMidFunc(midFunc)
}
