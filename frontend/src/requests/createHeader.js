import { store } from "@/store";

function createHeader(headerType) {
  let header;
  let contentType;
  if (headerType === "file") {
    contentType = "multipart/form-data";
  } else if (headerType === "json") {
    contentType = "application/json";
  } else {
    contentType = "application/x-www-form-urlencoded";
  }
  if (store.getters.loggedIn) {
    const token = "Bearer ".concat(store.getters.accessToken);
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
