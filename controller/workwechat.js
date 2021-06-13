import {getAccessToken, postMessage} from "../services/workwechat";

export default async function workerWechatHandler(request) {
    let token = request.headers.get("token") != null ? request.headers.get("token") : new URL(request.url).searchParams.get("token")
    if (!token) {
        return new Response("missing key")
    }
    try {
        let req_json = await request.json()
        let agentid = req_json.agentid

        const corpsecret = await WORKWECHAT.get(agentid + '_agentid')
        let access_token = await WORKWECHAT.get(agentid + '_accessToken')

        if (access_token === null) {
            access_token = await getAccessToken(corpsecret)
            await WORKWECHAT.put(agentid + "_accessToken", access_token, { expirationTtl: 7000 })
        }
        return await postMessage(access_token, JSON.stringify(req_json))
    }
    catch (err) {
         return new Response(err)
    }
}