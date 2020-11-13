package main

import (
	"fmt"
	"testing"

	"workWechat_api/pkg/wechatWorkApi"
)

func TestGettoken(t *testing.T)  {
	fmt.Println(wechatWorkApi.GetaccessToken("1000002"))
}

func TestSentMessage(t *testing.T)  {
	var accessToken ="0sr4fSWaR1C8wyokLFiTy5tBgxkRErdD8USuGf9ipc09yzUDVeqH-nRhAjZ_3TZk-wP0k-AHg41g2XTKHq9ShOIG8UkAumgRru_qDvprop3beiL4rIXyUYKkR3SUehKNohRRqjRQKhFCv8_iM-GvT-f51IzqlpEALGgNZCSKZBKNo8WN5_FQ-5aJHkSPy3_r6tUPGrHzuaf6pyoGaIxtdg"
	postStr := `
	{
		"touser": "@all",
		"msgtype": "text",
		"agentid": 1000002,
		"text": {
		"content": "你的快递已到，请携带工卡前往邮件中心领取。\n出发前可查看<a href=\"http://work.weixin.qq.com\">邮件中心视频实况</a>，聪明避开排队。"
		}
	}`

	wechatWorkApi.SendMessage(accessToken, []byte(postStr))
}