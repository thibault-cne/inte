import { getAPI } from "@/apis/axios-api";
import { authStore } from "@/store/authStore";

function refreshToken() {
  const token = "Bearer ".concat(authStore.getters.refreshToken);
  const header = {
    Authorization: token,
    "Content-Type": "application/json",
  };
  return new Promise((resolve, reject) => {
    getAPI
      .get("/login/refresh", {
        headers: header,
      })
      .then((response) => {
        authStore.commit("setAccessToken", response.data.access_token);
        location.reload();
        resolve();
      })
      .catch((error) => {
        if (error.response) {
          if (error.response.status === 401) {
            authStore.commit("destroyToken");
            location.reload();
          }
        }
        reject(error);
      });
  });
}

export { refreshToken };
