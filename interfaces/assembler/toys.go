package assembler

import (
	"gateway/infrastructure/util/util/highperf"
	"gateway/interfaces/dto"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func ToPixivImgDTOs(data string) (*oss.ListObjectsResult, error) {
	res := &oss.ListObjectsResult{}
	if err := sonic.Unmarshal(highperf.Str2bytes(data), res); err != nil {
		return nil, err
	}
	return res, nil
}

func ToFriendDTOs(data string) ([]dto.Friend, error) {
	res := make([]dto.Friend, 0)
	if err := sonic.Unmarshal(highperf.Str2bytes(data), &res); err != nil {
		return nil, err
	}
	return res, nil
}
