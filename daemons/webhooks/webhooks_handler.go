package webhooks

import (
	"fmt"
	"io"
	"net/http"

	goshopify "github.com/bold-commerce/go-shopify/v3"
	jsoniter "github.com/json-iterator/go"
	"github.com/unrolled/render"
	"github.com/whiskerside/myshopify/cache"
	"github.com/whiskerside/myshopify/conf"
	"github.com/whiskerside/myshopify/job"
	"github.com/whiskerside/myshopify/log"
	"github.com/whiskerside/myshopify/params"
)

const (
	AccessDeny = "401 Unauthorized"
)

var (
	logs = log.Logger()
)

type Dispatcher struct {
	Topic string
}

func (h Dispatcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rendering := render.New()

	shopifyApp := goshopify.App{ApiSecret: conf.Env.Shopify.SharedSecret}
	if conf.Env.Mode != conf.ModeDev && !shopifyApp.VerifyWebhookRequest(r) {
		rendering.Text(w, http.StatusUnauthorized, AccessDeny)
	} else {
		shopifyDomain := r.Header.Get("X-Shopify-Shop-Domain")
		tinyShop, tsAvailable := cache.GetShopByShopifyDomain(shopifyDomain)
		if tsAvailable {
			logs.Info().Msg(
				fmt.Sprintf("Loaded webhook request shop[%s] from the cache", shopifyDomain))
			tinyShop.ShopifyDomain = shopifyDomain

			var payload interface{}
			data, err := io.ReadAll(r.Body)
			if err == nil && len(data) != 0 {
				jsoniter.Unmarshal(data, &payload)
			} else {
				// In order to avoid the empty(nil) payload passs-in
				payload = make(map[string]interface{})
			}
			params := params.MsgBody{TinyShop: tinyShop, Payload: payload}
			_, err = job.Enqueue(h.Topic, params)
			if err != nil {
				// TODO maybe should raise the alert to monitoring system in real-time
				logs.Err(err).Msg("Push event to message queue raise exception")
			} else {
				rendering.Text(w, http.StatusOK, "OK")
			}
		} else {
			rendering.Text(w, http.StatusUnauthorized, AccessDeny)
		}
	}
}
