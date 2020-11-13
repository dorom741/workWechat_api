package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"workWechat_api/pkg/logging"
	"workWechat_api/web/services"
)

func SendMsgHandler(w http.ResponseWriter, r *http.Request){
	var (
		token = r.URL.Query().Get("token")
		agentId = r.URL.Query().Get("agentid")
	)
	if r.Method != http.MethodPost ||  agentId == "" || token == "" {
		FallbackHandler(w,r)
		return
	}
	bodyByte,err := ioutil.ReadAll(r.Body)
	if err != nil{
		logging.Error("reading request body error: %T",err)
	}
	w.Header().Set("content-type","application/json")
	_,_ = fmt.Fprintln(w, services.MsgSender(agentId, bodyByte))
}

func FallbackHandler(w http.ResponseWriter, r *http.Request){
	_,err := fmt.Fprintln(w,"hello")
	if err != nil{
		logging.Error("writer Response error")
	}
	
}

