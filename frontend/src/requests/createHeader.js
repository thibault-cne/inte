import { authStore } from "@/store/authStore";

function createHeader(headerType) {
  let header;
  let contentType;
  if (headerType === "file") {
    contentType = "multipart/form-data";
  } else if (headerType === "json") {
    contentType = "application/json";
  } else if (headerType === "data") {
    contentType = "multipart/form-data";
  } else {
    contentType = "application/x-www-form-urlencoded";
  }
  if (authStore.getters.loggedIn) {
    const token = "Bearer ".concat(authStore.getters.accessToken);
    header = {
      Authorization: token,
      "Content-Type": contentType,
    };
  } else {
    header = {
      "Content-Type": contentType,
    };
  }

  return header;
}

export { createHeader };
