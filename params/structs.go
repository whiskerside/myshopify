package params

// Shop is a context wrapper for shop with a few required attributes.
// This data structure will be consistent with the data in the cache system.
type Shop struct {
	ShopID        int    `redis:"shop_id"`
	ShopifyDomain string `redis:"shopify_domain"`
	ShopifyToken  string `redis:"shopify_token"`
	CurrentPlan   string `redis:"current_plan"`
}

type MsgBody struct {
	TinyShop Shop
	Payload  interface{}
}
