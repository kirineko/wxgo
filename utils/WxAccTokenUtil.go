package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"net/http"
	"strings"
	"wxgo/models"
)

type AccessTokenResponse struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
}

func init() {
	logs.SetLogger("console")
}

func FetchAccessToken(appID, appSecret, accessTokenFetchUrl string) (string, error) {

	requestLine := strings.Join([]string{accessTokenFetchUrl,
		"?grant_type=client_credential&appid=",
		appID,
		"&secret=",
		appSecret}, "")

	resp, err := http.Get(requestLine)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("发送get请求获取 atoken 错误", err)
		logs.Error("发送get请求获取 atoken 错误", err)
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("发送get请求获取 atoken 读取返回body错误", err)
		logs.Error("发送get请求获取 atoken 读取返回body错误", err)
		return "", err
	}

	if bytes.Contains(body, []byte("access_token")) {
		atr := AccessTokenResponse{}
		err = json.Unmarshal(body, &atr)
		if err != nil {
			fmt.Println("发送get请求获取 atoken 返回数据json解析错误", err)
			logs.Error("发送get请求获取 atoken 返回数据json解析错误", err)
			return "", err
		}
		return atr.AccessToken, nil
	} else {
		fmt.Println("发送get请求获取 微信返回 err")
		ater := models.AccessTokenErrorResponse{}
		err = json.Unmarshal(body, &ater)
		fmt.Printf("发送get请求获取 微信返回 的错误信息 %+v\n", ater)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf("%s", ater.Errmsg)
	}
}

func GetAndUpdateDBWxAToken(o orm.Ormer) error {

	at := models.WxAccessToken{Id: 1}
	o.ReadOrCreate(&at, "id")

	wxBase := models.WxBase{Id: 1}
	err := orm.NewOrm().Read(&wxBase)
	if err != nil {
		fmt.Println("从数据库查询WxBase失败", err)
		return err
	}

	//向微信服务器发送获取accessToken的get请求
	accessToken, err := FetchAccessToken(wxBase.AppID, wxBase.AppSecret, "https://api.weixin.qq.com/cgi-bin/token")
	if err != nil {
		fmt.Println("向微信服务器发送获取accessToken的get请求失败", err)
		logs.Error("向微信服务器发送获取accessToken的get请求失败", err)
		return err
	}

	at.AccessToken = accessToken
	o.Update(&at, "access_token")

	return nil
}
