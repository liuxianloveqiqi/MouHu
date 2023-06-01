package logic

import (
	"MouHu/service/app/qa/model"
	"context"
	"encoding/json"
	"errors"
	"github.com/nsqio/go-nsq"
	"time"

	"MouHu/service/app/qa/rpc/internal/svc"
	"MouHu/service/app/qa/rpc/types/qa"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnsQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnsQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnsQuestionLogic {
	return &AnsQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 实现NSQ的HandleMessage方法
func (l *AnsQuestionLogic) HandleMessage(message *nsq.Message) error {
	// 解析接收到的消息为Answer结构体
	var answer qa.AnsQuestionReq
	err := json.Unmarshal(message.Body, &answer)
	if err != nil {
		logx.Error("解析消息失败 ", err)
		return errors.New("解析消息失败 :" + err.Error())
	}
	// 处理问题的逻辑
	ans := model.Answer{
		UserId:     answer.UserId,
		QuestionId: answer.QuestionId,
		Content:    answer.Content,
		ParentId:   0,
		IsAnswer:   1,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	if err = l.svcCtx.Mdb.Create(&ans).Error; err != nil {
		logx.Error("新建问题err: ", err)
		return err
	}
	message.Finish()
	return nil
}

// 初始化消费者并启动监听
func InitConsumerAns(l *AnsQuestionLogic) {
	// 创建NSQ消费者
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("answer_topic", "answer_channel", config)
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
func (l *AnsQuestionLogic) AnsQuestion(in *qa.AnsQuestionReq) (*qa.CommonResp, error) {
	// todo: add your logic here and delete this line
	// 将问题转换为JSON格式
	body, err := json.Marshal(in)
	if err != nil {
		logx.Error("解析消息失败 ", err)
		return nil, errors.New("解析消息失败 :" + err.Error())
	}
	// 发送问题消息到NSQ的指定主题
	err = l.svcCtx.Product.Publish("answer_topic", body)
	if err != nil {
		logx.Error("publish消息失败:", err)
		return nil, errors.New("publish消息失败 :" + err.Error())
	}
	done := make(chan bool)

	go func() {
		InitConsumerAns(l)
		done <- true
	}()
	<-done
	return &qa.CommonResp{
		Code:    0,
		Message: "Success!",
	}, nil

}
