package shangou

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/easynet-cn/go-meituan/isv"
	"github.com/easynet-cn/go-meituan/medicine"
)

const (
	AppId        = ""
	AppSecret    = ""
	AppPoiCode   = ""
	AccessToken  = ""
	RefreshToken = ""
)

func Test_GetAccessToken(t *testing.T) {
	shanGouClient := getClient()

	bytes, err := shanGouClient.DoRequest("https://waimaiopen.meituan.com/api/v1/oauth/authorize", "GET", "", &isv.GetAccessTokenRequest{AppPoiCode: AppPoiCode, ResponseType: "token"})

	fmt.Println(string(bytes), err)
}

func Test_RefreshAccessToken(t *testing.T) {
	shanGouClient := getClient()

	bytes, err := shanGouClient.DoRequest("https://waimaiopen.meituan.com/api/v1/oauth/token", "POST", "", &isv.RefreshAccessTokenRequest{GrantType: "refresh_token", RefreshToken: RefreshToken})

	fmt.Println(string(bytes), err)
}

func Test_GetStoreCategories(t *testing.T) {
	shanGouClient := getClient()

	bytes, err := shanGouClient.DoRequest("https://waimaiopen.meituan.com/api/v1/medicineCat/list", "GET", "", &medicine.MedicineCatListRequest{AppPoiCode: AppPoiCode})

	fmt.Println(string(bytes), err)
}

func Test_GetStoreCategoriesByMapParam(t *testing.T) {
	shanGouClient := getClient()

	bytes, err := shanGouClient.DoRequest("https://waimaiopen.meituan.com/api/v1/medicineCat/list", "GET", "", map[string]interface{}{"app_poi_code": AppPoiCode})

	fmt.Println(string(bytes), err)
}

func Test_GetMedicineList(t *testing.T) {
	shanGouClient := getClient()

	bytes, err := shanGouClient.DoRequest("https://waimaiopen.meituan.com/api/v1/medicine/list", "GET", AccessToken, &medicine.MedicineListRequest{AppPoiCode: AppPoiCode, Offset: 0, Limit: 200})

	fmt.Println(string(bytes), err)
}

func Test_CateMedicineCategory(t *testing.T) {
	shanGouClient := getClient()

	request := &medicine.MedicineCatRequest{
		AppPoiCode:   AppPoiCode,
		CategoryCode: "092400",
		CategoryName: "其他",
		Sequence:     24,
	}

	bytes, err := shanGouClient.DoRequest("https://waimaiopen.meituan.com/api/v1/medicineCat/save", "POST", AccessToken, request)

	fmt.Println(string(bytes), err)
}

func getClient() *ShanGouClient {
	return &ShanGouClient{
		HttpClient: &http.Client{},
		AppId:      AppId,
		AppSecret:  AppSecret,
	}
}

type GetOrderDaySeqRequest struct {
	AppPoiCode string `url:"app_poi_code"`
}
