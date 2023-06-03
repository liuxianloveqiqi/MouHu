package svc

import (
	"MouHu/service/app/like/model"
	"MouHu/service/app/like/rpc/internal/config"
	"MouHu/service/common/init_db"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config      config.Config
	FollowModel model.FollowModel
	Rdb         *redis.Client
	Mdb         *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	coon := sqlx.NewMysql(c.Mysql.DataSource)
	MysqlDb := init_db.InitGorm(c.Mysql.DataSource)
	MysqlDb.AutoMigrate(&model.Follow{})
	redisDb := init_db.InitRedis(c.RedisClient.Host)
	return &ServiceContext{
		Config:      c,
		Mdb:         MysqlDb,
		FollowModel: model.NewFollowModel(coon, c.CacheRedis),
		Rdb:         redisDb,
	}
}
