// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: qa.proto

package qa

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Qa_PubQuestion_FullMethodName       = "/question.qa/PubQuestion"
	Qa_AnsQuestion_FullMethodName       = "/question.qa/AnsQuestion"
	Qa_ComAnswer_FullMethodName         = "/question.qa/ComAnswer"
	Qa_GetQuestion_FullMethodName       = "/question.qa/GetQuestion"
	Qa_GetAnswer_FullMethodName         = "/question.qa/GetAnswer"
	Qa_GetCommit_FullMethodName         = "/question.qa/GetCommit"
	Qa_DelQuestion_FullMethodName       = "/question.qa/DelQuestion"
	Qa_DelAnswerOrCommit_FullMethodName = "/question.qa/DelAnswerOrCommit"
	Qa_AltQuestion_FullMethodName       = "/question.qa/AltQuestion"
	Qa_AltAnswerOrCommit_FullMethodName = "/question.qa/AltAnswerOrCommit"
)

// QaClient is the client API for Qa service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QaClient interface {
	PubQuestion(ctx context.Context, in *PubQuestionReq, opts ...grpc.CallOption) (*CommonResp, error)
	AnsQuestion(ctx context.Context, in *AnsQuestionReq, opts ...grpc.CallOption) (*CommonResp, error)
	ComAnswer(ctx context.Context, in *ComAnswerReq, opts ...grpc.CallOption) (*CommonResp, error)
	GetQuestion(ctx context.Context, in *GetAnswerReq, opts ...grpc.CallOption) (*QuestionList, error)
	GetAnswer(ctx context.Context, in *GetAnswerReq, opts ...grpc.CallOption) (*GetAnswerResp, error)
	GetCommit(ctx context.Context, in *GetCommitReq, opts ...grpc.CallOption) (*GetCommitResp, error)
	DelQuestion(ctx context.Context, in *DelQuestionReq, opts ...grpc.CallOption) (*CommonResp, error)
	DelAnswerOrCommit(ctx context.Context, in *DelAnswerOrCommitReq, opts ...grpc.CallOption) (*CommonResp, error)
	AltQuestion(ctx context.Context, in *AltQuestionReq, opts ...grpc.CallOption) (*CommonResp, error)
	AltAnswerOrCommit(ctx context.Context, in *AltAnswerOrCommitReq, opts ...grpc.CallOption) (*CommonResp, error)
}

type qaClient struct {
	cc grpc.ClientConnInterface
}

func NewQaClient(cc grpc.ClientConnInterface) QaClient {
	return &qaClient{cc}
}

