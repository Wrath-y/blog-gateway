package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
	"log"
)

var OSS *oss.Bucket

func Setup() {
	var err error
	OSS, err = Bucket(viper.GetString("aliyun.oss.bucket"))
	if err != nil {
		log.Fatal(err.Error())
	}
}

func List(nextMarker string, page int) (*oss.ListObjectsResult, error) {
	list, err := OSS.ListObjects(oss.MaxKeys(page), oss.Marker(nextMarker))
	if err != nil {
		return &list, err
	}

	return &list, nil
}

func Delete(name string) (int, error) {
	err := OSS.DeleteObject(name)
	if err != nil {
		return 0, err
	}

	return 1, nil
}

func Count() (int, error) {
	var lsRes oss.ListObjectsResult
	var err error

	lsRes, err = OSS.ListObjects(oss.MaxKeys(1000))
	if err != nil {
		return 0, err
	}
	count := len(lsRes.Objects)

	return count, nil
}

// Bucket b = viper.GetString("aliyun.oss.bucket_name")
func Bucket(b string) (*oss.Bucket, error) {
	clientSer, err := oss.New(viper.GetString("aliyun.oss.endpoint"),
		viper.GetString("aliyun.oss.access_keyid"),
		viper.GetString("aliyun.oss.access_keysecret"))
	if err != nil {
		return nil, err
	}
	// 获取存储空间。
	return clientSer.Bucket(b)
}
