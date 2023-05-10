package cache

import (
	"context"
	"fmt"

	"github.com/thoas/go-funk"
	"github.com/whiskerside/myshopify/log"
	"github.com/whiskerside/myshopify/params"
	"github.com/whiskerside/myshopify/storage"
)

func ShopCacheKey(shopifyDomain string) string {
	return fmt.Sprintf("myshopify:shop:%s", shopifyDomain)
}

// GetShopByShopifyDomain retrieves a Shop object from Redis cache by its Shopify domain.
//
// shopifyDomain: a string representing the Shopify domain.
// (params.Shop, bool): a tuple containing a Shop object and a boolean indicating if it was found in cache.
func GetShopByShopifyDomain(shopifyDomain string) (params.Shop, bool) {
	var ts params.Shop
	rdb := storage.RedisClient()
	if err := rdb.HGetAll(context.Background(), ShopCacheKey(shopifyDomain)).Scan(&ts); err != nil {
		log.Logger.Err(err).Msg("GetCachedShop err")
		return ts, false
	}
	if funk.NotEmpty(ts.ShopifyDomain) {
		return ts, true
	}
	return ts, false
}
