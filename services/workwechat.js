
export async function getAccessToken(corpsecret) {
    const url = `https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=${corpid}&corpsecret=${corpsecret}`
    var response = await fetch(url)
    const resp_json = await response.json();
    return resp_json.access_token
}


export async function postMessage(accesstoken, postData) {
    const url = `https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=${accesstoken}`
    return await fetch(url, {method: "POST", headers: {'Content-Type': 'application/json'}, body: postData})
}
