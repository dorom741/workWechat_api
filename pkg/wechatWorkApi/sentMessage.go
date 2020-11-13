package wechatWorkApi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"workWechat_api/pkg/config"
	"workWechat_api/pkg/logging"
)

var (
	getAccesstokenUrl = "https://qyapi.weixin.qq.com/cgi-bin/gettoken"
	sentMessageUrl = "https://qyapi.weixin.qq.com/cgi-bin/message/send"
	)

func GetaccessToken(agentId string)  string{

	params := url.Values{}
	urlRaw, err := url.Parse(getAccesstokenUrl)
	if err != nil {
		return ""
	}
	params.Set("corpid", config.QiyeInfo.Corpid)
	params.Set("corpsecret", config.QiyeInfo.AppList[agentId])
	urlRaw.RawQuery = params.Encode()

	resp,err := http.Get(urlRaw.String())
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var mapResult map[string]interface{}
	err = json.Unmarshal(body, &mapResult)
	if err != nil {
		logging.Error("JsonToMap err: ", err)
	}
	return mapResult["access_token"].(string)
}

func SendMessage(accessToken string, postByte []byte)  string{
	params := url.Values{}
	urlRaw, _ := url.Parse(sentMessageUrl)
	params.Set("access_token",accessToken)
	urlRaw.RawQuery = params.Encode()
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlRaw.String(), bytes.NewBuffer(postByte))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset","UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		logging.Error("sending message error: %T",err)
		return "sending message error"
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}