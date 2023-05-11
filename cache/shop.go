package cache

import (
	"context"
	"fmt"

	"github.com/thoas/go-funk"
	"github.com/whiskerside/myshopify/log"
	"github.com/whiskerside/myshopify/params"
	"github.com/whiskerside/myshopify/storage"
)

var (
	logs = log.Logger()
)

// ShopCacheKey returns a cache key for the given Shopify domain.
//
// shopifyDomain: the domain of the Shopify store.
// string: a string representation of the cache key.
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
		logs.Err(err).Msg("GetCachedShop err")
		return ts, false
	}
	if funk.NotEmpty(ts.ShopifyDomain) {
		return ts, true
	}
	return ts, false
}
