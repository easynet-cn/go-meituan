package meituan

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func Test_Singnature(t *testing.T) {
	url := "https://waimaiopen.meituan.com/api/v1/order/getOrderDaySeq?app_id=0000&app_poi_code=31号测试店&timestamp=1389751221a00ba58f000001aa697ab000006d52d"

	sig := fmt.Sprintf("%x", md5.Sum([]byte(url)))

	if sig != "dbb4d444d68596d03dbcb79a4e4a4e3f" {
		t.Error(sig)
	}

	fmt.Println(sig)
}
