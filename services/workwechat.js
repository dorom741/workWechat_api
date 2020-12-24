
async function getAccessToken(corpsecret) {
    const url = `https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=${corpid}&corpsecret=${corpsecret}`
    var response = await fetch(url)
    const resp_json = await response.json();
    return resp_json.access_token
}


function postMessage(accesstoken, postData) {
    const url = `https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=${accesstoken}`
    var response = await fetch(url, { method: "POST", headers: { 'Content-Type': 'application/json' }, body: postData })
    return response

}

export default {
    getAccessToken,
    postMessage,
}