package tranapi

import (
	"net/http"

	"github.com/wmuizelaar/myproject/api/sdk/http/mid"
	"github.com/wmuizelaar/myproject/app/domain/tranapp"
	"github.com/wmuizelaar/myproject/app/sdk/auth"
	"github.com/wmuizelaar/myproject/app/sdk/authclient"
	"github.com/wmuizelaar/myproject/business/domain/productbus"
	"github.com/wmuizelaar/myproject/business/domain/userbus"
	"github.com/wmuizelaar/myproject/business/sdk/sqldb"
	"github.com/wmuizelaar/myproject/foundation/logger"
	"github.com/wmuizelaar/myproject/foundation/web"
	"github.com/jmoiron/sqlx"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log        *logger.Logger
	DB         *sqlx.DB
	UserBus    *userbus.Business
	ProductBus *productbus.Business
	AuthClient *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := mid.Authenticate(cfg.Log, cfg.AuthClient)
	transaction := mid.BeginCommitRollback(cfg.Log, sqldb.NewBeginner(cfg.DB))
	ruleAdmin := mid.Authorize(cfg.Log, cfg.AuthClient, auth.RuleAdminOnly)

	api := newAPI(tranapp.NewApp(cfg.UserBus, cfg.ProductBus))
	app.HandlerFunc(http.MethodPost, version, "/tranexample", api.create, authen, ruleAdmin, transaction)
}
