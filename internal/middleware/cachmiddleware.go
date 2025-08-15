package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/bradfitz/gomemcache/memcache"
	
)
var MC = memcache.New("127.0.0.1:11211")
func CacheMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		



		

		



	}
} 


func CacheExchangerates(key string, quotes map[string]float64) error {
	jsonData, err := json.Marshal(quotes)
	if err != nil {
		return err
	}
	return MC.Set(&memcache.Item{
		Key: key,
		Value: jsonData,
		Expiration: 3600,
	})
}

func GetCachedRates(key string) (map[string]float64, error) {
	items, err := MC.Get(key)
	if err != nil {
		return nil, err
	} 

	var quotes map[string]float64
	err = json.Unmarshal(items.Value, &quotes)
	if err != nil {
		return nil, err

	}
	return quotes, nil
}
