package profiler

import (
	"maps"
	"time"

	"github.com/grafana/pyroscope-go"
)

type Option func(c *pyroscope.Config)

func WithAuth(user, pass string) Option {
	return func(c *pyroscope.Config) {
		c.BasicAuthUser = user
		c.BasicAuthPassword = pass
	}
}

func WithTags(tags map[string]string, merge bool) Option {
	return func(c *pyroscope.Config) {
		if tags != nil {
			if merge {
				maps.Copy(c.Tags, tags)
			} else {
				c.Tags = tags
			}
		}
	}
}

func WithProfileTypes(types []pyroscope.ProfileType) Option {
	return func(c *pyroscope.Config) {
		if types != nil {
			c.ProfileTypes = types
		}
	}
}

func WithDisableGCRuns(v bool) Option {
	return func(c *pyroscope.Config) {
		c.DisableGCRuns = v
	}
}

func WithUploadRate(rate time.Duration) Option {
	return func(c *pyroscope.Config) {
		// default 15s
		c.UploadRate = rate
	}
}

func WithDebug(v bool) Option {
	return func(c *pyroscope.Config) {
		if v {
			c.Logger = pyroscope.StandardLogger
		} else {
			c.Logger = nil
		}
	}
}
