package shangou

import "net/http"

type ShanGouClient struct {
	HttpClient  *http.Client
	AppId       string
	AppSecret   string
	AccessToken string
}
