package consts

const (
	// Here defined all supports webhooks topic from the Shopify
	// https://shopify.dev/docs/api/admin-rest/2023-01/resources/webhook
	TopicAppSubscriptionsUpdate                              = "app_subscriptions/update"
	TopicAppUninstalled                                      = "app/uninstalled"
	TopicBulkOperationsFinish                                = "bulk_operations/finish"
	TopicCartsCreate                                         = "carts/create"
	TopicCartsUpdate                                         = "carts/update"
	TopicCheckoutsCreate                                     = "checkouts/create"
	TopicCheckoutsDelete                                     = "checkouts/delete"
	TopicCheckoutsUpdate                                     = "checkouts/update"
	TopicCollectionListingsAdd                               = "collection_listings/add"
	TopicCollectionListingsRemove                            = "collection_listings/remove"
	TopicCollectionListingsUpdate                            = "collection_listings/update"
	TopicCollectionsCreate                                   = "collections/create"
	TopicCollectionsDelete                                   = "collections/delete"
	TopicCollectionsUpdate                                   = "collections/update"
	TopicCustomerGroupsCreate                                = "customer_groups/create"
	TopicCustomerGroupsDelete                                = "customer_groups/delete"
	TopicCustomerGroupsUpdate                                = "customer_groups/update"
	TopicCustomerPaymentMethodsCreate                        = "customer_payment_methods/create"
	TopicCustomerPaymentMethodsRevoke                        = "customer_payment_methods/revoke"
	TopicCustomerPaymentMethodsUpdate                        = "customer_payment_methods/update"
	TopicCustomersMarketingConsentUpdate                     = "customers_marketing_consent/update"
	TopicCustomersCreate                                     = "customers/create"
	TopicCustomersDelete                                     = "customers/delete"
	TopicCustomersDisable                                    = "customers/disable"
	TopicCustomersEnable                                     = "customers/enable"
	TopicCustomersUpdate                                     = "customers/update"
	TopicDisputesCreate                                      = "disputes/create"
	TopicDisputesUpdate                                      = "disputes/update"
	TopicDomainsCreate                                       = "domains/create"
	TopicDomainsDestroy                                      = "domains/destroy"
	TopicDomainsUpdate                                       = "domains/update"
	TopicDraftOrdersCreate                                   = "draft_orders/create"
	TopicDraftOrdersDelete                                   = "draft_orders/delete"
	TopicDraftOrdersUpdate                                   = "draft_orders/update"
	TopicFulfillmentEventsCreate                             = "fulfillment_events/create"
	TopicFulfillmentEventsDelete                             = "fulfillment_events/delete"
	TopicFulfillmentOrdersCancellationRequestAccepted        = "fulfillment_orders/cancellation_request_accepted"
	TopicFulfillmentOrdersCancellationRequestRejected        = "fulfillment_orders/cancellation_request_rejected"
	TopicFulfillmentOrdersCancellationRequestSubmitted       = "fulfillment_orders/cancellation_request_submitted"
	TopicFulfillmentOrdersCancelled                          = "fulfillment_orders/cancelled"
	TopicFulfillmentOrdersFulfillmentRequestAccepted         = "fulfillment_orders/fulfillment_request_accepted"
	TopicFulfillmentOrdersFulfillmentRequestRejected         = "fulfillment_orders/fulfillment_request_rejected"
	TopicFulfillmentOrdersFulfillmentRequestSubmitted        = "fulfillment_orders/fulfillment_request_submitted"
	TopicFulfillmentOrdersFulfillmentServiceFailedToComplete = "fulfillment_orders/fulfillment_service_failed_to_complete"
	TopicFulfillmentOrdersHoldReleased                       = "fulfillment_orders/hold_released"
	TopicFulfillmentOrdersLineItemsPreparedForLocalDelivery  = "fulfillment_orders/line_items_prepared_for_local_delivery"
	TopicFulfillmentOrdersLineItemsPreparedForPickup         = "fulfillment_orders/line_items_prepared_for_pickup"
	TopicFulfillmentOrdersMoved                              = "fulfillment_orders/moved"
	TopicFulfillmentOrdersOrderRoutingComplete               = "fulfillment_orders/order_routing_complete"
	TopicFulfillmentOrdersPlacedOnHold                       = "fulfillment_orders/placed_on_hold"
	TopicFulfillmentOrdersRescheduled                        = "fulfillment_orders/rescheduled"
	TopicFulfillmentOrdersScheduledFulfillmentOrderReady     = "fulfillment_orders/scheduled_fulfillment_order_ready"
	TopicFulfillmentsCreate                                  = "fulfillments/create"
	TopicFulfillmentsUpdate                                  = "fulfillments/update"
	TopicInventoryItemsCreate                                = "inventory_items/create"
	TopicInventoryItemsDelete                                = "inventory_items/delete"
	TopicInventoryItemsUpdate                                = "inventory_items/update"
	TopicInventoryLevelsConnect                              = "inventory_levels/connect"
	TopicInventoryLevelsDisconnect                           = "inventory_levels/disconnect"
	TopicInventoryLevelsUpdate                               = "inventory_levels/update"
	TopicLocalesCreate                                       = "locales/create"
	TopicLocalesUpdate                                       = "locales/update"
	TopicLocationsActivate                                   = "locations/activate"
	TopicLocationsCreate                                     = "locations/create"
	TopicLocationsDeactivate                                 = "locations/deactivate"
	TopicLocationsDelete                                     = "locations/delete"
	TopicLocationsUpdate                                     = "locations/update"
	TopicMarketsCreate                                       = "markets/create"
	TopicMarketsDelete                                       = "markets/delete"
	TopicMarketsUpdate                                       = "markets/update"
	TopicOrderTransactionsCreate                             = "order_transactions/create"
	TopicOrdersCancelled                                     = "orders/cancelled"
	TopicOrdersCreate                                        = "orders/create"
	TopicOrdersDelete                                        = "orders/delete"
	TopicOrdersEdited                                        = "orders/edited"
	TopicOrdersFulfilled                                     = "orders/fulfilled"
	TopicOrdersPaid                                          = "orders/paid"
	TopicOrdersPartiallyFulfilled                            = "orders/partially_fulfilled"
	TopicOrdersUpdated                                       = "orders/updated"
	TopicPaymentSchedulesDue                                 = "payment_schedules/due"
	TopicProductListingsAdd                                  = "product_listings/add"
	TopicProductListingsRemove                               = "product_listings/remove"
	TopicProductListingsUpdate                               = "product_listings/update"
	TopicProductsCreate                                      = "products/create"
	TopicProductsDelete                                      = "products/delete"
	TopicProductsUpdate                                      = "products/update"
	TopicProfilesCreate                                      = "profiles/create"
	TopicProfilesDelete                                      = "profiles/delete"
	TopicProfilesUpdate                                      = "profiles/update"
	TopicRefundsCreate                                       = "refunds/create"
	TopicScheduledProductListingsAdd                         = "scheduled_product_listings/add"
	TopicScheduledProductListingsRemove                      = "scheduled_product_listings/remove"
	TopicScheduledProductListingsUpdate                      = "scheduled_product_listings/update"
	TopicSellingPlanGroupsCreate                             = "selling_plan_groups/create"
	TopicSellingPlanGroupsDelete                             = "selling_plan_groups/delete"
	TopicSellingPlanGroupsUpdate                             = "selling_plan_groups/update"
	TopicShopUpdate                                          = "shop/update"
	TopicSubscriptionBillingAttemptsChallenged               = "subscription_billing_attempts/challenged"
	TopicSubscriptionBillingAttemptsFailure                  = "subscription_billing_attempts/failure"
	TopicSubscriptionBillingAttemptsSuccess                  = "subscription_billing_attempts/success"
	TopicSubscriptionBillingCycleEditsCreate                 = "subscription_billing_cycle_edits/create"
	TopicSubscriptionBillingCycleEditsDelete                 = "subscription_billing_cycle_edits/delete"
	TopicSubscriptionBillingCycleEditsUpdate                 = "subscription_billing_cycle_edits/update"
	TopicSubscriptionContractsCreate                         = "subscription_contracts/create"
	TopicSubscriptionContractsUpdate                         = "subscription_contracts/update"
	TopicTenderTransactionsCreate                            = "tender_transactions/create"
	TopicThemesCreate                                        = "themes/create"
	TopicThemesDelete                                        = "themes/delete"
	TopicThemesPublish                                       = "themes/publish"
	TopicThemesUpdate                                        = "themes/update"
	// Inner topics only used to send message to queue, it won't triggered by webhook events
	TopicInnerShopInstalled = "shops/installed"
)

