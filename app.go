package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"gopkg.in/redis.v5"
)

var clientRedis = redis.NewClient(&redis.Options{
	Addr:     "db_redis:6379",
	Password: "",
	DB:       0,
})

var (
	errorEmptyKeyAssing = errors.New("Key does not empty ")
	errorNotFoundErrMsg = errors.New("Cache Info Not Found ")
)

// Cache interface
type Cache interface {
	// cache provider NAME ( redis, memcache etc...)
	Name() string

	Set(key string, data interface{}) error

	Get(key string) (interface{}, error)
}

type RedisCache struct {
	data map[string]interface{}
}

type MemoryCache struct {
	data map[string]interface{}
}

// just struct name
func (c *MemoryCache) Name() string {
	return "MemoryCache"
}
func (c *RedisCache) Name() string {
	return "RedisCache"
}

// Get from RedisCache
func (c *RedisCache) Get(key string) (interface{}, error) {
	return clientRedis.Get(key).Result()
}

// Get from MemoryCache
func (c *MemoryCache) Get(key string) (interface{}, error) {
	if value, ok := c.data[key]; ok == true {
		return value, nil
	}
	return interface{}(""), errorNotFoundErrMsg
}

// Set key -> MemoryCache
func (c *MemoryCache) Set(key string, data interface{}) error {
	if key == "" {
		return errorEmptyKeyAssing
	}
	c.data[key] = data
	return nil
}

// Set key -> RedisCache
func (c *RedisCache) Set(key string, data interface{}) error {
	if key == "" {
		return errorEmptyKeyAssing
	}
	return clientRedis.Set(key, data, time.Second*100).Err()
}

func getDataFromCache(key string, cache Cache) (interface{}, error) {
	// may be run another logic
	return cache.Get(key)
}

func setDataToCache(key string, data interface{}, cache Cache) error {
	// may be run another logic
	return cache.Set(key, data)
}

func getNew_MemoryCache() *MemoryCache {
	return &MemoryCache{data: make(map[string]interface{})}
}
func getNew_RedisCache() *RedisCache {
	return &RedisCache{}
}

func main() {

	cacheRedis := getNew_RedisCache()
	cacheMC := getNew_MemoryCache()
	cache_key := "first_bike"

	/*
		pong, err := client.Ping().Result()
		if err != nil {
			panic(err)
		}

		print("Redis ping response ; ", pong)*/

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pong, err := clientRedis.Ping().Result()
		if err != nil {
			panic(err)
		}

		if err := setDataToCache(cache_key, interface{}("Kawasaki KLE 500"), cacheMC); err != nil {
			fmt.Println("/MemoryCache/ return err : ", err)
		} else {
			fmt.Println("/MemoryCache/ set to data with: ", string(cache_key))
		}

		if err := setDataToCache(cache_key, interface{}("Kawasaki KLE 650"), cacheRedis); err != nil {
			fmt.Println("/Redis/ return err : ", err)
		} else {
			fmt.Println("/Redis/ set to data with: ", string(cache_key))
		}

		fmt.Fprintf(w, "Hello from cache interface implementation. Redis ping response = %s \n", pong)
	})

	http.HandleFunc("/getRedis", func(w http.ResponseWriter, r *http.Request) {

		dataFromRedis, err := getDataFromCache(cache_key, cacheRedis)
		if err != nil {
			fmt.Println("return err : ", err)
		}

		fmt.Fprintf(w, "Data from Redis : %s \n", dataFromRedis)
	})

	http.HandleFunc("/getMC", func(w http.ResponseWriter, r *http.Request) {

		dataFromCache, err := getDataFromCache(cache_key, cacheMC)
		if err != nil {
			fmt.Println("return err : ", err)
		}

		fmt.Fprintf(w, "Data from Lokal Memory : %s \n", dataFromCache)
	})

	log.Fatal(http.ListenAndServe(":5000", nil))

}
