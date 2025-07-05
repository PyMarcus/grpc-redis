package repository

import (
	"context"
	"testing"

	"github.com/go-redis/redismock/v9"
)

func TestSet(t *testing.T){
	ctx := context.Background()

	k := "test_k"
	v := "test_v"

	db, mockRedisClient := redismock.NewClientMock()

	mockRedisClient.ExpectSet(k, v, 0).SetVal("OK")

	err := db.Set(ctx, k, v, 0).Err()

	if err != nil{
		t.Fatalf("Fail to test set %v", err)
	}

	if err := mockRedisClient.ExpectationsWereMet(); err != nil {
		t.Errorf("Fail to run mock: %v", err)
	}

}

func TestGet(t *testing.T){
	ctx := context.Background()

	k := "test_k"
	v := "test_v"

	db, mockRedisClient := redismock.NewClientMock()

	mockRedisClient.ExpectGet(k).SetVal(v)

	val, err := db.Get(ctx, k).Result()

	if err != nil{
		t.Fatalf("Fail to test set %v", err)
	}

	if val != v{
		t.Fatalf("Fail! Expected: %s - Received: %s", v, val)
	}

	if err := mockRedisClient.ExpectationsWereMet(); err != nil {
		t.Errorf("Fail to run mock: %v", err)
	}

}

func TestDel(t *testing.T){
	ctx := context.Background()

	k := "test_k"

	db, mockRedisClient := redismock.NewClientMock()

	mockRedisClient.ExpectDel(k).SetVal(1)

	deleted, err := db.Del(ctx, k).Result()
	if err != nil {
		t.Fatalf("DEL failed: %v", err)
	}
	if deleted != 1 {
		t.Errorf("Expected 1 key deleted, but %d", deleted)
	}

	if err := mockRedisClient.ExpectationsWereMet(); err != nil {
		t.Errorf("Fail to run mock: %v", err)
	}

}