var Topics = []string{
	TopicAppSubscriptionsUpdate,
	TopicAppUninstalled,
	TopicBulkOperationsFinish,
	TopicCartsCreate,
	TopicCartsUpdate,
	TopicCheckoutsCreate,
	TopicCheckoutsDelete,
	TopicCheckoutsUpdate,
	TopicCollectionListingsAdd,
	TopicCollectionListingsRemove,
	TopicCollectionListingsUpdate,
	TopicCollectionsCreate,
	TopicCollectionsDelete,
	TopicCollectionsUpdate,
	TopicCustomerGroupsCreate,
	TopicCustomerGroupsDelete,
	TopicCustomerGroupsUpdate,
	TopicCustomerPaymentMethodsCreate,
	TopicCustomerPaymentMethodsRevoke,
	TopicCustomerPaymentMethodsUpdate,
	TopicCustomersMarketingConsentUpdate,
	TopicCustomersCreate,
	TopicCustomersDelete,
	TopicCustomersDisable,
	TopicCustomersEnable,
	TopicCustomersUpdate,
	TopicDisputesCreate,
	TopicDisputesUpdate,
	TopicDomainsCreate,
	TopicDomainsDestroy,
	TopicDomainsUpdate,
	TopicDraftOrdersCreate,
	TopicDraftOrdersDelete,
	TopicDraftOrdersUpdate,
	TopicFulfillmentEventsCreate,
	TopicFulfillmentEventsDelete,
	TopicFulfillmentOrdersCancellationRequestAccepted,
	TopicFulfillmentOrdersCancellationRequestRejected,
	TopicFulfillmentOrdersCancellationRequestSubmitted,
	TopicFulfillmentOrdersCancelled,
	TopicFulfillmentOrdersFulfillmentRequestAccepted,
	TopicFulfillmentOrdersFulfillmentRequestRejected,
	TopicFulfillmentOrdersFulfillmentRequestSubmitted,
	TopicFulfillmentOrdersFulfillmentServiceFailedToComplete,
	TopicFulfillmentOrdersHoldReleased,
	TopicFulfillmentOrdersLineItemsPreparedForLocalDelivery,
	TopicFulfillmentOrdersLineItemsPreparedForPickup,
	TopicFulfillmentOrdersMoved,
	TopicFulfillmentOrdersOrderRoutingComplete,
	TopicFulfillmentOrdersPlacedOnHold,
	TopicFulfillmentOrdersRescheduled,
	TopicFulfillmentOrdersScheduledFulfillmentOrderReady,
	TopicFulfillmentsCreate,
	TopicFulfillmentsUpdate,
	TopicInventoryItemsCreate,
	TopicInventoryItemsDelete,
	TopicInventoryItemsUpdate,
	TopicInventoryLevelsConnect,
	TopicInventoryLevelsDisconnect,
	TopicInventoryLevelsUpdate,
	TopicLocalesCreate,
	TopicLocalesUpdate,
	TopicLocationsActivate,
	TopicLocationsCreate,
	TopicLocationsDeactivate,
	TopicLocationsDelete,
	TopicLocationsUpdate,
	TopicMarketsCreate,
	TopicMarketsDelete,
	TopicMarketsUpdate,
	TopicOrderTransactionsCreate,
	TopicOrdersCancelled,
	TopicOrdersCreate,
	TopicOrdersDelete,
	TopicOrdersEdited,
	TopicOrdersFulfilled,
	TopicOrdersPaid,
	TopicOrdersPartiallyFulfilled,
	TopicOrdersUpdated,
	TopicPaymentSchedulesDue,
	TopicProductListingsAdd,
	TopicProductListingsRemove,
	TopicProductListingsUpdate,
	TopicProductsCreate,
	TopicProductsDelete,
	TopicProductsUpdate,
	TopicProfilesCreate,
	TopicProfilesDelete,
	TopicProfilesUpdate,
	TopicRefundsCreate,
	TopicScheduledProductListingsAdd,
	TopicScheduledProductListingsRemove,
	TopicScheduledProductListingsUpdate,
	TopicSellingPlanGroupsCreate,
	TopicSellingPlanGroupsDelete,
	TopicSellingPlanGroupsUpdate,
	TopicShopUpdate,
	TopicSubscriptionBillingAttemptsChallenged,
	TopicSubscriptionBillingAttemptsFailure,
	TopicSubscriptionBillingAttemptsSuccess,
	TopicSubscriptionBillingCycleEditsCreate,
	TopicSubscriptionBillingCycleEditsDelete,
	TopicSubscriptionBillingCycleEditsUpdate,
	TopicSubscriptionContractsCreate,
	TopicSubscriptionContractsUpdate,
	TopicTenderTransactionsCreate,
	TopicThemesCreate,
	TopicThemesDelete,
	TopicThemesPublish,
	TopicThemesUpdate,
}

