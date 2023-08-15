package facade

import (
	"gateway/application/service"
	"gateway/infrastructure/common/context"
	"gateway/infrastructure/common/errcode"
	"gateway/interfaces/assembler"
	"strconv"
)

func GetPixivs(c *context.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page_size", "15"))
	if err != nil {
		c.FailWithErrCode(errcode.BlogNetworkBusy, err.Error())
		return
	}
	rpcResp, err := service.NewToysApplicationService(c).GetPixivs(c.DefaultQuery("next_marker", ""), page)
	if err != nil {
		c.FailWithErrCode(errcode.BlogNetworkBusy, err.Error())
		return
	}

	res, err := assembler.ToPixivImgDTOs(rpcResp.Data)
	if err != nil {
		c.FailWithErrCode(errcode.BlogNetworkBusy, err.Error())
		return
	}

	c.Success(res)
}

func GetFriends(c *context.Context) {
	rpcResp, err := service.NewToysApplicationService(c).GetFriends()
	if err != nil {
		c.FailWithErrCode(errcode.BlogNetworkBusy, err.Error())
		return
	}

	res, err := assembler.ToFriendDTOs(rpcResp.Data)
	if err != nil {
		c.FailWithErrCode(errcode.BlogNetworkBusy, err.Error())
		return
	}

	c.Success(res)
}
