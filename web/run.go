package web

import (
	"net/http"

	"workWechat_api/pkg/config"
	"workWechat_api/pkg/logging"
	"workWechat_api/web/controller"
	"workWechat_api/web/interceptor"
)

func Run()  {

	http.Handle("/",interceptor.TimeMiddleware(http.HandlerFunc(controller.FallbackHandler)))
	http.Handle("/sendmessage", interceptor.TimeMiddleware(http.HandlerFunc(controller.SendMsgHandler)))

	if config.General.Ssl  && (config.General.CertPath != "" && config.General.KeyPath != ""){
		logging.Info("running  server with SSL")
		err := http.ListenAndServeTLS(config.General.Listen,config.General.CertPath,config.General.KeyPath,nil)
		if err != nil{
			logging.Panic("run  server with SSL error:",err)
		}
	} else {
		logging.Info("running  server ")
		err := http.ListenAndServe(config.General.Listen,nil)
		if err != nil{
			logging.Panic("run  server error:",err)
		}
	}


}
