package assembler

import (
	"encoding/json"
	"gateway/infrastructure/util/util/highperf"
	"gateway/interfaces/dto"
	"gateway/interfaces/proto"
	"strconv"
)

func ToCommentDTOs(comments string) ([]*dto.CommentsItem, error) {
	res := make([]*dto.CommentsItem, 0)
	if err := json.Unmarshal(highperf.Str2bytes(comments), &res); err != nil {
		return nil, err
	}
	return res, nil
}

func ToCommentCountDTO(count string) (dto.CommentCount, error) {
	res := dto.CommentCount{}
	data, err := strconv.Atoi(count)
	if err != nil {
		return res, err
	}
	res.Count = data

	return res, nil
}

func ToPBAddComment(req dto.AddComment) *proto.AddCommentReq {
	return &proto.AddCommentReq{
		ArticleId: int64(req.ArticleId),
		Pid:       int64(req.Pid),
		Name:      req.Name,
		Email:     req.Email,
		Url:       req.Url,
		Content:   req.Content,
	}
}
