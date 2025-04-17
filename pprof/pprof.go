package pprof

import (
	"net/http"
	"net/http/pprof"
)

func Register(mux *http.ServeMux, getTokenFn func() string) {
	if getTokenFn == nil {
		return
	}

	prefix := "/debug/pprof"

	mux.HandleFunc(prefix+"/", handleFuncMiddleware(pprof.Index, getTokenFn))
	mux.HandleFunc(prefix+"/cmdline", handleFuncMiddleware(pprof.Cmdline, getTokenFn))
	mux.HandleFunc(prefix+"/profile", handleFuncMiddleware(pprof.Profile, getTokenFn))
	mux.HandleFunc(prefix+"/symbol", handleFuncMiddleware(pprof.Symbol, getTokenFn))
	mux.HandleFunc(prefix+"/trace", handleFuncMiddleware(pprof.Trace, getTokenFn))
	mux.Handle(prefix+"/allocs", handleMiddleware(pprof.Handler("allocs"), getTokenFn))
	mux.Handle(prefix+"/block", handleMiddleware(pprof.Handler("block"), getTokenFn))
	mux.Handle(prefix+"/goroutine", handleMiddleware(pprof.Handler("goroutine"), getTokenFn))
	mux.Handle(prefix+"/heap", handleMiddleware(pprof.Handler("heap"), getTokenFn))
	mux.Handle(prefix+"/mutex", handleMiddleware(pprof.Handler("mutex"), getTokenFn))
	mux.Handle(prefix+"/threadcreate", handleMiddleware(pprof.Handler("threadcreate"), getTokenFn))
}

func handleMiddleware(next http.Handler, getTokenFn func() string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if checkToken(w, r, getTokenFn) {
			next.ServeHTTP(w, r)
		}
	})
}

func handleFuncMiddleware(
	next func(http.ResponseWriter, *http.Request), getTokenFn func() string,
) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if checkToken(w, r, getTokenFn) {
			next(w, r)
		}
	}
}

func checkToken(w http.ResponseWriter, r *http.Request, getTokenFn func() string) bool {
	if getTokenFn == nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"code":404,"msg":"Not Found"}`))
		return false
	}

	tk := getTokenFn()
	if tk == "" {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"code":404,"msg":"Not Found"}`))
		return false
	}

	t := r.URL.Query().Get("token")
	if tk != t {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"code":403,"msg":"Forbidden"}`))
		return false
	}

	return true
}
