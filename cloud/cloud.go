package cloud

import "io"

// uploader    文件上传的统一接口
type Uploader interface {
	Upload(objectKey string, r io.Reader) (string, error)
}
