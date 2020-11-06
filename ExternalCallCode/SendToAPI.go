package ExternalCallCode

import (
	"PrintYun/Config"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendWxAuthAPI(CODE string) (string, error) {
	url := fmt.Sprintf(Config.Code2SessiOnURL, Config.APPID, Config.SECRET, CODE)
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return "", err
	}
	var wxMap map[string]string
	err = json.NewDecoder(resp.Body).Decode(&wxMap)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// wxMap {
	//		openid
	// 		unionid
	// }
	return wxMap["openid"], nil
}