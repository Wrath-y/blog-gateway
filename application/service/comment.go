package service

import (
	"context"
	ctx "gateway/infrastructure/common/context"
	"gateway/infrastructure/common/errcode"
	"gateway/infrastructure/util/consul"
	"gateway/infrastructure/util/grpcclient"
	"gateway/interfaces/proto"
	"time"
)

type CommentApplicationService struct {
	*ctx.Context
}

func NewCommentApplicationService(ctx *ctx.Context) *CommentApplicationService {
	return &CommentApplicationService{
		ctx,
	}
}

func (a *CommentApplicationService) GetComments(id, articleId int64) (*proto.Response, error) {
	c, grpcClient, closeFunc, err := a.getClientAndContext("comment")
	if err != nil {
		a.Logger.ErrorL("获取client失败", "", err.Error())
		return nil, err
	}
	defer closeFunc()

	req := &proto.GetCommentsReq{
		Id:        id,
		ArticleId: articleId,
	}
	a.Logger.Info("rpc请求参数", req, "")
	rpcResp, err := grpcClient.GetComments(c, req)
	if err != nil {
		a.Logger.ErrorL("rpc请求失败", req, err.Error())
		return nil, err
	}
	if rpcResp.Code != 0 {
		a.Logger.ErrorL("评论列表返回异常", req, rpcResp)
		return nil, errcode.BlogNetworkBusy
	}
	if rpcResp.Data == "" {
		a.Logger.ErrorL("评论列表返回空数据", req, rpcResp)
		return nil, errcode.BlogNetworkBusy
	}

	return rpcResp, nil
}

func (a *CommentApplicationService) GetCountByArticleId(articleId int64) (*proto.Response, error) {
	c, grpcClient, closeFunc, err := a.getClientAndContext("comment")
	if err != nil {
		a.Logger.ErrorL("获取client失败", "", err.Error())
		return nil, err
	}
	defer closeFunc()

	req := &proto.OnlyArticleIdReq{
		ArticleId: articleId,
	}
	a.Logger.Info("rpc请求参数", req, "")
	rpcResp, err := grpcClient.GetCountByArticleId(c, req)
	if err != nil {
		a.Logger.ErrorL("rpc请求失败", req, err.Error())
		return nil, err
	}
	if rpcResp.Code != 0 {
		a.Logger.ErrorL("评论数量返回异常", req, rpcResp)
		return nil, errcode.BlogNetworkBusy
	}
	if rpcResp.Data == "" {
		a.Logger.ErrorL("评论数量返回空数据", req, rpcResp)
		return nil, errcode.BlogNetworkBusy
	}

	return rpcResp, nil
}

func (a *CommentApplicationService) AddComment(req *proto.AddCommentReq) (*proto.Response, error) {
	c, grpcClient, closeFunc, err := a.getClientAndContext("comment")
	if err != nil {
		a.Logger.ErrorL("获取client失败", "", err.Error())
		return nil, err
	}
	defer closeFunc()

	a.Logger.Info("rpc请求参数", req, "")
	rpcResp, err := grpcClient.AddComment(c, req)
	if err != nil {
		a.Logger.ErrorL("rpc请求失败", req, err.Error())
		return nil, err
	}
	if rpcResp.Code != 0 {
		a.Logger.ErrorL("添加评论返回异常", req, rpcResp)
		return nil, errcode.BlogNetworkBusy
	}
	if rpcResp.Data == "" {
		a.Logger.ErrorL("添加评论返回空数据", req, rpcResp)
		return nil, errcode.BlogNetworkBusy
	}

	return rpcResp, nil
}

func (a *CommentApplicationService) getClientAndContext(serviceName string) (context.Context, proto.CommentClient, func(), error) {
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

	grpcClient := proto.NewCommentClient(conn)

	c, cancel := context.WithTimeout(context.Background(), time.Second*3)
	closeFunc := func() {
		connCloseFunc()
		cancel()
	}

	return c, grpcClient, closeFunc, nil
}
