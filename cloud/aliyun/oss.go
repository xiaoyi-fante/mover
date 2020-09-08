package aliyun

import (
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// OSS 结构体定义
type AliOSS struct {
	Endpoint, AccessKey, AccessSecret string
	Bucket                            *oss.Bucket
}

// NewAliOSS  新建阿里云 OSS 对象
func NewAliOSS(bucket, endpoint, key, secret string) (*AliOSS, error) {
	client, err := oss.New(endpoint, key, secret)
	if err != nil {
		return nil, err
	}
	bucketObject, err := client.Bucket(bucket)
	if err != nil {
		return nil, err
	}
	return &AliOSS{
		Endpoint:     endpoint,
		AccessKey:    key,
		AccessSecret: secret,
		Bucket:       bucketObject,
	}, nil
}

// 实现Upload方法
func (ali *AliOSS) Upload(objectKey string, r io.Reader) (string, error) {
	err := ali.Bucket.PutObject(objectKey, r)
	if err != nil {
		return "", err
	}
	return objectKey, nil
}
