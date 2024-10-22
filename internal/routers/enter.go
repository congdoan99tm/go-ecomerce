package routers

import (
	"github.com/dinos/go-ecommerce-be-api/internal/routers/manager"
	"github.com/dinos/go-ecommerce-be-api/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
