
export async function getAccessToken(corpSecret) {
    const url = `https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=${corpid}&corpsecret=${corpSecret}`
    const response = await fetch(url);
    const resp_json = await response.json();
    return resp_json.access_token
}


export async function postMessage(accessToken, postData) {
    const url = `https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=${accessToken}`
    return await fetch(url, {method: "POST", headers: {'Content-Type': 'application/json'}, body: postData})
}
