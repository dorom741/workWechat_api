package main

import (
	"testing"

	"workWechat_api/web/services"
)

func  TestSenter(t *testing.T) {
	postStr := `
	{
		"touser": "@all",
		"msgtype": "text",
		"agentid": 1000002,
		"text": {
		"content": "api test,\n<a href=\"http://\">hello</a>"
		}
	}`
	services.MsgSender("1000002",[]byte(postStr))
}
