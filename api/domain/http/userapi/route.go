package userapi

import (
	"net/http"

	"github.com/wmuizelaar/myproject/api/sdk/http/mid"
	"github.com/wmuizelaar/myproject/app/domain/userapp"
	"github.com/wmuizelaar/myproject/app/sdk/auth"
	"github.com/wmuizelaar/myproject/app/sdk/authclient"
	"github.com/wmuizelaar/myproject/business/domain/userbus"
	"github.com/wmuizelaar/myproject/foundation/logger"
	"github.com/wmuizelaar/myproject/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log        *logger.Logger
	UserBus    *userbus.Business
	AuthClient *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := mid.Authenticate(cfg.Log, cfg.AuthClient)
	ruleAdmin := mid.Authorize(cfg.Log, cfg.AuthClient, auth.RuleAdminOnly)
	ruleAuthorizeUser := mid.AuthorizeUser(cfg.Log, cfg.AuthClient, cfg.UserBus, auth.RuleAdminOrSubject)
	ruleAuthorizeAdmin := mid.AuthorizeUser(cfg.Log, cfg.AuthClient, cfg.UserBus, auth.RuleAdminOnly)

	api := newAPI(userapp.NewApp(cfg.UserBus))
	app.HandlerFunc(http.MethodGet, version, "/users", api.query, authen, ruleAdmin)
	app.HandlerFunc(http.MethodGet, version, "/users/{user_id}", api.queryByID, authen, ruleAuthorizeUser)
	app.HandlerFunc(http.MethodPost, version, "/users", api.create, authen, ruleAdmin)
	app.HandlerFunc(http.MethodPut, version, "/users/role/{user_id}", api.updateRole, authen, ruleAuthorizeAdmin)
	app.HandlerFunc(http.MethodPut, version, "/users/{user_id}", api.update, authen, ruleAuthorizeUser)
	app.HandlerFunc(http.MethodDelete, version, "/users/{user_id}", api.delete, authen, ruleAuthorizeUser)
}
