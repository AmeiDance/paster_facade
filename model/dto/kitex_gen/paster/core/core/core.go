// Code generated by Kitex v0.0.5. DO NOT EDIT.

package core

import (
	"context"
	"fmt"
	"github.com/ameidance/paster_facade/model/dto/kitex_gen/paster/core"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return coreServiceInfo
}

var coreServiceInfo = newServiceInfo()

func newServiceInfo() *kitex.ServiceInfo {
	serviceName := "Core"
	handlerType := (*core.Core)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetPost":     kitex.NewMethodInfo(getPostHandler, newGetPostArgs, newGetPostResult, false),
		"SavePost":    kitex.NewMethodInfo(savePostHandler, newSavePostArgs, newSavePostResult, false),
		"DeletePost":  kitex.NewMethodInfo(deletePostHandler, newDeletePostArgs, newDeletePostResult, false),
		"GetComments": kitex.NewMethodInfo(getCommentsHandler, newGetCommentsArgs, newGetCommentsResult, false),
		"SaveComment": kitex.NewMethodInfo(saveCommentHandler, newSaveCommentArgs, newSaveCommentResult, false),
		"Check":       kitex.NewMethodInfo(checkHandler, newCheckArgs, newCheckResult, false),
		"Watch":       kitex.NewMethodInfo(watchHandler, newWatchArgs, newWatchResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "core",
	}
	extra["streaming"] = true
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.0.5",
		Extra:           extra,
	}
	return svcInfo
}

func getPostHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core.GetPostRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core.Core).GetPost(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetPostArgs:
		success, err := handler.(core.Core).GetPost(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetPostResult)
		realResult.Success = success
	}
	return nil
}
func newGetPostArgs() interface{} {
	return &GetPostArgs{}
}

func newGetPostResult() interface{} {
	return &GetPostResult{}
}

type GetPostArgs struct {
	Req *core.GetPostRequest
}