func (c *qaClient) PubQuestion(ctx context.Context, in *PubQuestionReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Qa_PubQuestion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qaClient) AnsQuestion(ctx context.Context, in *AnsQuestionReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Qa_AnsQuestion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qaClient) ComAnswer(ctx context.Context, in *ComAnswerReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Qa_ComAnswer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qaClient) GetQuestion(ctx context.Context, in *GetAnswerReq, opts ...grpc.CallOption) (*QuestionList, error) {
	out := new(QuestionList)
	err := c.cc.Invoke(ctx, Qa_GetQuestion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qaClient) GetAnswer(ctx context.Context, in *GetAnswerReq, opts ...grpc.CallOption) (*GetAnswerResp, error) {
	out := new(GetAnswerResp)
	err := c.cc.Invoke(ctx, Qa_GetAnswer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qaClient) GetCommit(ctx context.Context, in *GetCommitReq, opts ...grpc.CallOption) (*GetCommitResp, error) {
	out := new(GetCommitResp)
	err := c.cc.Invoke(ctx, Qa_GetCommit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qaClient) DelQuestion(ctx context.Context, in *DelQuestionReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Qa_DelQuestion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qaClient) DelAnswerOrCommit(ctx context.Context, in *DelAnswerOrCommitReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Qa_DelAnswerOrCommit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qaClient) AltQuestion(ctx context.Context, in *AltQuestionReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Qa_AltQuestion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qaClient) AltAnswerOrCommit(ctx context.Context, in *AltAnswerOrCommitReq, opts ...grpc.CallOption) (*CommonResp, error) {
	out := new(CommonResp)
	err := c.cc.Invoke(ctx, Qa_AltAnswerOrCommit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QaServer is the server API for Qa service.
// All implementations must embed UnimplementedQaServer
// for forward compatibility
type QaServer interface {
	PubQuestion(context.Context, *PubQuestionReq) (*CommonResp, error)
	AnsQuestion(context.Context, *AnsQuestionReq) (*CommonResp, error)
	ComAnswer(context.Context, *ComAnswerReq) (*CommonResp, error)
	GetQuestion(context.Context, *GetAnswerReq) (*QuestionList, error)
	GetAnswer(context.Context, *GetAnswerReq) (*GetAnswerResp, error)
	GetCommit(context.Context, *GetCommitReq) (*GetCommitResp, error)
	DelQuestion(context.Context, *DelQuestionReq) (*CommonResp, error)
	DelAnswerOrCommit(context.Context, *DelAnswerOrCommitReq) (*CommonResp, error)
	AltQuestion(context.Context, *AltQuestionReq) (*CommonResp, error)
	AltAnswerOrCommit(context.Context, *AltAnswerOrCommitReq) (*CommonResp, error)
	mustEmbedUnimplementedQaServer()
}

// UnimplementedQaServer must be embedded to have forward compatible implementations.
type UnimplementedQaServer struct {
}

func (UnimplementedQaServer) PubQuestion(context.Context, *PubQuestionReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PubQuestion not implemented")
}
func (UnimplementedQaServer) AnsQuestion(context.Context, *AnsQuestionReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnsQuestion not implemented")
}
func (UnimplementedQaServer) ComAnswer(context.Context, *ComAnswerReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComAnswer not implemented")
}
func (UnimplementedQaServer) GetQuestion(context.Context, *GetAnswerReq) (*QuestionList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuestion not implemented")
}
func (UnimplementedQaServer) GetAnswer(context.Context, *GetAnswerReq) (*GetAnswerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAnswer not implemented")
}
func (UnimplementedQaServer) GetCommit(context.Context, *GetCommitReq) (*GetCommitResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommit not implemented")
}
func (UnimplementedQaServer) DelQuestion(context.Context, *DelQuestionReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelQuestion not implemented")
}
func (UnimplementedQaServer) DelAnswerOrCommit(context.Context, *DelAnswerOrCommitReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelAnswerOrCommit not implemented")
}
func (UnimplementedQaServer) AltQuestion(context.Context, *AltQuestionReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AltQuestion not implemented")
}
func (UnimplementedQaServer) AltAnswerOrCommit(context.Context, *AltAnswerOrCommitReq) (*CommonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AltAnswerOrCommit not implemented")
}
func (UnimplementedQaServer) mustEmbedUnimplementedQaServer() {}

// UnsafeQaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QaServer will
// result in compilation errors.
type UnsafeQaServer interface {
	mustEmbedUnimplementedQaServer()
}

func RegisterQaServer(s grpc.ServiceRegistrar, srv QaServer) {
	s.RegisterService(&Qa_ServiceDesc, srv)
}

func _Qa_PubQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PubQuestionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QaServer).PubQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Qa_PubQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QaServer).PubQuestion(ctx, req.(*PubQuestionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qa_AnsQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnsQuestionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QaServer).AnsQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Qa_AnsQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QaServer).AnsQuestion(ctx, req.(*AnsQuestionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qa_ComAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComAnswerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QaServer).ComAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Qa_ComAnswer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QaServer).ComAnswer(ctx, req.(*ComAnswerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qa_GetQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnswerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QaServer).GetQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Qa_GetQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QaServer).GetQuestion(ctx, req.(*GetAnswerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qa_GetAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAnswerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QaServer).GetAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Qa_GetAnswer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QaServer).GetAnswer(ctx, req.(*GetAnswerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qa_GetCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QaServer).GetCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Qa_GetCommit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QaServer).GetCommit(ctx, req.(*GetCommitReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qa_DelQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelQuestionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QaServer).DelQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Qa_DelQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QaServer).DelQuestion(ctx, req.(*DelQuestionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qa_DelAnswerOrCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelAnswerOrCommitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QaServer).DelAnswerOrCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Qa_DelAnswerOrCommit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QaServer).DelAnswerOrCommit(ctx, req.(*DelAnswerOrCommitReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qa_AltQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AltQuestionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QaServer).AltQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Qa_AltQuestion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QaServer).AltQuestion(ctx, req.(*AltQuestionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Qa_AltAnswerOrCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AltAnswerOrCommitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QaServer).AltAnswerOrCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Qa_AltAnswerOrCommit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QaServer).AltAnswerOrCommit(ctx, req.(*AltAnswerOrCommitReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Qa_ServiceDesc is the grpc.ServiceDesc for Qa service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Qa_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "question.qa",
	HandlerType: (*QaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PubQuestion",
			Handler:    _Qa_PubQuestion_Handler,
		},
		{
			MethodName: "AnsQuestion",
			Handler:    _Qa_AnsQuestion_Handler,
		},
		{
			MethodName: "ComAnswer",
			Handler:    _Qa_ComAnswer_Handler,
		},
		{
			MethodName: "GetQuestion",
			Handler:    _Qa_GetQuestion_Handler,
		},
		{
			MethodName: "GetAnswer",
			Handler:    _Qa_GetAnswer_Handler,
		},
		{
			MethodName: "GetCommit",
			Handler:    _Qa_GetCommit_Handler,
		},
		{
			MethodName: "DelQuestion",
			Handler:    _Qa_DelQuestion_Handler,
		},
		{
			MethodName: "DelAnswerOrCommit",
			Handler:    _Qa_DelAnswerOrCommit_Handler,
		},
		{
			MethodName: "AltQuestion",
			Handler:    _Qa_AltQuestion_Handler,
		},
		{
			MethodName: "AltAnswerOrCommit",
			Handler:    _Qa_AltAnswerOrCommit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qa.proto",
}
