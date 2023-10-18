import { redirect, type Handle } from "@sveltejs/kit";
import { verifyToken } from "./lib/verifyToken";

export const handle: Handle = async ({event, resolve}) => {

    let session: boolean = false;
    let sessionToken: string | undefined;
    if(event.cookies.get("session"))  {
        sessionToken = verifyToken(event.cookies.get("session"));
        if(sessionToken) {
            session = true;
        };
    };

    if(session) {
        console.log(sessionToken);

        if(event.url.pathname.startsWith("/auth")) {
            throw redirect(303, "/");
        };
    }

    return await resolve(event);
}
