package tag

import (
	"encoding/json"
	"oppapi/internal/cache"
	"oppapi/internal/model"
)

const PaymentEntityName = "payment"

type PaymentCache struct {
	cache *cache.Redis
}

// NewPaymentCache uses the connection allocated in the init of the cache module.
// It does not open each connection per request but it reuses the initial one.
func NewPaymentCache() (*PaymentCache, error) {
	cache, err := cache.WithRedis()
	if err != nil {
		return nil, err
	}
	return &PaymentCache{
		cache: cache,
	}, nil
}

// Check does a dive into the redis cache for an id.
func (tc *PaymentCache) Check(id string) (model.Payment, bool, error) {
	val, err := tc.cache.Client.Get(PaymentEntityName).Result()
	if err != nil {
		return model.Payment{}, false, err
	}
	var payment model.Payment
	json.Unmarshal([]byte(val), &payment)
	return payment, true, nil
}

// Flush purges cache for single element, when it was modified or deleted.
func (tc *PaymentCache) Flush(id string) error {
	return nil
}

// Purge makes empty the whole cache of the collection.
func (tc *PaymentCache) Purge() error {
	return nil
}
