package assembler

import (
	"encoding/json"
	"gateway/infrastructure/util/util/highperf"
	"gateway/interfaces/dto"
)

func ToArticleDTOs(articles string) ([]*dto.ArticlesItem, error) {
	res := make([]*dto.ArticlesItem, 0)
	if err := json.Unmarshal(highperf.Str2bytes(articles), &res); err != nil {
		return nil, err
	}
	return res, nil
}

func ToArticleDTO(article string) (*dto.Article, error) {
	res := &dto.Article{}
	if err := json.Unmarshal(highperf.Str2bytes(article), res); err != nil {
		return nil, err
	}
	return res, nil
}
