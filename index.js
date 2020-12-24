const Router = require('./router')
import { getAccessToken, postMessage } from "./services/workwechat"
/**
 * Example of how router can be used in an application
 *  */
addEventListener('fetch', event => {
    event.respondWith(handleRequest(event.request))
})

function handler(request) {
    const init = {
        headers: { 'content-type': 'application/json' },
    }
    const body = JSON.stringify({ some: 'json' })
    return new Response(body, init)
}

async function handleRequest(request) {
    const r = new Router()
    // Replace with the appropriate paths and handlers
    // r.get('.*/bar', () => new Response('responding for /bar'))
    r.post('/message', request => workerWechatHandler(request))
    // r.post('.*/foo.*', request => handler(request))
    // r.get('/demos/router/foo', request => fetch(request)) // return the response from the origin

    r.get('/', () => new Response('Hello worker!')) // return a default message for the root route

    return r.route(request);
}



async function workerWechatHandler(request) {
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
        const result = await postMessage(access_token, JSON.stringify(req_json))
        return result
    }
    catch (err) {
         return new Response(err)
    }
}