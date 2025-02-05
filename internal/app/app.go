package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/sahirrrr/PARI-Test/internal/app/controllers/rest"
	"github.com/sahirrrr/PARI-Test/internal/app/infra"
	"github.com/typical-go/typical-go/pkg/errkit"
	"go.uber.org/dig"
)

func Start(
	di *dig.Container,
	cfg *infra.AppCfg,
	Echo *echo.Echo,
) (err error) {
	if err := di.Invoke(rest.SetRoute); err != nil {
		return err
	}

	// Create a channel to receive any potential errors
	errCh := make(chan error)

	// Launch the HTTP server in a goroutine
	go func() {
		if err := initHTTPServer(Echo, cfg); err != nil {
			errCh <- err
		}
	}()

	// Init Timezone in a goroutine
	go func() {
		if err := initTimezone(cfg); err != nil {
			errCh <- err
		}
	}()

	// Wait for any potential errors from the servers
	for i := 0; i < 2; i++ {
		err := <-errCh
		if err != nil {
			return err
		}
	}

	return nil
}

func initHTTPServer(e *echo.Echo, cfg *infra.AppCfg) error {
	return e.StartServer(&http.Server{
		Addr:         cfg.RESTPort,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
	})
}

func initTimezone(cfg *infra.AppCfg) error {
	local, err := time.LoadLocation(cfg.Timezone)
	if err != nil {
		log.Fatalf("failed to LoadLocation(): %v", err)
		<-time.After(time.Second * 5)
	}

	time.Local = local

	return err
}

// Shutdown infra.
func Shutdown(p struct {
	dig.In
	Pg   *sqlx.DB
	Echo *echo.Echo
},
) error {
	fmt.Printf("Shutdown at %s\n", time.Now().String())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errs := errkit.Errors{
		p.Echo.Shutdown(ctx),
		p.Pg.Close(),
	}

	return errs.Unwrap()
}
