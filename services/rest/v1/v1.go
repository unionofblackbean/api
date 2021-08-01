package v1

import (
	"github.com/unionofblackbean/api/app"
	"github.com/unionofblackbean/api/common/web"
	"github.com/unionofblackbean/api/services/rest/v1/controllers"
	"github.com/unionofblackbean/api/services/rest/v1/controllers/session"
	"github.com/unionofblackbean/api/services/rest/v1/controllers/user"
)

func RegisterEndpoints(srv *web.Server, deps *app.Deps) {
	v1Group := srv.Group("/v1")
	{
		sessionGroup := v1Group.Group("/session")
		{
			sessionGroup.Any("/login", session.NewLoginController(deps).Any)
		}

		userGroup := v1Group.Group("/user")
		{
			userGroup.Any("/register", user.NewRegisterController(deps).Any)
		}

		v1Group.Any("/health", controllers.Health)
	}
}
