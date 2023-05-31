package svc

import (
	"MouHu/service/app/qa/model"
	"MouHu/service/app/qa/rpc/internal/config"
	"MouHu/service/common/init_db"
	"github.com/nsqio/go-nsq"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config        config.Config
	QuestionModel model.QuestionModel
	AnswerModel   model.AnswerModel
	Rdb           *redis.Client
	Mdb           *gorm.DB
	Product       *nsq.Producer
}

func NewServiceContext(c config.Config) *ServiceContext {
	coon := sqlx.NewMysql(c.Mysql.DataSource)
	MysqlDb := init_db.InitGorm(c.Mysql.DataSource)
	MysqlDb.AutoMigrate(&model.Question{}, &model.Answer{})
	redisDb := init_db.InitRedis(c.RedisClient.Host)
	product := init_db.InitNsqProduct(c.Nsq.Host)
	return &ServiceContext{
		Config:        c,
		Mdb:           MysqlDb,
		QuestionModel: model.NewQuestionModel(coon, c.CacheRedis),
		AnswerModel:   model.NewAnswerModel(coon, c.CacheRedis),
		Rdb:           redisDb,
		Product:       product,
	}
}
