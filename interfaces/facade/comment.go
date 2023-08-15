package facade

import (
	"gateway/application/service"
	"gateway/infrastructure/common/context"
	"gateway/infrastructure/common/errcode"
	"gateway/interfaces/assembler"
	"gateway/interfaces/dto"
	"github.com/gin-gonic/gin"
)

func GetComments(c *context.Context) {
	articleIdStr, ok := c.GetQuery("article_id")
	if !ok {
		c.FailWithErrCode(errcode.BlogNetworkBusy, ok)
		return
	}
	lastIdStr, ok := c.GetQuery("last_id")
	if !ok {
		c.FailWithErrCode(errcode.BlogNetworkBusy, ok)
		return
	}

	rpcResp, err := service.NewCommentApplicationService(c).GetComments(assembler.StrToInt64(lastIdStr), assembler.StrToInt64(articleIdStr))
	if err != nil {
		c.FailWithErrCode(errcode.BlogNetworkBusy, err.Error())
		return
	}

	res, err := assembler.ToCommentDTOs(rpcResp.Data)
	if err != nil {
		c.FailWithErrCode(errcode.BlogNetworkBusy, err.Error())
		return
	}

	c.Success(res)
}

func GetCommentCount(c *context.Context) {
	articleIdStr, ok := c.GetQuery("article_id")
	if !ok {
		c.FailWithErrCode(errcode.BlogNetworkBusy, ok)
		return
	}

	rpcResp, err := service.NewCommentApplicationService(c).GetCountByArticleId(assembler.StrToInt64(articleIdStr))
	if err != nil {
		c.FailWithErrCode(errcode.BlogNetworkBusy, err.Error())
		return
	}

	res, err := assembler.ToCommentCountDTO(rpcResp.Data)
	if err != nil {
		c.FailWithErrCode(errcode.BlogNetworkBusy, err.Error())
		return
	}

	c.Success(res)
}

func AddComment(c *context.Context) {
	var req dto.AddComment
	if err := c.ShouldBindJSON(&req); err != nil {
		c.FailWithErrCode(errcode.BlogInvalidParam, err.Error())
		return
	}

	_, err := service.NewCommentApplicationService(c).AddComment(assembler.ToPBAddComment(req))
	if err != nil {
		c.FailWithErrCode(errcode.BlogNetworkBusy, err.Error())
		return
	}

	c.Success(gin.H{})
}
