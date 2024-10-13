package start2

import (
	"bitbucket.org/alibaba-international/ghs-common/adapter"
	"bitbucket.org/alibaba-international/go-pkg/env"
	"bitbucket.org/alibaba-international/go-pkg/logger"
	"bitbucket.org/alibaba-international/go-pkg/rest/api"
	"bitbucket.org/alibaba-international/go-pkg/rest/router"
	"bitbucket.org/alibaba-international/go-pkg/rest/swagger"
	"bitbucket.org/alibaba-international/go-pkg/rest/typescript"
	"bitbucket.org/alibaba-international/go-pkg/shutdown"
	"bitbucket.org/alibaba-international/go-pkg/types/locale"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Start(apis []api.Api) {
	logger.Init()
	g := gin.New()
	for _, env1 := range []adapter.Environment{adapter.EnvTest, adapter.EnvProd} {
		r := router.NewFromGin(g, string(env1))
		r.Use(func(c *api.Context) {
			c.Set("env", env1)
			c.Next()
		})
		r.Use(router.HandleLogger(), router.HandleResponse())
		r.Use(func(c *api.Context) {
			var lo = locale.Default
			if val := c.GetHeader("Locale"); val != "" {
				l := locale.Locale(val)
				if err := l.Validate(); err == nil {
					lo = l
				}
			}
			c.SetLocale(lo)
		})
		r.HandleApis(apis...)
	}
	switch env.Get() {
	case env.DEV, env.TEST:
		serveSwagger(g, apis)
		serveTypescript(g, apis)
	default:
	}
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: g,
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return
			}
			logrus.WithError(err).Error("client api http server start error")
		}
	}()
	shutdown.OnShutdown(httpServer.Shutdown)
	shutdown.Wait()
}

func serveSwagger(r gin.IRouter, apis []api.Api) {
	o := swagger.Options{
		BasePath:  "",
		ServePath: "swagger",
		Apis:      apis,
		Title:     "API",
	}
	if err := swagger.Route(r, o); err != nil {
		logrus.WithError(err).Fatal("Failed to register swagger route")
	}
}

func serveTypescript(r gin.IRouter, apis []api.Api) {
	c := typescript.Options{
		RequestImportStr: "import request from \"@/core/request\";",
		ServePath:        "typescript",
		Apis:             apis,
	}
	if err := typescript.RegisterService(r, c); err != nil {
		logrus.WithError(err).Fatal("Failed to register typescript route")
	}
}
