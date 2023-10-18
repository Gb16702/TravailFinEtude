import jwt from "jsonwebtoken"
import { JWT_SECRET_KEY } from "$env/static/private";

export function verifyToken(token: string |undefined) {
    try {
        return jwt.verify(token, JWT_SECRET_KEY)
    }

    catch(err) {
        console.log(err);
        throw new Error("Invalid token");
    }
}