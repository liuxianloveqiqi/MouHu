// Code generated by goctl. DO NOT EDIT.
package types

type QuestionList struct {
	QuestionId int64  `json:"questionId"`
	UserID     int64  `json:"userId"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	DeleteTime string `json:"deleteTime"`
}

type AnswerList struct {
	AnswerId   int64  `json:"answerId"`
	UserId     int64  `json:"userId"`
	QuestionId int64  `json:"questionId"`
	ParentId   int64  `json:"parentId"`
	Content    string `json:"content"`
	IsAnswer   bool   `json:"isAnswer"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
	DeleteTime string `json:"deleteTime"`
}

type PubQuestionReq struct {
	Title   string `json:"title" validate:"required,max=30"`
	Content string `json:"content" validate:"required,max=1000"`
}

type AnsQuestionReq struct {
	QuestionId int64  `json:"questionId"`
	Content    string `json:"content"`
}

type ComAnswerReq struct {
	ParentId int64  `json:"answerId"`
	Content  string `json:"content"`
}

type GetQuestionResp struct {
	Question []QuestionList `json:"qa"`
}

type GetAnswerResp struct {
	Answer []AnswerList `json:"answer"`
}

type GetCommitResp struct {
	Commit []AnswerList `json:"commit"`
}

type DelQuestionReq struct {
	QuestionId int64 `json:"questionId"`
}

type DelAnswerOrCommitReq struct {
	AnsOrComId int64 `json:"ansOrComId"`
}

type AltQuestionReq struct {
	QuestionId int64  `json:"questionId"`
	Title      string `json:"title,omitempty" `
	Content    string `json:"content,omitempty" `
}

type AltAnswerOrCommitReq struct {
	AnsOrComId int64  `json:"ansOrComId"`
	Content    string `json:"content"`
}
