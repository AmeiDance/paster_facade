// Code generated by Kitex v0.0.5. DO NOT EDIT.

package facade

import (
	"context"
	"fmt"
	"github.com/ameidance/paster_facade/model/vo/kitex_gen/paster/facade"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return facadeServiceInfo
}

var facadeServiceInfo = newServiceInfo()

func newServiceInfo() *kitex.ServiceInfo {
	serviceName := "Facade"
	handlerType := (*facade.Facade)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetPost":     kitex.NewMethodInfo(getPostHandler, newGetPostArgs, newGetPostResult, false),
		"SavePost":    kitex.NewMethodInfo(savePostHandler, newSavePostArgs, newSavePostResult, false),
		"GetComments": kitex.NewMethodInfo(getCommentsHandler, newGetCommentsArgs, newGetCommentsResult, false),
		"SaveComment": kitex.NewMethodInfo(saveCommentHandler, newSaveCommentArgs, newSaveCommentResult, false),
		"Check":       kitex.NewMethodInfo(checkHandler, newCheckArgs, newCheckResult, false),
		"Watch":       kitex.NewMethodInfo(watchHandler, newWatchArgs, newWatchResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "facade",
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
		req := new(facade.GetPostRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(facade.Facade).GetPost(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetPostArgs:
		success, err := handler.(facade.Facade).GetPost(ctx, s.Req)
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
	Req *facade.GetPostRequest
}

func (p *GetPostArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetPostArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetPostArgs) Unmarshal(in []byte) error {
	msg := new(facade.GetPostRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetPostArgs_Req_DEFAULT *facade.GetPostRequest

func (p *GetPostArgs) GetReq() *facade.GetPostRequest {
	if !p.IsSetReq() {
		return GetPostArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetPostArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetPostResult struct {
	Success *facade.GetPostResponse
}

var GetPostResult_Success_DEFAULT *facade.GetPostResponse

func (p *GetPostResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetPostResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetPostResult) Unmarshal(in []byte) error {
	msg := new(facade.GetPostResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetPostResult) GetSuccess() *facade.GetPostResponse {
	if !p.IsSetSuccess() {
		return GetPostResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetPostResult) SetSuccess(x interface{}) {
	p.Success = x.(*facade.GetPostResponse)
}

func (p *GetPostResult) IsSetSuccess() bool {
	return p.Success != nil
}

func savePostHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(facade.SavePostRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(facade.Facade).SavePost(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *SavePostArgs:
		success, err := handler.(facade.Facade).SavePost(ctx, s.Req)
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
	Req *facade.SavePostRequest
}

func (p *SavePostArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SavePostArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *SavePostArgs) Unmarshal(in []byte) error {
	msg := new(facade.SavePostRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SavePostArgs_Req_DEFAULT *facade.SavePostRequest

func (p *SavePostArgs) GetReq() *facade.SavePostRequest {
	if !p.IsSetReq() {
		return SavePostArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SavePostArgs) IsSetReq() bool {
	return p.Req != nil
}

type SavePostResult struct {
	Success *facade.SavePostResponse
}

var SavePostResult_Success_DEFAULT *facade.SavePostResponse

func (p *SavePostResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SavePostResult")
	}
	return proto.Marshal(p.Success)
}

func (p *SavePostResult) Unmarshal(in []byte) error {
	msg := new(facade.SavePostResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SavePostResult) GetSuccess() *facade.SavePostResponse {
	if !p.IsSetSuccess() {
		return SavePostResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SavePostResult) SetSuccess(x interface{}) {
	p.Success = x.(*facade.SavePostResponse)
}

func (p *SavePostResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getCommentsHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(facade.GetCommentsRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(facade.Facade).GetComments(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetCommentsArgs:
		success, err := handler.(facade.Facade).GetComments(ctx, s.Req)
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
	Req *facade.GetCommentsRequest
}

func (p *GetCommentsArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetCommentsArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetCommentsArgs) Unmarshal(in []byte) error {
	msg := new(facade.GetCommentsRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetCommentsArgs_Req_DEFAULT *facade.GetCommentsRequest

func (p *GetCommentsArgs) GetReq() *facade.GetCommentsRequest {
	if !p.IsSetReq() {
		return GetCommentsArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetCommentsArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetCommentsResult struct {
	Success *facade.GetCommentsResponse
}

var GetCommentsResult_Success_DEFAULT *facade.GetCommentsResponse

func (p *GetCommentsResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetCommentsResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetCommentsResult) Unmarshal(in []byte) error {
	msg := new(facade.GetCommentsResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetCommentsResult) GetSuccess() *facade.GetCommentsResponse {
	if !p.IsSetSuccess() {
		return GetCommentsResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetCommentsResult) SetSuccess(x interface{}) {
	p.Success = x.(*facade.GetCommentsResponse)
}

func (p *GetCommentsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func saveCommentHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(facade.SaveCommentRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(facade.Facade).SaveComment(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *SaveCommentArgs:
		success, err := handler.(facade.Facade).SaveComment(ctx, s.Req)
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
	Req *facade.SaveCommentRequest
}

func (p *SaveCommentArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in SaveCommentArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *SaveCommentArgs) Unmarshal(in []byte) error {
	msg := new(facade.SaveCommentRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var SaveCommentArgs_Req_DEFAULT *facade.SaveCommentRequest

func (p *SaveCommentArgs) GetReq() *facade.SaveCommentRequest {
	if !p.IsSetReq() {
		return SaveCommentArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *SaveCommentArgs) IsSetReq() bool {
	return p.Req != nil
}

type SaveCommentResult struct {
	Success *facade.SaveCommentResponse
}

var SaveCommentResult_Success_DEFAULT *facade.SaveCommentResponse

func (p *SaveCommentResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in SaveCommentResult")
	}
	return proto.Marshal(p.Success)
}

func (p *SaveCommentResult) Unmarshal(in []byte) error {
	msg := new(facade.SaveCommentResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *SaveCommentResult) GetSuccess() *facade.SaveCommentResponse {
	if !p.IsSetSuccess() {
		return SaveCommentResult_Success_DEFAULT
	}
	return p.Success
}

func (p *SaveCommentResult) SetSuccess(x interface{}) {
	p.Success = x.(*facade.SaveCommentResponse)
}

func (p *SaveCommentResult) IsSetSuccess() bool {
	return p.Success != nil
}

func checkHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(facade.HealthCheckRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(facade.Facade).Check(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CheckArgs:
		success, err := handler.(facade.Facade).Check(ctx, s.Req)
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
	Req *facade.HealthCheckRequest
}

func (p *CheckArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CheckArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CheckArgs) Unmarshal(in []byte) error {
	msg := new(facade.HealthCheckRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CheckArgs_Req_DEFAULT *facade.HealthCheckRequest

func (p *CheckArgs) GetReq() *facade.HealthCheckRequest {
	if !p.IsSetReq() {
		return CheckArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CheckArgs) IsSetReq() bool {
	return p.Req != nil
}

type CheckResult struct {
	Success *facade.HealthCheckResponse
}

var CheckResult_Success_DEFAULT *facade.HealthCheckResponse

func (p *CheckResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CheckResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CheckResult) Unmarshal(in []byte) error {
	msg := new(facade.HealthCheckResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CheckResult) GetSuccess() *facade.HealthCheckResponse {
	if !p.IsSetSuccess() {
		return CheckResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CheckResult) SetSuccess(x interface{}) {
	p.Success = x.(*facade.HealthCheckResponse)
}

func (p *CheckResult) IsSetSuccess() bool {
	return p.Success != nil
}

func watchHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	st := arg.(*streaming.Args).Stream
	stream := &facadeWatchServer{st}
	req := new(facade.HealthCheckRequest)
	if err := st.RecvMsg(req); err != nil {
		return err
	}
	return handler.(facade.Facade).Watch(req, stream)
}

type facadeWatchClient struct {
	streaming.Stream
}

func (x *facadeWatchClient) Recv() (*facade.HealthCheckResponse, error) {
	m := new(facade.HealthCheckResponse)
	return m, x.Stream.RecvMsg(m)
}

type facadeWatchServer struct {
	streaming.Stream
}

func (x *facadeWatchServer) Send(m *facade.HealthCheckResponse) error {
	return x.Stream.SendMsg(m)
}

func newWatchArgs() interface{} {
	return &WatchArgs{}
}

func newWatchResult() interface{} {
	return &WatchResult{}
}

type WatchArgs struct {
	Req *facade.HealthCheckRequest
}

func (p *WatchArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in WatchArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *WatchArgs) Unmarshal(in []byte) error {
	msg := new(facade.HealthCheckRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var WatchArgs_Req_DEFAULT *facade.HealthCheckRequest

func (p *WatchArgs) GetReq() *facade.HealthCheckRequest {
	if !p.IsSetReq() {
		return WatchArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *WatchArgs) IsSetReq() bool {
	return p.Req != nil
}

type WatchResult struct {
	Success *facade.HealthCheckResponse
}

var WatchResult_Success_DEFAULT *facade.HealthCheckResponse

func (p *WatchResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in WatchResult")
	}
	return proto.Marshal(p.Success)
}

func (p *WatchResult) Unmarshal(in []byte) error {
	msg := new(facade.HealthCheckResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *WatchResult) GetSuccess() *facade.HealthCheckResponse {
	if !p.IsSetSuccess() {
		return WatchResult_Success_DEFAULT
	}
	return p.Success
}

func (p *WatchResult) SetSuccess(x interface{}) {
	p.Success = x.(*facade.HealthCheckResponse)
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

func (p *kClient) GetPost(ctx context.Context, Req *facade.GetPostRequest) (r *facade.GetPostResponse, err error) {
	var _args GetPostArgs
	_args.Req = Req
	var _result GetPostResult
	if err = p.c.Call(ctx, "GetPost", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SavePost(ctx context.Context, Req *facade.SavePostRequest) (r *facade.SavePostResponse, err error) {
	var _args SavePostArgs
	_args.Req = Req
	var _result SavePostResult
	if err = p.c.Call(ctx, "SavePost", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetComments(ctx context.Context, Req *facade.GetCommentsRequest) (r *facade.GetCommentsResponse, err error) {
	var _args GetCommentsArgs
	_args.Req = Req
	var _result GetCommentsResult
	if err = p.c.Call(ctx, "GetComments", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SaveComment(ctx context.Context, Req *facade.SaveCommentRequest) (r *facade.SaveCommentResponse, err error) {
	var _args SaveCommentArgs
	_args.Req = Req
	var _result SaveCommentResult
	if err = p.c.Call(ctx, "SaveComment", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Check(ctx context.Context, Req *facade.HealthCheckRequest) (r *facade.HealthCheckResponse, err error) {
	var _args CheckArgs
	_args.Req = Req
	var _result CheckResult
	if err = p.c.Call(ctx, "Check", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Watch(ctx context.Context, req *facade.HealthCheckRequest) (Facade_WatchClient, error) {
	streamClient, ok := p.c.(client.Streaming)
	if !ok {
		return nil, fmt.Errorf("client not support streaming")
	}
	res := new(streaming.Result)
	err := streamClient.Stream(ctx, "Watch", nil, res)
	if err != nil {
		return nil, err
	}
	stream := &facadeWatchClient{res.Stream}
	if err := stream.Stream.SendMsg(req); err != nil {
		return nil, err
	}
	if err := stream.Stream.Close(); err != nil {
		return nil, err
	}
	return stream, nil
}
