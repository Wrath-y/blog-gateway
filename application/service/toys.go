package service

import (
	"context"
	ctx "gateway/infrastructure/common/context"
	"gateway/infrastructure/common/errcode"
	"gateway/infrastructure/util/consul"
	"gateway/infrastructure/util/grpcclient"
	"gateway/interfaces/proto"
	"github.com/golang/protobuf/ptypes/empty"
	"time"
)

type ToysApplicationService struct {
	*ctx.Context
}

func NewToysApplicationService(ctx *ctx.Context) *ToysApplicationService {
	return &ToysApplicationService{
		ctx,
	}
}

func (a *ToysApplicationService) GetPixivs(nextMarker string, page int) (*proto.Response, error) {
	c, grpcClient, closeFunc, err := a.getClientAndContext("toys")
	if err != nil {
		a.Logger.ErrorL("获取client失败", "", err.Error())
		return nil, err
	}
	defer closeFunc()

	req := &proto.GetPixivsReq{
		NextMarker: nextMarker,
		Page:       int32(page),
	}
	a.Logger.Info("rpc请求参数", req, "")
	rpcResp, err := grpcClient.GetPixivs(c, req)
	if err != nil {
		a.Logger.ErrorL("rpc请求失败", req, err.Error())
		return nil, err
	}
	if rpcResp.Code != 0 {
		a.Logger.ErrorL("友链返回异常", req, rpcResp)
		return nil, errcode.BlogNetworkBusy
	}
	if rpcResp.Data == "" {
		a.Logger.ErrorL("友链返回空数据", req, rpcResp)
		return nil, errcode.BlogNetworkBusy
	}

	return rpcResp, nil
}

func (a *ToysApplicationService) GetFriends() (*proto.Response, error) {
	c, grpcClient, closeFunc, err := a.getClientAndContext("toys")
	if err != nil {
		a.Logger.ErrorL("获取client失败", "", err.Error())
		return nil, err
	}
	defer closeFunc()

	req := &empty.Empty{}
	a.Logger.Info("rpc请求参数", req, "")
	rpcResp, err := grpcClient.GetFriends(c, req)
	if err != nil {
		a.Logger.ErrorL("rpc请求失败", req, err.Error())
		return nil, err
	}
	if rpcResp.Code != 0 {
		a.Logger.ErrorL("友链返回异常", req, rpcResp)
		return nil, errcode.BlogNetworkBusy
	}
	if rpcResp.Data == "" {
		a.Logger.ErrorL("友链返回空数据", req, rpcResp)
		return nil, errcode.BlogNetworkBusy
	}

	return rpcResp, nil
}

func (a *ToysApplicationService) getClientAndContext(serviceName string) (context.Context, proto.ToysClient, func(), error) {
	instance, err := consul.Client.GetHealthRandomInstance(serviceName)
	if err != nil {
		return nil, nil, nil, err
	}

	conn, err := grpcclient.NewClient(instance).GetHealthConn()
	if err != nil {
		return nil, nil, nil, err
	}

	connCloseFunc := func() {
		if err := conn.Close(); err != nil {
			a.Logger.ErrorL("grpc链接关闭失败", "", err.Error())
		}
	}

	grpcClient := proto.NewToysClient(conn)

	c, cancel := context.WithTimeout(context.Background(), time.Second*3)
	closeFunc := func() {
		connCloseFunc()
		cancel()
	}

	return c, grpcClient, closeFunc, nil
}
