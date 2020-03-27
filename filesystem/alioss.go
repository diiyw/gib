package filesystem

import (
	"bytes"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/diiyw/gib/http"
)

// AliOss the aliyun oss storage
type AliOss struct {
	URL             string
	Endpoint        string
	BucketName      string
	AccesskeyID     string
	AccesskeySecret string
	ContentType     string
}

// Has if exits the file
func (o *AliOss) Has(filename string) (bool, error) {
	c, err := o.GetBucket()
	if err != nil {
		return false, err
	}
	return c.IsObjectExist(filename)
}

// UploadFromURL upload an image from URL to oss
func (o *AliOss) UploadFromURL(imgURL string) ([]byte, error) {
	c, err := o.GetBucket()
	if err != nil {
		return nil, err
	}
	img, err := http.Do(http.Url(imgURL))
	if err != nil {
		return nil, err
	}
	filename := GetFilename(imgURL)
	if err := c.PutObject(filename, bytes.NewReader(img), oss.ContentType(o.ContentType)); err != nil {
		time.Sleep(time.Second)
		if err := c.PutObject(filename, bytes.NewReader(img), oss.ContentType(o.ContentType)); err != nil {
			return nil, err
		}
	}
	return []byte(o.URL + filename), nil
}

// Upload upload to oss with binary
func (o *AliOss) Upload(filename string, data []byte) ([]byte, error) {
	c, err := o.GetBucket()
	if err != nil {
		return nil, err
	}
	if err := c.PutObject(filename, bytes.NewReader(data), oss.ContentType(o.ContentType)); err != nil {
		time.Sleep(time.Second)
		if err := c.PutObject(filename, bytes.NewReader(data), oss.ContentType(o.ContentType)); err != nil {
			return nil, err
		}
	}
	return []byte(o.URL + filename), nil
}

// GetBucket Get the oss bucket instance.
func (o *AliOss) GetBucket() (*oss.Bucket, error) {
	client, err := oss.New(o.Endpoint, o.AccesskeyID, o.AccesskeySecret)
	if err != nil {
		return nil, err
	}
	// get bucket
	bucket, err := client.Bucket(o.BucketName)
	if err != nil {
		return nil, err
	}
	return bucket, nil
}