func (p *GetPostArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetPostArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetPostArgs) Unmarshal(in []byte) error {
	msg := new(core.GetPostRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetPostArgs_Req_DEFAULT *core.GetPostRequest

func (p *GetPostArgs) GetReq() *core.GetPostRequest {
	if !p.IsSetReq() {
		return GetPostArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPostArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetPostResult struct {
	Success *core.GetPostResponse
}

var GetPostResult_Success_DEFAULT *core.GetPostResponse

func (p *GetPostResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetPostResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetPostResult) Unmarshal(in []byte) error {
	msg := new(core.GetPostResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPostResult) GetSuccess() *core.GetPostResponse {
	if !p.IsSetSuccess() {
		return GetPostResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPostResult) SetSuccess(x interface{}) {
	p.Success = x.(*core.GetPostResponse)
}

func (p *GetPostResult) IsSetSuccess() bool {
	return p.Success != nil
}

func savePostHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core.SavePostRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core.Core).SavePost(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *SavePostArgs:
		success, err := handler.(core.Core).SavePost(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SavePostResult)
		realResult.Success = success
	}
	return nil
}
func newSavePostArgs() interface{} {
	return &SavePostArgs{}
}

func newSavePostResult() interface{} {
	return &SavePostResult{}
}

type SavePostArgs struct {
	Req *core.SavePostRequest
}

func (p *SavePostArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SavePostArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *SavePostArgs) Unmarshal(in []byte) error {
	msg := new(core.SavePostRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SavePostArgs_Req_DEFAULT *core.SavePostRequest

func (p *SavePostArgs) GetReq() *core.SavePostRequest {
	if !p.IsSetReq() {
		return SavePostArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SavePostArgs) IsSetReq() bool {
	return p.Req != nil
}

type SavePostResult struct {
	Success *core.SavePostResponse
}

var SavePostResult_Success_DEFAULT *core.SavePostResponse

func (p *SavePostResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SavePostResult")
	}
	return proto.Marshal(p.Success)
}

func (p *SavePostResult) Unmarshal(in []byte) error {
	msg := new(core.SavePostResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SavePostResult) GetSuccess() *core.SavePostResponse {
	if !p.IsSetSuccess() {
		return SavePostResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SavePostResult) SetSuccess(x interface{}) {
	p.Success = x.(*core.SavePostResponse)
}

func (p *SavePostResult) IsSetSuccess() bool {
	return p.Success != nil
}

func deletePostHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core.DeletePostRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core.Core).DeletePost(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *DeletePostArgs:
		success, err := handler.(core.Core).DeletePost(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeletePostResult)
		realResult.Success = success
	}
	return nil
}
func newDeletePostArgs() interface{} {
	return &DeletePostArgs{}
}

func newDeletePostResult() interface{} {
	return &DeletePostResult{}
}

type DeletePostArgs struct {
	Req *core.DeletePostRequest
}

func (p *DeletePostArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeletePostArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *DeletePostArgs) Unmarshal(in []byte) error {
	msg := new(core.DeletePostRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeletePostArgs_Req_DEFAULT *core.DeletePostRequest

func (p *DeletePostArgs) GetReq() *core.DeletePostRequest {
	if !p.IsSetReq() {
		return DeletePostArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeletePostArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeletePostResult struct {
	Success *core.DeletePostResponse
}

var DeletePostResult_Success_DEFAULT *core.DeletePostResponse

func (p *DeletePostResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeletePostResult")
	}
	return proto.Marshal(p.Success)
}

func (p *DeletePostResult) Unmarshal(in []byte) error {
	msg := new(core.DeletePostResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeletePostResult) GetSuccess() *core.DeletePostResponse {
	if !p.IsSetSuccess() {
		return DeletePostResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeletePostResult) SetSuccess(x interface{}) {
	p.Success = x.(*core.DeletePostResponse)
}

func (p *DeletePostResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getCommentsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core.GetCommentsRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core.Core).GetComments(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetCommentsArgs:
		success, err := handler.(core.Core).GetComments(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetCommentsResult)
		realResult.Success = success
	}
	return nil
}
func newGetCommentsArgs() interface{} {
	return &GetCommentsArgs{}
}

func newGetCommentsResult() interface{} {
	return &GetCommentsResult{}
}

type GetCommentsArgs struct {
	Req *core.GetCommentsRequest
}

func (p *GetCommentsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetCommentsArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetCommentsArgs) Unmarshal(in []byte) error {
	msg := new(core.GetCommentsRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetCommentsArgs_Req_DEFAULT *core.GetCommentsRequest

func (p *GetCommentsArgs) GetReq() *core.GetCommentsRequest {
	if !p.IsSetReq() {
		return GetCommentsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetCommentsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetCommentsResult struct {
	Success *core.GetCommentsResponse
}

var GetCommentsResult_Success_DEFAULT *core.GetCommentsResponse

func (p *GetCommentsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetCommentsResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetCommentsResult) Unmarshal(in []byte) error {
	msg := new(core.GetCommentsResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetCommentsResult) GetSuccess() *core.GetCommentsResponse {
	if !p.IsSetSuccess() {
		return GetCommentsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetCommentsResult) SetSuccess(x interface{}) {
	p.Success = x.(*core.GetCommentsResponse)
}

func (p *GetCommentsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func saveCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core.SaveCommentRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core.Core).SaveComment(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *SaveCommentArgs:
		success, err := handler.(core.Core).SaveComment(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*SaveCommentResult)
		realResult.Success = success
	}
	return nil
}
func newSaveCommentArgs() interface{} {
	return &SaveCommentArgs{}
}

func newSaveCommentResult() interface{} {
	return &SaveCommentResult{}
}

type SaveCommentArgs struct {
	Req *core.SaveCommentRequest
}

func (p *SaveCommentArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SaveCommentArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *SaveCommentArgs) Unmarshal(in []byte) error {
	msg := new(core.SaveCommentRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SaveCommentArgs_Req_DEFAULT *core.SaveCommentRequest

func (p *SaveCommentArgs) GetReq() *core.SaveCommentRequest {
	if !p.IsSetReq() {
		return SaveCommentArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SaveCommentArgs) IsSetReq() bool {
	return p.Req != nil
}

type SaveCommentResult struct {
	Success *core.SaveCommentResponse
}

var SaveCommentResult_Success_DEFAULT *core.SaveCommentResponse

func (p *SaveCommentResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SaveCommentResult")
	}
	return proto.Marshal(p.Success)
}

func (p *SaveCommentResult) Unmarshal(in []byte) error {
	msg := new(core.SaveCommentResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SaveCommentResult) GetSuccess() *core.SaveCommentResponse {
	if !p.IsSetSuccess() {
		return SaveCommentResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SaveCommentResult) SetSuccess(x interface{}) {
	p.Success = x.(*core.SaveCommentResponse)
}

func (p *SaveCommentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func checkHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(core.HealthCheckRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(core.Core).Check(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CheckArgs:
		success, err := handler.(core.Core).Check(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CheckResult)
		realResult.Success = success
	}
	return nil
}
func newCheckArgs() interface{} {
	return &CheckArgs{}
}

func newCheckResult() interface{} {
	return &CheckResult{}
}

type CheckArgs struct {
	Req *core.HealthCheckRequest
}

func (p *CheckArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CheckArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CheckArgs) Unmarshal(in []byte) error {
	msg := new(core.HealthCheckRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CheckArgs_Req_DEFAULT *core.HealthCheckRequest

func (p *CheckArgs) GetReq() *core.HealthCheckRequest {
	if !p.IsSetReq() {
		return CheckArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckArgs) IsSetReq() bool {
	return p.Req != nil
}

type CheckResult struct {
	Success *core.HealthCheckResponse
}

var CheckResult_Success_DEFAULT *core.HealthCheckResponse

func (p *CheckResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CheckResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CheckResult) Unmarshal(in []byte) error {
	msg := new(core.HealthCheckResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckResult) GetSuccess() *core.HealthCheckResponse {
	if !p.IsSetSuccess() {
		return CheckResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckResult) SetSuccess(x interface{}) {
	p.Success = x.(*core.HealthCheckResponse)
}

func (p *CheckResult) IsSetSuccess() bool {
	return p.Success != nil
}

func watchHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	st := arg.(*streaming.Args).Stream
	stream := &coreWatchServer{st}
	req := new(core.HealthCheckRequest)
	if err := st.RecvMsg(req); err != nil {
		return err
	}
	return handler.(core.Core).Watch(req, stream)
}

type coreWatchClient struct {
	streaming.Stream
}

func (x *coreWatchClient) Recv() (*core.HealthCheckResponse, error) {
	m := new(core.HealthCheckResponse)
	return m, x.Stream.RecvMsg(m)
}

type coreWatchServer struct {
	streaming.Stream
}

func (x *coreWatchServer) Send(m *core.HealthCheckResponse) error {
	return x.Stream.SendMsg(m)
}

func newWatchArgs() interface{} {
	return &WatchArgs{}
}

func newWatchResult() interface{} {
	return &WatchResult{}
}

type WatchArgs struct {
	Req *core.HealthCheckRequest
}

func (p *WatchArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in WatchArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *WatchArgs) Unmarshal(in []byte) error {
	msg := new(core.HealthCheckRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var WatchArgs_Req_DEFAULT *core.HealthCheckRequest

func (p *WatchArgs) GetReq() *core.HealthCheckRequest {
	if !p.IsSetReq() {
		return WatchArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *WatchArgs) IsSetReq() bool {
	return p.Req != nil
}

type WatchResult struct {
	Success *core.HealthCheckResponse
}

var WatchResult_Success_DEFAULT *core.HealthCheckResponse

func (p *WatchResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in WatchResult")
	}
	return proto.Marshal(p.Success)
}

func (p *WatchResult) Unmarshal(in []byte) error {
	msg := new(core.HealthCheckResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *WatchResult) GetSuccess() *core.HealthCheckResponse {
	if !p.IsSetSuccess() {
		return WatchResult_Success_DEFAULT
	}
	return p.Success
}

func (p *WatchResult) SetSuccess(x interface{}) {
	p.Success = x.(*core.HealthCheckResponse)
}

func (p *WatchResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetPost(ctx context.Context, Req *core.GetPostRequest) (r *core.GetPostResponse, err error) {
	var _args GetPostArgs
	_args.Req = Req
	var _result GetPostResult
	if err = p.c.Call(ctx, "GetPost", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SavePost(ctx context.Context, Req *core.SavePostRequest) (r *core.SavePostResponse, err error) {
	var _args SavePostArgs
	_args.Req = Req
	var _result SavePostResult
	if err = p.c.Call(ctx, "SavePost", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeletePost(ctx context.Context, Req *core.DeletePostRequest) (r *core.DeletePostResponse, err error) {
	var _args DeletePostArgs
	_args.Req = Req
	var _result DeletePostResult
	if err = p.c.Call(ctx, "DeletePost", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetComments(ctx context.Context, Req *core.GetCommentsRequest) (r *core.GetCommentsResponse, err error) {
	var _args GetCommentsArgs
	_args.Req = Req
	var _result GetCommentsResult
	if err = p.c.Call(ctx, "GetComments", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SaveComment(ctx context.Context, Req *core.SaveCommentRequest) (r *core.SaveCommentResponse, err error) {
	var _args SaveCommentArgs
	_args.Req = Req
	var _result SaveCommentResult
	if err = p.c.Call(ctx, "SaveComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Check(ctx context.Context, Req *core.HealthCheckRequest) (r *core.HealthCheckResponse, err error) {
	var _args CheckArgs
	_args.Req = Req
	var _result CheckResult
	if err = p.c.Call(ctx, "Check", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Watch(ctx context.Context, req *core.HealthCheckRequest) (Core_WatchClient, error) {
	streamClient, ok := p.c.(client.Streaming)
	if !ok {
		return nil, fmt.Errorf("client not support streaming")
	}
	res := new(streaming.Result)
	err := streamClient.Stream(ctx, "Watch", nil, res)
	if err != nil {
		return nil, err
	}
	stream := &coreWatchClient{res.Stream}
	if err := stream.Stream.SendMsg(req); err != nil {
		return nil, err
	}
	if err := stream.Stream.Close(); err != nil {
		return nil, err
	}
	return stream, nil
}
