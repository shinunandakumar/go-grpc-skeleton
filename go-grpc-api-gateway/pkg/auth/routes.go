package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/pkshahid/go-grpc-api-gateway/pkg/auth/routes"
	"github.com/pkshahid/go-grpc-api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	router := r.Group("/auth")
	router.POST("/register", svc.Register)
	router.POST("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
