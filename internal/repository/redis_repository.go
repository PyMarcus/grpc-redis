package repository

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(addr string) *RedisRepository {
	client := redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     "", 
		DB:           0,
		PoolSize:     100, 
		MinIdleConns: 10,  
	})
	return &RedisRepository{client: client}
}

func (r *RedisRepository) Connect(ctx context.Context) error {
	if err := r.client.Ping(ctx).Err(); err != nil {
		return err 
	}

	log.Println("Conectado com sucesso ao Redis!")
	return nil
}

func (r *RedisRepository) Set(ctx context.Context, key, value string) error {
	err := r.client.Set(ctx, key, value, 0).Err()

	if err != nil {
		log.Printf("Falha ao salvar no Redis => key: %s", key)
		return err
	}
	return nil
}

func (r *RedisRepository) Get(ctx context.Context, key string) (string, error) {
	value, err := r.client.Get(ctx, key).Result()

	if err != nil {
		if err != redis.Nil {
			log.Printf("Falha ao buscar no Redis => key: %s", key)
		}
		return "", err
	}
	return value, nil
}

func (r *RedisRepository) Del(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()

	if err != nil {
		log.Printf("Falha ao deletar no Redis => key: %s", key)
		return err
	}

	return nil
}