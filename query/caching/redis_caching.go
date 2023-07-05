package caching

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

const urlPrefix = "url:"

var redisClusterClient *redis.ClusterClient
var urlCachingDuration = getUrlCachingDuration()

func ConnectToRedis() {
	log.Println("Connecting to Redis...")
	redisCluster := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    getRedisNodesFromEnv(),
		Password: os.Getenv("REDIS_PASSWORD"),
	})
	_, err := redisCluster.Ping().Result()
	if err != nil {
		log.Fatalln(err)
	}
	redisClusterClient = redisCluster
	log.Println("Connected to Redis!")
}

func CacheURL(path, url string) error {
	err := redisClusterClient.Set(urlPrefix+path, url, urlCachingDuration).Err()
	return err
}

func GetURLFromCache(path string) (string, error) {
	cachedURL, err := redisClusterClient.Get(urlPrefix + path).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return cachedURL, nil
}

func getUrlCachingDuration() time.Duration {
	urlCachingTime, err := strconv.Atoi(os.Getenv("URL_CACHING_TIME"))
	if err != nil {
		return 180 * time.Second
	}
	return time.Duration(urlCachingTime) * time.Second
}

func getRedisNodesFromEnv() []string {
	nodes := os.Getenv("REDIS_HOSTS")
	return strings.Split(nodes, ",")
}
