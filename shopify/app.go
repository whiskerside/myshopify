package shopify

import (
	goshopify "github.com/bold-commerce/go-shopify/v3"
	"github.com/whiskerside/myshopify/conf"
)

var (
	App = goshopify.App{
		ApiKey:      conf.Env.Shopify.APIKey,
		ApiSecret:   conf.Env.Shopify.SharedSecret,
		RedirectUrl: conf.Env.Shopify.RedirectURI,
		Scope:       conf.Env.Shopify.Scopes,
	}
)
