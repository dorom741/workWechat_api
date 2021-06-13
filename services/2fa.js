// const tfa = require('2fa');
import tfa from '2fa'

const node2fa = require("node-2fa")

export function generateToken(secret) {
    return node2fa.generateToken(secret)
}

