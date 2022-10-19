import { LoggedIn } from "@/models/LoggedIn";
import { getRequest } from "./getRequests";

async function isLogged(): Promise<LoggedIn> {
    let resp = await getRequest("/api/auth/status", "");
    return <LoggedIn>resp.data;
}

export { isLogged };


  