var QueueForProducts = []string{
	TopicProductListingsAdd,
	TopicProductListingsRemove,
	TopicProductListingsUpdate,
	TopicProductsCreate,
	TopicProductsDelete,
	TopicProductsUpdate,
	TopicInventoryItemsCreate,
	TopicInventoryItemsDelete,
	TopicInventoryItemsUpdate,
	TopicInventoryLevelsConnect,
	TopicInventoryLevelsDisconnect,
	TopicInventoryLevelsUpdate,
	// low priority
	TopicCollectionListingsAdd,
	TopicCollectionListingsRemove,
	TopicCollectionListingsUpdate,
	TopicCollectionsCreate,
	TopicCollectionsDelete,
	TopicCollectionsUpdate,
}
var QueueForOrders = []string{
	TopicOrderTransactionsCreate,
	TopicOrdersCancelled,
	TopicOrdersCreate,
	TopicOrdersDelete,
	TopicOrdersEdited,
	TopicOrdersFulfilled,
	TopicOrdersPaid,
	TopicOrdersPartiallyFulfilled,
	TopicOrdersUpdated,
	// low priority
	TopicDraftOrdersCreate,
	TopicDraftOrdersDelete,
	TopicDraftOrdersUpdate,
	TopicFulfillmentEventsCreate,
	TopicFulfillmentEventsDelete,
	TopicFulfillmentOrdersCancellationRequestAccepted,
	TopicFulfillmentOrdersCancellationRequestRejected,
	TopicFulfillmentOrdersCancellationRequestSubmitted,
	TopicFulfillmentOrdersCancelled,
	TopicFulfillmentOrdersFulfillmentRequestAccepted,
	TopicFulfillmentOrdersFulfillmentRequestRejected,
	TopicFulfillmentOrdersFulfillmentRequestSubmitted,
	TopicFulfillmentOrdersFulfillmentServiceFailedToComplete,
	TopicFulfillmentOrdersHoldReleased,
	TopicFulfillmentOrdersLineItemsPreparedForLocalDelivery,
	TopicFulfillmentOrdersLineItemsPreparedForPickup,
	TopicFulfillmentOrdersMoved,
	TopicFulfillmentOrdersOrderRoutingComplete,
	TopicFulfillmentOrdersPlacedOnHold,
	TopicFulfillmentOrdersRescheduled,
	TopicFulfillmentOrdersScheduledFulfillmentOrderReady,
	TopicFulfillmentsCreate,
	TopicFulfillmentsUpdate,
}

