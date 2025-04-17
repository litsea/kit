package graceful

import (
	"time"

	"github.com/litsea/kit/log"
)

type Option func(*Graceful)

func WithService(s Service) Option {
	return func(g *Graceful) {
		g.srv = s
	}
}

func WithStopTimeout(dur time.Duration) Option {
	return func(c *Graceful) {
		c.stopTimeout = dur
	}
}

func WithLogger(l log.Logger) Option {
	return func(c *Graceful) {
		if l != nil {
			c.l = l
		}
	}
}
