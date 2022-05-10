import { getAPI } from "@/apis/axios-api";
import { refreshToken } from "@/requests/refreshRequests";
import { createHeader } from "@/requests/createHeader";

function getRequest(url, headerType) {
  let header = createHeader(headerType);
  return new Promise((resolve, reject) => {
    getAPI
      .get(url, { headers: header })
      .then((response) => {
        resolve(response);
      })
      .catch((error) => {
        if (error.response) {
          if (error.response.status === 401) {
            refreshToken();
          }
        }
        reject(error);
      });
  });
}

export { getRequest };
