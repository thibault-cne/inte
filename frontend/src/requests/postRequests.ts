import { getAPI } from "./axios";
import { createHeader } from "./createHeader";

function postRequest(data: Record<string, string>, url: string, headerType: string) {
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
              // refreshToken();
            }
          }
          reject(error);
        });
    });
  }
  
  export { postRequest };
  

