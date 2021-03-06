// © 2019 Imhotep Software LLC. All rights reserved.

package main

import (
	"flag"
	"net/http"
	"os"
	"time"

	"github.com/gopherland/ntnx_labs/deploy/picker/internal/web"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
)

const httpPort = ":4500"

var (
	// Version set via build
	version  = "dev"
	logDebug bool
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}

func main() {
	flag.BoolVar(&logDebug, "d", logDebug, "Turn off debugging logs")
	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if logDebug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	h := web.New("assets")

	c := setupMiddleware()
	r := mux.NewRouter()
	r.Handle("/api/v1/load", c.Then(http.HandlerFunc(h.LoadHandler))).Methods("GET")
	r.Handle("/api/v1/picker", c.Then(http.HandlerFunc(h.PickHandler))).Methods("GET")

	log.Info().Msgf("Picker[%s] is listening on port [%s]", version, httpPort)
	log.Panic().Err(http.ListenAndServe(httpPort, r))
}

func setupMiddleware() alice.Chain {
	c := alice.New()
	c = c.Append(hlog.NewHandler(log.Logger))
	c = c.Append(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Debug().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("Dic")
	}))
	c = c.Append(hlog.RemoteAddrHandler("ip"))
	c = c.Append(hlog.UserAgentHandler("user_agent"))
	c = c.Append(hlog.RefererHandler("referer"))
	c = c.Append(hlog.RequestIDHandler("req_id", "Request-Id"))
	return c
}
