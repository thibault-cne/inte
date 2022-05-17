import axios from "axios";

let base_backend_url = `http://${process.env.VUE_APP_BACKEND_DOMAIN}:${process.env.VUE_APP_BACKEND_PORT}/`;

const getAPI = axios.create({
  baseURL: base_backend_url,
  headers: {
    "Content-Type": "application/json",
  },
});

export { getAPI };
