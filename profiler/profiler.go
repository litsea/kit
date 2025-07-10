package profiler

import (
	"errors"
	"os"
	"runtime"
	"slices"

	"github.com/grafana/pyroscope-go"
)

var ErrEmptyNameOrServerAddress = errors.New("empty name or server address")

func Start(name, addr string, opts ...Option) (*pyroscope.Profiler, error) {
	if name == "" || addr == "" {
		return nil, ErrEmptyNameOrServerAddress
	}

	hn, _ := os.Hostname()
	c := pyroscope.Config{
		ApplicationName: name,
		ServerAddress:   addr,
		DisableGCRuns:   true,
		Tags: map[string]string{
			"hostname": hn,
		},
		ProfileTypes: []pyroscope.ProfileType{
			// these profile types are enabled by default:
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
			// these profile types are optional:
			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	}

	for _, opt := range opts {
		opt(&c)
	}

	if slices.Contains(c.ProfileTypes, pyroscope.ProfileMutexCount) ||
		slices.Contains(c.ProfileTypes, pyroscope.ProfileMutexDuration) {
		runtime.SetMutexProfileFraction(5)
	}

	if slices.Contains(c.ProfileTypes, pyroscope.ProfileBlockCount) ||
		slices.Contains(c.ProfileTypes, pyroscope.ProfileBlockDuration) {
		runtime.SetBlockProfileRate(5)
	}

	return pyroscope.Start(c)
}
