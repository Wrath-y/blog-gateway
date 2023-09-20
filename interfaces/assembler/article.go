package assembler

import (
	"gateway/infrastructure/util/util/highperf"
	"gateway/interfaces/dto"
	"github.com/bytedance/sonic"
	"strconv"
)

func ToArticleDTOs(articles string) ([]*dto.ArticlesItem, error) {
	res := make([]*dto.ArticlesItem, 0)
	if err := sonic.Unmarshal(highperf.Str2bytes(articles), &res); err != nil {
		return nil, err
	}
	return res, nil
}

func ToArticleDTO(article string) (*dto.Article, error) {
	res := &dto.Article{}
	if err := sonic.Unmarshal(highperf.Str2bytes(article), res); err != nil {
		return nil, err
	}
	return res, nil
}

func StrToInt64(str string) int64 {
	id, _ := strconv.Atoi(str)
	return int64(id)
}
