package application

import (
	"context"
	"demo-backend/internal/api"
	"demo-backend/internal/platform/db"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"time"
)

type DemoApp struct {
	router http.Handler
	config DemoApiServiceProperty
}

func NewDemoApp(config DemoApiServiceProperty) *DemoApp {
	// Init DB connection
	dataSource := db.NewDataSource(config.PgDbConfig)
	handlerFactory := api.NewHandlerFactory(dataSource)

	// Map request to corresponding handler
	router := chi.NewRouter()

	// Basic CORS settings
	c := cors.New(cors.Options{
		// AllowedOrigins is a list of origins a cross-domain request can be executed from
		// If you want to allow any origin, you can use the wildcard "*"
		AllowedOrigins: []string{"http://localhost:3000", "https://example.com"}, // Adjust according to your needs
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		// AllowedMethods is a list of methods the client is allowed to use with cross-domain requests
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// AllowedHeaders is a list of non-simple headers the client is allowed to use with cross-domain requests
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		// AllowCredentials indicates whether the request can include user credentials like cookies
		AllowCredentials: true,
		// ExposedHeaders indicates which headers are safe to expose to the API of a CORS API specification
		ExposedHeaders: []string{"Link"},
		// MaxAge indicates how long (in seconds) the results of a preflight request can be cached
		MaxAge: 300, // 5 minutes
	})

	router.Use(middleware.Logger, c.Handler)
	api.NewRequestRouteConfigurer(handlerFactory).Configure(router)

	demoApp := &DemoApp{
		config: config,
		router: router,
	}

	return demoApp
}

func (d *DemoApp) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", d.config.ServerPort),
		Handler: d.router,
	}

	var err error

	ch := make(chan error, 1)
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
	return nil
}