var QueueForMiscs = []string{
	TopicShopUpdate,
	TopicAppSubscriptionsUpdate,
	TopicAppUninstalled,
	TopicThemesCreate,
	TopicThemesDelete,
	TopicThemesPublish,
	TopicThemesUpdate,
	// medium priority
	TopicProfilesCreate,
	TopicProfilesDelete,
	TopicProfilesUpdate,
	TopicRefundsCreate,
	TopicPaymentSchedulesDue,
	TopicDisputesCreate,
	TopicBulkOperationsFinish,
	TopicCartsCreate,
	TopicCartsUpdate,
	TopicCheckoutsCreate,
	TopicCheckoutsDelete,
	TopicCheckoutsUpdate,
	TopicSellingPlanGroupsCreate,
	TopicSellingPlanGroupsDelete,
	TopicSellingPlanGroupsUpdate,
	TopicScheduledProductListingsAdd,
	TopicScheduledProductListingsRemove,
	TopicScheduledProductListingsUpdate,
	TopicSubscriptionBillingAttemptsChallenged,
	TopicSubscriptionBillingAttemptsFailure,
	TopicSubscriptionBillingAttemptsSuccess,
	TopicSubscriptionBillingCycleEditsCreate,
	TopicSubscriptionBillingCycleEditsDelete,
	TopicSubscriptionBillingCycleEditsUpdate,
	TopicSubscriptionContractsCreate,
	TopicSubscriptionContractsUpdate,
	// low priority
	TopicLocalesCreate,
	TopicLocalesUpdate,
	TopicLocationsActivate,
	TopicLocationsCreate,
	TopicLocationsDeactivate,
	TopicLocationsDelete,
	TopicLocationsUpdate,
	TopicMarketsCreate,
	TopicMarketsDelete,
	TopicMarketsUpdate,
	TopicCustomerGroupsCreate,
	TopicCustomerGroupsDelete,
	TopicCustomerGroupsUpdate,
	TopicCustomerPaymentMethodsCreate,
	TopicCustomerPaymentMethodsRevoke,
	TopicCustomerPaymentMethodsUpdate,
	TopicCustomersMarketingConsentUpdate,
	TopicCustomersCreate,
	TopicCustomersDelete,
	TopicCustomersDisable,
	TopicCustomersEnable,
	TopicCustomersUpdate,
	TopicTenderTransactionsCreate,
	TopicDisputesUpdate,
	TopicDomainsCreate,
	TopicDomainsDestroy,
	TopicDomainsUpdate,
	// Innter topics
	TopicInnerShopInstalled,
}
