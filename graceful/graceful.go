package graceful

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/litsea/kit/log"
)

const defaultStopTimeout = 30 * time.Second

var ErrEmptyService = errors.New("empty service")

type Graceful struct {
	srv         Service
	l           log.Logger
	stopTimeout time.Duration
}

type Service interface {
	Name() string
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

func New(opts ...Option) *Graceful {
	g := &Graceful{
		l:           &log.EmptyLogger{},
		stopTimeout: defaultStopTimeout,
	}

	for _, opt := range opts {
		opt(g)
	}

	return g
}

func (g *Graceful) Run(ctx context.Context) error {
	if g.srv == nil {
		return fmt.Errorf("graceful.Run: %w", ErrEmptyService)
	}

	// kill -INT <pid>
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	errCh := make(chan error, 1)
	go func() {
		errCh <- g.srv.Start(ctx)
	}()

	select {
	case <-sigs:
		g.l.Info("graceful.Run: service shutting down gracefully", "service", g.srv.Name())

		sdCtx, sdCancel := context.WithTimeout(context.Background(), g.stopTimeout)
		defer sdCancel()

		cleanCh := make(chan error, 1)
		go func() {
			cleanCh <- g.srv.Stop(sdCtx)
		}()

		for {
			select {
			case <-sdCtx.Done():
				return fmt.Errorf("graceful.Run: service %s gracefully shutdown timeout: %w",
					g.srv.Name(), sdCtx.Err())
			case err := <-cleanCh:
				if err != nil {
					return fmt.Errorf("graceful.Run: an error occurred while shutting down the service %s: %w",
						g.srv.Name(), err)
				}
				g.l.Info("graceful.Run: service gracefully stopped", "service", g.srv.Name())
				return nil
			}
		}
	case err := <-errCh:
		return fmt.Errorf("graceful.Run: service %s start failed: %w",
			g.srv.Name(), err)
	}
}
