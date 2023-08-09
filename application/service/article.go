package service

import (
	"context"
	ctx "gateway/infrastructure/common/context"
	"gateway/infrastructure/common/errcode"
	"gateway/infrastructure/util/consul"
	"gateway/infrastructure/util/grpcclient"
	"gateway/interfaces/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type ArticleApplicationService struct {
	*ctx.Context
}

func NewArticleApplicationService(ctx *ctx.Context) *ArticleApplicationService {
	return &ArticleApplicationService{
		ctx,
	}
}

func (a *ArticleApplicationService) GetById(id int64) (*proto.Response, error) {
	c, grpcClient, closeFunc, err := a.getClientAndContext("article")
	if err != nil {
		a.Logger.ErrorL("获取client失败", "", err.Error())
		return nil, err
	}
	defer closeFunc()

	req := &proto.GetByIdReq{
		Id: id,
	}
	rpcResp, err := grpcClient.GetById(c, req)
	if err != nil {
		return nil, err
	}
	if rpcResp.Code != 0 {
		a.Logger.ErrorL("文章详情返回异常", req, rpcResp)
		return nil, errcode.BlogNetworkBusy
	}
	if rpcResp.Data == "" {
		a.Logger.ErrorL("文章详情返回空数据", req, rpcResp)
		return nil, errcode.BlogNetworkBusy
	}

	return rpcResp, nil
}

func (a *ArticleApplicationService) GetArticles(id int64, size int32) (*proto.Response, error) {
	c, grpcClient, closeFunc, err := a.getClientAndContext("article")
	if err != nil {
		a.Logger.ErrorL("获取client失败", "", err.Error())
		return nil, err
	}
	defer closeFunc()

	req := &proto.FindByIdReq{
		Id:   id,
		Size: size,
	}
	rpcResp, err := grpcClient.FindById(c, req)
	if err != nil {
		return nil, err
	}
	if rpcResp.Code != 0 {
		a.Logger.ErrorL("文章列表返回异常", req, rpcResp)
		return nil, errcode.BlogNetworkBusy
	}
	if rpcResp.Data == "" {
		a.Logger.ErrorL("文章列表返回空数据", req, rpcResp)
		return nil, errcode.BlogNetworkBusy
	}

	return rpcResp, nil
}

func (a *ArticleApplicationService) GetAllArticles() (*proto.Response, error) {
	c, grpcClient, closeFunc, err := a.getClientAndContext("article")
	if err != nil {
		a.Logger.ErrorL("获取client失败", "", err.Error())
		return nil, err
	}
	defer closeFunc()

	rpcResp, err := grpcClient.FindAll(c, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	if rpcResp.Code != 0 {
		a.Logger.ErrorL("所有文章列表返回异常", "", rpcResp)
		return nil, errcode.BlogNetworkBusy
	}
	if rpcResp.Data == "" {
		a.Logger.ErrorL("所有文章列表返回空数据", "", rpcResp)
		return nil, errcode.BlogNetworkBusy
	}

	return rpcResp, nil
}

func (a *ArticleApplicationService) getClientAndContext(serviceName string) (context.Context, proto.ArticleClient, func(), error) {
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

	grpcClient := proto.NewArticleClient(conn)

	c, cancel := context.WithTimeout(context.Background(), time.Second*3)
	closeFunc := func() {
		connCloseFunc()
		cancel()
	}

	return c, grpcClient, closeFunc, nil
}
