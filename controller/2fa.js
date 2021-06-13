import {generateToken} from "../services/2fa";

const twoFaDataKey = '2fa_data'

export async function getTokenHandler(request) {
    const searchParams = new URL(request.url).searchParams
    const allData = JSON.parse(await twoFA.get(twoFaDataKey)) || {}
    if (searchParams.get('password') !== twoFA_password) return new Response('', {status: 403})
    if (searchParams.get('name') === null) return new Response(JSON.stringify(Object.keys(allData)))

    const data = allData[searchParams.get('name')]
    if (data === undefined) return new Response('', {status: 404})

    const password_key = await getCryptoKey(searchParams.get("password"))
    const decrypt = await crypto.subtle.decrypt(
        {name: "AES-GCM", iv: Buffer.from(data.iv)},
        password_key,
        base64ToUint8Array(data.secret)
    );
    const secret = Buffer.from(decrypt).toString('utf-8')
    let result = generateToken(secret)
    return new Response(JSON.stringify(result))

}


export async function add2FAHandler(request) {
    const requestData = await request.json()
    const checkingKeys = ['name', 'secret']
    for (const key of checkingKeys) {
        if (requestData[key] === undefined) return new Response(`field "${key}" missing`, {status: 400})
    }

    const iv = crypto.getRandomValues(new Uint8Array(12));
    const password_key = await getCryptoKey()

    const encrypted_secret = await crypto.subtle.encrypt(
        {name: "AES-GCM", iv},
        password_key,
        new TextEncoder().encode(requestData['secret']), // The data you want to hash as an ArrayBuffer
    )
    const encrypt_str = arrayBufferToBase64(encrypted_secret)
    let result = {"secret": encrypt_str, iv: [...iv], name: requestData.name}
    let allData = JSON.parse(await twoFA.get(twoFaDataKey)) || {}
    allData[requestData.name] = result
    await twoFA.put(twoFaDataKey, JSON.stringify(allData))
    return new Response(JSON.stringify(result))
}


function getCryptoKey(password = twoFA_password) {
    const padding = '0'.repeat((128 - password.length * 8) / 8);
    return crypto.subtle.importKey(
        "raw",
        new TextEncoder().encode(password + padding),
        "AES-GCM",
        true,
        ["encrypt", "decrypt"])
}

function arrayBufferToBase64(buffer) {
    let binary = '';
    const bytes = new Uint8Array(buffer);
    const len = bytes.byteLength;
    for (let i = 0; i < len; i++) {
        binary += String.fromCharCode(bytes[i]);
    }
    return btoa(binary);
}

function base64ToUint8Array(base64String) {
    const padding = '='.repeat((4 - base64String.length % 4) % 4);
    const base64 = (base64String + padding)
        .replace(/\-/g, '+')
        .replace(/_/g, '/');

    const rawData = atob(base64);
    const outputArray = new Uint8Array(rawData.length);

    for (let i = 0; i < rawData.length; ++i) {
        outputArray[i] = rawData.charCodeAt(i);
    }
    return outputArray;
}
