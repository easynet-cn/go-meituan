package shangou

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/easynet-cn/go-meituan"
)

type ShanGouClient struct {
	HttpClient  *http.Client
	AppId       string
	AppSecret   string
	AccessToken string
}

func (s *ShanGouClient) DoRequest(requestUrl string, method string, request interface{}) ([]byte, error) {
	timestamp := time.Now().Unix()

	url, urlValues := meituan.Sign(requestUrl, s.AppId, s.AppSecret, timestamp, request)

	var resp *http.Response
	var err error

	if method == "POST" {
		resp, err = s.HttpClient.PostForm(url, urlValues)
	} else {
		resp, err = s.HttpClient.Get(fmt.Sprintf("%s&%s", url, urlValues.Encode()))
	}

	if err != nil {
		return nil, err
	}

	if bytes, err := ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	} else {
		return bytes, nil
	}
}
