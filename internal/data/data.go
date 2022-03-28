package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kratos-layout/internal/conf"
)

// ProviderDataSet is data providers.
var ProviderDataSet = wire.NewSet(NewData, NewDb, NewRedis, NewOpenRepo)

// Data .
type Data struct {
	db    *gorm.DB
	redis *redis.Client
}

// NewData .
func NewData(c *conf.Data, db *gorm.DB, rdb *redis.Client, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db:    db,
		redis: rdb,
	}, cleanup, nil
}

// NewDb .
func NewDb(c *conf.Data, logger log.Logger) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(c.Mysql.Dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return
}

// NewRedis .
func NewRedis(c *conf.Data, logger log.Logger) (rdb *redis.Client, err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
	})
	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		panic("failed to connect redis")
	}

	return
}
