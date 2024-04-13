package httpclient

import (
	"fmt"
	"testing"
)

func TestGetV2(t *testing.T) {
	url := "https://api.myip.com"
	resp, err := GetV2(url, nil, nil)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(resp))
	t.Log(resp)

}
