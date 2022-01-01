package echo

import (
	"fmt"
	"net/http"

	"github.com/Khmer495/go-templete/internal/app/api/v1/config"
	"github.com/Khmer495/go-templete/internal/app/api/v1/openapi"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Printf("Request Body: %v\nResponse Body: %v\n", string(reqBody), string(resBody))
}

func openapiRegisterHandlers(router openapi.EchoRouter, si openapi.ServerInterface) {
	wrapper := openapi.ServerInterfaceWrapper{
		Handler: si,
	}
	_ = wrapper
}

func openapiRegisterHandlersWithAuth(router openapi.EchoRouter, si openapi.ServerInterface) {
	wrapper := openapi.ServerInterfaceWrapper{
		Handler: si,
	}
	router.GET("/profile", wrapper.GetProfile)
	router.PUT("/profile", wrapper.PutProfile)
	router.POST("/users", wrapper.PostUser)
	router.GET("/teams", wrapper.GetTeams)
	router.POST("/teams", wrapper.PostTeams)
	router.PUT("/teams/:teamId", wrapper.PutTeamsTeamId)
	router.DELETE("/teams/:teamId", wrapper.DeleteTeamsTeamId)
	router.POST("/teams/:teamId/join", wrapper.PostTeamsTeamIdJoin)
}

func NewEchoServer(authValidator middleware.KeyAuthValidator, openapiHandler openapi.ServerInterface) {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/health", func(c echo.Context) error { return c.NoContent(http.StatusOK) })

	api := e.Group("/api")
	api.Use(middleware.Logger())
	api.Use(middleware.BodyDump(bodyDumpHandler))

	apiv1 := api.Group("/v1")
	openapiRegisterHandlers(apiv1, openapiHandler)

	apiv1auth := apiv1.Group("")
	apiv1auth.Use(middleware.KeyAuth(authValidator))
	openapiRegisterHandlersWithAuth(apiv1auth, openapiHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%s", config.RestPort)))
}
