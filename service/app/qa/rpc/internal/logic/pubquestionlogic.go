package logic

import (
	"MouHu/service/app/qa/model"
	"MouHu/service/app/qa/rpc/internal/svc"
	"MouHu/service/app/qa/rpc/types/qa"
	"context"
	"encoding/json"
	"errors"
	"github.com/nsqio/go-nsq"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type PubQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPubQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PubQuestionLogic {
	return &PubQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 实现NSQ的HandleMessage方法
func (l *PubQuestionLogic) HandleMessage(message *nsq.Message) error {
	// 解析接收到的消息为Question结构体
	var question qa.PubQuestionReq
	err := json.Unmarshal(message.Body, &question)
	if err != nil {
		logx.Error("解析消息失败 ", err)
		return errors.New("解析消息失败 :" + err.Error())
	}
	// 处理问题的逻辑
	que := model.Question{
		UserId:     question.UserId,
		Title:      question.Title,
		Content:    question.Content,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	if err = l.svcCtx.Mdb.Create(&que).Error; err != nil {
		logx.Error("新建问题err: ", err)
		return err
	}
	message.Finish()
	return nil
}

// 初始化消费者并启动监听
func InitConsumerQue(l *PubQuestionLogic) {
	// 创建NSQ消费者
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("question_topic", "question_channel", config)
	if err != nil {
		logx.Error("不能创建NSQ consumer:", err)
		panic(err)
	}

	// 设置消息处理函数
	consumer.AddHandler(l)

	// 连接到NSQD服务
	err = consumer.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		logx.Error("无法连接到NSQLookupd:", err)
		panic(err)
	}

}
func (l *PubQuestionLogic) PubQuestion(in *qa.PubQuestionReq) (*qa.CommonResp, error) {
	// todo: add your logic here and delete this line
	// 将问题转换为JSON格式
	body, err := json.Marshal(in)
	if err != nil {
		logx.Error("解析消息失败 ", err)
		return nil, errors.New("解析消息失败 :" + err.Error())
	}
	// 发送问题消息到NSQ的指定主题
	err = l.svcCtx.Product.Publish("question_topic", body)
	if err != nil {
		logx.Error("publish消息失败:", err)
		return nil, errors.New("publish消息失败 :" + err.Error())
	}
	// 并发消费者
	go func() {
		InitConsumerQue(l)
	}()
	return &qa.CommonResp{
		Code:    0,
		Message: "Success!",
	}, nil
}
