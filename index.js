import {add2FAHandler, getTokenHandler} from "./controller/2fa";

const Router = require('./router')
import workerWechatHandler from "./controller/workwechat"


addEventListener('fetch', event => {
    event.respondWith(handleRequest(event.request).catch(
      (err) => new Response(err.stack, { status: 500 })
    ))
})

async function handleRequest(request) {
    const r = new Router()
    // Replace with the appropriate paths and handlers
    // r.get('.*/bar', () => new Response('responding for /bar'))
    r.post('/message', request => workerWechatHandler(request))
    r.get('/2fa', request => getTokenHandler(request))
    r.post('/2fa', request => add2FAHandler(request))

    // r.post('.*/foo.*', request => handler(request))
    // r.get('/demos/router/foo', request => fetch(request)) // return the response from the origin
    r.get('/', () => new Response('Hello worker!')) // return a default message for the root route
    return r.route(request);
}


