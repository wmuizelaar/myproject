// Package reporting binds the reporting domain set of routes into the specified app.
package reporting

import (
	"time"

	"github.com/wmuizelaar/myproject/api/domain/http/checkapi"
	"github.com/wmuizelaar/myproject/api/domain/http/vproductapi"
	"github.com/wmuizelaar/myproject/api/sdk/http/mux"
	"github.com/wmuizelaar/myproject/business/domain/userbus"
	"github.com/wmuizelaar/myproject/business/domain/userbus/stores/usercache"
	"github.com/wmuizelaar/myproject/business/domain/userbus/stores/userdb"
	"github.com/wmuizelaar/myproject/business/domain/vproductbus"
	"github.com/wmuizelaar/myproject/business/domain/vproductbus/stores/vproductdb"
	"github.com/wmuizelaar/myproject/business/sdk/delegate"
	"github.com/wmuizelaar/myproject/foundation/web"
)

// Routes constructs the add value which provides the implementation of
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (add) Add(app *web.App, cfg mux.Config) {

	// Construct the business domain packages we need here so we are using the
	// sames instances for the different set of domain apis.
	delegate := delegate.New(cfg.Log)
	userBus := userbus.NewBusiness(cfg.Log, delegate, usercache.NewStore(cfg.Log, userdb.NewStore(cfg.Log, cfg.DB), time.Minute))
	vproductBus := vproductbus.NewBusiness(vproductdb.NewStore(cfg.Log, cfg.DB))

	checkapi.Routes(app, checkapi.Config{
		Build: cfg.Build,
		Log:   cfg.Log,
		DB:    cfg.DB,
	})

	vproductapi.Routes(app, vproductapi.Config{
		UserBus:     userBus,
		VProductBus: vproductBus,
		AuthClient:  cfg.AuthClient,
	})
}
