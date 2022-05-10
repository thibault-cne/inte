import { getAPI } from "@/apis/axios-api";
import { refreshToken } from "@/requests/refreshRequests";
import { createHeader } from "@/requests/createHeader";

function postRequest(data, url, headerType) {
  const header = createHeader(headerType);
  return new Promise((resolve, reject) => {
    getAPI
      .post(url, data, { headers: header })
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

export { postRequest };
