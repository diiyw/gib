package alioss

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/diiyw/goutils/http"
	"net/url"
	"strings"
	"time"
)

type AliOss struct {
	Url             string
	Endpoint        string
	BucketName      string
	AccesskeyId     string
	AccesskeySecret string
}

func (o *AliOss) UploadFromUrl(imgUrl string) ([]byte, error) {
	c, err := o.GetBucket()
	if err != nil {
		return nil, err
	}
	img, err := http.Get(imgUrl, nil)
	if err != nil {
		return nil, err
	}
	filename := GetFilename(imgUrl)
	if err := c.PutObject(filename, bytes.NewReader(img), nil); err != nil {
		time.Sleep(time.Second)
		if err := c.PutObject(filename, bytes.NewReader(img), nil); err != nil {
			return nil, err
		}
	}
	return []byte(o.Url + filename), nil
}

func (o *AliOss) GetBucket() (*oss.Bucket, error) {
	client, err := oss.New(o.Endpoint, o.AccesskeyId, o.AccesskeySecret)
	if err != nil {
		return nil, err
	}
	// 创建存储空间。
	bucket, err := client.Bucket(o.BucketName)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}

func GetFilename(uri string) string {
	u, err := url.Parse(uri)
	if err != nil {
		return ""
	}
	u.Path = strings.Replace(u.Path, "/", "_", -1)

	return strings.Trim(u.Path, "_")
}
