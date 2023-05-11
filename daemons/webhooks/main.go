package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/whiskerside/myshopify/cache"
	"github.com/whiskerside/myshopify/conf"
	"github.com/whiskerside/myshopify/consts"
	"github.com/whiskerside/myshopify/job"
	"github.com/whiskerside/myshopify/log"
	"github.com/whiskerside/myshopify/params"
	"github.com/whiskerside/myshopify/shopify"
)

var (
	port int
	logs = log.Logger()
)

func init() {
	flag.IntVar(&port, "port", 8080, "Server port")
}

func main() {
	flag.Parse()
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	gin.SetMode(conf.Env.Mode)

	router := gin.Default()
	rg := router.Group(conf.Env.Webhooks.RouterPrefix)
	for _, t := range consts.Topics {
		rg.GET(fmt.Sprintf("/%s", t), webhooksHandler(t))
	}

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logs.Fatal().Err(err).Msg("listen: %s")
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()
	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	logs.Info().Msg("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logs.Fatal().Err(err).Msg("Server forced to shutdown")
	}
	logs.Info().Msg("Server exiting")
}

func webhooksHandler(topic string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logs.Info().
			Str("topic", topic).
			Msg("Received webhook request topic")

		if conf.Env.Mode != conf.ModeDev && !shopify.App.VerifyWebhookRequest(ctx.Request) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized, "error": consts.ErrMsgAccessDenyed,
			})
			return
		} else {
			shopifyDomain := ctx.Request.Header.Get("X-Shopify-Shop-Domain")
			tinyShop, tsAvailable := cache.GetShopByShopifyDomain(shopifyDomain)
			if tsAvailable {
				logs.Info().Msg(
					fmt.Sprintf("Loaded webhook request shop[%s] from the cache", shopifyDomain))
				tinyShop.ShopifyDomain = shopifyDomain

				var payload interface{}
				data, err := io.ReadAll(ctx.Request.Body)
				if err == nil && len(data) != 0 {
					jsoniter.Unmarshal(data, &payload)
				} else {
					// In order to avoid the empty(nil) payload passs-in
					payload = make(map[string]interface{})
				}
				params := params.MsgBody{TinyShop: tinyShop, Payload: payload}
				_, err = job.Enqueue(topic, params)
				if err != nil {
					// TODO maybe should raise the alert to monitoring system in real-time
					logs.Err(err).Msg("Push event to message queue raise exception")
				} else {
					ctx.String(http.StatusOK, "OK")
				}
			} else {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code":  http.StatusUnauthorized,
					"error": consts.ErrMsgAccessDenyed,
				})
				return
			}
		}
		ctx.String(http.StatusOK, "OK")
	}
}
