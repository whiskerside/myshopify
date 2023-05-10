package funtion

import (
	goshopify "github.com/bold-commerce/go-shopify/v3"
	"github.com/whiskerside/myshopify/conf"
)

var (
	ShopifyApp = goshopify.App{
		ApiKey:      conf.Env.ShopifyAPIKey,
		ApiSecret:   conf.Env.ShopifySharedSecret,
		RedirectUrl: conf.Env.ShopifyRedirectURL,
		Scope:       conf.Env.ShopifyScopes,
	}
)
