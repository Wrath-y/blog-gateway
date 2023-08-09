package facade

import (
	"gateway/application/service"
	"gateway/infrastructure/common/context"
	"gateway/infrastructure/common/errcode"
	"gateway/interfaces/assembler"
	"strconv"
)

func GetArticles(c *context.Context) {
	lastId := c.GetInt64("last_id")

	rpcResp, err := service.NewArticleApplicationService(c).GetArticles(lastId, 6)
	if err != nil {
		c.FailWithErrCode(errcode.BlogNetworkBusy, nil)
		return
	}

	res, err := assembler.ToArticleDTOs(rpcResp.Data)
	if err != nil {
		c.FailWithErrCode(errcode.BlogNetworkBusy, err.Error())
		return
	}

	c.Success(res)
}

func GetAllArticles(c *context.Context) {
	rpcResp, err := service.NewArticleApplicationService(c).GetAllArticles()
	if err != nil {
		c.FailWithErrCode(errcode.BlogNetworkBusy, nil)
		return
	}

	res, err := assembler.ToArticleDTOs(rpcResp.Data)
	if err != nil {
		c.FailWithErrCode(errcode.BlogNetworkBusy, err.Error())
		return
	}

	c.Success(res)
}

func GetArticle(c *context.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.FailWithErrCode(errcode.BlogInvalidParam, err.Error())
		return
	}

	rpcResp, err := service.NewArticleApplicationService(c).GetById(int64(id))
	if err != nil {
		c.FailWithErrCode(errcode.BlogNetworkBusy, nil)
		return
	}

	res, err := assembler.ToArticleDTO(rpcResp.Data)
	if err != nil {
		c.Logger.ErrorL("解析dto失败", rpcResp, err.Error())
		c.FailWithErrCode(errcode.BlogNetworkBusy, err.Error())
		return
	}

	c.Success(res)
}
