import { LoggedIn } from "@/models/LoggedIn";
import { getRequest } from "./getRequests";

async function isLogged(): Promise<LoggedIn> {
  const resp = await getRequest("/auth/status", "");
  return <LoggedIn>resp.data;
}

export { isLogged };
