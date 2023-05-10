package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/facebookgo/pidfile"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/whiskerside/myshopify/conf"
	"github.com/whiskerside/myshopify/consts"
	"github.com/whiskerside/myshopify/daemons/webhooks"
	"github.com/whiskerside/myshopify/log"
	"github.com/whiskerside/myshopify/middlewares"
)

var (
	logfile string
	port    int
)

func init() {
	flag.IntVar(&port, "port", 8000, "Server port")
	flag.StringVar(&logfile, "log", "webhook.log", "Log file path")
}

func main() {

	flag.Parse()

	logFile, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer logFile.Close()
	log.Logger.Output(logFile)
	pidfile.Write()

	r := mux.NewRouter()
	router := r.PathPrefix("/webhooks").Subrouter()

	for _, t := range consts.Topics {
		router.Handle(fmt.Sprintf("/%s", t),
			webhooks.Dispatcher{Topic: t}).Methods(consts.MethodPost)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handlers.LoggingHandler(logFile, &middlewares.CORSRouter{R: r}),
	}

	go func() {
		pid, _ := pidfile.Read()
		log.Logger.Info().
			Str("redis-host", conf.Env.Redis.Host).
			Int("redis-port", conf.Env.Redis.Port).
			Int("redis-db", conf.Env.Redis.Db).
			Str("log-file", logfile).
			Int("pid", pid).
			Msg("Server running")

		log.Logger.Printf("%v", server.ListenAndServe())
	}()
	// make sure idle connections returned
	processed := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-c

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); nil != err {
			log.Logger.Fatal().Err(err).Msg("Server shutdown failed")
		}
		log.Logger.Info().Msg("Server gracefully shutdown")
		close(processed)
	}()

	// serve
	err = server.ListenAndServe()
	log.Logger.Info().Msg("Server shutdown")
	if http.ErrServerClosed != err {
		log.Logger.Fatal().Err(err).Msg("Server not gracefully shutdown")
	}
	// waiting for goroutine above processed
	<-processed
}
