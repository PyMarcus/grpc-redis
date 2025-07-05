package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct{
	Host string 
	Port string 
	redisClient *redis.Client 
}

func NewRedisRepository(ip, port string) *RedisRepository{
	return &RedisRepository{ip, port, redis.NewClient(&redis.Options{
		Addr:fmt.Sprintf("%s:%s",ip,port),
		Password: "",
		DB: 0,
		PoolSize: 100,
		MinIdleConns: 10,
	})}
}

func (r *RedisRepository) Connect(ctx context.Context){
	_, err := r.redisClient.Ping(ctx).Result()

	if err != nil{
		log.Fatalf("Fail to connect with redis: %v", err)
	}

	log.Println("Connected with redis!")
}

// Set: salva chave e valor no redis, tempo por padrao, de expiracao, nao ha
func (r *RedisRepository) Set(ctx *context.Context, key, value string) error{
	err := r.redisClient.Set(*ctx, key, value, 0).Err()

	if err != nil{
		log.Printf("Fail to set=> key: %s with value: %s ", key, value)
		return err 
	}
	return nil 
}

// Get: obtem o dado do redis com base na chave
func (r *RedisRepository) Get(ctx *context.Context, key string) (string, error){
	value, err := r.redisClient.Get(*ctx, key).Result()

	if err != nil{
		log.Printf("Fail to get=> key: %s", key)
		return "", err 
	}
	return value, nil 
}

// Del: remove o dado do redis com base na chave
func (r *RedisRepository) Del(ctx *context.Context, key string) (error){
	value, err := r.redisClient.Del(*ctx, key).Result()

	if err != nil{
		log.Printf("Fail to delete=> key: %s", key)
		return err 
	}

	log.Println("Removed from redis:", value)
	return nil 
}
