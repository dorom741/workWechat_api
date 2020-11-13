package services

import (
	"workWechat_api/pkg/logging"
	"workWechat_api/pkg/storage"
	"workWechat_api/pkg/wechatWorkApi"
)

func  MsgSender(agentid string,postStr []byte)  string{
	accessToken, err:= storage.RedisPool().Get(agentid)
	if len(accessToken) == 0 {
		logging.Warning("error get accessToken from redis: %s ,error: %T",agentid, err)
		accessToken = wechatWorkApi.GetaccessToken(agentid)
		storage.RedisPool().Set(agentid,accessToken)
	}
	return wechatWorkApi.SendMessage(accessToken,postStr)
}
