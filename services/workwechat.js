import {kv} from "storage/KVstorage"

function getAccessToken(corpid, corpsecret) {
    let url = `https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=${corpid}&corpsecret=${corpsecret}`
    fetch(url).then(async response => {
        const resp_json = await response.json();
        if (resp_json.errcode) {
            return 0
        }
        return resp_json.access_token
    })

}


function postMessage(accesstoken, postData) {
    let url = `https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=${accesstoken}`
    fetch(url, { method: "POST", headers: { 'Content-Type': 'application/json' }, body: JSON.stringify(postData) })
    .then(response=>{
      return response
    })

}

async function storeAccessToken(corpid,accesstoken){
   return await kv.put(corpid,accesstoken,{expirationTtl:7200})

}

export default {
    getAccessToken,
    postMessage,
    storeAccessToken
}