import { getAPI } from "./axios";
import { getRequest } from "./getRequests";

async function isLogged(): Object {
    let resp = await getRequest("/api/auth/status", "");
    return resp.data;
}

export { getRequest };


  